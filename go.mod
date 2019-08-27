module hello

go 1.12

replace github.com/golang/crypto v0.0.0-20190820162420-60c769a6c586 => golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586

replace github.com/golang/sys v0.0.0-20190813064441-fde4db37ae7a => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a

require (
	github.com/adiabat/bech32 v0.0.0-20170505011816-6289d404861d
	github.com/alecthomas/kong v0.2.1-0.20190721020729-f7d3d9bfb5ed
	github.com/alecthomas/participle v0.3.0
	github.com/alecthomas/repr v0.0.0-20181024024818-d37bc2a10ba1
	github.com/dchest/blake2b v1.0.0
	github.com/decred/dcrd/dcrec/secp256k1 v1.0.2
	github.com/onsi/gomega v1.6.0 // indirect
	github.com/ybbus/jsonrpc v2.1.2+incompatible
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 // indirect
	golang.org/x/sys v0.0.0-20190412213103-97732733099d // indirect
)
