package mcblock

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"ton-event-idx/pkg/psql"
	"ton-event-idx/pkg/utils/crypt"

	"github.com/jackc/pgconn"
)

type repo struct {
	client psql.Client
}

func (r *repo) Create(ctx context.Context, mc *McBlock) error {
	q := `
		INSERT INTO mc_block 	(id, work_chain, shard, seqno, root_hash, file_hash)
		VALUES 					($1::bytea, $2, $3::bytea, $4, $5::bytea, $6::bytea)
		RETURNING id
	`

	shardBs := make([]byte, 8)
	binary.LittleEndian.PutUint64(shardBs, mc.Block.Shard)

	if err := r.client.QueryRow(ctx, q,
		crypt.NewBlockID(mc.Block),
		mc.Block.Workchain,
		shardBs,
		mc.Block.SeqNo,
		mc.Block.RootHash,
		mc.Block.FileHash,
	).Scan(&mc.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(
				"SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState(),
			)
			return newErr
		}
		return err
	}

	return nil
}

func NewRepo(client psql.Client) Repo {
	return &repo{
		client: client,
	}
}
