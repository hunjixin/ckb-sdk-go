package main

import (
	"strings"
)

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func LowToHeigh(str string) string {
	phases := strings.Split(str,"_")
	result := ""
	for _, phase := range phases {
		result += Capitalize(phase)
	}
	return result
}