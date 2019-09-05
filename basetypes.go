package ckb_sdk_go


type Status string

const (
	Pending   Status = "pending"
	Proposed  Status = "proposed"
	Committed Status = "committed"
)

type RpcScriptHashType string

const (
	RpcData RpcScriptHashType = "data"
	RpcType RpcScriptHashType = "type"
)

type ScriptHashType uint32

const (
	Data ScriptHashType = 0
	Type ScriptHashType = 1
)

type DepType string

/*const (
	Code     ScriptHashType = "code"
	DepGroup ScriptHashType = "depGroup"
)*/

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














