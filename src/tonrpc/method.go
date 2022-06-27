package tonrpc

import (
	"encoding/json"
	"ton-event-idx/src/requtil"
)

type TonRPC struct{ JsonRpcURL string }

func processReq[P RpcParams](url string, method string, params P) ([]byte, error) {
	data, err := json.Marshal(newRpcReq(method, params))
	if err != nil {
		return nil, err
	}

	return requtil.SendPostReq(url, data)
}

func (t TonRPC) GetBlockTransactions(params GetBlockTransactions) (*BasicResp[BlockTransactions], error) {
	resp, err := processReq(t.JsonRpcURL, "getBlockTransactions", params)
	if err != nil {
		return nil, err
	}

	var data BasicResp[BlockTransactions]
	if err := json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (t TonRPC) GetTransactions(params GetTransactions) (*BasicResp[RawTransaction], error) {
	resp, err := processReq(t.JsonRpcURL, "getTransactions", params)
	if err != nil {
		return nil, err
	}
	var data BasicResp[RawTransaction]
	if err := json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
