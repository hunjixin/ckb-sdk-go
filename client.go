package ckb_sdk_go

import (
	"encoding/json"
	"github.com/ybbus/jsonrpc"
	"strconv"
)

type CkbClient struct {
	url string
	client jsonrpc.RPCClient
}

func NewCkbClient(url string) *CkbClient {
	client := jsonrpc.NewClient(url)
	return &CkbClient{
		url:url,
		client:client,
	}
}


func (ckbClient *CkbClient)  GetBlockByNumber(number int) (*RpcBlock, error){
	rpcBlock := &RpcBlock{}
	err := ckbClient.client.CallFor(rpcBlock, "get_block_by_number", strconv.Itoa(number))
	if err != nil {
		return nil, err
	}
	return rpcBlock, nil
}

func (ckbClient *CkbClient)  GetBlock(hash Hash) (*RpcBlock, error){
	rpcBlock := &RpcBlock{}
	err := ckbClient.client.CallFor(rpcBlock, "get_block", hash)
	if err != nil {
		return nil, err
	}
	return rpcBlock, nil
}

func (ckbClient *CkbClient)  GetTransaction(hash Hash) (*RpcTransaction, error){
	rpcTx := &RpcTransaction{}
	err := ckbClient.client.CallFor(rpcTx, "get_transaction", hash)
	if err != nil {
		return nil, err
	}
	return rpcTx, nil
}

