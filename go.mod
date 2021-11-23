module ckb-sdk-go

go 1.12

replace github.com/golang/crypto v0.0.0-20190820162420-60c769a6c586 => golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586

replace github.com/golang/sys v0.0.0-20190813064441-fde4db37ae7a => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a

require (
	github.com/adiabat/bech32 v0.0.0-20170505011816-6289d404861d
	github.com/alecthomas/participle v0.3.0
	github.com/dchest/blake2b v1.0.0
	github.com/decred/dcrd/dcrec/secp256k1 v1.0.2
	github.com/ethereum/go-ethereum v1.10.10
	github.com/hraban/lrucache v0.0.0-20150309112052-a1cd14943f73 // indirect
	github.com/hunjixin/automapper v0.0.0-20190819023506-f5120f46a84f
	github.com/ybbus/jsonrpc v2.1.2+incompatible
)
