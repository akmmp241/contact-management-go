package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		if errorRollback != nil {
			panic(err)
		}
		panic(err)
	}
	errorCommit := tx.Commit()
	if errorCommit != nil {
		panic(err)
	}
}
