package scan

import (
	"github.com/sirupsen/logrus"
	"github.com/xssnick/tonutils-go/liteclient/tlb"
)

func processMcBlock(block *tlb.BlockInfo) {
	logrus.Infof("new masterchain block: %d", block.SeqNo)
}
