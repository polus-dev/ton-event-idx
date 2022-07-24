package mcblock

import "encoding/binary"

type MCBlockModel struct {
	ID    []byte
	Shard []byte
	SeqNo uint32
}

type MCBlockDTO struct {
	Shard uint64
	SeqNo uint32
}

func (mcm *MCBlockModel) ToDTO() *MCBlockDTO {
	return &MCBlockDTO{
		Shard: binary.LittleEndian.Uint64(mcm.Shard),
		SeqNo: mcm.SeqNo,
	}
}

func (mcdto *MCBlockDTO) ToModel() *MCBlockModel {
	shardBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(shardBytes, mcdto.Shard)

	return &MCBlockModel{
		Shard: shardBytes,
		SeqNo: mcdto.SeqNo,
	}
}
