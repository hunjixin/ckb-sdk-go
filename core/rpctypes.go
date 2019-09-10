package core

import "reflect"

var TRpcAlert = reflect.TypeOf(RpcAlert{})
var TRpcAlertMessage = reflect.TypeOf(RpcAlertMessage{})
var TRpcBlockTemplate = reflect.TypeOf(RpcBlockTemplate{})
var TRpcUncleTemplate = reflect.TypeOf(RpcUncleTemplate{})
var TRpcCellbaseTemplate = reflect.TypeOf(RpcCellbaseTemplate{})
var TRpcTransactionTemplate = reflect.TypeOf(RpcTransactionTemplate{})
var TRpcScript = reflect.TypeOf(RpcScript{})
var TRpcCellOutput = reflect.TypeOf(RpcCellOutput{})
var TRpcOutPoint = reflect.TypeOf(RpcOutPoint{})
var TRpcCellInput = reflect.TypeOf(RpcCellInput{})
var TRpcWitness = reflect.TypeOf(RpcWitness{})
var TRpcCellDep = reflect.TypeOf(RpcCellDep{})
var TRpcTransaction = reflect.TypeOf(RpcTransaction{})
var TRpcTransactionView = reflect.TypeOf(RpcTransactionView{})
var TRpcTransactionWithStatus = reflect.TypeOf(RpcTransactionWithStatus{})
var TRpcTxStatus = reflect.TypeOf(RpcTxStatus{})
var TRpcHeader = reflect.TypeOf(RpcHeader{})
var TRpcHeaderView = reflect.TypeOf(RpcHeaderView{})
var TRpcUncleBlock = reflect.TypeOf(RpcUncleBlock{})
var TRpcUncleBlockView = reflect.TypeOf(RpcUncleBlockView{})
var TRpcBlock = reflect.TypeOf(RpcBlock{})
var TRpcBlockView = reflect.TypeOf(RpcBlockView{})
var TRpcEpochView = reflect.TypeOf(RpcEpochView{})
var TRpcBlockReward = reflect.TypeOf(RpcBlockReward{})
var TRpcCellOutputWithOutPoint = reflect.TypeOf(RpcCellOutputWithOutPoint{})
var TRpcCellWithStatus = reflect.TypeOf(RpcCellWithStatus{})
var TRpcChainInfo = reflect.TypeOf(RpcChainInfo{})
var TRpcDryRunResult = reflect.TypeOf(RpcDryRunResult{})
var TRpcLiveCell = reflect.TypeOf(RpcLiveCell{})
var TRpcCellTransaction = reflect.TypeOf(RpcCellTransaction{})
var TRpcTransactionPoint = reflect.TypeOf(RpcTransactionPoint{})
var TRpcLockHashIndexState = reflect.TypeOf(RpcLockHashIndexState{})
var TRpcNode = reflect.TypeOf(RpcNode{})
var TRpcNodeAddress = reflect.TypeOf(RpcNodeAddress{})
var TRpcBannedAddress = reflect.TypeOf(RpcBannedAddress{})
var TRpcTxPoolInfo = reflect.TypeOf(RpcTxPoolInfo{})
var TRpcPeerState = reflect.TypeOf(RpcPeerState{})

type RpcAlert struct {
	Id           string   `json:"id"`
	Cancel       string   `json:"cancel"`
	Min_version  *string  `json:"min_version"`
	Max_version  *string  `json:"max_version"`
	Priority     string   `json:"priority"`
	Notice_until string   `json:"notice_until"`
	Message      string   `json:"message"`
	Signatures   []string `json:"signatures"`
}

type RpcAlertMessage struct {
	Id           string `json:"id"`
	Priority     string `json:"priority"`
	Notice_until string `json:"notice_until"`
	Message      string `json:"message"`
}

type RpcBlockTemplate struct {
	Version            string                   `json:"version"`
	Difficulty         string                   `json:"difficulty"`
	Current_time       string                   `json:"current_time"`
	Number             string                   `json:"number"`
	Epoch              string                   `json:"epoch"`
	Parent_hash        string                   `json:"parent_hash"`
	Cycles_limit       string                   `json:"cycles_limit"`
	Bytes_limit        string                   `json:"bytes_limit"`
	Uncles_count_limit string                   `json:"uncles_count_limit"`
	Uncles             []RpcUncleTemplate       `json:"uncles"`
	Transactions       []RpcTransactionTemplate `json:"transactions"`
	Proposals          []string                 `json:"proposals"`
	Cellbase           RpcCellbaseTemplate      `json:"cellbase"`
	Work_id            string                   `json:"work_id"`
	Dao                string                   `json:"dao"`
}

type RpcUncleTemplate struct {
	Hash      string    `json:"hash"`
	Required  bool      `json:"required"`
	Proposals []string  `json:"proposals"`
	Header    RpcHeader `json:"header"`
}

type RpcCellbaseTemplate struct {
	Hash   string         `json:"hash"`
	Cycles *string        `json:"cycles"`
	Data   RpcTransaction `json:"data"`
}

type RpcTransactionTemplate struct {
	Hash     string         `json:"hash"`
	Required bool           `json:"required"`
	Cycles   *string        `json:"cycles"`
	Depends  *[]string      `json:"depends"`
	Data     RpcTransaction `json:"data"`
}

type RpcScript struct {
	Args      []string          `json:"args"`
	Code_hash string            `json:"code_hash"`
	Hash_type RpcScriptHashType `json:"hash_type"`
}

type RpcCellOutput struct {
	Capacity string     `json:"capacity"`
	Lock     RpcScript  `json:"lock"`
	Type_    *RpcScript `json:"type_"`
}

type RpcOutPoint struct {
	Tx_hash string `json:"tx_hash"`
	Index   string `json:"index"`
}

