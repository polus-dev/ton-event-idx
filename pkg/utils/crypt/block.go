package crypt

import (
	"encoding/binary"

	"github.com/xssnick/tonutils-go/liteclient/tlb"
)

func NewBlockID(b *tlb.BlockInfo) []byte {
	total := make([]byte, 16)

	binary.LittleEndian.PutUint32(total[0:4], uint32(b.Workchain))
	binary.LittleEndian.PutUint64(total[4:12], b.Shard)
	binary.LittleEndian.PutUint32(total[12:16], b.SeqNo)

	return total
}
