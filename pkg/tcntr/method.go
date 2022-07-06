package tcntr

import (
	"encoding/json"
)

type TonCenterSerializer struct {
	JsonRpcURL string
}

func serializeReq[P RpcParams](method string, params *P) ([]byte, error) {
	data, err := json.Marshal(newRpcReq(method, params))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t TonCenterSerializer) GetBlockTransactions(params *GetBlockTransactions) ([]byte, error) {
	return serializeReq("getBlockTransactions", params)
}

func (t TonCenterSerializer) GetTransactions(params *GetTransactions) ([]byte, error) {
	return serializeReq("getTransactions", params)
}
