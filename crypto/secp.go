package crypto

import (
	"crypto/ecdsa"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignMesage(hash []byte, priv *secp256k1.PrivateKey) []byte {
	sig, _ := crypto.Sign(hash, (*ecdsa.PrivateKey)(priv))
	return sig
}
