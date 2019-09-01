package ckb_sdk_go


type Status string

const (
	Pending   Status = "pending"
	Proposed  Status = "proposed"
	Committed Status = "committed"
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













