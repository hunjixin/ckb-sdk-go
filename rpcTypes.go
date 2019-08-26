package ckb_sdk_go

type RpcScript struct {
	args      []Bytes
	code_hash Hash256
	hash_type ScriptHashType
}

type RpcOutPoint struct {
	tx_hash Hash256
	index   Index
}

type RpcCellInput struct {
	previous_output RpcOutPoint
	since           Since
}

type RpcCellOutput struct {
	capacity Capacity
	lock     RpcScript
	type_    RpcScript `json:"type"`
}

type RpcCell = RpcCellOutput

type RpcCellWithStatus struct {
	RpcCell
	status CellStatus
}

type RpcCellDep struct {
	out_point RpcOutPoint
	dep_type  DepType
}

type RpcCellIncludingOutPoint struct {
	capacity  Capacity
	lock      RpcScript
	out_point RpcOutPoint
}

type RpcRawTransaction struct {
	version      Version
	cell_deps    []RpcCellDep
	header_deps  []Hash256
	inputs       []RpcCellInput
	outputs      []RpcCellOutput
	witnesses    []Witness
	outputs_data []Bytes
}

type RpcTransaction struct {
	RpcRawTransaction
	hash Hash256
}

type RpcTransactionWithStatus struct {
	transaction RpcTransaction
	tx_status   TxStatus
}

type RpcTransactionPoint struct {
	block_number BlockNumber
	index        Index
	tx_hash      Hash256
}

type RpcTransactionByLockHash struct {
	consumed_by RpcTransactionPoint
	created_by  RpcTransactionPoint
}
type RpcTransactionsByLockHash []RpcTransactionByLockHash

type RpcLiveCellByLockHash struct {
	cell_output RpcCellOutput
	created_by  RpcTransactionPoint
}
type RpcLiveCellsByLockHash = []RpcLiveCellByLockHash

type RpcHeader struct {
	dao               DAO
	difficulty        Difficulty
	epoch             EpochInHeader
	hash              Hash256
	number            BlockNumber
	parent_hash       Hash256
	proposals_hash    Hash256
	nonce             Nonce
	timestamp         Timestamp
	transactions_root Hash256
	uncles_count      Count
	uncles_hash       Hash256
	witnesses_root    Hash256
	version           Version
}

type RpcUncleBlock struct {
	header    RpcHeader
	proposals []ProposalShortId
}

type RpcBlock struct {
	header       RpcHeader
	uncles       []RpcUncleBlock
	transactions []RpcTransaction
	proposals    []ProposalShortId
}

type RpcAlertMessage struct {
	id           string
	priority     string
	notice_until Timestamp
	message      string
}

type RpcBlockchainInfo struct {
	is_initial_block_download bool
	epoch                     string
	difficulty                string
	median_time               string
	chain                     string
	alerts                    []RpcAlertMessage
}

type RpcNodeInfo struct {
	addresses   []AddressInfo
	node_id     string
	is_outbound bool
	version     string
}

type RpcPeersState struct {
	last_updated     string
	blocks_in_flight string
	peer             string
}

type RpcTxPoolInfo struct {
	orphan              Count
	pending             Count
	proposed            Count
	last_txs_updated_at Timestamp
	total_tx_cycles     Cycles
	total_tx_size       Size
}

type RpcEpoch struct {
	difficulty   string
	length       string
	number       string
	start_number string
}

type RpcLockHashIndexState struct {
	block_hash   Hash256
	block_number BlockNumber
	lock_hash    Hash256
}

type RpcLockHashIndexStates []RpcLockHashIndexState

type RpcBannedAddress struct {
	address    string
	ban_reason string
	ban_until  Timestamp
	created_at Timestamp
}
type RpcBannedAddresses []RpcBannedAddress

type RpcCellbaseOutputCapacityDetails struct {
	primary         string
	proposal_reward string
	secondary       string
	total           string
	tx_fee          string
}
