package ckb_sdk_go

type RpcAlert struct {
	Id           string
	Cancel       string
	Signatures   []string
	Min_version  *string
	Max_version  *string
	Priority     string
	Notice_until string
	Message      string
}

type RpcAlertMessage struct {
	Id           string
	Priority     string
	Notice_until string
	Message      string
}

type RpcBlockTemplate struct {
	Version            string
	Difficulty         string
	Current_time       string
	Number             string
	Epoch              string
	Parent_hash        string
	Cycles_limit       string
	Bytes_limit        string
	Uncles_count_limit string
	Uncles             []RpcUncleTemplate
	Transactions       []RpcTransactionTemplate
	Proposals          []string
	Cellbase           RpcCellbaseTemplate
	Work_id            string
	Dao                string
}

type RpcUncleTemplate struct {
	Hash      string
	Required  bool
	Proposals []string
	Header    RpcHeader
}

type RpcCellbaseTemplate struct {
	Hash   string
	Cycles *string
	Data   RpcTransaction
}

type RpcTransactionTemplate struct {
	Hash     string
	Required bool
	Cycles   *string
	Depends  *[]string
	Data     RpcTransaction
}

type RpcScript struct {
	Args      []string
	Code_hash string
	Hash_type ScriptHashType
}

type RpcCellOutput struct {
	Capacity string
	Data     string
	Lock     RpcScript
	Type_    *RpcScript
}

type RpcCellOutPoint struct {
	Tx_hash string
	Index   string
}

type RpcOutPoint struct {
	Cell       *RpcCellOutPoint
	Block_hash *string
}

type RpcCellInput struct {
	Previous_output RpcOutPoint
	Since           string
}

type RpcWitness struct {
	data []string
}

type RpcTransaction struct {
	Version   string
	Deps      []RpcOutPoint
	Inputs    []RpcCellInput
	Outputs   []RpcCellOutput
	Witnesses []RpcWitness
}

type RpcTransactionView struct {
	Inner RpcTransaction
	Hash  string
}

type RpcTransactionWithStatus struct {
	Transaction RpcTransactionView
	Tx_status   RpcTxStatus
}

type RpcTxStatus struct {
	Status     Status
	Block_hash *string
}

type RpcSeal struct {
	Nonce string
	Proof string
}

type RpcHeader struct {
	Version           string
	Parent_hash       string
	Timestamp         string
	Number            string
	Epoch             string
	Transactions_root string
	Witnesses_root    string
	Proposals_hash    string
	Difficulty        string
	Uncles_hash       string
	Uncles_count      string
	Dao               string
	Seal              RpcSeal
}

type RpcHeaderView struct {
	Inner RpcHeader
	Hash  string
}

type RpcUncleBlock struct {
	Header    RpcHeader
	Proposals []string
}

type RpcUncleBlockView struct {
	Header    RpcHeaderView
	Proposals []string
}

type RpcBlock struct {
	Header       RpcHeader
	Uncles       []RpcUncleBlock
	Transactions []RpcTransaction
	Proposals    []string
}

type RpcBlockView struct {
	Header       RpcHeaderView
	Uncles       []RpcUncleBlockView
	Transactions []RpcTransactionView
	Proposals    []string
}

type RpcEpochView struct {
	Number       string
	Epoch_reward string
	Start_number string
	Length       string
	Difficulty   string
}

type RpcBlockRewardView struct {
	Total           string
	Primary         string
	Secondary       string
	Tx_fee          string
	Proposal_reward string
}

type RpcCellOutputWithOutPoint struct {
	Out_point RpcOutPoint
	Capacity  string
	Lock      RpcScript
}

type RpcCellWithStatus struct {
	Cell   *RpcCellOutput
	Status string
}

type RpcChainInfo struct {
	Chain                     string
	Median_time               string
	Epoch                     string
	Difficulty                string
	Is_initial_block_download bool
	Alerts                    []RpcAlertMessage
}

type RpcDryRunResult struct {
	Cycles string
}

type RpcLiveCell struct {
	Created_by  RpcTransactionPoint
	Cell_output RpcCellOutput
}

type RpcCellTransaction struct {
	Created_by  RpcTransactionPoint
	Consumed_by *RpcTransactionPoint
}

type RpcTransactionPoint struct {
	Block_number string
	Tx_hash      string
	Index        string
}

type RpcLockHashIndexState struct {
	Lock_hash    string
	Block_number string
	Block_hash   string
}

type RpcNode struct {
	Version     string
	Node_id     string
	Addresses   []RpcNodeAddress
	Is_outbound *bool
}

type RpcNodeAddress struct {
	Address string
	Score   string
}

type RpcBannedAddress struct {
	Address    string
	Ban_until  string
	Ban_reason string
	Created_at string
}

type RpcTxPoolInfo struct {
	Pending             string
	Proposed            string
	Orphan              string
	Total_tx_size       string
	Total_tx_cycles     string
	Last_txs_updated_at string
}

type RpcPeerState struct {
	peer             string
	last_updated     string
	blocks_in_flight string
}
