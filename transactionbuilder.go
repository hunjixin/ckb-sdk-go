package ckb_sdk_go

import (
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/dchest/blake2b"
	)

type Witness [][]byte

type  TransactionBuilder struct {
	Version uint32
	Deps []OutPoint
	Inputs []CellInput
	Outputs []CellOutput
	Witnesses []Witness
}

func FromTransction(tx Transaction) *TransactionBuilder {
	return &TransactionBuilder{
		Version:tx.Version,
		Deps:tx.Deps,
		Inputs:tx.Inputs,
		Outputs:tx.Outputs,
		Witnesses:tx.Witnesses,
	}
}

func (builder *TransactionBuilder) SetVersion(version uint32) *TransactionBuilder{
	builder.Version= version
	return builder
}

func (builder *TransactionBuilder) AppendOutPoint(dep OutPoint) *TransactionBuilder{
	builder.Deps= append(builder.Deps, dep)
	return builder
}

func (builder *TransactionBuilder) ClearDeps(dep OutPoint) *TransactionBuilder{
	builder.Deps = []OutPoint{}
	return builder
}

func (builder *TransactionBuilder) AppendInput(input CellInput) *TransactionBuilder{
	builder.Inputs = append(builder.Inputs,input)
	return builder
}

func (builder *TransactionBuilder) ClearInput(input CellInput) *TransactionBuilder{
	builder.Inputs = []CellInput{}
	return builder
}

func (builder *TransactionBuilder) AppendOutput(output CellOutput) *TransactionBuilder{
	builder.Outputs = append(builder.Outputs,output)
	return builder
}

func (builder *TransactionBuilder) ClearOutput() *TransactionBuilder{
	builder.Outputs = []CellOutput{}
	return builder
}

func (builder *TransactionBuilder) AppendWitness(witness Witness) *TransactionBuilder{
	builder.Witnesses = append(builder.Witnesses, witness)
	return builder
}

func (builder *TransactionBuilder) ClearWitness() *TransactionBuilder{
	//state.write(self.witness_hash().as_fixed_bytes())
}


type RawTransaction struct{
	Version uint32
	Deps  []OutPoint
	Inputs []CellInput
	Outputs []CellOutput
}


func (builder *TransactionBuilder) Hash() []byte{
	builder.Witnesses = []Witness{}
	return builder
}

func  (builder *TransactionBuilder)Build() Transaction {
	return Transaction{
		Version:builder.Version,
		Deps:builder.Deps,
		Inputs:builder.Inputs,
		Outputs:builder.Outputs,
		Witnesses:builder.Witnesses,
	}
}

func (tx *Transaction)  TxHash() [32]byte {
	rawTx := &RawTransaction{
		Version:tx.Version,
		Deps: tx.Deps,
		Inputs:tx.Inputs,
		Outputs:tx.Outputs,
	}
	bytes, _ := Marshal(rawTx)
	hashBytes := blake2b.Sum256(bytes)
	return hashBytes
}
func (tx *Transaction)  WitnessHash() [32]byte {
	bytes, _ := Marshal(tx)
	hashBytes := blake2b.Sum256(bytes)
	return hashBytes
}



func sign(tx TransactionBuilder, priv *secp256k1.PrivateKey) Transaction {

	raw := tx.Build()
	data := []byte{}
	hash := blake2b.Sum256(data)
	sig, _  := priv.Sign(hash[:])

	tx.AppendWitness([][]byte{sig.Serialize()})
	return tx.Build()
}

