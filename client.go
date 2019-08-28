package ckb_sdk_go

import (
	"github.com/ybbus/jsonrpc"
)

type CkbClient struct {
	url    string
	client jsonrpc.RPCClient
}

func (ckbClient *CkbClient) SendAlert(_alert RpcAlert) error {
	res, err := ckbClient.client.Call("send_alert", _alert)

	if err != nil {
		return err
	}
	if res.Error != nil {
		return res.Error
	}
	return nil

}

func (ckbClient *CkbClient) GetBlock(_hash string) (*RpcBlockView, error) {
	rpcblockview := &RpcBlockView{}
	err := ckbClient.client.CallFor(rpcblockview, "get_block", _hash)

	if err != nil {
		return nil, err
	}
	return rpcblockview, nil

}

func (ckbClient *CkbClient) GetBlockByNumber(_number string) (*RpcBlockView, error) {
	rpcblockview := &RpcBlockView{}
	err := ckbClient.client.CallFor(rpcblockview, "get_block_by_number", _number)

	if err != nil {
		return nil, err
	}
	return rpcblockview, nil

}

func (ckbClient *CkbClient) GetHeader(_hash string) (*RpcHeaderView, error) {
	rpcheaderview := &RpcHeaderView{}
	err := ckbClient.client.CallFor(rpcheaderview, "get_header", _hash)

	if err != nil {
		return nil, err
	}
	return rpcheaderview, nil

}

func (ckbClient *CkbClient) GetHeaderByNumber(_number string) (*RpcHeaderView, error) {
	rpcheaderview := &RpcHeaderView{}
	err := ckbClient.client.CallFor(rpcheaderview, "get_header_by_number", _number)

	if err != nil {
		return nil, err
	}
	return rpcheaderview, nil

}

func (ckbClient *CkbClient) GetTransaction(_hash string) (*RpcTransactionWithStatus, error) {
	rpctransactionwithstatus := &RpcTransactionWithStatus{}
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

func (ckbClient *CkbClient) GetTipHeader() (*RpcHeaderView, error) {
	rpcheaderview := &RpcHeaderView{}
	err := ckbClient.client.CallFor(rpcheaderview, "get_tip_header")

	if err != nil {
		return nil, err
	}
	return rpcheaderview, nil

}

func (ckbClient *CkbClient) GetCellsByLockHash(_lock_hash string, _from string, _to string) (*RpcCellOutputWithOutPoint, error) {
	rpccelloutputwithoutpoint := &RpcCellOutputWithOutPoint{}
	err := ckbClient.client.CallFor(rpccelloutputwithoutpoint, "get_cells_by_lock_hash", _lock_hash, _from, _to)

	if err != nil {
		return nil, err
	}
	return rpccelloutputwithoutpoint, nil

}

func (ckbClient *CkbClient) GetLiveCell(_out_point RpcOutPoint) (*RpcCellWithStatus, error) {
	rpccellwithstatus := &RpcCellWithStatus{}
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

func (ckbClient *CkbClient) GetCurrentEpoch() (*RpcEpochView, error) {
	rpcepochview := &RpcEpochView{}
	err := ckbClient.client.CallFor(rpcepochview, "get_current_epoch")

	if err != nil {
		return nil, err
	}
	return rpcepochview, nil

}

func (ckbClient *CkbClient) GetEpochByNumber(number string) (*RpcEpochView, error) {
	rpcepochview := &RpcEpochView{}
	err := ckbClient.client.CallFor(rpcepochview, "get_epoch_by_number", number)

	if err != nil {
		return nil, err
	}
	return rpcepochview, nil

}

func (ckbClient *CkbClient) GetCellbaseOutputCapacityDetails(_hash string) (*RpcBlockRewardView, error) {
	rpcblockrewardview := &RpcBlockRewardView{}
	err := ckbClient.client.CallFor(rpcblockrewardview, "get_cellbase_output_capacity_details", _hash)

	if err != nil {
		return nil, err
	}
	return rpcblockrewardview, nil

}

func (ckbClient *CkbClient) ComputeTransactionHash(tx RpcTransaction) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "_compute_transaction_hash", tx)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) ComputeScriptHash(script RpcScript) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "_compute_script_hash", script)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) DryRunTransaction(_tx RpcTransaction) (*RpcDryRunResult, error) {
	rpcdryrunresult := &RpcDryRunResult{}
	err := ckbClient.client.CallFor(rpcdryrunresult, "dry_run_transaction", _tx)

	if err != nil {
		return nil, err
	}
	return rpcdryrunresult, nil

}