func (ckbClient *CkbClient) GetBlockHash(number int) (Hash, error){
	hash := ""
	err := ckbClient.client.CallFor(&hash, "get_block_hash", number)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (ckbClient *CkbClient) GetTipHeader() (*RpcHeader, error){
    tipHeader := &RpcHeader{}
	err := ckbClient.client.CallFor(tipHeader, "get_tip_header")
	if err != nil {
		return nil, err
	}
	return tipHeader, nil
}

func (ckbClient *CkbClient) GetCellsByLockHash() ([]*RpcCell, error){
	cell := []*RpcCell{}
	err := ckbClient.client.CallFor(&cell, "get_cells_by_lock_hash")
	if err != nil {
		return nil, err
	}
	return cell, nil
}

func (ckbClient *CkbClient) GetLiveCell(hash *Hash, rpcPoint *RpcOutPoint) (*RpcCellWithStatus, error){
	cell := &RpcCellWithStatus{}
	err := ckbClient.client.CallFor(cell, "get_live_cell", hash, rpcPoint)
	if err != nil {
		return nil, err
	}
	return cell, nil
}

func (ckbClient *CkbClient) GetTipBlockNumber() (int, error){
	blockNumber := ""
	err := ckbClient.client.CallFor(&blockNumber, "get_tip_block_number")
	if err != nil {
		return -1, err
	}
	num, err := strconv.Atoi(blockNumber)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func (ckbClient *CkbClient) GetBlockchainInfo() (*RpcBlockchainInfo, error){
	blockChainInfo := &RpcBlockchainInfo{}
	err := ckbClient.client.CallFor(blockChainInfo, "get_blockchain_info")
	if err != nil {
		return nil, err
	}
	return blockChainInfo, nil
}

func (ckbClient *CkbClient) SendTransaction(transaction *RawTransaction) (*Hash, error){
	hash := Hash("")
	err := ckbClient.client.CallFor(&hash, "send_transaction", transaction)
	if err != nil {
		return nil, err
	}
	return &hash, nil
}

func (ckbClient *CkbClient) LocalNodeInfo() (*RpcNodeInfo, error){
	nodeInfo := &RpcNodeInfo{}
	err := ckbClient.client.CallFor(nodeInfo, "local_node_info")
	if err != nil {
		return nil, err
	}
	return nodeInfo, nil
}

func (ckbClient *CkbClient) TxPoolInfo() (*RpcTxPoolInfo, error){
	poolInfo := &RpcTxPoolInfo{}
	err := ckbClient.client.CallFor(poolInfo, "tx_pool_info")
	if err != nil {
		return nil, err
	}
	return poolInfo, nil
}

func (ckbClient *CkbClient) GetPeers() ( []*RpcNodeInfo, error){
	peers := []*RpcNodeInfo{}
	err := ckbClient.client.CallFor(&peers, "get_peers")
	if err != nil {
		return nil, err
	}
	return peers, nil
}

func (ckbClient *CkbClient) GetPeersState() ( []*PeersState, error){
	peers := []*PeersState{}
	err := ckbClient.client.CallFor(&peers, "get_peers_state")
	if err != nil {
		return nil, err
	}
	return peers, nil
}

func (ckbClient *CkbClient) GetCurrentEpoch() (*RpcEpoch, error){
	epoch := &RpcEpoch{}
	err := ckbClient.client.CallFor(epoch, "get_current_epoch")
	if err != nil {
		return nil, err
	}
	return epoch, nil
}

func (ckbClient *CkbClient) GetEpochByNumber(number int) (*RpcEpoch, error){
	epoch := &RpcEpoch{}
	err := ckbClient.client.CallFor(epoch, "get_epoch_by_number", strconv.Itoa(number))
	if err != nil {
		return nil, err
	}
	return epoch, nil
}

func (ckbClient *CkbClient) DryRunTransaction(rpcRawTransaction *RpcRawTransaction) error{
	_, err := ckbClient.client.Call("dry_run_transaction", rpcRawTransaction)
	if err != nil {
		return  err
	}
	return nil
}

func (ckbClient *CkbClient) DeindexLockHash(hash *Hash) error{
	_, err := ckbClient.client.Call("deindex_lock_hash", hash)
	if err != nil {
		return err
	}
	return nil
}

func (ckbClient *CkbClient) GetLiveCellsByLockHash(hash Hash, pageNumber int, pageSize int) ( []*RpcCell, error){
	cells := []*RpcCell{}
	err := ckbClient.client.CallFor(&cells, "get_live_cells_by_lock_hash", hash, pageNumber, pageSize)
	if err != nil {
		return nil, err
	}
	return cells, nil
}

func (ckbClient *CkbClient) GetLockHashIndexStates() (*RpcLockHashIndexStates, error){
	rpcLockHashIndexStates := &RpcLockHashIndexStates{}
	err := ckbClient.client.CallFor(rpcLockHashIndexStates, "get_lock_hash_index_states")
	if err != nil {
		return nil, err
	}
	return rpcLockHashIndexStates, nil
}

func (ckbClient *CkbClient) GetTransactionsByLockHash(hash Hash, pageNumber int, pageSize int) (*RpcTransactionByLockHash, error){
	rpcTransactionByLockHash := &RpcTransactionByLockHash{}
	err := ckbClient.client.CallFor(rpcTransactionByLockHash, "get_transactions_by_lock_hash", hash, strconv.Itoa(pageNumber), strconv.Itoa(pageSize))
	if err != nil {
		return nil, err
	}
	return rpcTransactionByLockHash, nil
}

func (ckbClient *CkbClient) IndexLockHash(hash Hash) (*RpcLockHashIndexState, error){
	rpcLockHashIndexState := &RpcLockHashIndexState{}
	err := ckbClient.client.CallFor(rpcLockHashIndexState, "index_lock_hash", hash)
	if err != nil {
		return nil, err
	}
	return rpcLockHashIndexState, nil
}

func (ckbClient *CkbClient) GetBannedAddresses() ([]*RpcBannedAddress, error){
	rpcBannedAddress := []*RpcBannedAddress{}
	err := ckbClient.client.CallFor(&rpcBannedAddress, "get_banned_addresses")
	if err != nil {
		return nil, err
	}
	return rpcBannedAddress, nil
}

func (ckbClient *CkbClient) GetHeaderByNumber(number int) (*RpcHeader, error){
	header := &RpcHeader{}
	err := ckbClient.client.CallFor(header, " get_header_by_number", strconv.Itoa(number))
	if err != nil {
		return nil, err
	}
	return header, nil
}

func (ckbClient *CkbClient) GetHeader(hash Hash) (*RpcHeader, error){
	header := &RpcHeader{}
	err := ckbClient.client.CallFor(header, "get_header", hash)
	if err != nil {
		return nil, err
	}
	return header, nil
}



func (ckbClient *CkbClient) GetCellbaseOutputCapacityDetails(hash Hash) (*RpcCellbaseOutputCapacityDetails, error){
	rpcCellbaseOutputCapacityDetails := &RpcCellbaseOutputCapacityDetails{}
	err := ckbClient.client.CallFor(rpcCellbaseOutputCapacityDetails, "get_cellbase_output_capacity_details", hash)
	if err != nil {
		return nil, err
	}
	return rpcCellbaseOutputCapacityDetails, nil
}




func (ckbClient *CkbClient) ComputeTransactionHashMethod(rawtx *RpcRawTransaction) (Hash, error){
    hash := Hash("")
	err := ckbClient.client.CallFor(&hash, "_compute_transaction_hash",rawtx)
	if err != nil {
		return "", err
	}
	return hash, nil
}







