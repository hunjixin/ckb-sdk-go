pub struct Alert {
    pub id: AlertId,
    pub cancel: AlertId,
    pub signatures: Vec<JsonBytes>,
    pub min_version: Option<String>,
    pub max_version: Option<String>,
    pub priority: AlertPriority,
    pub notice_until: Timestamp,
    pub message: String,

}

pub struct AlertMessage {
    pub id: AlertId,
    pub priority: AlertPriority,
    pub notice_until: Timestamp,
    pub message: String,
}

pub struct BlockTemplate {
    pub version: Version,
    pub difficulty: U256,
    pub current_time: Timestamp,
    pub number: BlockNumber,
    pub epoch: EpochNumber,
    pub parent_hash: H256,
    pub cycles_limit: Cycle,
    pub bytes_limit: Unsigned,
    pub uncles_count_limit: Unsigned,
    pub uncles: Vec<UncleTemplate>,
    pub transactions: Vec<TransactionTemplate>,
    pub proposals: Vec<ProposalShortId>,
    pub cellbase: CellbaseTemplate,
    pub work_id: Unsigned,
    pub dao: JsonBytes,
}

pub struct UncleTemplate {
    pub hash: H256,
    pub required: bool,
    pub proposals: Vec<ProposalShortId>,
    pub header: Header, // temporary
}

pub struct CellbaseTemplate {
    pub hash: H256,
    pub cycles: Option<Cycle>,
    pub data: Transaction, // temporary
}

pub struct TransactionTemplate {
    pub hash: H256,
    pub required: bool,
    pub cycles: Option<Cycle>,
    pub depends: Option<Vec<Unsigned>>,
    pub data: Transaction, // temporary
}

pub struct Script {
    pub args: Vec<JsonBytes>,
    pub code_hash: H256,
    pub hash_type: ScriptHashType,
}

pub struct CellOutput {
    pub capacity: Capacity,
    pub data: JsonBytes,
    pub lock: Script,
    pub type_: Option<Script>,
}

pub struct CellOutPoint {
    pub tx_hash: H256,
    pub index: Unsigned,
}

pub struct OutPoint {
    pub cell: Option<CellOutPoint>,
    pub block_hash: Option<H256>,
}

pub struct CellInput {
    pub previous_output: OutPoint,
    pub since: Unsigned,
}

pub struct Witness {
    data: Vec<JsonBytes>,
}

pub struct Transaction {
    pub version: Version,
    pub deps: Vec<OutPoint>,
    pub inputs: Vec<CellInput>,
    pub outputs: Vec<CellOutput>,
    pub witnesses: Vec<Witness>,
}

pub struct TransactionView {
    pub inner: Transaction,
    pub hash: H256,
}

pub struct TransactionWithStatus {
    pub transaction: TransactionView,
    /// Indicate the Transaction status
    pub tx_status: TxStatus,
}

pub struct TxStatus {
    pub status: Status,
    pub block_hash: Option<H256>,
}

pub struct Seal {
    pub nonce: Unsigned,
    pub proof: JsonBytes,
}

pub struct Header {
    pub version: Version,
    pub parent_hash: H256,
    pub timestamp: Timestamp,
    pub number: BlockNumber,
    pub epoch: EpochNumber,
    pub transactions_root: H256,
    pub witnesses_root: H256,
    pub proposals_hash: H256,
    pub difficulty: U256,
    pub uncles_hash: H256,
    pub uncles_count: Unsigned,
    pub dao: JsonBytes,
    pub seal: Seal,
}

pub struct HeaderView {
    pub inner: Header,
    pub hash: H256,
}

pub struct UncleBlock {
    pub header: Header,
    pub proposals: Vec<ProposalShortId>,
}

pub struct UncleBlockView {
    pub header: HeaderView,
    pub proposals: Vec<ProposalShortId>,
}

pub struct Block {
    pub header: Header,
    pub uncles: Vec<UncleBlock>,
    pub transactions: Vec<Transaction>,
    pub proposals: Vec<ProposalShortId>,
}

pub struct BlockView {
    pub header: HeaderView,
    pub uncles: Vec<UncleBlockView>,
    pub transactions: Vec<TransactionView>,
    pub proposals: Vec<ProposalShortId>,
}

pub struct EpochView {
    pub number: EpochNumber,
    pub epoch_reward: Capacity,
    pub start_number: BlockNumber,
    pub length: BlockNumber,
    pub difficulty: U256,
}

pub struct BlockRewardView {
    pub total: Capacity,
    pub primary: Capacity,
    pub secondary: Capacity,
    pub tx_fee: Capacity,
    pub proposal_reward: Capacity,
}

pub struct CellOutputWithOutPoint {
    pub out_point: OutPoint,
    pub capacity: Capacity,
    pub lock: Script,
}

