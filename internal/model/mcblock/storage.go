package mcblock

import "context"

type Repo interface {
	Create(ctx context.Context, block *McBlock) error
}
