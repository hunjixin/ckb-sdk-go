package ckb_sdk_go

import (
	"ckb-sdk-go/core"
	"github.com/ybbus/jsonrpc"
)

type CkbClient struct {
	url    string
	client jsonrpc.RPCClient
}

func NewCkbClient(url string) *CkbClient {
	return &CkbClient{
		url:    url,
		client: jsonrpc.NewClient(url),
	}
}
func (ckbClient *CkbClient) SendAlert(_alert core.RpcAlert) error {
	res, err := ckbClient.client.Call("send_alert", _alert)

	if err != nil {
		return err
	}
	if res.Error != nil {
		return res.Error
	}
	return nil

}

func (ckbClient *CkbClient) GetBlock(_hash string) (*core.RpcBlockView, error) {
	rpcblockview := &core.RpcBlockView{}
	err := ckbClient.client.CallFor(rpcblockview, "get_block", _hash)

	if err != nil {
		return nil, err
	}
	return rpcblockview, nil

}

func (ckbClient *CkbClient) GetBlockByNumber(_number string) (*core.RpcBlockView, error) {
	rpcblockview := &core.RpcBlockView{}
	err := ckbClient.client.CallFor(rpcblockview, "get_block_by_number", _number)

	if err != nil {
		return nil, err
	}
	return rpcblockview, nil

}

func (ckbClient *CkbClient) GetHeader(_hash string) (*core.RpcHeaderView, error) {
	rpcheaderview := &core.RpcHeaderView{}
	err := ckbClient.client.CallFor(rpcheaderview, "get_header", _hash)

	if err != nil {
		return nil, err
	}
	return rpcheaderview, nil

}

func (ckbClient *CkbClient) GetHeaderByNumber(_number string) (*core.RpcHeaderView, error) {
	rpcheaderview := &core.RpcHeaderView{}
	err := ckbClient.client.CallFor(rpcheaderview, "get_header_by_number", _number)

	if err != nil {
		return nil, err
	}
	return rpcheaderview, nil

}

func (ckbClient *CkbClient) GetTransaction(_hash string) (*core.RpcTransactionWithStatus, error) {
	rpctransactionwithstatus := &core.RpcTransactionWithStatus{}
	err := ckbClient.client.CallFor(rpctransactionwithstatus, "get_transaction", _hash)

	if err != nil {
		return nil, err
	}
	return rpctransactionwithstatus, nil

}

func (ckbClient *CkbClient) GetBlockHash(_number string) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "get_block_hash", _number)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) GetTipHeader() (*core.RpcHeaderView, error) {
	rpcheaderview := &core.RpcHeaderView{}
	err := ckbClient.client.CallFor(rpcheaderview, "get_tip_header")

	if err != nil {
		return nil, err
	}
	return rpcheaderview, nil

}

func (ckbClient *CkbClient) GetCellsByLockHash(_lock_hash string, _from string, _to string) (*core.RpcCellOutputWithOutPoint, error) {
	rpccelloutputwithoutpoint := &core.RpcCellOutputWithOutPoint{}
	err := ckbClient.client.CallFor(rpccelloutputwithoutpoint, "get_cells_by_lock_hash", _lock_hash, _from, _to)

	if err != nil {
		return nil, err
	}
	return rpccelloutputwithoutpoint, nil

}

func (ckbClient *CkbClient) GetLiveCell(_out_point core.RpcOutPoint) (*core.RpcCellWithStatus, error) {
	rpccellwithstatus := &core.RpcCellWithStatus{}
	err := ckbClient.client.CallFor(rpccellwithstatus, "get_live_cell", _out_point)

	if err != nil {
		return nil, err
	}
	return rpccellwithstatus, nil

}

func (ckbClient *CkbClient) GetTipBlockNumber() (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "get_tip_block_number")

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) GetCurrentEpoch() (*core.RpcEpochView, error) {
	rpcepochview := &core.RpcEpochView{}
	err := ckbClient.client.CallFor(rpcepochview, "get_current_epoch")

	if err != nil {
		return nil, err
	}
	return rpcepochview, nil

}

func (ckbClient *CkbClient) GetEpochByNumber(number string) (*core.RpcEpochView, error) {
	rpcepochview := &core.RpcEpochView{}
	err := ckbClient.client.CallFor(rpcepochview, "get_epoch_by_number", number)

	if err != nil {
		return nil, err
	}
	return rpcepochview, nil

}

func (ckbClient *CkbClient) GetCellbaseOutputCapacityDetails(_hash string) (*core.RpcBlockReward, error) {
	rpcblockreward := &core.RpcBlockReward{}
	err := ckbClient.client.CallFor(rpcblockreward, "get_cellbase_output_capacity_details", _hash)

	if err != nil {
		return nil, err
	}
	return rpcblockreward, nil

}

func (ckbClient *CkbClient) ComputeTransactionHash(tx core.RpcTransaction) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "_compute_transaction_hash", tx)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) ComputeScriptHash(script core.RpcScript) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "_compute_script_hash", script)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) DryRunTransaction(_tx core.RpcTransaction) (*core.RpcDryRunResult, error) {
	rpcdryrunresult := &core.RpcDryRunResult{}
	err := ckbClient.client.CallFor(rpcdryrunresult, "dry_run_transaction", _tx)

	if err != nil {
		return nil, err
	}
	return rpcdryrunresult, nil

}

