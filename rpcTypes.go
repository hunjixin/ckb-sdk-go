package ckb_sdk_go

type RpcScript struct {
	Args      []Bytes
	Code_hash Hash256
	Hash_type ScriptHashType
}

type RpcOutPoint struct {
	Tx_hash Hash256
	Index   Index
}

type RpcCellInput struct {
	Previous_output RpcOutPoint
	Since           Since
}

type RpcCellOutput struct {
	Capacity Capacity
	Lock     RpcScript
	Type_    RpcScript `json:"type"`
}

type RpcCell = RpcCellOutput

type RpcCellWithStatus struct {
	RpcCell
	Status CellStatus
}

type RpcCellDep struct {
	Out_point RpcOutPoint
	Dep_type  DepType
}

type RpcCellIncludingOutPoint struct {
	Capacity  Capacity
	Lock      RpcScript
	Out_point RpcOutPoint
}

type RpcRawTransaction struct {
	Version      Version
	Cell_deps    []RpcCellDep
	Header_deps  []Hash256
	Inputs       []RpcCellInput
	Outputs      []RpcCellOutput
	//Witnesses    []Witness
	Outputs_data []Bytes
}

type RpcTransaction struct {
	RpcRawTransaction
	Hash Hash256
}

type RpcTransactionWithStatus struct {
	Transaction RpcTransaction
	Tx_status   TxStatus
}

type RpcTransactionPoint struct {
	Block_number BlockNumber
	Index        Index
	Tx_hash      Hash256
}

type RpcTransactionByLockHash struct {
	Consumed_by RpcTransactionPoint
	Created_by  RpcTransactionPoint
}
type RpcTransactionsByLockHash []RpcTransactionByLockHash

type RpcLiveCellByLockHash struct {
	Cell_output RpcCellOutput
	Created_by  RpcTransactionPoint
}
type RpcLiveCellsByLockHash = []RpcLiveCellByLockHash

type RpcHeader struct {
	Dao               DAO
	Difficulty        Difficulty
	Epoch             EpochInHeader
	Hash              Hash256
	Number            BlockNumber
	Parent_hash       Hash256
	Proposals_hash    Hash256
	Nonce             Nonce
	Timestamp         Timestamp
	Transactions_root Hash256
	Uncles_count      Count
	Uncles_hash       Hash256
	Witnesses_root    Hash256
	Version           Version
}

type RpcUncleBlock struct {
	Header    RpcHeader
	Proposals []ProposalShortId
}

type RpcBlock struct {
	Header       RpcHeader
	Uncles       []RpcUncleBlock
	Transactions []RpcTransaction
	Proposals    []ProposalShortId
}

type RpcAlertMessage struct {
	Id           string
	Priority     string
	Notice_until Timestamp
	Message      string
}

type RpcBlockchainInfo struct {
	Is_initial_block_download bool
	Epoch                     string
	Difficulty                string
	Median_time               string
	Chain                     string
	Alerts                    []RpcAlertMessage
}

type RpcNodeInfo struct {
	Addresses   []AddressInfo
	Node_id     string
	Is_outbound bool
	Version     string
}

type RpcPeersState struct {
	Last_updated     string
	Blocks_in_flight string
	Peer             string
}

type RpcTxPoolInfo struct {
	Orphan              Count
	Pending             Count
	Proposed            Count
	Last_txs_updated_at Timestamp
	Total_tx_cycles     Cycles
	Total_tx_size       Size
}

type RpcEpoch struct {
	Difficulty   string
	Length       string
	Number       string
	Start_number string
}

type RpcLockHashIndexState struct {
	Block_hash   Hash256
	Block_number BlockNumber
	Lock_hash    Hash256
}

type RpcLockHashIndexStates []RpcLockHashIndexState

type RpcBannedAddress struct {
	Address    string
	Ban_reason string
	Ban_until  Timestamp
	Created_at Timestamp
}
type RpcBannedAddresses []RpcBannedAddress

type RpcCellbaseOutputCapacityDetails struct {
	Primary         string
	Proposal_reward string
	Secondary       string
	Total           string
	Tx_fee          string
}
