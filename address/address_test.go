package address

import (
	"github.com/decred/dcrd/dcrec/secp256k1"
	"reflect"
	"testing"
)

func TestEnsurePrivateKey(t *testing.T) {
	type args struct {
		priStr string
	}
	tests := []struct {
		name string
		args args
		want *secp256k1.PrivateKey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnsurePrivateKey(tt.args.priStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnsurePrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrivKeyToAddress(t *testing.T) {
	type args struct {
		priv   *secp256k1.PrivateKey
		option *AddressOptions
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrivKeyToAddress(tt.args.priv, tt.args.option); got != tt.want {
				t.Errorf("PrivKeyToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublicKeyToAddress(t *testing.T) {
	type args struct {
		pubkey *secp256k1.PublicKey
		option *AddressOptions
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PublicKeyToAddress(tt.args.pubkey, tt.args.option); got != tt.want {
				t.Errorf("PublicKeyToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bech32Address(t *testing.T) {
	type args struct {
		identifier []byte
		addrOption *AddressOptions
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bech32Address(tt.args.identifier, tt.args.addrOption); got != tt.want {
				t.Errorf("bech32Address() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateIdentifier(t *testing.T) {
	type args struct {
		pubkey *secp256k1.PublicKey
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateIdentifier(tt.args.pubkey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toAddressPayload(t *testing.T) {
	type args struct {
		identifier    []byte
		addressType   AddressType
		codeHashIndex CodeHashIndex
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toAddressPayload(tt.args.identifier, tt.args.addressType, tt.args.codeHashIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toAddressPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}