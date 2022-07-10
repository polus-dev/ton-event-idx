package client

import (
	"context"
	"ton-event-idx/internal/app"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

func NewLiteApiClient() (*ton.APIClient, error) {
	client := liteclient.NewClient()

	err := client.Connect(
		context.Background(),
		app.CFG.LITE_SERVER_HOST,
		app.CFG.LITE_SERVER_PKEY,
	)

	if err != nil {
		return nil, err
	}

	return ton.NewAPIClient(client), nil
}
