package core

import (
	"reflect"
)

var TAlert = reflect.TypeOf(Alert{})
var TAlertMessage = reflect.TypeOf(AlertMessage{})
var TBlockTemplate = reflect.TypeOf(BlockTemplate{})
var TUncleTemplate = reflect.TypeOf(UncleTemplate{})
var TCellbaseTemplate = reflect.TypeOf(CellbaseTemplate{})
var TTransactionTemplate = reflect.TypeOf(TransactionTemplate{})
var TScript = reflect.TypeOf(Script{})
var TCellOutput = reflect.TypeOf(CellOutput{})
var TOutPoint = reflect.TypeOf(OutPoint{})
var TCellInput = reflect.TypeOf(CellInput{})
var TWitness = reflect.TypeOf(Witness{})
var TCellDep = reflect.TypeOf(CellDep{})
var TTransaction = reflect.TypeOf(Transaction{})
var TTransactionView = reflect.TypeOf(TransactionView{})
var TTransactionWithStatus = reflect.TypeOf(TransactionWithStatus{})
var TTxStatus = reflect.TypeOf(TxStatus{})
var THeader = reflect.TypeOf(Header{})
var THeaderView = reflect.TypeOf(HeaderView{})
var TUncleBlock = reflect.TypeOf(UncleBlock{})
var TUncleBlockView = reflect.TypeOf(UncleBlockView{})
var TBlock = reflect.TypeOf(Block{})
var TBlockView = reflect.TypeOf(BlockView{})
var TEpochView = reflect.TypeOf(EpochView{})
var TBlockReward = reflect.TypeOf(BlockReward{})
var TCellOutputWithOutPoint = reflect.TypeOf(CellOutputWithOutPoint{})
var TCellWithStatus = reflect.TypeOf(CellWithStatus{})
var TChainInfo = reflect.TypeOf(ChainInfo{})
var TDryRunResult = reflect.TypeOf(DryRunResult{})
var TLiveCell = reflect.TypeOf(LiveCell{})
var TCellTransaction = reflect.TypeOf(CellTransaction{})
var TTransactionPoint = reflect.TypeOf(TransactionPoint{})
var TLockHashIndexState = reflect.TypeOf(LockHashIndexState{})
var TNode = reflect.TypeOf(Node{})
var TNodeAddress = reflect.TypeOf(NodeAddress{})
var TBannedAddress = reflect.TypeOf(BannedAddress{})
var TTxPoolInfo = reflect.TypeOf(TxPoolInfo{})
var TPeerState = reflect.TypeOf(PeerState{})

type Alert struct {
	Id           uint32
	Cancel       uint32
	Min_version  *string
	Max_version  *string
	Priority     uint32
	Notice_until uint64
	Message      string
	Signatures   [][]byte
}

type AlertMessage struct {
	Id           uint32
	Priority     uint32
	Notice_until uint64
	Message      string
}

type BlockTemplate struct {
	Version            uint32
	Difficulty         U256
	Current_time       uint64
	Number             uint64
	Epoch              uint64
	Parent_hash        H256
	Cycles_limit       uint64
	Bytes_limit        uint64
	Uncles_count_limit uint64
	Uncles             []UncleTemplate
	Transactions       []TransactionTemplate
	Proposals          [][10]byte
	Cellbase           CellbaseTemplate
	Work_id            uint64
	Dao                []byte
}

type UncleTemplate struct {
	Hash      H256
	Required  bool
	Proposals [][10]byte
	Header    Header
}

type CellbaseTemplate struct {
	Hash   H256
	Cycles *uint64
	Data   Transaction
}

type TransactionTemplate struct {
	Hash     H256
	Required bool
	Cycles   *uint64
	Depends  *[]uint64
	Data     Transaction
}

type Script struct {
	Args      [][]byte
	Code_hash H256
	Hash_type ScriptHashType
}

type CellOutput struct {
	Capacity uint64
	Lock     Script
	Type_    *Script
}

type OutPoint struct {
	Tx_hash H256
	Index   uint64
}

type CellInput struct {
	Previous_output OutPoint
	Since           uint64
}


type CellDep struct {
	out_point OutPoint
	dep_type  DepType
}

type Transaction struct {
	Version      uint32
	Cell_deps    []CellDep
	Header_deps  []H256
	Inputs       []CellInput
	Outputs      []CellOutput
	Witnesses    []Witness
	Outputs_data [][]byte
}

type TransactionView struct {
	Inner Transaction
	Hash  H256
}

type TransactionWithStatus struct {
	Transaction TransactionView
	Tx_status   TxStatus
}

type TxStatus struct {
	Status     Status
	Block_hash *H256
}

type Header struct {
	Version           uint32
	Parent_hash       H256
	Timestamp         uint64
	Number            uint64
	Epoch             uint64
	Transactions_root H256
	Witnesses_root    H256
	Proposals_hash    H256
	Difficulty        U256
	Uncles_hash       H256
	Uncles_count      uint64
	Dao               []byte
	Nonce             uint64
}

type HeaderView struct {
	Inner Header
	Hash  H256
}

type UncleBlock struct {
	Header    Header
	Proposals [][10]byte
}

type UncleBlockView struct {
	Header    HeaderView
	Proposals [][10]byte
}

type Block struct {
	Header       Header
	Uncles       []UncleBlock
	Transactions []Transaction
	Proposals    [][10]byte
}

type BlockView struct {
	Header       HeaderView
	Uncles       []UncleBlockView
	Transactions []TransactionView
	Proposals    [][10]byte
}

type EpochView struct {
	Number       uint64
	Start_number uint64
	Length       uint64
	Difficulty   U256
}

type BlockReward struct {
	Total           uint64
	Primary         uint64
	Secondary       uint64
	Tx_fee          uint64
	Proposal_reward uint64
}

type CellOutputWithOutPoint struct {
	Out_point  OutPoint
	Block_hash H256
	Capacity   uint64
	Lock       Script
}

type CellWithStatus struct {
	Cell   *CellOutput
	Status string
}

type ChainInfo struct {
	Chain                     string
	Median_time               uint64
	Epoch                     uint64
	Difficulty                U256
	Is_initial_block_download bool
	Alerts                    []AlertMessage
}

type DryRunResult struct {
	Cycles uint64
}

type LiveCell struct {
	Created_by  TransactionPoint
	Cell_output CellOutput
}

type CellTransaction struct {
	Created_by  TransactionPoint
	Consumed_by *TransactionPoint
}

type TransactionPoint struct {
	Block_number uint64
	Tx_hash      H256
	Index        uint64
}

type LockHashIndexState struct {
	Lock_hash    H256
	Block_number uint64
	Block_hash   H256
}

type Node struct {
	Version     string
	Node_id     string
	Addresses   []NodeAddress
	Is_outbound *bool
}

type NodeAddress struct {
	Address string
	Score   uint64
}

type BannedAddress struct {
	Address    string
	Ban_until  uint64
	Ban_reason string
	Created_at uint64
}

type TxPoolInfo struct {
	Pending             uint64
	Proposed            uint64
	Orphan              uint64
	Total_tx_size       uint64
	Total_tx_cycles     uint64
	Last_txs_updated_at uint64
}

type PeerState struct {
	peer             uint64
	last_updated     uint64
	blocks_in_flight uint64
}
