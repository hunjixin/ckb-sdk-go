package ckb_sdk_go

import (
	"reflect"
)

var TAlert = reflect.TypeOf(Alert{})

type Alert struct {
	Id           uint32
	Cancel       uint32
	Signatures   [][]byte
	Min_version  *string
	Max_version  *string
	Priority     uint32
	Notice_until uint64
	Message      string
}

var TAlertMessage = reflect.TypeOf(AlertMessage{})

type AlertMessage struct {
	Id           uint32
	Priority     uint32
	Notice_until uint64
	Message      string
}

var TBlockTemplate = reflect.TypeOf(BlockTemplate{})

type BlockTemplate struct {
	Version            uint32
	Difficulty         [32]byte
	Current_time       uint64
	Number             uint64
	Epoch              uint64
	Parent_hash        [32]byte
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

var TUncleTemplate = reflect.TypeOf(UncleTemplate{})

type UncleTemplate struct {
	Hash      [32]byte
	Required  bool
	Proposals [][10]byte
	Header    Header
}

var TCellbaseTemplate = reflect.TypeOf(CellbaseTemplate{})

type CellbaseTemplate struct {
	Hash   [32]byte
	Cycles *uint64
	Data   Transaction
}

var TTransactionTemplate = reflect.TypeOf(TransactionTemplate{})

type TransactionTemplate struct {
	Hash     [32]byte
	Required bool
	Cycles   *uint64
	Depends  *[]uint64
	Data     Transaction
}

var TScript = reflect.TypeOf(Script{})

type Script struct {
	Args      [][]byte
	Code_hash [32]byte
	Hash_type ScriptHashType
}

var TCellOutput = reflect.TypeOf(CellOutput{})

type CellOutput struct {
	Capacity uint64
	Data     []byte
	Lock     Script
	Type_    *Script
}
CellOutput
var TCellOutPoint = reflect.TypeOf(CellOutPoint{})

type CellOutPoint struct {
	Tx_hash [32]byte
	Index   uint64
}

var TOutPoint = reflect.TypeOf(OutPoint{})

type OutPoint struct {
	Cell       *CellOutPoint
	Block_hash *[32]byte
}

var TCellInput = reflect.TypeOf(CellInput{})

type CellInput struct {
	Previous_output OutPoint
	Since           uint64
}

var TWitness = reflect.TypeOf(Witness{})


var TTransaction = reflect.TypeOf(Transaction{})

type Transaction struct {
	Version   uint32
	Deps      []OutPoint
	Inputs    []CellInput
	Outputs   []CellOutput
	Witnesses []Witness
}

var TTransactionView = reflect.TypeOf(TransactionView{})

type TransactionView struct {
	Inner Transaction
	Hash  [32]byte
}

var TTransactionWithStatus = reflect.TypeOf(TransactionWithStatus{})

type TransactionWithStatus struct {
	Transaction TransactionView
	Tx_status   TxStatus
}

var TTxStatus = reflect.TypeOf(TxStatus{})

type TxStatus struct {
	Status     Status
	Block_hash *[32]byte
}

var TSeal = reflect.TypeOf(Seal{})

type Seal struct {
	Nonce uint64
	Proof []byte
}

var THeader = reflect.TypeOf(Header{})

type Header struct {
	Version           uint32
	Parent_hash       [32]byte
	Timestamp         uint64
	Number            uint64
	Epoch             uint64
	Transactions_root [32]byte
	Witnesses_root    [32]byte
	Proposals_hash    [32]byte
	Difficulty        [32]byte
	Uncles_hash       [32]byte
	Uncles_count      uint64
	Dao               []byte
	Seal              Seal
}

var THeaderView = reflect.TypeOf(HeaderView{})

type HeaderView struct {
	Inner Header
	Hash  [32]byte
}

var TUncleBlock = reflect.TypeOf(UncleBlock{})

type UncleBlock struct {
	Header    Header
	Proposals [][10]byte
}

var TUncleBlockView = reflect.TypeOf(UncleBlockView{})

type UncleBlockView struct {
	Header    HeaderView
	Proposals [][10]byte
}

var TBlock = reflect.TypeOf(Block{})

type Block struct {
	Header       Header
	Uncles       []UncleBlock
	Transactions []Transaction
	Proposals    [][10]byte
}

var TBlockView = reflect.TypeOf(BlockView{})

type BlockView struct {
	Header       HeaderView
	Uncles       []UncleBlockView
	Transactions []TransactionView
	Proposals    [][10]byte
}

var TEpochView = reflect.TypeOf(EpochView{})

type EpochView struct {
	Number       uint64
	Epoch_reward uint64
	Start_number uint64
	Length       uint64
	Difficulty   [32]byte
}

var TBlockRewardView = reflect.TypeOf(BlockRewardView{})

type BlockRewardView struct {
	Total           uint64
	Primary         uint64
	Secondary       uint64
	Tx_fee          uint64
	Proposal_reward uint64
}

var TCellOutputWithOutPoint = reflect.TypeOf(CellOutputWithOutPoint{})

type CellOutputWithOutPoint struct {
	Out_point OutPoint
	Capacity  uint64
	Lock      Script
}

var TCellWithStatus = reflect.TypeOf(CellWithStatus{})

type CellWithStatus struct {
	Cell   *CellOutput
	Status string
}

var TChainInfo = reflect.TypeOf(ChainInfo{})

type ChainInfo struct {
	Chain                     string
	Median_time               uint64
	Epoch                     uint64
	Difficulty                [32]byte
	Is_initial_block_download bool
	Alerts                    []AlertMessage
}

var TDryRunResult = reflect.TypeOf(DryRunResult{})

type DryRunResult struct {
	Cycles uint64
}

var TLiveCell = reflect.TypeOf(LiveCell{})

type LiveCell struct {
	Created_by  TransactionPoint
	Cell_output CellOutput
}

var TCellTransaction = reflect.TypeOf(CellTransaction{})

type CellTransaction struct {
	Created_by  TransactionPoint
	Consumed_by *TransactionPoint
}

var TTransactionPoint = reflect.TypeOf(TransactionPoint{})

type TransactionPoint struct {
	Block_number uint64
	Tx_hash      [32]byte
	Index        uint64
}

var TLockHashIndexState = reflect.TypeOf(LockHashIndexState{})

type LockHashIndexState struct {
	Lock_hash    [32]byte
	Block_number uint64
	Block_hash   [32]byte
}

var TNode = reflect.TypeOf(Node{})

type Node struct {
	Version     string
	Node_id     string
	Addresses   []NodeAddress
	Is_outbound *bool
}

var TNodeAddress = reflect.TypeOf(NodeAddress{})

type NodeAddress struct {
	Address string
	Score   uint64
}

var TBannedAddress = reflect.TypeOf(BannedAddress{})

type BannedAddress struct {
	Address    string
	Ban_until  uint64
	Ban_reason string
	Created_at uint64
}

var TTxPoolInfo = reflect.TypeOf(TxPoolInfo{})

type TxPoolInfo struct {
	Pending             uint64
	Proposed            uint64
	Orphan              uint64
	Total_tx_size       uint64
	Total_tx_cycles     uint64
	Last_txs_updated_at uint64
}

var TPeerState = reflect.TypeOf(PeerState{})

type PeerState struct {
	peer             uint64
	last_updated     uint64
	blocks_in_flight uint64
}