type RpcCellInput struct {
	Previous_output RpcOutPoint `json:"previous_output"`
	Since           string      `json:"since"`
}

type RpcWitness struct {
	Data []string
}

type RpcCellDep struct {
	out_point RpcOutPoint
	dep_type  RpcDepType
}

type RpcTransaction struct {
	Version      string          `json:"version"`
	Cell_deps    []RpcCellDep    `json:"cell_deps"`
	Header_deps  []string        `json:"header_deps"`
	Inputs       []RpcCellInput  `json:"inputs"`
	Outputs      []RpcCellOutput `json:"outputs"`
	Witnesses    []RpcWitness    `json:"witnesses"`
	Outputs_data []string        `json:"outputs_data"`
}

type RpcTransactionView struct {
	Inner RpcTransaction `json:"inner"`
	Hash  string         `json:"hash"`
}

type RpcTransactionWithStatus struct {
	Transaction RpcTransactionView `json:"transaction"`
	Tx_status   RpcTxStatus        `json:"tx_status"`
}

type RpcTxStatus struct {
	Status     Status  `json:"status"`
	Block_hash *string `json:"block_hash"`
}

type RpcHeader struct {
	Version           string `json:"version"`
	Parent_hash       string `json:"parent_hash"`
	Timestamp         string `json:"timestamp"`
	Number            string `json:"number"`
	Epoch             string `json:"epoch"`
	Transactions_root string `json:"transactions_root"`
	Witnesses_root    string `json:"witnesses_root"`
	Proposals_hash    string `json:"proposals_hash"`
	Difficulty        string `json:"difficulty"`
	Uncles_hash       string `json:"uncles_hash"`
	Uncles_count      string `json:"uncles_count"`
	Dao               string `json:"dao"`
	Nonce             string `json:"nonce"`
}

type RpcHeaderView struct {
	Inner RpcHeader `json:"inner"`
	Hash  string    `json:"hash"`
}

type RpcUncleBlock struct {
	Header    RpcHeader `json:"header"`
	Proposals []string  `json:"proposals"`
}

type RpcUncleBlockView struct {
	Header    RpcHeaderView `json:"header"`
	Proposals []string      `json:"proposals"`
}

type RpcBlock struct {
	Header       RpcHeader        `json:"header"`
	Uncles       []RpcUncleBlock  `json:"uncles"`
	Transactions []RpcTransaction `json:"transactions"`
	Proposals    []string         `json:"proposals"`
}

type RpcBlockView struct {
	Header       RpcHeaderView        `json:"header"`
	Uncles       []RpcUncleBlockView  `json:"uncles"`
	Transactions []RpcTransactionView `json:"transactions"`
	Proposals    []string             `json:"proposals"`
}

type RpcEpochView struct {
	Number       string `json:"number"`
	Start_number string `json:"start_number"`
	Length       string `json:"length"`
	Difficulty   string `json:"difficulty"`
}

type RpcBlockReward struct {
	Total           string `json:"total"`
	Primary         string `json:"primary"`
	Secondary       string `json:"secondary"`
	Tx_fee          string `json:"tx_fee"`
	Proposal_reward string `json:"proposal_reward"`
}

type RpcCellOutputWithOutPoint struct {
	Out_point  RpcOutPoint `json:"out_point"`
	Block_hash string      `json:"block_hash"`
	Capacity   string      `json:"capacity"`
	Lock       RpcScript   `json:"lock"`
}

type RpcCellWithStatus struct {
	Cell   *RpcCellOutput `json:"cell"`
	Status string         `json:"status"`
}

type RpcChainInfo struct {
	Chain                     string            `json:"chain"`
	Median_time               string            `json:"median_time"`
	Epoch                     string            `json:"epoch"`
	Difficulty                string            `json:"difficulty"`
	Is_initial_block_download bool              `json:"is_initial_block_download"`
	Alerts                    []RpcAlertMessage `json:"alerts"`
}

type RpcDryRunResult struct {
	Cycles string `json:"cycles"`
}

type RpcLiveCell struct {
	Created_by  RpcTransactionPoint `json:"created_by"`
	Cell_output RpcCellOutput       `json:"cell_output"`
}

type RpcCellTransaction struct {
	Created_by  RpcTransactionPoint  `json:"created_by"`
	Consumed_by *RpcTransactionPoint `json:"consumed_by"`
}

type RpcTransactionPoint struct {
	Block_number string `json:"block_number"`
	Tx_hash      string `json:"tx_hash"`
	Index        string `json:"index"`
}

type RpcLockHashIndexState struct {
	Lock_hash    string `json:"lock_hash"`
	Block_number string `json:"block_number"`
	Block_hash   string `json:"block_hash"`
}

type RpcNode struct {
	Version     string           `json:"version"`
	Node_id     string           `json:"node_id"`
	Addresses   []RpcNodeAddress `json:"addresses"`
	Is_outbound *bool            `json:"is_outbound"`
}

type RpcNodeAddress struct {
	Address string `json:"address"`
	Score   string `json:"score"`
}

type RpcBannedAddress struct {
	Address    string `json:"address"`
	Ban_until  string `json:"ban_until"`
	Ban_reason string `json:"ban_reason"`
	Created_at string `json:"created_at"`
}

type RpcTxPoolInfo struct {
	Pending             string `json:"pending"`
	Proposed            string `json:"proposed"`
	Orphan              string `json:"orphan"`
	Total_tx_size       string `json:"total_tx_size"`
	Total_tx_cycles     string `json:"total_tx_cycles"`
	Last_txs_updated_at string `json:"last_txs_updated_at"`
}

type RpcPeerState struct {
	peer             string
	last_updated     string
	blocks_in_flight string
}
