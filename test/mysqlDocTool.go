package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"strings"
)

var TABLE_TEMPLATE = "<p>-----------------------------</p><span>表名称: {{table_name}}</span></br><span>行数: {{column_num}}</span></br><span>注释: {{table_comment}}</span><table cellspacing='0'><thead><tr><td>字段名</th><td>类型</td><td>说明</td></tr></thead>{{items}}</table></br>"
var TABLE_ELEMENT_TEMPLATE = "<tr><td>{{column_name}}<td>{{data_type}}</td><td>{{comment}}</td></th></tr>"
var USERNAME, PASSWORD, HOST, PORT, SCHEMA string

func init() {
	flag.StringVar(&USERNAME, "name", "root", "userName")
	flag.StringVar(&PASSWORD, "password", "W3Ipmm11L4pU", "password")
	flag.StringVar(&HOST, "host", "rm-2ze41lui7lm025267o.mysql.rds.aliyuncs.com", "host")
	flag.StringVar(&PORT, "port", "3306", "port")
	flag.StringVar(&SCHEMA, "schema", "db_saas_platform", "schema")

	flag.Parse()
}

// 初始化链接
func connect(userName, password, host, port, schema string) *sql.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", userName, password, host, port, schema)
	db, err := sql.Open("mysql", url)
	check(err)

	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(1)
	db.Ping()

	return db
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 读取数据库中所有的表
func readTables() (tableNames, tableComments []string) {
	db := connect(USERNAME, PASSWORD, HOST, PORT, SCHEMA)

	defer db.Close()

	rows, _ := db.Query("SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.`TABLES` WHERE table_schema = '" + SCHEMA + "' AND table_type = 'base table'")

	tableNames = make([]string, 0)
	tableComments = make([]string, 0)

	for rows.Next() {
		var tableName, tableComment string
		rows.Scan(&tableName, &tableComment)
		tableNames = append(tableNames, tableName)
		tableComments = append(tableComments, tableComment)
	}

	return
}

// 读取表的所有列
func readColumns(table string) (columnNames, columnTypes, columnComments []string) {
	db := connect(USERNAME, PASSWORD, HOST, PORT, SCHEMA)

	defer db.Close()

	rows, _ := db.Query("select COLUMN_NAME,COLUMN_TYPE,COLUMN_COMMENT from information_schema.columns where table_schema='" + SCHEMA + "' and table_name='" + table + "'")

	columnNames = make([]string, 0)
	columnTypes = make([]string, 0)
	columnComments = make([]string, 0)

	for rows.Next() {
		var columnName, columnType, columnComment string
		rows.Scan(&columnName, &columnType, &columnComment)

		columnNames = append(columnNames, columnName)
		columnTypes = append(columnTypes, columnType)
		columnComments = append(columnComments, columnComment)
	}

	return
}

func generateHtml(str string) {
	f, err := os.Create("./tables.html")
	check(err)
	defer f.Close()

	f.WriteString(str)
}

func main() {
	tableNames, tableComments := readTables()

	var tableStr string = ""
	for i := 0; i < len(tableNames); i++ {
		tableTmplate := strings.Replace(TABLE_TEMPLATE, "{{table_name}}", tableNames[i], -1)
		tableTmplate = strings.Replace(tableTmplate, "{{table_comment}}", tableComments[i], -1)

		columnNames, columnTypes, columnComments := readColumns(tableNames[i])

		tableTmplate = strings.Replace(tableTmplate, "{{column_num}}", strconv.Itoa(len(columnNames)+1), -1)

		items := ""
		for j := 0; j < len(columnNames); j++ {
			columnTmplate := strings.Replace(TABLE_ELEMENT_TEMPLATE, "{{column_name}}", columnNames[j], -1)
			columnTmplate = strings.Replace(columnTmplate, "{{data_type}}", columnTypes[j], -1)
			columnTmplate = strings.Replace(columnTmplate, "{{comment}}", columnComments[j], -1)

			items += columnTmplate
		}

		tableTmplate = strings.Replace(tableTmplate, "{{items}}", items, -1)

		tableStr += tableTmplate
	}

	fmt.Printf("共%d个表\n", len(tableNames))

	generateHtml(tableStr)
}