pub struct CellWithStatus {
    pub cell: Option<CellOutput>,
    pub status: String,
}

pub struct ChainInfo {
    // network name
    pub chain: String,
    // median time for the current tip block
    pub median_time: Timestamp,
    // the current epoch number
    pub epoch: EpochNumber,
    // the current difficulty
    pub difficulty: U256,
    // estimate of whether this node is in InitialBlockDownload mode
    pub is_initial_block_download: bool,
    // any network and blockchain warnings
    pub alerts: Vec<AlertMessage>,
}

pub struct DryRunResult {
    pub cycles: Cycle,
}

pub struct LiveCell {
    pub created_by: TransactionPoint,
    pub cell_output: CellOutput,
}

pub struct CellTransaction {
    pub created_by: TransactionPoint,
    pub consumed_by: Option<TransactionPoint>,
}

pub struct TransactionPoint {
    pub block_number: BlockNumber,
    pub tx_hash: H256,
    pub index: Unsigned,
}

pub struct LockHashIndexState {
    pub lock_hash: H256,
    pub block_number: BlockNumber,
    pub block_hash: H256,
}
pub struct Node {
    pub version: String,
    pub node_id: String,
    pub addresses: Vec<NodeAddress>,
    pub is_outbound: Option<bool>,
}

pub struct NodeAddress {
    pub address: String,
    pub score: Unsigned,
}

pub struct BannedAddress {
    pub address: String,
    pub ban_until: Timestamp,
    pub ban_reason: String,
    pub created_at: Timestamp,
}
pub struct TxPoolInfo {
    pub pending: Unsigned,
    pub proposed: Unsigned,
    pub orphan: Unsigned,
    pub total_tx_size: Unsigned,
    pub total_tx_cycles: Unsigned,
    pub last_txs_updated_at: Timestamp,
}

pub struct PeerState {
    // TODO use peer_id
    // peer session id
    peer: Unsigned,
    // last updated timestamp
    last_updated: Timestamp,
    // blocks count has request but not receive response yet
    blocks_in_flight: Unsigned,
}

pub trait AlertRpc {
    // curl -d '{"id": 2, "jsonrpc": "2.0", "method":"send_alert","params": [{}]}' -H 'content-type:application/json' 'http://localhost:8114'
    #[rpc(name = "send_alert")]
    fn send_alert(&self, _alert: Alert) -> Result<()>;
}

pub trait ChainRpc {
    #[rpc(name = "get_block")]
    fn get_block(&self, _hash: H256) -> Result<Option<BlockView>>;

    #[rpc(name = "get_block_by_number")]
    fn get_block_by_number(&self, _number: BlockNumber) -> Result<Option<BlockView>>;

    #[rpc(name = "get_header")]
    fn get_header(&self, _hash: H256) -> Result<Option<HeaderView>>;

    #[rpc(name = "get_header_by_number")]
    fn get_header_by_number(&self, _number: BlockNumber) -> Result<Option<HeaderView>>;

    #[rpc(name = "get_transaction")]
    fn get_transaction(&self, _hash: H256) -> Result<Option<TransactionWithStatus>>;

    #[rpc(name = "get_block_hash")]
    fn get_block_hash(&self, _number: BlockNumber) -> Result<Option<H256>>;

    #[rpc(name = "get_tip_header")]
    fn get_tip_header(&self) -> Result<HeaderView>;

    #[rpc(name = "get_cells_by_lock_hash")]
    fn get_cells_by_lock_hash(
        &self,
        _lock_hash: H256,
        _from: BlockNumber,
        _to: BlockNumber,
    ) -> Result<Vec<CellOutputWithOutPoint>>;

    #[rpc(name = "get_live_cell")]
    fn get_live_cell(&self, _out_point: OutPoint) -> Result<CellWithStatus>;

    #[rpc(name = "get_tip_block_number")]
    fn get_tip_block_number(&self) -> Result<BlockNumber>;

    #[rpc(name = "get_current_epoch")]
    fn get_current_epoch(&self) -> Result<EpochView>;

    #[rpc(name = "get_epoch_by_number")]
    fn get_epoch_by_number(&self, number: EpochNumber) -> Result<Option<EpochView>>;

    #[rpc(name = "get_cellbase_output_capacity_details")]
    fn get_cellbase_output_capacity_details(&self, _hash: H256) -> Result<Option<BlockRewardView>>;
}

pub trait ExperimentRpc {
    #[rpc(name = "_compute_transaction_hash")]
    fn compute_transaction_hash(&self, tx: Transaction) -> Result<H256>;

    #[rpc(name = "_compute_script_hash")]
    fn compute_script_hash(&self, script: Script) -> Result<H256>;

    #[rpc(name = "dry_run_transaction")]
    fn dry_run_transaction(&self, _tx: Transaction) -> Result<DryRunResult>;

