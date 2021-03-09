package mysqltools

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/caoshuyu/kit/dlog"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MysqlClient struct {
	Client *sql.DB
	Conf   *MySqlConf
}

type MySqlConf struct {
	DbDsn       string        `json:"db_dsn"`
	MaxOpen     int           `json:"max_open"`
	MaxIdle     int           `json:"max_idle"`
	DbName      string        `json:"db_name"`
	MaxLifetime time.Duration `json:"max_lifetime"`
}

//connect mysql server
func (mc *MysqlClient) Connect() error {
	db, err := sql.Open("mysql", mc.Conf.DbDsn)
	if nil != err {
		return err
	}
	if 0 != mc.Conf.MaxOpen {
		db.SetMaxOpenConns(mc.Conf.MaxOpen)
	}
	if 0 != mc.Conf.MaxIdle {
		db.SetMaxIdleConns(mc.Conf.MaxIdle)
	}
	if 0 == mc.Conf.MaxLifetime {
		mc.Conf.MaxLifetime = time.Second * time.Duration(300)
	}
	db.SetConnMaxLifetime(mc.Conf.MaxLifetime)
	err = db.Ping()
	if nil != err {
		return err
	}
	mc.Client = db
	return nil
}

//check mysql connect , eq connect number and max connect number
func (mc *MysqlClient) CheckMonitor() error {
	openConn := mc.Client.Stats().OpenConnections
	if 0 == openConn {
		err := errors.New("sql: database is closed")
		dlog.ERROR("_db_OpenConnections", "db not open",
			"dbDsn", mc.Conf.DbDsn)
		return err
	}
	if 5*openConn > 4*mc.Conf.MaxOpen {
		dlog.WARN("_db_OpenConnections", "open conn is more", "open conn", openConn, "max open conn", mc.Conf.MaxOpen)
	} else {
		dlog.INFO("_db_OpenConnections", "conn is ok", "open conn", openConn, "max open conn", mc.Conf.MaxOpen)
	}
	return nil
}

func WriteDbError(tableName string, err error, sql string, addParams ...interface{}) {
	var infoList []interface{}
	for k, _ := range addParams {
		infoList = append(infoList, addParams[k])
	}
	type ErrorType struct {
		ErrorParams []interface{}
	}
	paramsJson, _ := json.Marshal(ErrorType{
		ErrorParams: infoList,
	})
	dlog.ERROR("msg", "dbError", "tableName", tableName, "error", err.Error(), "sql", sql, "params", string(paramsJson))
}

func UpdateTableValue(db *sql.DB, sqlText string, params []interface{}) (number int64, err error) {
	stmt, err := db.Prepare(sqlText)
	if nil != err {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		params...,
	)
	number, err = res.RowsAffected()
	if nil != err {
		return 0, err
	}
	return number, nil
}