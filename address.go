package ckb_sdk_go

import (
	"encoding/hex"
	"github.com/adiabat/bech32"
	"github.com/dchest/blake2b"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"strings"
)

var (
	config = &blake2b.Config{
		Size:   32,
		Person: []byte("ckb-default-hash"),
	}
)

type AddressPrefix string

const (
	Mainnet AddressPrefix = "ckb"
	Testnet AddressPrefix = "ckt"
)

type AddressType byte

const (
	BinHash AddressType = 0x00
	HashIdx AddressType = 0x01
)

//export type CodeHashIndex = '0x00' | string

type CodeHashIndex byte
type AddressOptions struct {
	prefix        AddressPrefix
	type_         AddressType
	codeHashIndex CodeHashIndex
}

var (
	defaultAddressOptions = &AddressOptions{
		prefix:        Mainnet,
		type_:         HashIdx,
		codeHashIndex: 0x00,
	}
)

func EnsurePrivateKey(priStr string) *secp256k1.PrivateKey {
	priStr = strings.Replace(priStr, "0x", "", -1)
	if len(priStr)%2 != 0 {
		priStr = "0" + priStr
	}
	priBytes, _ := hex.DecodeString(priStr)
	pri, _ := secp256k1.PrivKeyFromBytes(priBytes)
	return pri
}

func PrivKeyToAddress(priv *secp256k1.PrivateKey, option *AddressOptions) string {
	return PublicKeyToAddress(priv.PubKey(), option)
}

func PublicKeyToAddress(pubkey *secp256k1.PublicKey, option *AddressOptions) string {
	if option == nil {
		option = defaultAddressOptions
	}
	identifier := generateIdentifier(pubkey)
	return bech32Address(identifier, option)
}

func generateIdentifier(pubkey *secp256k1.PublicKey) []byte {
	bytes := pubkey.SerializeCompressed()
	hash, _ := blake2b.New(config)
	hash.Write(bytes)
	return hash.Sum(nil)[0:20]
}

func bech32Address(identifier []byte, addrOption *AddressOptions) string {
	if addrOption == nil {
		addrOption = defaultAddressOptions
	}
	payload := toAddressPayload(identifier, addrOption.type_, addrOption.codeHashIndex)
	return bech32.Encode(string(addrOption.prefix), payload)
}

func toAddressPayload(identifier []byte, addressType AddressType, codeHashIndex CodeHashIndex) []byte {
	result := []byte{}
	result = append(result, byte(addressType))
	result = append(result, byte(codeHashIndex))
	result = append(result, identifier...)
	return result
}
