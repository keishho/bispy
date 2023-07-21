package repository

import (
	"bispy-agent/database"
	"bispy-agent/database/table"
)

type position struct{}

var Position position

func (position) IsIDInPositions(ID string, positions []table.Position) bool {
	for _, position := range positions {
		if position.ID == ID {
			return true
		}
	}

	return false
}

type UpdateLiveValuesInput struct {
	ID        string  `db:"id"`
	MarkPrice float64 `db:"markPrice"`
	Pnl       float64 `db:"pnl"`
	Roe       float64 `db:"roe"`
}

func (position) Insert(newPosition *table.Position) {
	tx := database.DB.MustBegin()

	tx.NamedExec(`
	INSERT INTO position (
		id, 
		traderID, 
		traderName, 
		entryPrice, 
		markPrice, 
		pnl, 
		roe, 
		amount
	) 
	VALUES (
		:id, 
		:traderID, 
		:traderName, 
		:entryPrice, 
		:markPrice, 
		:pnl, 
		:roe, 
		:amount
	)`,
		newPosition)

	tx.Commit()
}

func (position) Delete(id string) {
	tx := database.DB.MustBegin()

	tx.MustExec(`DELETE FROM position WHERE id = ?`, id)

	tx.Commit()
}

func (position) UpdateLiveValues(input *UpdateLiveValuesInput) {
	tx := database.DB.MustBegin()

	tx.NamedExec(`
	UPDATE position
	SET
		markPrice = :markPrice,
		pnl = :pnl,
		roe = :roe
	WHERE id = :id`, input)

	tx.Commit()
}
