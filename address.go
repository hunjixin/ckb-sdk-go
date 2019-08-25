package ckb_sdk_go

import (
	"github.com/decred/dcrd/dcrec/secp256k1"
	"golang.org/x/crypto/blake2b"
	"github.com/adiabat/bech32"
)

type AddressPrefix string
const(
	Mainnet AddressPrefix = "ckb"
	Testnet AddressPrefix = "kt"
)

type AddressType byte
const(
	BinHash AddressType = 0x00
	HashIdx AddressType = 0x01
)


//export type CodeHashIndex = '0x00' | string

type CodeHashIndex byte
type AddressOptions struct  {
	prefix AddressPrefix
	type_ AddressType
	codeHashIndex  CodeHashIndex
}

var (
	defaultAddressOptions = &AddressOptions {
		prefix:Testnet,
		type_: HashIdx,
		codeHashIndex : 0x00,
	}
)

func PublicKeyToAddress(pubkey *secp256k1.PublicKey, option *AddressOptions) string {
	if option == nil {
		option = defaultAddressOptions
	}
	identifier := blake2b.Sum256(pubkey.Serialize())
	return bech32Address(identifier[0:20], option)
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