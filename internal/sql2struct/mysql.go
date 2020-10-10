package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

var DBTypeToStructType = map[string]string{
	"tinyint":           "int8",
	"smallint":          "int16",
	"integer":           "int32",
	"int":               "int32",
	"bigint":            "int64",
	"integer unsigned":  "uint",
	"tinyint unsigned":  "uint8",
	"smallint unsigned": "uint16",
	"bigint unsigned":   "uint64",
	"double precision":  "float32",
	"float":             "float32",
	"bool":              "bool",
	"text":              "string",
	"longtext":          "string",
	"mediumtext":        "string",
	"varchar":           "string",
	"char":              "string",
	"enum":              "string",
	"set":               "string",
	"date":              "time.Time",
	"datetime":          "time.Time",
	"timestamp":         "time.Time",
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{
		DBEngine: nil,
		DBInfo:   info,
	}
}

func (m *DBModel) Connect() error {
	var err error
	format := "%s:%s@tcp(%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		format,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)

	if m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn); err != nil {
		return err
	}

	return nil
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	querySql := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, " +
		"COLUMN_TYPE, COLUMN_COMMENT FROM COLUMNS where TABLE_SCHEMA = ? " +
		"AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(querySql, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()
	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(
			&column.ColumnName,
			&column.DataType,
			&column.ColumnKey,
			&column.IsNullable,
			&column.ColumnType,
			&column.ColumnComment,
		)

		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}
