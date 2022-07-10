package scan

import (
	"context"
	"ton-event-idx/internal/app"
	"ton-event-idx/internal/model/mcblock"

	"github.com/sirupsen/logrus"
	"github.com/xssnick/tonutils-go/liteclient/tlb"
)

func processMcBlock(block *tlb.BlockInfo) {
	logrus.Infof("new masterchain block: %d", block.SeqNo)

	repo := mcblock.NewRepo(app.DBCONN)

	dbblock := &mcblock.McBlock{
		ID:    block.FileHash,
		Block: block,
	}

	err := repo.Create(context.Background(), dbblock)
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Info("MC ID FROM DB: ", dbblock)
}
