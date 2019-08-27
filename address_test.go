package ckb_sdk_go

import (
	"encoding/hex"
	"testing"
)

var (
		testAddresOption = &AddressOptions {
		prefix:Testnet,
		type_: HashIdx,
		codeHashIndex : 0x00,
	}

)
func Test_addressgenerate(t *testing.T) {
	priv := EnsurePrivateKey("2a73ebc64fe94a5fe22b06cb5bd322758da2650fd22d80ae7ecc9efe12703047")
	addr := PublicKeyToAddress(priv.PubKey(), testAddresOption)
	if addr != "ckt1qyqqa49550zjkj4z7x4zqqdxdkp4czwuqu6q92237h" {
		t.Errorf("expect %s but got %s", "ckt1qyqqa49550zjkj4z7x4zqqdxdkp4czwuqu6q92237h", addr)
	}
}

func Test_addressgenerate2(t *testing.T) {
	priv := EnsurePrivateKey("e79f3207ea4980b7fed79956d5934249ceac4751a4fae01a0f7c4a96884bc4e3")
	addr := PublicKeyToAddress(priv.PubKey(), testAddresOption)
	if addr != "ckt1qyqrdsefa43s6m882pcj53m4gdnj4k440axqswmu83" {
		t.Errorf("expect %s but got %s", "e79f3207ea4980b7fed79956d5934249ceac4751a4fae01a0f7c4a96884bc4e3", addr)
	}
}

func Test_GenerateAddress(t *testing.T) {
	priv :=  EnsurePrivateKey("e79f3207ea4980b7fed79956d5934249ceac4751a4fae01a0f7c4a96884bc4e3")
	addr := "ckt1qyqrdsefa43s6m882pcj53m4gdnj4k440axqswmu83"
	option :=  &AddressOptions{
		prefix:        Testnet,
		type_:         HashIdx,
		codeHashIndex: 0x00,
	}
	newAddr := PublicKeyToAddress(priv.PubKey(), option)
	if addr != newAddr {
		t.Errorf("expect %s but got %s", newAddr, addr)
	}
}

func Test_GenerateAddressType(t *testing.T) {
	priv :=  EnsurePrivateKey("e79f3207ea4980b7fed79956d5934249ceac4751a4fae01a0f7c4a96884bc4e3")
	addr := "ckt1qyqrdsefa43s6m882pcj53m4gdnj4k440axqswmu83"
	option :=  &AddressOptions{
		prefix:        Testnet,
		type_:         0x01,
		codeHashIndex: 0x00,
	}
	newAddr := PublicKeyToAddress(priv.PubKey(), option)
	if addr != newAddr {
		t.Errorf("expect %s but got %s", newAddr, addr)
	}
}

func Test_GenerateAddressCodeHash(t *testing.T) {
	priv :=  EnsurePrivateKey("e79f3207ea4980b7fed79956d5934249ceac4751a4fae01a0f7c4a96884bc4e3")
	addr := "ckt1qyqrdsefa43s6m882pcj53m4gdnj4k440axqswmu83"
	option :=  &AddressOptions{
		prefix:        Testnet,
		type_:         0x01,
		codeHashIndex: 0x00,
	}
	newAddr := PublicKeyToAddress(priv.PubKey(), option)
	if addr != newAddr {
		t.Errorf("expect %s but got %s", newAddr, addr)
	}
}


func Test_Bech32(t *testing.T) {
	priv := EnsurePrivateKey("0e79f3207ea4980b7fed79956d5934249ceac4751a4fae01a0f7c4a96884bc4e")
	identifier := "2f663ae60e00153d223657c685a15604255b168b"

	 newIdentifier := hex.EncodeToString(generateIdentifier(priv.PubKey()))
	if newIdentifier != identifier {
		t.Errorf("expect %s but got %s", identifier, newIdentifier)
	}
}
