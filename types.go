package ckb_sdk_go

/**
 * @see https //github.com/nervosnetwork/ckb/blob/develop/protocol/src/protocol.fbs for more infGomation
 */

// TransactionStatus

type TransactionStatus string

const (
	Pending   TransactionStatus = "pending"
	Proposed  TransactionStatus = "proposed"
	Committed TransactionStatus = "committed"
)

type ScriptHashType string

const (
	Data ScriptHashType = "data"
	Type ScriptHashType = "type"
)

type DepType string

const (
	Code     ScriptHashType = "code"
	DepGroup ScriptHashType = "depGroup"
)

type CellStatus string

const (
	Live    CellStatus = "live"
	Unknown CellStatus = "unknown"
)

type CapacityUnit int

const (
	Shannon CapacityUnit = 1
	Byte    CapacityUnit = 100000000
)

type DAO string
type Hash = string
type Hash256 = string
type UInt32 = uint32
type Index = string
type Version = string
type Count = string
type Difficulty = string
type BlockNumber = string
type EpochInHeader = string
type Capacity = string
type ProposalShortId = string
type Timestamp = string
type Nonce = string
type Cycles = string
type Size = string
type Bytes = string
type Since = string

/**
 * RPC Units
 */

/* eslint-disable max-len */
/**
 * @typedef Script, lock or type script
 * @description Script, the script model in CKB. CKB scripts use UNIX standard execution environment. Each script binary should contain a main function with the following signature `int main(int argc, char* argv[]);`. CKB will concat `signed_args` and `args`, then use the concatenated array to fill `argc/argv` part, then start the script execution. Upon termination, the executed `main` function here will provide a return code, `0` means the script execution succeeds, other values mean the execution fails.
 * @property args, arguments.
 * @property codeHash, point to its dependency, if the referred dependency is listed in the deps field in a transaction, the codeHash means the hash of the referred cell's data.
 * @property hashType, a enumerate indicates the type of the code which is referened by the code hash
 */
/* eslint-enable max-len */
type Script struct {
	args     []Bytes
	codeHash Hash256
	hashType ScriptHashType
}

/**
 * @typedef CellInput, cell input in a transaction
 * @property previousOutput, point to its P1 cell
 * @property since, a parameter to prevent a cell to be spent before a centain block timestamp or a block number,
 *           [RFC](https //github.com/nervosnetwork/rfcs/blob/master/rfcs/0017-tx-valid-since/0017-tx-valid-since.md)
 */
type CellInput struct {
	previousOutput OutPoint
	since          Since
}

/**
 * @typedef CellOutput, cell output in a transaction
 * @property capacity, the capacity of the genereated P1 cell
 * @property lock, lock script
 * @property type, type script
 */
type CellOutput struct {
	capacity Capacity
	lock     Script
	type_    Script
}

/**
 * @typedef OutPoint, used to refer a generated cell by transaction hash and output index
 * @property hash, transaction hash
 * @property index, index of cell output
 */
type OutPoint struct {
	txHash Hash256
	index  Index
}

/**
 * @typeof CellDep, cell dependencies in a transaction
 * @property outPoint, the out point of the cell dependency
 * @property depType, indicate if the data of the cell containing a group of dependencies
 */
type CellDep struct {
	outPoint OutPoint
	depType  DepType
}

type Witness struct {
	data []Hash
}

/**
 * @typedef RawTransaction, raw transaction object
 * @property version, transaction version
 * @property cellDeps, cell deps used in the transaction
 * @property headerDeps, header deps referenced to a specific block used in the transaction
 * @property inputs, cell inputs in the transaction
 * @property outputs, cell outputs in the transaction
 * @property witnesses, segrated witnesses
 * @property outputsData, data referenced by scripts
 */
type RawTransaction struct {
	version     Version
	cellDeps    []CellDep
	headerDeps  []Hash256
	inputs      []CellInput
	outputs     []CellOutput
	witnesses   []Witness
	outputsData []Bytes
}

/**
 * @typedef Transaction, transaction object
 * @extends RawTransaction
 * @property hash, transaction hash
 */
