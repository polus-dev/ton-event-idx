package tonrpc

type TonType struct {
	Type string `json:"@type"`
}

type TonBlockIdExt struct {
	TonType          // "ton.blockIdExt"
	Workchain int    `json:"workchain"`
	Shard     string `json:"shard"`
	Seqno     int    `json:"seqno"`
	RootHash  string `json:"root_hash"`
	FileHash  string `json:"file_hash"`
}

type BlocksShortTxId struct {
	TonType        // "blocks.shortTxId
	Mode    int    `json:"mode"`
	Account string `json:"account"`
	LT      string `json:"lt"`
	Hash    string `json:"hash"`
}

type MsgDataRaw struct {
	TonType          // "msg.dataRaw"
	Body      string `json:"body"`
	InitState string `json:"init_state"`
}

type RawMessage struct {
	TonType                // "raw.message"
	Source      string     `json:"source"`
	Destination string     `json:"destination"`
	Value       string     `json:"value"`
	FwdFee      string     `json:"fwd_fee"`
	IhrFee      string     `json:"ihr_fee"`
	CreatedLT   string     `json:"created_lt"`
	BodyHash    string     `json:"body_hash"`
	Message     string     `json:"message"`
	MsgData     MsgDataRaw `json:"msg_data"`
}

type InternalTransactionId struct {
	TonType        // "internal.transactionId"
	LT      string `json:"lt"`
	Hash    string `json:"hash"`
}

type RawTransaction struct {
	TonType                             // "raw.transaction"
	Utime         int                   `json:"utime"`
	Data          string                `json:"data"`
	Fee           string                `json:"fee"`
	StorageFee    string                `json:"storage_fee"`
	OtherFee      string                `json:"other_fee"`
	TransactionId InternalTransactionId `json:"transaction_id"`
	InMsg         RawMessage            `json:"in_msg"`
	OutMsgs       []RawMessage          `json:"out_msgs"`
}

type BlockTransactions struct {
	Type         string            `json:"@type"` // "blocks.transactions"
	Id           TonBlockIdExt     `json:"id"`
	ReqCount     int               `json:"req_count"`
	Incomplete   bool              `json:"incomplete"`
	Transactions []BlocksShortTxId `json:"transactions"`
	Extra        string            `json:"@extra"`
}

type RpcResult interface {
	BlockTransactions | RawTransaction
}

type BasicResp[R RpcResult] struct {
	OK      bool   `json:"ok"`
	JsonRpc string `json:"jsonrpc"`
	Result  []R    `json:"result"`
}
