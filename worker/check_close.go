package worker

import (
	"bispy-agent/database/repository"
	"bispy-agent/database/table"
	"bispy-agent/query"
	"fmt"
	"strconv"
)

type CheckCloseInput struct {
	UserUid         string
	UserName        string
	positions       query.GetPositionResponse
	storedPositions []table.Position
}

func isPositionStillOpen(ID string, positions query.GetPositionResponse) bool {
	for _, position := range positions.Data.OtherPositionRetList {
		positionID := strconv.Itoa(int(position.UpdateTimeStamp)) + position.Symbol

		if positionID == ID {
			return true
		}
	}

	return false
}

func CheckClose(input *CheckCloseInput) {
	for _, position := range input.storedPositions {
		if isPositionStillOpen(position.ID, input.positions) == true {
			continue
		}

		fmt.Println("Closing position", position.ID, "for", input.UserName)
		repository.Position.Delete(position.ID)
	}
}
