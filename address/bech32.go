package address

import (
	"bytes"
	"fmt"
	"strings"
)

var charset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

var generator = []int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}

func polymod(values []int) int {
	chk := 1
	for _, v := range values {
		top := chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ v
		for i := 0; i < 5; i++ {
			if (top>>uint(i))&1 == 1 {
				chk ^= generator[i]
			}
		}
	}
	return chk
}

func hrpExpand(hrp string) []int {
	ret := []int{}
	for _, c := range hrp {
		ret = append(ret, int(c>>5))
	}
	ret = append(ret, 0)
	for _, c := range hrp {
		ret = append(ret, int(c&31))
	}
	return ret
}

func verifyChecksum(hrp string, data []int) bool {
	return polymod(append(hrpExpand(hrp), data...)) == 1
}

func createChecksum(hrp string, data []int) []int {
	values := append(append(hrpExpand(hrp), data...), []int{0, 0, 0, 0, 0, 0}...)
	mod := polymod(values) ^ 1
	ret := make([]int, 6)
	for p := 0; p < len(ret); p++ {
		ret[p] = (mod >> uint(5*(5-p))) & 31
	}
	return ret
}

// Encode encodes hrp(human-readable part) and data(32bit data array), returns Bech32 / or error
// if hrp is uppercase, return uppercase Bech32
func Encode(hrp string, data []int) (string, error) {
	if (len(hrp) + len(data) + 7) > 90 {
		return "", fmt.Errorf("too long : hrp length=%d, data length=%d", len(hrp), len(data))
	}
	if len(hrp) < 1 {
		return "", fmt.Errorf("invalid hrp : hrp=%v", hrp)
	}
	for p, c := range hrp {
		if c < 33 || c > 126 {
			return "", fmt.Errorf("invalid character human-readable part : hrp[%d]=%d", p, c)
		}
	}
	if strings.ToUpper(hrp) != hrp && strings.ToLower(hrp) != hrp {
		return "", fmt.Errorf("mix case : hrp=%v", hrp)
	}
	lower := strings.ToLower(hrp) == hrp
	hrp = strings.ToLower(hrp)
	combined := append(data, createChecksum(hrp, data)...)
	var ret bytes.Buffer
	ret.WriteString(hrp)
	ret.WriteString("1")
	for idx, p := range combined {
		if p < 0 || p >= len(charset) {
			return "", fmt.Errorf("invalid data : data[%d]=%d", idx, p)
		}
		ret.WriteByte(charset[p])
	}
	if lower {
		return ret.String(), nil
	}
	return strings.ToUpper(ret.String()), nil
}

// Decode decodes bechString(Bech32) returns hrp(human-readable part) and data(32bit data array) / or error
func Decode(bechString string) (string, []int, error) {
	if len(bechString) > 90 {
		return "", nil, fmt.Errorf("too long : len=%d", len(bechString))
	}
	if strings.ToLower(bechString) != bechString && strings.ToUpper(bechString) != bechString {
		return "", nil, fmt.Errorf("mixed case")
	}
	bechString = strings.ToLower(bechString)
	pos := strings.LastIndex(bechString, "1")
	if pos < 1 || pos+7 > len(bechString) {
		return "", nil, fmt.Errorf("separator '1' at invalid position : pos=%d , len=%d", pos, len(bechString))
	}
	hrp := bechString[0:pos]
	for p, c := range hrp {
		if c < 33 || c > 126 {
			return "", nil, fmt.Errorf("invalid character human-readable part : bechString[%d]=%d", p, c)
		}
	}
	data := []int{}
	for p := pos + 1; p < len(bechString); p++ {
		d := strings.Index(charset, fmt.Sprintf("%c", bechString[p]))
		if d == -1 {
			return "", nil, fmt.Errorf("invalid character data part : bechString[%d]=%d", p, bechString[p])
		}
		data = append(data, d)
	}
	if !verifyChecksum(hrp, data) {
		return "", nil, fmt.Errorf("invalid checksum")
	}
	return hrp, data[:len(data)-6], nil
}
