package main

import (
	"ckb-sdk-go"
	address2 "ckb-sdk-go/address"
	"ckb-sdk-go/core"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/hunjixin/automapper"
	"github.com/ybbus/jsonrpc"
	"reflect"
)

func main() {
	privBytes,_ := hex.DecodeString("2a73ebc64fe94a5fe22b06cb5bd322758da2650fd22d80ae7ecc9efe12703047")
	priv,_ := secp256k1.PrivKeyFromBytes(privBytes)
	address := address2.PublicKeyToAddress(priv.PubKey(), address2.TestnetAddressOptions)
	fmt.Println(address)  //ckt1qyqqa49550zjkj4z7x4zqqdxdkp4czwuqu6q92237h

	builder:= ckb_sdk_go.NewTransactionBuilder()
	builder.Deps = []core.OutPoint{
		{
			Cell:       &core.CellOutPoint{
				Tx_hash :*core.StringToHash("41ec4d87557b87e9e5c6ff942e8f0e34e40f8347e3a84587ab4c04eebeed7791"),
				Index   :1,
			},
			Block_hash:  core.StringToHash("1e673afbff501eaf0b43f6f5d1f505374a5db4c689855668e51376885c4d0ec6"),
		},
	}

	builder.Inputs=[]core.CellInput{
		{
			Previous_output : core.OutPoint{
				Cell:       &core.CellOutPoint{
					Tx_hash :*core.StringToHash("8353ffa58e5b13001ee5c480db19b92e5f5e300862dc865bf73a895b8e56a4c7"),
					Index   :0,
				},
				Block_hash: nil,
			},
			Since      :     0,
		},
	}
	arg, _ := hex.DecodeString("1d36cd579406bf915cb266ada98b5dc22a82f13a")
//ckt1qyqp6dkd272qd0u3tjexdtdf3dwuy25z7yaqragwxr
	builder.Outputs = []core.CellOutput{
		{
			Capacity: 70*100000000,
			Data:     []byte{},
			Lock:     core.Script{
				Args    :   [][]byte{arg},
				Code_hash : *core.StringToHash("54811ce986d5c3e57eaafab22cdd080e32209e39590e204a99b32935f835a13c"),
				Hash_type : core.Data,
			},
			Type_:    nil,
		},
	}
//	builder.AppendWitness([][]byte{[]byte{1,2,3}})
	ckb_sdk_go.SignTx(builder, priv)
	tx := builder.Build()
	hash := tx.TxHash()
	fmt.Println(hex.EncodeToString(hash[:]))
	t := reflect.TypeOf(core.RpcTransaction{})
	rpcTx, _ := automapper.Mapper(builder.Build(), t)
	client := ckb_sdk_go.NewCkbClient("http://127.0.0.1:8114")
	str, err := client.SendTransaction(rpcTx.(core.RpcTransaction))
	fmt.Println(err)
	fmt.Println(str)
	request := &jsonrpc.RPCRequest{
		Method:  "send_transaction",
		Params: []core.RpcTransaction{rpcTx.(core.RpcTransaction)},
		JSONRPC: "2.0",
	}

	xxxx, _ := json.Marshal(request)
	fmt.Println(string(xxxx))

}
