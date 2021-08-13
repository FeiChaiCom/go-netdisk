package mysqlclient

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type MySql struct {
	host       string
	port       int
	usr        string
	pwd        string
	database   string
	driverName string
	db         *sql.DB
}

func NewMySql() (*MySql, error) {
	mysql := new(MySql)
	mysql.db = nil

	return mysql, nil
}

func (m *MySql) Open(host, usr, pwd, database string, port, maxOpenConns, maxIdleConns int) error {
	m.host = host
	m.port = port
	m.usr = usr
	m.pwd = pwd
	m.database = database

	//driver: usr:pwd@tcp(host:port)/database
	m.driverName = m.usr + ":" + m.pwd + "@tcp(" + m.host + ":" + strconv.Itoa(m.port) + ")/" + m.database

	db, err := sql.Open("mysql", m.driverName)
	if err != nil {
		return err
	}

	m.db = db

	if maxOpenConns < 0 {
		maxOpenConns = 0
	}

	if maxIdleConns < 0 {
		maxIdleConns = 0
	}

	m.db.SetMaxOpenConns(maxOpenConns)
	m.db.SetMaxIdleConns(maxIdleConns)

	return m.db.Ping()
}

func (m *MySql) Close() {
	if m.db != nil {
		m.db.Close()
	}
}

func (m *MySql) Query(sql string) (map[string]map[string]string, error) {
	rows, err := m.db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cols, _ := rows.Columns()
	valCols := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range valCols {
		scans[i] = &valCols[i]
	}

	results := make(map[string]map[string]string)
	i := 0
	for rows.Next() {
		if err := rows.Scan(scans...); err != nil {
			return nil, err
		}

		row := make(map[string]string)
		for k, v := range valCols {
			key := cols[k]
			row[key] = string(v)
		}

		results[strconv.Itoa(i)] = row
		i++
	}

	return results, nil
}

func (m *MySql) Insert(sql string, data []interface{}) (int64, error) {
	stmt, err := m.db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(data...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MySql) Update(sql string, data []interface{}) (int64, error) {
	stmt, err := m.db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(data...)
	if err != nil {
		return 0, err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return num, nil
}

func (m *MySql) Remove(sql string, data []interface{}) (int64, error) {
	stmt, err := m.db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(data...)
	if err != nil {
		return 0, err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return num, nil
}
