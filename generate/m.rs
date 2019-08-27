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