func (ckbClient *CkbClient) CalculateDaoMaximumWithdraw(_out_point core.RpcOutPoint, _hash string) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "calculate_dao_maximum_withdraw", _out_point, _hash)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) GetLiveCellsByLockHash(_lock_hash string, _page string, _per_page string, _reverse_order *bool) (*core.RpcLiveCell, error) {
	rpclivecell := &core.RpcLiveCell{}
	err := ckbClient.client.CallFor(rpclivecell, "get_live_cells_by_lock_hash", _lock_hash, _page, _per_page, _reverse_order)

	if err != nil {
		return nil, err
	}
	return rpclivecell, nil

}

func (ckbClient *CkbClient) GetTransactionsByLockHash(_lock_hash string, _page string, _per_page string, _reverse_order *bool) (*core.RpcCellTransaction, error) {
	rpccelltransaction := &core.RpcCellTransaction{}
	err := ckbClient.client.CallFor(rpccelltransaction, "get_transactions_by_lock_hash", _lock_hash, _page, _per_page, _reverse_order)

	if err != nil {
		return nil, err
	}
	return rpccelltransaction, nil

}

func (ckbClient *CkbClient) IndexLockHash(_lock_hash string, _index_from *string) (*core.RpcLockHashIndexState, error) {
	rpclockhashindexstate := &core.RpcLockHashIndexState{}
	err := ckbClient.client.CallFor(rpclockhashindexstate, "index_lock_hash", _lock_hash, _index_from)

	if err != nil {
		return nil, err
	}
	return rpclockhashindexstate, nil

}

func (ckbClient *CkbClient) DeindexLockHash(_lock_hash string) error {
	res, err := ckbClient.client.Call("deindex_lock_hash", _lock_hash)

	if err != nil {
		return err
	}
	if res.Error != nil {
		return res.Error
	}
	return nil

}

func (ckbClient *CkbClient) GetLockHashIndexStates() (*core.RpcLockHashIndexState, error) {
	rpclockhashindexstate := &core.RpcLockHashIndexState{}
	err := ckbClient.client.CallFor(rpclockhashindexstate, "get_lock_hash_index_states")

	if err != nil {
		return nil, err
	}
	return rpclockhashindexstate, nil

}

func (ckbClient *CkbClient) GetBlockTemplate(bytes_limit *string, proposals_limit *string, max_version *string) (*core.RpcBlockTemplate, error) {
	rpcblocktemplate := &core.RpcBlockTemplate{}
	err := ckbClient.client.CallFor(rpcblocktemplate, "get_block_template", bytes_limit, proposals_limit, max_version)

	if err != nil {
		return nil, err
	}
	return rpcblocktemplate, nil

}

func (ckbClient *CkbClient) SubmitBlock(_work_id string, _data core.RpcBlock) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "submit_block", _work_id, _data)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) LocalNodeInfo() (*core.RpcNode, error) {
	rpcnode := &core.RpcNode{}
	err := ckbClient.client.CallFor(rpcnode, "local_node_info")

	if err != nil {
		return nil, err
	}
	return rpcnode, nil

}

func (ckbClient *CkbClient) GetPeers() (*core.RpcNode, error) {
	rpcnode := &core.RpcNode{}
	err := ckbClient.client.CallFor(rpcnode, "get_peers")

	if err != nil {
		return nil, err
	}
	return rpcnode, nil

}

func (ckbClient *CkbClient) GetBannedAddresses() (*core.RpcBannedAddress, error) {
	rpcbannedaddress := &core.RpcBannedAddress{}
	err := ckbClient.client.CallFor(rpcbannedaddress, "get_banned_addresses")

	if err != nil {
		return nil, err
	}
	return rpcbannedaddress, nil

}

func (ckbClient *CkbClient) SetBan(address string, command string, ban_time *string, absolute *bool, reason *string) error {
	res, err := ckbClient.client.Call("set_ban", address, command, ban_time, absolute, reason)

	if err != nil {
		return err
	}
	if res.Error != nil {
		return res.Error
	}
	return nil

}

func (ckbClient *CkbClient) SendTransaction(_tx core.RpcTransaction) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "send_transaction", _tx)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) TxPoolInfo() (*core.RpcTxPoolInfo, error) {
	rpctxpoolinfo := &core.RpcTxPoolInfo{}
	err := ckbClient.client.CallFor(rpctxpoolinfo, "tx_pool_info")

	if err != nil {
		return nil, err
	}
	return rpctxpoolinfo, nil

}

func (ckbClient *CkbClient) GetBlockchainInfo() (*core.RpcChainInfo, error) {
	rpcchaininfo := &core.RpcChainInfo{}
	err := ckbClient.client.CallFor(rpcchaininfo, "get_blockchain_info")

	if err != nil {
		return nil, err
	}
	return rpcchaininfo, nil

}

func (ckbClient *CkbClient) GetPeersState() (*core.RpcPeerState, error) {
	rpcpeerstate := &core.RpcPeerState{}
	err := ckbClient.client.CallFor(rpcpeerstate, "get_peers_state")

	if err != nil {
		return nil, err
	}
	return rpcpeerstate, nil

}

func (ckbClient *CkbClient) AddNode(peer_id string, address string) error {
	res, err := ckbClient.client.Call("add_node", peer_id, address)

	if err != nil {
		return err
	}
	if res.Error != nil {
		return res.Error
	}
	return nil

}

func (ckbClient *CkbClient) RemoveNode(peer_id string) error {
	res, err := ckbClient.client.Call("remove_node", peer_id)

	if err != nil {
		return err
	}
	if res.Error != nil {
		return res.Error
	}
	return nil

}

func (ckbClient *CkbClient) ProcessBlockWithoutVerify(data core.RpcBlock) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "process_block_without_verify", data)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) BroadcastTransaction(transaction core.RpcTransaction) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "broadcast_transaction", transaction)

	if err != nil {
		return nil, err
	}
	return &string, nil

}
