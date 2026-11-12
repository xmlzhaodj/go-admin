// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"context"
	"database/sql"
	"regexp"
	"strings"
)

// CommonQuery is a common method of query.
func CommonQuery(db *sql.DB, query string, args ...interface{}) ([]map[string]interface{}, error) {

	rs, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return readQueryRows(rs)
}

var queryColumnTypePattern = regexp.MustCompile(`\\((.*)\\)`)

func readQueryRows(rs *sql.Rows) ([]map[string]interface{}, error) {
	defer rs.Close()

	col, colErr := rs.Columns()

	if colErr != nil {
		return nil, colErr
	}

	typeVal, err := rs.ColumnTypes()
	if err != nil {
		return nil, err
	}

	typeNames := make([]string, len(typeVal))
	for i, column := range typeVal {
		typeNames[i] = strings.ToUpper(queryColumnTypePattern.ReplaceAllString(column.DatabaseTypeName(), ""))
	}

	results := make([]map[string]interface{}, 0)
	for rs.Next() {
		var colVar = make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			SetColVarType(&colVar, i, typeNames[i])
		}
		result := make(map[string]interface{})
		if scanErr := rs.Scan(colVar...); scanErr != nil {
			return nil, scanErr
		}
		for j := 0; j < len(col); j++ {
			SetResultValue(&result, col[j], colVar[j], typeNames[j])
		}
		results = append(results, result)
	}
	if err := rs.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

// CommonExec is a common method of exec.
func CommonExec(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {

	rs, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

// CommonQueryWithTx is a common method of query.
func CommonQueryWithTx(tx *sql.Tx, query string, args ...interface{}) ([]map[string]interface{}, error) {

	rs, err := tx.Query(query, args...)

	if err != nil {
		panic(err)
	}

	return readQueryRows(rs)
}

// CommonExecWithTx is a common method of exec.
func CommonExecWithTx(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	rs, err := tx.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

// CommonBeginTxWithLevel starts a transaction with given transaction isolation level and db connection.
func CommonBeginTxWithLevel(db *sql.DB, level sql.IsolationLevel) *sql.Tx {
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: level})
	if err != nil {
		panic(err)
	}
	return tx
}
