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
		
type RpcAlert struct {
Id  AlertId
Cancel  AlertId
Signatures  []JsonBytes
Min_version  *string
Max_version  *string
Priority  AlertPriority
Notice_until  Timestamp
Message  string

}
type RpcAlertMessage struct {
Id  AlertId
Priority  AlertPriority
Notice_until  Timestamp
Message  string

}
type RpcBlockTemplate struct {
Version  Version
Difficulty  U256
Current_time  Timestamp
Number  BlockNumber
Epoch  EpochNumber
Parent_hash  H256
Cycles_limit  Cycle
Bytes_limit  Unsigned
Uncles_count_limit  Unsigned
Uncles  []RpcUncleTemplate
Transactions  []RpcTransactionTemplate
Proposals  []ProposalShortId
Cellbase  RpcCellbaseTemplate
Work_id  Unsigned
Dao  JsonBytes

}
type RpcUncleTemplate struct {
Hash  H256
Required  bool
Proposals  []ProposalShortId
Header  RpcHeader

}
type RpcCellbaseTemplate struct {
Hash  H256
Cycles  *Cycle
Data  RpcTransaction

}
type RpcTransactionTemplate struct {
Hash  H256
Required  bool
Cycles  *Cycle
Depends  *[]Unsigned
Data  RpcTransaction

}
type RpcScript struct {
Args  []JsonBytes
Code_hash  H256
Hash_type  ScriptHashType

}
type RpcCellOutput struct {
Capacity  Capacity
Data  JsonBytes
Lock  RpcScript
Type_  *RpcScript

}
type RpcCellOutPoint struct {
Tx_hash  H256
Index  Unsigned

}
type RpcOutPoint struct {
Cell  *RpcCellOutPoint
Block_hash  *H256

}
type RpcCellInput struct {
Previous_output  RpcOutPoint
Since  Unsigned

}
type RpcWitness struct {
data  []JsonBytes

}
type RpcTransaction struct {
Version  Version
Deps  []RpcOutPoint
Inputs  []RpcCellInput
Outputs  []RpcCellOutput
Witnesses  []RpcWitness

}
type RpcTransactionView struct {
Inner  RpcTransaction
Hash  H256

}
type RpcTransactionWithStatus struct {
Transaction  RpcTransactionView
Tx_status  RpcTxStatus

}
type RpcTxStatus struct {
Status  Status
Block_hash  *H256

}
type RpcSeal struct {
Nonce  Unsigned
Proof  JsonBytes

}
type RpcHeader struct {
Version  Version
Parent_hash  H256
Timestamp  Timestamp
Number  BlockNumber
Epoch  EpochNumber
Transactions_root  H256
Witnesses_root  H256
Proposals_hash  H256
Difficulty  U256
Uncles_hash  H256
Uncles_count  Unsigned
Dao  JsonBytes
Seal  RpcSeal

}
type RpcHeaderView struct {
Inner  RpcHeader
Hash  H256

}
type RpcUncleBlock struct {
Header  RpcHeader
Proposals  []ProposalShortId

}
type RpcUncleBlockView struct {
Header  RpcHeaderView
Proposals  []ProposalShortId

}
type RpcBlock struct {
Header  RpcHeader
Uncles  []RpcUncleBlock
Transactions  []RpcTransaction
Proposals  []ProposalShortId

}
type RpcBlockView struct {
Header  RpcHeaderView
Uncles  []RpcUncleBlockView
Transactions  []RpcTransactionView
Proposals  []ProposalShortId

}
type RpcEpochView struct {
Number  EpochNumber
Epoch_reward  Capacity
Start_number  BlockNumber
Length  BlockNumber
Difficulty  U256

}
type RpcBlockRewardView struct {
Total  Capacity
Primary  Capacity
Secondary  Capacity
Tx_fee  Capacity
Proposal_reward  Capacity

}
type RpcCellOutputWithOutPoint struct {
Out_point  RpcOutPoint
Capacity  Capacity
Lock  RpcScript

}
type RpcCellWithStatus struct {
Cell  *RpcCellOutput
Status  string

}
type RpcChainInfo struct {
Chain  string
Median_time  Timestamp
Epoch  EpochNumber
Difficulty  U256
Is_initial_block_download  bool
Alerts  []RpcAlertMessage

}
type RpcDryRunResult struct {
Cycles  Cycle

}
type RpcLiveCell struct {
Created_by  RpcTransactionPoint
Cell_output  RpcCellOutput

}
type RpcCellTransaction struct {
Created_by  RpcTransactionPoint
Consumed_by  *RpcTransactionPoint

}
type RpcTransactionPoint struct {
Block_number  BlockNumber
Tx_hash  H256
Index  Unsigned

}
type RpcLockHashIndexState struct {
Lock_hash  H256
Block_number  BlockNumber
Block_hash  H256

}
type RpcNode struct {
Version  string
Node_id  string
Addresses  []RpcNodeAddress
Is_outbound  *bool

}
type RpcNodeAddress struct {
Address  string
Score  Unsigned

}
type RpcBannedAddress struct {
Address  string
Ban_until  Timestamp
Ban_reason  string
Created_at  Timestamp

}
type RpcTxPoolInfo struct {
Pending  Unsigned
Proposed  Unsigned
Orphan  Unsigned
Total_tx_size  Unsigned
Total_tx_cycles  Unsigned
Last_txs_updated_at  Timestamp

}
type RpcPeerState struct {
peer  Unsigned
last_updated  Timestamp
blocks_in_flight  Unsigned

}
//AlertRpc
func (ckbClient *CkbClient) SendAlert(_alert RpcAlert,  ) error {

}
//ChainRpc
func (ckbClient *CkbClient) GetBlock(_hash H256,  ) (*RpcBlockView , error){
rpcblockview:= &RpcBlockView{}
err := ckbClient.client.CallFor(rpcblockview,"get_block",_hash)

		if err != nil {
			return nil, err
		}
return rpcblockview, nil

}
func (ckbClient *CkbClient) GetBlockByNumber(_number BlockNumber,  ) (*RpcBlockView , error){
rpcblockview:= &RpcBlockView{}
err := ckbClient.client.CallFor(rpcblockview,"get_block_by_number",_number)

		if err != nil {
			return nil, err
		}
return rpcblockview, nil

}
func (ckbClient *CkbClient) GetHeader(_hash H256,  ) (*RpcHeaderView , error){
rpcheaderview:= &RpcHeaderView{}
err := ckbClient.client.CallFor(rpcheaderview,"get_header",_hash)

		if err != nil {
			return nil, err
		}
return rpcheaderview, nil

}
func (ckbClient *CkbClient) GetHeaderByNumber(_number BlockNumber,  ) (*RpcHeaderView , error){
rpcheaderview:= &RpcHeaderView{}
err := ckbClient.client.CallFor(rpcheaderview,"get_header_by_number",_number)

		if err != nil {
			return nil, err
		}
return rpcheaderview, nil

}
func (ckbClient *CkbClient) GetTransaction(_hash H256,  ) (*RpcTransactionWithStatus , error){
rpctransactionwithstatus:= &RpcTransactionWithStatus{}
err := ckbClient.client.CallFor(rpctransactionwithstatus,"get_transaction",_hash)

		if err != nil {
			return nil, err
		}
return rpctransactionwithstatus, nil

}
func (ckbClient *CkbClient) GetBlockHash(_number BlockNumber,  ) (*H256 , error){
h256:= &H256{}
err := ckbClient.client.CallFor(h256,"get_block_hash",_number)

		if err != nil {
			return nil, err
		}
return h256, nil

}
func (ckbClient *CkbClient) GetTipHeader() (RpcHeaderView , error){
rpcheaderview:= &RpcHeaderView{}
err := ckbClient.client.CallFor(rpcheaderview,"get_tip_header")

		if err != nil {
			return nil, err
		}
return rpcheaderview, nil

}
func (ckbClient *CkbClient) GetCellsByLockHash(_lock_hash H256,  _from BlockNumber,  _to BlockNumber,  ) ([]RpcCellOutputWithOutPoint , error){
rpccelloutputwithoutpoint:= &RpcCellOutputWithOutPoint{}
err := ckbClient.client.CallFor(rpccelloutputwithoutpoint,"get_cells_by_lock_hash",_lock_hash,_from,_to)

		if err != nil {
			return nil, err
		}
return rpccelloutputwithoutpoint, nil

}
func (ckbClient *CkbClient) GetLiveCell(_out_point RpcOutPoint,  ) (RpcCellWithStatus , error){
rpccellwithstatus:= &RpcCellWithStatus{}
err := ckbClient.client.CallFor(rpccellwithstatus,"get_live_cell",_out_point)

		if err != nil {
			return nil, err
		}
return rpccellwithstatus, nil

}
func (ckbClient *CkbClient) GetTipBlockNumber() (BlockNumber , error){
blocknumber:= &BlockNumber{}
err := ckbClient.client.CallFor(blocknumber,"get_tip_block_number")

		if err != nil {
			return nil, err
		}
return blocknumber, nil

}
func (ckbClient *CkbClient) GetCurrentEpoch() (RpcEpochView , error){
rpcepochview:= &RpcEpochView{}
err := ckbClient.client.CallFor(rpcepochview,"get_current_epoch")

		if err != nil {
			return nil, err
		}
return rpcepochview, nil

}
func (ckbClient *CkbClient) GetEpochByNumber(number EpochNumber,  ) (*RpcEpochView , error){
rpcepochview:= &RpcEpochView{}
err := ckbClient.client.CallFor(rpcepochview,"get_epoch_by_number",number)

		if err != nil {
			return nil, err
		}
return rpcepochview, nil

}
func (ckbClient *CkbClient) GetCellbaseOutputCapacityDetails(_hash H256,  ) (*RpcBlockRewardView , error){
rpcblockrewardview:= &RpcBlockRewardView{}
err := ckbClient.client.CallFor(rpcblockrewardview,"get_cellbase_output_capacity_details",_hash)

		if err != nil {
			return nil, err
		}
return rpcblockrewardview, nil

}
//ExperimentRpc
func (ckbClient *CkbClient) ComputeTransactionHash(tx RpcTransaction,  ) (H256 , error){
h256:= &H256{}
err := ckbClient.client.CallFor(h256,"_compute_transaction_hash",tx)

		if err != nil {
			return nil, err
		}
return h256, nil

}
func (ckbClient *CkbClient) ComputeScriptHash(script RpcScript,  ) (H256 , error){
h256:= &H256{}
err := ckbClient.client.CallFor(h256,"_compute_script_hash",script)

		if err != nil {
			return nil, err
		}
return h256, nil

}
func (ckbClient *CkbClient) DryRunTransaction(_tx RpcTransaction,  ) (RpcDryRunResult , error){
rpcdryrunresult:= &RpcDryRunResult{}
err := ckbClient.client.CallFor(rpcdryrunresult,"dry_run_transaction",_tx)

		if err != nil {
			return nil, err
		}
return rpcdryrunresult, nil

}
func (ckbClient *CkbClient) CalculateDaoMaximumWithdraw(_out_point RpcOutPoint,  _hash H256,  ) (Capacity , error){
capacity:= &Capacity{}
err := ckbClient.client.CallFor(capacity,"calculate_dao_maximum_withdraw",_out_point,_hash)

		if err != nil {
			return nil, err
		}
return capacity, nil

}
//IndexerRpc
func (ckbClient *CkbClient) GetLiveCellsByLockHash(_lock_hash H256,  _page Unsigned,  _per_page Unsigned,  _reverse_order *bool,  ) ([]RpcLiveCell , error){
rpclivecell:= &RpcLiveCell{}
err := ckbClient.client.CallFor(rpclivecell,"get_live_cells_by_lock_hash",_lock_hash,_page,_per_page,_reverse_order)

		if err != nil {
			return nil, err
		}
return rpclivecell, nil

}
func (ckbClient *CkbClient) GetTransactionsByLockHash(_lock_hash H256,  _page Unsigned,  _per_page Unsigned,  _reverse_order *bool,  ) ([]RpcCellTransaction , error){
rpccelltransaction:= &RpcCellTransaction{}
err := ckbClient.client.CallFor(rpccelltransaction,"get_transactions_by_lock_hash",_lock_hash,_page,_per_page,_reverse_order)

		if err != nil {
			return nil, err
		}
return rpccelltransaction, nil

}
func (ckbClient *CkbClient) IndexLockHash(_lock_hash H256,  _index_from *BlockNumber,  ) (RpcLockHashIndexState , error){
rpclockhashindexstate:= &RpcLockHashIndexState{}
err := ckbClient.client.CallFor(rpclockhashindexstate,"index_lock_hash",_lock_hash,_index_from)

		if err != nil {
			return nil, err
		}
return rpclockhashindexstate, nil

}
func (ckbClient *CkbClient) DeindexLockHash(_lock_hash H256,  ) error {

}
func (ckbClient *CkbClient) GetLockHashIndexStates() ([]RpcLockHashIndexState , error){
rpclockhashindexstate:= &RpcLockHashIndexState{}
err := ckbClient.client.CallFor(rpclockhashindexstate,"get_lock_hash_index_states")

		if err != nil {
			return nil, err
		}
return rpclockhashindexstate, nil

}
//MinerRpc
func (ckbClient *CkbClient) GetBlockTemplate(bytes_limit *Unsigned,  proposals_limit *Unsigned,  max_version *Version,  ) (RpcBlockTemplate , error){
rpcblocktemplate:= &RpcBlockTemplate{}
err := ckbClient.client.CallFor(rpcblocktemplate,"get_block_template",bytes_limit,proposals_limit,max_version)

		if err != nil {
			return nil, err
		}
return rpcblocktemplate, nil

}
func (ckbClient *CkbClient) SubmitBlock(_work_id string,  _data RpcBlock,  ) (*H256 , error){
h256:= &H256{}
err := ckbClient.client.CallFor(h256,"submit_block",_work_id,_data)

		if err != nil {
			return nil, err
		}
return h256, nil

}
//NetworkRpc
func (ckbClient *CkbClient) LocalNodeInfo() (RpcNode , error){
rpcnode:= &RpcNode{}
err := ckbClient.client.CallFor(rpcnode,"local_node_info")

		if err != nil {
			return nil, err
		}
return rpcnode, nil

}
func (ckbClient *CkbClient) GetPeers() ([]RpcNode , error){
rpcnode:= &RpcNode{}
err := ckbClient.client.CallFor(rpcnode,"get_peers")

		if err != nil {
			return nil, err
		}
return rpcnode, nil

}
func (ckbClient *CkbClient) GetBannedAddresses() ([]RpcBannedAddress , error){
rpcbannedaddress:= &RpcBannedAddress{}
err := ckbClient.client.CallFor(rpcbannedaddress,"get_banned_addresses")

		if err != nil {
			return nil, err
		}
return rpcbannedaddress, nil

}
func (ckbClient *CkbClient) SetBan(address string,  command string,  ban_time *Timestamp,  absolute *bool,  reason *string,  ) error {

}
//PoolRpc
func (ckbClient *CkbClient) SendTransaction(_tx RpcTransaction,  ) (H256 , error){
h256:= &H256{}
err := ckbClient.client.CallFor(h256,"send_transaction",_tx)

		if err != nil {
			return nil, err
		}
return h256, nil

}
func (ckbClient *CkbClient) TxPoolInfo() (RpcTxPoolInfo , error){
rpctxpoolinfo:= &RpcTxPoolInfo{}
err := ckbClient.client.CallFor(rpctxpoolinfo,"tx_pool_info")

		if err != nil {
			return nil, err
		}
return rpctxpoolinfo, nil

}
//StatsRpc
func (ckbClient *CkbClient) GetBlockchainInfo() (RpcChainInfo , error){
rpcchaininfo:= &RpcChainInfo{}
err := ckbClient.client.CallFor(rpcchaininfo,"get_blockchain_info")

		if err != nil {
			return nil, err
		}
return rpcchaininfo, nil

}
func (ckbClient *CkbClient) GetPeersState() ([]RpcPeerState , error){
rpcpeerstate:= &RpcPeerState{}
err := ckbClient.client.CallFor(rpcpeerstate,"get_peers_state")

		if err != nil {
			return nil, err
		}
return rpcpeerstate, nil

}
