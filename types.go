package ckb_sdk_go

/**
 * @see https //github.com/nervosnetwork/ckb/blob/develop/protocol/src/protocol.fbs for more infGomation
 */

// TransactionStatus

type Status string

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

type JsonBytes Bytes
type AlertId string
type DAO string
type Hash = string
type H256 = string
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
type Cycle = string
type Size = string
type Unsigned = string
type Bytes = string
type Since = string
type AlertPriority = string
type EpochNumber = string
type U256 = string
