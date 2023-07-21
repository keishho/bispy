package table

const PositionTableSchema = `
CREATE TABLE if not exists position (
	id VARCHAR PRIMARY KEY,
	traderID VARCHAR NOT NULL,
	traderName VARCHAR NOT NULL,
	entryPrice FLOAT NOT NULL,
	markPrice FLOAT NOT NULL,
	pnl FLOAT NOT NULL,
	roe FLOAT NOT NULL,
	amount FLOAT NOT NULL
)`

type Position struct {
	ID         string  `db:"id"`
	TraderID   string  `db:"traderID"`
	TraderName string  `db:"traderName"`
	EntryPrice float64 `db:"entryPrice"`
	MarkPrice  float64 `db:"markPrice"`
	Pnl        float64 `db:"pnl"`
	Roe        float64 `db:"roe"`
	Amount     float64 `db:"amount"`
}
