package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/hunjixin/automapper"
	"github.com/ybbus/jsonrpc"
	"hello"
	"reflect"
)

func main() {
	privBytes,_ := hex.DecodeString("2a73ebc64fe94a5fe22b06cb5bd322758da2650fd22d80ae7ecc9efe12703047")
	priv,_ := secp256k1.PrivKeyFromBytes(privBytes)
	address := ckb_sdk_go.PublicKeyToAddress(priv.PubKey(), ckb_sdk_go.TestnetAddressOptions)
	fmt.Println(address)  //ckt1qyqqa49550zjkj4z7x4zqqdxdkp4czwuqu6q92237h

	builder:= ckb_sdk_go.NewTransactionBuilder()
	builder.Inputs=[]ckb_sdk_go.CellInput{
		{
			Previous_output : ckb_sdk_go.OutPoint{
				Cell:       &ckb_sdk_go.CellOutPoint{
					Tx_hash :*ckb_sdk_go.StringToHash("795b53f8b701514c0bdc558cb1396c1ba4ab31d382ec0bda3813aa9cd5327af6"),
					Index   :0,
				},
				Block_hash: nil,
			},
			Since      :     0,
		},
	}
	arg, _ := hex.DecodeString("1d36cd579406bf915cb266ada98b5dc22a82f13a")
//ckt1qyqp6dkd272qd0u3tjexdtdf3dwuy25z7yaqragwxr
	builder.Outputs = []ckb_sdk_go.CellOutput{
		{
			Capacity: 70*100000000,
			Data:     []byte{},
			Lock:     ckb_sdk_go.Script{
				Args    :   [][]byte{arg},
				Code_hash : *ckb_sdk_go.StringToHash("54811ce986d5c3e57eaafab22cdd080e32209e39590e204a99b32935f835a13c"),
				Hash_type : ckb_sdk_go.Data,
			},
			Type_:    nil,
		},
	}

	sig := ckb_sdk_go.SignTx(builder, priv)
	fmt.Println(sig)
	t := reflect.TypeOf(ckb_sdk_go.RpcTransaction{})
	rpcTx, _ := automapper.Mapper(builder.Build(), t)
	client := ckb_sdk_go.NewCkbClient("http://127.0.0.1:8114")
	str, err := client.SendTransaction(rpcTx.(ckb_sdk_go.RpcTransaction))
	fmt.Println(err)
	fmt.Println(str)
	request := &jsonrpc.RPCRequest{
		Method:  "send_transaction",
		Params: []ckb_sdk_go.RpcTransaction {rpcTx.(ckb_sdk_go.RpcTransaction)},
		JSONRPC: "2.0",
	}

	xxxx, _ := json.Marshal(request)
	fmt.Println(string(xxxx))

}
