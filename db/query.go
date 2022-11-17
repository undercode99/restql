package db

import (
	"database/sql"
	"reflect"
)

type Query struct {
	Db          *sql.DB
	Sql         string
	BindData    []interface{}
	ColumnTypes []*sql.ColumnType
	ResultRows  [][]interface{}
}

func NewQuery(db *sql.DB, sql string) *Query {
	return &Query{
		Db:  db,
		Sql: sql,
	}
}

func (q *Query) Do() error {
	rows, err := q.Db.Query(q.Sql, q.BindData...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// get column type info
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	// get column names

	// used for allocation & dereferencing
	rowValues := make([]reflect.Value, len(columnTypes))
	for i := 0; i < len(columnTypes); i++ {
		// allocate reflect.Value representing a **T value
		rowValues[i] = reflect.New(reflect.PtrTo(columnTypes[i].ScanType()))
	}

	resultList := [][]interface{}{}
	for rows.Next() {
		// initially will hold pointers for Scan, after scanning the
		// pointers will be dereferenced so that the slice holds actual values
		rowResult := make([]interface{}, len(columnTypes))
		for i := 0; i < len(columnTypes); i++ {
			// get the **T value from the reflect.Value
			rowResult[i] = rowValues[i].Interface()
		}

		// scan each column value into the corresponding **T value
		if err := rows.Scan(rowResult...); err != nil {
			return err
		}

		// dereference pointers
		for i := 0; i < len(rowValues); i++ {
			// first pointer deref to get reflect.Value representing a *T value,
			// if rv.IsNil it means column value was NULL
			if rv := rowValues[i].Elem(); rv.IsNil() {
				rowResult[i] = nil
			} else {
				// second deref to get reflect.Value representing the T value
				// and call Interface to get T value from the reflect.Value
				rowResult[i] = rv.Elem().Interface()
			}
		}

		resultList = append(resultList, rowResult)

	}
	if err := rows.Err(); err != nil {
		return err
	}

	q.ColumnTypes = columnTypes
	q.ResultRows = resultList

	return nil

}

func (q *Query) GetColumnNames() []string {
	columnNames := []string{}
	for _, columnType := range q.ColumnTypes {
		columnNames = append(columnNames, columnType.Name())
	}
	return columnNames
}

func (q *Query) GetColumnTypes() []string {
	columnTypes := []string{}
	for _, columnType := range q.ColumnTypes {
		columnTypes = append(columnTypes, columnType.DatabaseTypeName())
	}
	return columnTypes
}

func (q *Query) GetMapColumnNameAndType() map[string]string {
	columnMap := map[string]string{}
	for _, columnType := range q.ColumnTypes {
		columnMap[columnType.Name()] = columnType.DatabaseTypeName()
	}
	return columnMap
}

func (q *Query) GetColumnNamesAndTypes() []string {
	columnNamesAndTypes := []string{}
	for _, columnType := range q.ColumnTypes {
		columnNamesAndTypes = append(columnNamesAndTypes, columnType.Name()+" ("+columnType.DatabaseTypeName()+")")
	}
	return columnNamesAndTypes
}

func (q *Query) GetResultRows() [][]interface{} {
	return q.ResultRows
}

func (q *Query) GetResultMapRows() []map[string]interface{} {
	resultMap := []map[string]interface{}{}
	for _, row := range q.ResultRows {
		rowMap := make(map[string]interface{})
		for i, columnType := range q.ColumnTypes {
			rowMap[columnType.Name()] = row[i]
		}
		resultMap = append(resultMap, rowMap)
	}
	return resultMap
}
