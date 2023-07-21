package worker

import (
	"bispy-agent/database/repository"
	"bispy-agent/database/table"
	"bispy-agent/query"
	"fmt"
	"strconv"
)

type CheckOpenInput struct {
	UserUid         string
	UserName        string
	positions       query.GetPositionResponse
	storedPositions []table.Position
}

func CheckOpen(input *CheckOpenInput) {
	for _, position := range input.positions.Data.OtherPositionRetList {
		positionID := strconv.Itoa(int(position.UpdateTimeStamp)) + position.Symbol

		alreadyStored := repository.Position.IsIDInPositions(positionID, input.storedPositions)
		if alreadyStored == true {
			repository.Position.UpdateLiveValues(&repository.UpdateLiveValuesInput{
				ID:        positionID,
				MarkPrice: position.MarkPrice,
				Pnl:       position.Pnl,
				Roe:       position.Roe,
			})
			fmt.Println("Position already stored, updating live values...")
			continue
		}

		fmt.Println("Inserting new position", position.Symbol, "for", input.UserName)
		repository.Position.Insert(&table.Position{
			ID:         positionID,
			TraderID:   input.UserUid,
			TraderName: input.UserName,
			EntryPrice: position.EntryPrice,
			MarkPrice:  position.MarkPrice,
			Pnl:        position.Pnl,
			Roe:        position.Roe,
			Amount:     position.Amount,
		})

	}

}