type Transaction struct {
	RawTransaction
	hash Hash256
}

type TxStatus struct {
	status    TransactionStatus
	blockHash Hash256
}

type TransactionWithStatus struct {
	transaction Transaction
	txStatus    TxStatus
}

/**
 * @typeof TransactionPoint
 * @property blockNumber
 * @property index
 * @property txHash
 */
type TransactionPoint struct {
	blockNumber BlockNumber
	index       Index
	txHash      Hash256
}

/**
 * @TransactionByLockHash
 * @property consumedBy
 * @property createdBy
 */
type TransactionByLockHash struct {
	consumedBy TransactionPoint
	createdBy  TransactionPoint
}

type TransactionsByLockHash []TransactionByLockHash

/**
 * @typedef BlockHeader, header of a block
 * @property dao
 * @property difficulty
 * @property epoch
 * @property hash
 * @property number
 * @property parentHash
 * @property proposalsHash
 * @property nonce
 * @property timestamp
 * @property transactionsRoot
 * @property unclesCount
 * @property unclesHash
 * @property witnessesRoot
 * @property version
 */
type BlockHeader struct {
	dao              DAO
	difficulty       Difficulty
	epoch            EpochInHeader
	hash             Hash256
	number           BlockNumber
	parentHash       Hash256
	proposalsHash    Hash256
	nonce            Nonce
	timestamp        Timestamp
	transactionsRoot Hash256
	unclesCount      Count
	unclesHash       Hash256
	witnessesRoot    Hash256
	version          Version
}

/**
 * @typedef UncleBlock, uncle block object
 * @property header, block header
 * @property proposals
 */

type UncleBlock struct {
	header    BlockHeader
	proposals []ProposalShortId
}

/**
 * @typedef Block, block object
 * @property header, block header
 * @property uncles, uncle blocks
 * @property transactions
 * @property proposals
 */
type Block struct {
	header       BlockHeader
	uncles       []UncleBlock
	transactions []Transaction
	proposals    []ProposalShortId
}

/**
 * @typedef Cell, cell object
 * @property capacty, cell capacity
 * @property lock, lock hash
 */
type Cell CellOutput

/**
 * @typedef Cell, cell object
 * @property capacty, cell capacity
 * @property lock, lock hash
 * @property outPoint
 */

type CellIncludingOutPoint struct {
	capacity Capacity
	lock     Script
	outPoint OutPoint
}

type Action struct {
	action string
	info   string
	time   Timestamp
}
type TransactionTrace = []Action

type LiveCellByLockHash struct {
	cellOutput CellOutput
	createdBy  TransactionPoint
}

type LiveCellsByLockHash []LiveCellByLockHash

type AlertMessage struct {
	id          string
	priority    string
	noticeUntil Timestamp
	message     string
}

type BlockchainInfo struct {
	isInitialBlockDownload bool
	epoch                  string
	difficulty             string
	medianTime             string
	chain                  string
	alerts                 []AlertMessage
}

type AddressInfo struct {
	address string
	score   string
}

type NodeInfo struct {
	version    string
	nodeId     string
	addresses  []AddressInfo
	isOutbound bool
}

type PeersState struct {
	lastUpdated    string
	blocksInFlight string
	peer           string
}

type TxPoolInfo struct {
	orphan           Count
	pending          Count
	proposed         Count
	lastTxsUpdatedAt Timestamp
	totalTxCycles    Cycles
	totalTxSize      Size
}

type Epoch struct {
	difficulty  string
	length      string
	number      string
	startNumber string
}

type RunDryResult struct {
	cycles Cycles
}

type LockHashIndexState struct {
	blockHash   Hash256
	blockNumber BlockNumber
	lockHash    Hash256
}

type LockHashIndexStates = []LockHashIndexState

type BannedAddress struct {
	address   string
	banReason string
	banUntil  Timestamp
	createdAt Timestamp
}

type BannedAddresses = []BannedAddress

type CellbaseOutputCapacityDetails struct {
	primary        string
	proposalReward string
	secondary      string
	total          string
	txFee          string
}