    // Calculate the maximum withdraw one can get, given a referenced DAO cell,
    // and a withdraw block hash
    #[rpc(name = "calculate_dao_maximum_withdraw")]
    fn calculate_dao_maximum_withdraw(&self, _out_point: OutPoint, _hash: H256)
        -> Result<Capacity>;
}

pub trait IndexerRpc {
    #[rpc(name = "get_live_cells_by_lock_hash")]
    fn get_live_cells_by_lock_hash(
        &self,
        _lock_hash: H256,
        _page: Unsigned,
        _per_page: Unsigned,
        _reverse_order: Option<bool>,
    ) -> Result<Vec<LiveCell>>;

    #[rpc(name = "get_transactions_by_lock_hash")]
    fn get_transactions_by_lock_hash(
        &self,
        _lock_hash: H256,
        _page: Unsigned,
        _per_page: Unsigned,
        _reverse_order: Option<bool>,
    ) -> Result<Vec<CellTransaction>>;

    #[rpc(name = "index_lock_hash")]
    fn index_lock_hash(
        &self,
        _lock_hash: H256,
        _index_from: Option<BlockNumber>,
    ) -> Result<LockHashIndexState>;

    #[rpc(name = "deindex_lock_hash")]
    fn deindex_lock_hash(&self, _lock_hash: H256) -> Result<()>;

    #[rpc(name = "get_lock_hash_index_states")]
    fn get_lock_hash_index_states(&self) -> Result<Vec<LockHashIndexState>>;
}

pub trait MinerRpc {
    // curl -d '{"id": 2, "jsonrpc": "2.0", "method":"get_block_template","params": ["0x1b1c832d02fdb4339f9868c8a8636c3d9dd10bd53ac7ce99595825bd6beeffb3", 1000, 1000]}' -H 'content-type:application/json' 'http://localhost:8114'
    #[rpc(name = "get_block_template")]
    fn get_block_template(
        &self,
        bytes_limit: Option<Unsigned>,
        proposals_limit: Option<Unsigned>,
        max_version: Option<Version>,
    ) -> Result<BlockTemplate>;

    // curl -d '{"id": 2, "jsonrpc": "2.0", "method":"submit_block","params": [{"header":{}, "uncles":[], "transactions":[], "proposals":[]}]}' -H 'content-type:application/json' 'http://localhost:8114'
    #[rpc(name = "submit_block")]
    fn submit_block(&self, _work_id: String, _data: Block) -> Result<Option<H256>>;
}

pub trait NetworkRpc {
    // curl -d '{"id": 2, "jsonrpc": "2.0", "method":"local_node_info","params": []}' -H 'content-type:application/json' 'http://localhost:8114'
    #[rpc(name = "local_node_info")]
    fn local_node_info(&self) -> Result<Node>;

    // curl -d '{"id": 2, "jsonrpc": "2.0", "method":"get_peers","params": []}' -H 'content-type:application/json' 'http://localhost:8114'
    #[rpc(name = "get_peers")]
    fn get_peers(&self) -> Result<Vec<Node>>;

    // curl -d '{"id": 2, "jsonrpc": "2.0", "method":"get_banned_addresses","params": []}' -H 'content-type:application/json' 'http://localhost:8114'
    #[rpc(name = "get_banned_addresses")]
    fn get_banned_addresses(&self) -> Result<Vec<BannedAddress>>;

    // curl -d '{"id": 2, "jsonrpc": "2.0", "method":"set_ban","params": ["192.168.0.0/24", "insert"]}' -H 'content-type:application/json' 'http://localhost:8114'
    #[rpc(name = "set_ban")]
    fn set_ban(
        &self,
        address: String,
        command: String,
        ban_time: Option<Timestamp>,
        absolute: Option<bool>,
        reason: Option<String>,
    ) -> Result<()>;
}

pub trait PoolRpc {
    // curl -d '{"id": 2, "jsonrpc": "2.0", "method":"send_transaction","params": [{"version":2, "deps":[], "inputs":[], "outputs":[]}]}' -H 'content-type:application/json' 'http://localhost:8114'
    #[rpc(name = "send_transaction")]
    fn send_transaction(&self, _tx: Transaction) -> Result<H256>;

    // curl -d '{"params": [], "method": "tx_pool_info", "jsonrpc": "2.0", "id": 2}' -H 'content-type:application/json' http://localhost:8114
    #[rpc(name = "tx_pool_info")]
    fn tx_pool_info(&self) -> Result<TxPoolInfo>;
}

pub trait StatsRpc {
    #[rpc(name = "get_blockchain_info")]
    fn get_blockchain_info(&self) -> Result<ChainInfo>;

    #[rpc(name = "get_peers_state")]
    fn get_peers_state(&self) -> Result<Vec<PeerState>>;
}
