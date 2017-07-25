package client

type MiningInfo struct {
	Status      string `json:"status"`
	Height      uint64 `json:"height,string,omitempty""`
	Rate        int    `json:"rate,string,omitempty""`
	Difficultly uint64 `json:"difficulty,string,omitempty""`
}

type Block struct {
	Header *BlockHeader
	Txs    []*Transaction
}

type BlockHeader struct {
	Bits              string `json:"bits"`
	Hash              string `json:"hash"`
	MerkleTreeHash    string `json:"merkle_tree_hash"`
	Nonce             uint64 `json:"nonce,string,omitempty"`
	PreviousBlockHash string `json:"previous_block_hash"`
	TimeStamp         uint64 `json:"time_stamp,string,omitempty"`
	Version           int    `json:"version,string,omitempty"`
	MixHash           string `json:"mixhash"`
	Number            uint64 `json:"number,string,omitempty"`
	TxCount           uint64 `json:"transaction_count,string,omitempty"`
}

type AccountInfo struct {
	Name           string   `json:"name"`
	Mnemonic       string   `json:"mnemonic"`
	HdIndex        int64    `json:"hd_index,string,omitempty"`
	Addresses      []string `json:"addresses,string,omitempty"`
	DefaultAddress string   `json:"default-address"`
	AddressCnt     int64    `json:"address-count,string,omitempty"`
	UserStatus     int      `json:"user-status,string,omitempty"`
}

type BalanceStatistic struct {
	TotalConfirmed int64 `json:"total-confirmed,string,omitempty"`
	TotalReceived  int64 `json:"total-received,string,omitempty"`
	TotalUnspent   int64 `json:"total-unspent,string,omitempty"`
	TotalAvailable int64 `json:"total-available,string,omitempty"`
	TotalFrozen    int64 `json:"total-frozen,string,omitempty"`
}

type BalanceSet struct {
	Balance *Balance `json:"balance"`
}

type Balance struct {
	Address   string `json:"address"`
	Confirmed int64  `json:"confirmed,string,omitempty"`
	Received  int64  `json:"received,string,omitempty"`
	Unspent   int64  `json:"unspent,string,omitempty"`
	Available int64  `json:"available,string,omitempty"`
	Fronzen   int64  `json:"fronzen,string,omitempty"`
}

type Transaction struct {
	Hash      string    `json:"hash"`
	Height    uint64    `json:"height,string,omitempty"`
	Timestamp uint64    `json:"timestamp,string,omitempty"`
	Direction string    `json:"direction"`
	Inputs    []*Input  `json:"inputs"`
	Outputs   []*Output `json:"outputs"`
	LockTime  uint64    `json:"lock_time,string,omitempty"`
	Version   string    `json:"version"`
}

type Input struct {
	Address   string  `json:"address"`
	PreOutput *Output `json:"previous_output"`
	Sequence  uint64  `json:"sequence,string,omitempty"`
}

type Output struct {
	Hash       string      `json:"hash"`
	Index      uint64      `json:"index,string,omitempty"`
	Script     string      `json:"script"`
	Value      uint64      `json:"value,string,omitempty"`
	Own        string      `json:"own"`
	Address    string      `json:"address"`
	EtpValue   uint64      `json:"etp-value,string,omitempty"`
	Attachment *Attachment `json:"attachment"`
}

type Attachment struct {
	Type string `json:"type"`
}

type UTXO struct {
	Points []*Point `json:"points"`
	Change uint64   `json:"change,string,omitempty"`
}

type Point struct {
	Hash  string `json:"hash"`
	Index uint64 `json:"index,string,omitempty"`
}
