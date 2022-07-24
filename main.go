/*
	ton-event-idx – TON Blockchain smart contracts event indexing service

	Copyright (C) 2022 https://github.com/cryshado

	This file is part of ton-event-idx.

	ton-event-idx is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	ton-event-idx is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with ton-event-idx.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"ton-event-idx/internal/app"
	"ton-event-idx/internal/scan"
	"ton-event-idx/pkg/client/liteapi"
	_ "ton-event-idx/pkg/log"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting the \"ton-event-idx\"")
	app.Configure()

	api, err := liteapi.NewLiteApiClient()
	if err != nil {
		logrus.Fatalf("can't create new lite apo client: %s", err)
	}

	if err := scan.StartScanMasterChain(api); err != nil {
		logrus.Fatal(err)
	}
}
