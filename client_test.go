package ckb_sdk_go

import (
	"ckb-sdk-go/core"
	"github.com/ybbus/jsonrpc"
	"reflect"
	"testing"
)

func TestCkbClient_AddNode(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		peer_id string
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			if err := ckbClient.AddNode(tt.args.peer_id, tt.args.address); (err != nil) != tt.wantErr {
				t.Errorf("AddNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCkbClient_BroadcastTransaction(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		transaction core.RpcTransaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.BroadcastTransaction(tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("BroadcastTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BroadcastTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_CalculateDaoMaximumWithdraw(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_out_point core.RpcOutPoint
		_hash      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.CalculateDaoMaximumWithdraw(tt.args._out_point, tt.args._hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateDaoMaximumWithdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateDaoMaximumWithdraw() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_ComputeScriptHash(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		script core.RpcScript
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.ComputeScriptHash(tt.args.script)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComputeScriptHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeScriptHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_ComputeTransactionHash(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		tx core.RpcTransaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.ComputeTransactionHash(tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComputeTransactionHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeTransactionHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_DeindexLockHash(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_lock_hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			if err := ckbClient.DeindexLockHash(tt.args._lock_hash); (err != nil) != tt.wantErr {
				t.Errorf("DeindexLockHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCkbClient_DryRunTransaction(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_tx core.RpcTransaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcDryRunResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.DryRunTransaction(tt.args._tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("DryRunTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DryRunTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetBannedAddresses(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *core.RpcBannedAddress
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetBannedAddresses()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBannedAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBannedAddresses() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetBlock(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcBlockView
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetBlock(tt.args._hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlock() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetBlockByNumber(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_number string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcBlockView
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetBlockByNumber(tt.args._number)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockByNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlockByNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetBlockHash(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_number string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetBlockHash(tt.args._number)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlockHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetBlockTemplate(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		bytes_limit     *string
		proposals_limit *string
		max_version     *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcBlockTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetBlockTemplate(tt.args.bytes_limit, tt.args.proposals_limit, tt.args.max_version)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlockTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetBlockchainInfo(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *core.RpcChainInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetBlockchainInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockchainInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlockchainInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetCellbaseOutputCapacityDetails(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcBlockReward
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetCellbaseOutputCapacityDetails(tt.args._hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCellbaseOutputCapacityDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCellbaseOutputCapacityDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetCellsByLockHash(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_lock_hash string
		_from      string
		_to        string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcCellOutputWithOutPoint
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetCellsByLockHash(tt.args._lock_hash, tt.args._from, tt.args._to)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCellsByLockHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCellsByLockHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetCurrentEpoch(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *core.RpcEpochView
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetCurrentEpoch()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentEpoch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrentEpoch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetEpochByNumber(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		number string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcEpochView
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetEpochByNumber(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpochByNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEpochByNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetHeader(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcHeaderView
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetHeader(tt.args._hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHeader() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetHeaderByNumber(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_number string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcHeaderView
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetHeaderByNumber(tt.args._number)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHeaderByNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHeaderByNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetLiveCell(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_out_point core.RpcOutPoint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcCellWithStatus
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetLiveCell(tt.args._out_point)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLiveCell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLiveCell() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetLiveCellsByLockHash(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_lock_hash     string
		_page          string
		_per_page      string
		_reverse_order *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcLiveCell
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetLiveCellsByLockHash(tt.args._lock_hash, tt.args._page, tt.args._per_page, tt.args._reverse_order)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLiveCellsByLockHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLiveCellsByLockHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetLockHashIndexStates(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *core.RpcLockHashIndexState
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetLockHashIndexStates()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLockHashIndexStates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLockHashIndexStates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetPeers(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *core.RpcNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetPeers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPeers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetPeersState(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *core.RpcPeerState
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetPeersState()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPeersState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeersState() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetTipBlockNumber(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetTipBlockNumber()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTipBlockNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTipBlockNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetTipHeader(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *core.RpcHeaderView
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetTipHeader()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTipHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTipHeader() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetTransaction(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcTransactionWithStatus
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetTransaction(tt.args._hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_GetTransactionsByLockHash(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_lock_hash     string
		_page          string
		_per_page      string
		_reverse_order *bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcCellTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.GetTransactionsByLockHash(tt.args._lock_hash, tt.args._page, tt.args._per_page, tt.args._reverse_order)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransactionsByLockHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTransactionsByLockHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_IndexLockHash(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_lock_hash  string
		_index_from *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *core.RpcLockHashIndexState
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.IndexLockHash(tt.args._lock_hash, tt.args._index_from)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndexLockHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexLockHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_LocalNodeInfo(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *core.RpcNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.LocalNodeInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalNodeInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocalNodeInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_ProcessBlockWithoutVerify(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		data core.RpcBlock
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.ProcessBlockWithoutVerify(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessBlockWithoutVerify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessBlockWithoutVerify() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_RemoveNode(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		peer_id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			if err := ckbClient.RemoveNode(tt.args.peer_id); (err != nil) != tt.wantErr {
				t.Errorf("RemoveNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCkbClient_SendAlert(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_alert core.RpcAlert
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			if err := ckbClient.SendAlert(tt.args._alert); (err != nil) != tt.wantErr {
				t.Errorf("SendAlert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCkbClient_SendTransaction(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_tx core.RpcTransaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.SendTransaction(tt.args._tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_SetBan(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		address  string
		command  string
		ban_time *string
		absolute *bool
		reason   *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			if err := ckbClient.SetBan(tt.args.address, tt.args.command, tt.args.ban_time, tt.args.absolute, tt.args.reason); (err != nil) != tt.wantErr {
				t.Errorf("SetBan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCkbClient_SubmitBlock(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	type args struct {
		_work_id string
		_data    core.RpcBlock
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.SubmitBlock(tt.args._work_id, tt.args._data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubmitBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubmitBlock() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCkbClient_TxPoolInfo(t *testing.T) {
	type fields struct {
		url    string
		client jsonrpc.RPCClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *core.RpcTxPoolInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ckbClient := &CkbClient{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := ckbClient.TxPoolInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("TxPoolInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TxPoolInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCkbClient(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want *CkbClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCkbClient(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCkbClient() = %v, want %v", got, tt.want)
			}
		})
	}
}