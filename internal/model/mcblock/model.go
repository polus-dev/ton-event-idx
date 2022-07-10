package mcblock

import "github.com/xssnick/tonutils-go/liteclient/tlb"

type McBlock struct {
	ID    []byte // id
	Block *tlb.BlockInfo
}
