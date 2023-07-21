package worker

import (
	"bispy-agent/database"
	"bispy-agent/database/table"
	"bispy-agent/query"
	"time"
)

func Supply(tick *time.Time) {
	for _, target := range Targets {
		go operate(target)
	}
}

func operate(target Target) {
	positions := query.GetPosition(&query.GetPositionRequest{
		EncryptedUid: target.uid,
		TradeType:    "PERPETUAL",
	})

	storedPositions := []table.Position{}
	database.DB.Select(&storedPositions, "SELECT * FROM position WHERE traderID = $1", target.uid)

	CheckOpen(&CheckOpenInput{
		UserUid:         target.uid,
		UserName:        target.name,
		positions:       positions,
		storedPositions: storedPositions,
	})

	CheckClose(&CheckCloseInput{
		UserUid:         target.uid,
		UserName:        target.name,
		positions:       positions,
		storedPositions: storedPositions,
	})
}
