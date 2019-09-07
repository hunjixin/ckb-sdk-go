package crypto

import (
	"github.com/dchest/blake2b"
)

var (
	config = &blake2b.Config{
		Size:   32,
		Person: []byte("ckb-default-hash"),
	}
)

func Blake20(data []byte) []byte {
	hash, _ := blake2b.New(config)
	hash.Write(data)
	return hash.Sum(nil)[0:20]
}


func Black256(bytes []byte) []byte {
	hash, _ := blake2b.New(config)
	hash.Write(bytes)
	return hash.Sum(nil)
}
func Black256M(bytes ...[]byte) []byte {
	hash, _ := blake2b.New(config)
	for _, b := range bytes {
		hash.Write(b)
	}
	return hash.Sum(nil)
}