func (ckbClient *CkbClient) CalculateDaoMaximumWithdraw(_out_point RpcOutPoint, _hash string) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "calculate_dao_maximum_withdraw", _out_point, _hash)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) GetLiveCellsByLockHash(_lock_hash string, _page string, _per_page string, _reverse_order *bool) (*RpcLiveCell, error) {
	rpclivecell := &RpcLiveCell{}
	err := ckbClient.client.CallFor(rpclivecell, "get_live_cells_by_lock_hash", _lock_hash, _page, _per_page, _reverse_order)

	if err != nil {
		return nil, err
	}
	return rpclivecell, nil

}

func (ckbClient *CkbClient) GetTransactionsByLockHash(_lock_hash string, _page string, _per_page string, _reverse_order *bool) (*RpcCellTransaction, error) {
	rpccelltransaction := &RpcCellTransaction{}
	err := ckbClient.client.CallFor(rpccelltransaction, "get_transactions_by_lock_hash", _lock_hash, _page, _per_page, _reverse_order)

	if err != nil {
		return nil, err
	}
	return rpccelltransaction, nil

}

func (ckbClient *CkbClient) IndexLockHash(_lock_hash string, _index_from *string) (*RpcLockHashIndexState, error) {
	rpclockhashindexstate := &RpcLockHashIndexState{}
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

func (ckbClient *CkbClient) GetLockHashIndexStates() (*RpcLockHashIndexState, error) {
	rpclockhashindexstate := &RpcLockHashIndexState{}
	err := ckbClient.client.CallFor(rpclockhashindexstate, "get_lock_hash_index_states")

	if err != nil {
		return nil, err
	}
	return rpclockhashindexstate, nil

}

func (ckbClient *CkbClient) GetBlockTemplate(bytes_limit *string, proposals_limit *string, max_version *string) (*RpcBlockTemplate, error) {
	rpcblocktemplate := &RpcBlockTemplate{}
	err := ckbClient.client.CallFor(rpcblocktemplate, "get_block_template", bytes_limit, proposals_limit, max_version)

	if err != nil {
		return nil, err
	}
	return rpcblocktemplate, nil

}

func (ckbClient *CkbClient) SubmitBlock(_work_id string, _data RpcBlock) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "submit_block", _work_id, _data)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) LocalNodeInfo() (*RpcNode, error) {
	rpcnode := &RpcNode{}
	err := ckbClient.client.CallFor(rpcnode, "local_node_info")

	if err != nil {
		return nil, err
	}
	return rpcnode, nil

}

func (ckbClient *CkbClient) GetPeers() (*RpcNode, error) {
	rpcnode := &RpcNode{}
	err := ckbClient.client.CallFor(rpcnode, "get_peers")

	if err != nil {
		return nil, err
	}
	return rpcnode, nil

}

func (ckbClient *CkbClient) GetBannedAddresses() (*RpcBannedAddress, error) {
	rpcbannedaddress := &RpcBannedAddress{}
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

func (ckbClient *CkbClient) SendTransaction(_tx RpcTransaction) (*string, error) {
	string := string("")
	err := ckbClient.client.CallFor(&string, "send_transaction", _tx)

	if err != nil {
		return nil, err
	}
	return &string, nil

}

func (ckbClient *CkbClient) TxPoolInfo() (*RpcTxPoolInfo, error) {
	rpctxpoolinfo := &RpcTxPoolInfo{}
	err := ckbClient.client.CallFor(rpctxpoolinfo, "tx_pool_info")

	if err != nil {
		return nil, err
	}
	return rpctxpoolinfo, nil

}

func (ckbClient *CkbClient) GetBlockchainInfo() (*RpcChainInfo, error) {
	rpcchaininfo := &RpcChainInfo{}
	err := ckbClient.client.CallFor(rpcchaininfo, "get_blockchain_info")

	if err != nil {
		return nil, err
	}
	return rpcchaininfo, nil

}

func (ckbClient *CkbClient) GetPeersState() (*RpcPeerState, error) {
	rpcpeerstate := &RpcPeerState{}
	err := ckbClient.client.CallFor(rpcpeerstate, "get_peers_state")

	if err != nil {
		return nil, err
	}
	return rpcpeerstate, nil

}
