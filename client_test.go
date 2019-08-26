package ckb_sdk_go

import (
	"fmt"
	"testing"
)

var (
	client = NewCkbClient("http://127.0.0.1:8114")
)


func TestGetBlockByNumber(t *testing.T) {

	block, err := client.GetBlockByNumber(1)
	if err !=nil {
		t.Error(err)
	}
	if block.Header.Number != "1" {
		t.Errorf("export block number %s but got %s", "1", block.Header.Number)
	}
}

func TestGetBlock(t *testing.T) {
	hash := "0x9e28ef017a90d88f9a863c57a11a4bd36e216df1d8a989e45f367e377498bf38"
	block, err := client.GetBlock(hash)
	if err !=nil {
		t.Error(err)
	}
	if block.Header.Number != "1" {
		t.Errorf("export block hash %s but got %s", hash, block.Header.Hash)
	}
}

func TestGetTipHeader(t *testing.T) {
	_, err := client.GetTipHeader()
	if err !=nil {
		t.Error(err)
	}
}

func TestGetCellsByLockHash(t *testing.T) {
	_, err := client.GetCellsByLockHash()
	if err !=nil {
		t.Error(err)
	}
}

func TestGetLiveCells(t *testing.T) {
	_, err := client.GetLiveCell(nil, nil)
	if err !=nil {
		t.Error(err)
	}
}

func TestGetTipBlockNumber(t *testing.T) {
	_, err := client.GetTipBlockNumber()
	if err !=nil {
		t.Error(err)
	}
}

func TestGetBlockchainInfo(t *testing.T) {
	_, err := client.GetBlockchainInfo()
	if err !=nil {
		t.Error(err)
	}
}

func TestSendTransaction(t *testing.T) {
	_, err := client.SendTransaction(nil)
	if err !=nil {
		t.Error(err)
	}
}


func TestLocalNodeInfo(t *testing.T) {
	_, err := client.LocalNodeInfo()
	if err !=nil {
		t.Error(err)
	}
}

func TestTxPoolInfo(t *testing.T) {
	_, err := client.TxPoolInfo()
	if err !=nil {
		t.Error(err)
	}
}

func TestGetPeers(t *testing.T) {
	_, err := client.GetPeers()
	if err !=nil {
		t.Error(err)
	}
}

func TestGetPeersState(t *testing.T) {
	_, err := client.GetPeersState()
	if err !=nil {
		t.Error(err)
	}
}

func TestGetCurrentEpoch(t *testing.T) {
	_, err := client.GetCurrentEpoch()
	if err !=nil {
		t.Error(err)
	}
}

func TestGetEpochByNumber(t *testing.T) {
	_, err := client.GetEpochByNumber(100)
	if err !=nil {
		t.Error(err)
	}
}


func TestDryRunTransaction(t *testing.T) {
	err := client.DryRunTransaction(nil)
	if err !=nil {
		t.Error(err)
	}
}

func TestDeindexLockHash(t *testing.T) {
	err := client.DeindexLockHash(nil)
	if err !=nil {
		t.Error(err)
	}
}

func TestGetLiveCellsByLockHash(t *testing.T) {
	_, err := client.GetLiveCellsByLockHash("0x1d36cd579406bf915cb266ada98b5dc22a82f13a",0,100)
	if err !=nil {
		t.Error(err)
	}
}

func TestGetLockHashIndexStates(t *testing.T) {
	_, err := client.GetLockHashIndexStates()
	if err !=nil {
		t.Error(err)
	}
}

func TestGetTransactionsByLockHash(t *testing.T) {
	_, err := client.GetTransactionsByLockHash("0x1d36cd579406bf915cb266ada98b5dc22a82f13a",0,100)
	if err !=nil {
		t.Error(err)
	}
}

func TestIndexLockHash(t *testing.T) {
	_, err := client.IndexLockHash("0x1d36cd579406bf915cb266ada98b5dc22a82f13a")
	if err !=nil {
		t.Error(err)
	}
}

func TestGetBannedAddresses(t *testing.T) {
	_, err := client.GetBannedAddresses()
	if err !=nil {
		t.Error(err)
	}
}


func TestGetHeaderByNumber(t *testing.T) {
	header, err := client.GetHeaderByNumber(1)
	if err !=nil {
		t.Error(err)
		return
	}
	if header.Number != "1" {
		t.Errorf("export block number %s but got %s", "1", header.Number)
	}
}

func TestGetHeader(t *testing.T) {
	hash := "0x9e28ef017a90d88f9a863c57a11a4bd36e216df1d8a989e45f367e377498bf38"
	header, err := client.GetHeader(hash)
	if err !=nil {
		t.Error(err)
	}
	if header.Number != "1" {
		t.Errorf("export block hash %s but got %s", hash, header.Hash)
	}
}

func TestGetCellbaseOutputCapacityDetails(t *testing.T) {
	hash := "0x9e28ef017a90d88f9a863c57a11a4bd36e216df1d8a989e45f367e377498bf38"
	cellbase, err := client.GetCellbaseOutputCapacityDetails(hash)
	if err !=nil {
		t.Error(err)
	}
	fmt.Println(cellbase)
}

func TestComputeTransactionHashMethod(t *testing.T) {
	rawTx := &RpcRawTransaction{}
	hash, err := client.ComputeTransactionHashMethod(rawTx)
	if err !=nil {
		t.Error(err)
	}
	fmt.Println(hash)
}


