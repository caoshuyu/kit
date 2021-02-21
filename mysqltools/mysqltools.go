package mysqltools

import (
	"database/sql"
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
