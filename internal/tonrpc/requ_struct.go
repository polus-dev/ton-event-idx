package tonrpc

type GetBlockTransactions struct {
	Workchain int `json:"workchain"`
	Shard     int `json:"shard"`
	Seqno     int `json:"seqno"`
}

type GetTransactions struct {
	Address  string `json:"address"`
	Limit    int    `json:"limit,omitempty"`
	LT       int    `json:"lt,omitempty"`
	Hash     string `json:"hash,omitempty"`
	ToLT     int    `json:"to_lt,omitempty"`
	Archival bool   `json:"archival,omitempty"`
}

type RpcParams interface {
	GetBlockTransactions | GetTransactions
}

type RpcReq[P RpcParams] struct {
	Method  string `json:"method"`
	JsonRpc string `json:"jsonrpc"`
	Params  P      `json:"params"`
}

func newRpcReq[P RpcParams](method string, params P) RpcReq[P] {
	return RpcReq[P]{Method: method, JsonRpc: "2.0", Params: params}
}
