package poker

import (
	"database/sql"
	"fmt"
)

var updateStmt = `UPDATE players SET score=$1 WHERE name=$2`
var deleteStmt = `DELETE FROM players WHERE name=$1`
var selectStmt = `SELECT name, score from players where name=$1`

type PostgresPlayerStore struct {
	store PlayerStore
	db    *sql.DB
}

func (p *PostgresPlayerStore) RecordWin(name string) {
	rows, err := p.db.Query(selectStmt, name)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string
		var score int

		err = rows.Scan(&name, &score)
		CheckError(err)

		fmt.Print(name, score)
	}

}

func (p *PostgresPlayerStore) GetPlayerScore(name string) (score int) {
	return 3
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
