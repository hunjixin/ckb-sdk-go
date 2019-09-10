module ckb-sdk-go

go 1.12

require (
	github.com/adiabat/bech32 v0.0.0-20170505011816-6289d404861d
	github.com/alecthomas/participle v0.3.0
	github.com/btcsuite/btcd v0.0.0-20190824003749-130ea5bddde3 // indirect
	github.com/dchest/blake2b v1.0.0
	github.com/decred/dcrd/dcrec/secp256k1 v1.0.2
	github.com/ethereum/go-ethereum v1.9.3
	github.com/hraban/lrucache v0.0.0-20150309112052-a1cd14943f73 // indirect
	github.com/hunjixin/automapper v0.0.0-20190819023506-f5120f46a84f
	github.com/onsi/gomega v1.6.0 // indirect
	github.com/ybbus/jsonrpc v2.1.2+incompatible
	golang.org/x/crypto v0.0.0 // indirect
	golang.org/x/net v0.0.0 // indirect
	golang.org/x/sync v0.0.0 // indirect
	golang.org/x/sys v0.0.0 // indirect

)

replace golang.org/x/crypto v0.0.0 => github.com/golang/crypto v0.0.0-20190907121410-71b5226ff739

replace golang.org/x/sys v0.0.0 => github.com/golang/sys v0.0.0-20190907184412-d223b2b6db03

replace golang.org/x/net v0.0.0 => github.com/golang/net v0.0.0-20190909003024-a7b16738d86b

replace golang.org/x/text v0.0.0 => github.com/golang/text v0.3.2

replace golang.org/x/sync v0.0.0 => github.com/golang/sync v0.0.0-20190423024810-112230192c58

replace golang.org/x/tools v0.0.0 => github.com/golang/tools v0.0.0-20190909030654-5b82db07426d

replace golang.org/x/xerrors v0.0.0 => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
