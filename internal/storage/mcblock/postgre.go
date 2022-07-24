package mcblock

import (
	"context"
	"ton-event-idx/pkg/client/psql"
)

type repo struct {
	client psql.Client
}

func (r *repo) SelectLast(ctx context.Context) (*MCBlockModel, error) {
	q := "SELECT shard, seqno FROM mc_block ORDER BY id DESC LIMIT 1"

	var mc MCBlockModel

	err := r.client.QueryRow(ctx, q).Scan(
		&mc.Shard, &mc.SeqNo,
	)

	return &mc, err
}

func (r *repo) IsEmpty(ctx context.Context) (bool, error) {
	q := "SELECT CASE WHEN EXISTS(SELECT 1 FROM mc_block) THEN false ELSE true END"

	var status bool
	if err := r.client.QueryRow(ctx, q).Scan(&status); err != nil {
		return status, err
	}

	return status, nil
}

type Repo interface {
	SelectLast(ctx context.Context) (*MCBlockModel, error)
	IsEmpty(ctx context.Context) (bool, error)
}

func NewRepo(client psql.Client) Repo {
	return &repo{client: client}
}
