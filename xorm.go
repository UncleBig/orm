package orm

import (
	"errors"
	"fmt"
	"github.com/dlintw/goconf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Xorm *xorm.Engine
)

func Init(conf *goconf.ConfigFile) (err error) {
	var db_def string
	var maxidle, _ = conf.GetInt("def_db", "db_maxidle")
	var maxconn, _ = conf.GetInt("def_db", "db_maxconn")
	var dbHost, _ = conf.GetString("def_db", "db_host")
	var dbPort, _ = conf.GetInt("def_db", "db_port")
	var dbUser, _ = conf.GetString("def_db", "db_user")
	//	var dbSocket, _ = conf.GetString("def_db", "db_socket")
	var dbName, _ = conf.GetString("def_db", "db_name")
	var dbCharset, _ = conf.GetString("def_db", "db_charset")
	var dbPass, _ = conf.GetString("def_db", "db_pass")

	if maxidle < 1 {
		maxidle = 5
	}

	if maxconn < 1 {
		maxconn = 5
	}

	//	if dbHost == "localhost" || dbHost == "127.0.0.1" {
	//		db_def = fmt.Sprintf("%s:%s@unix(%s)/%s?charset=%s", dbHost, dbUser, dbSocket, dbName, dbCharset)
	//	} else {
	db_def = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", dbUser, dbPass, dbHost, dbPort, dbName,
		dbCharset)
	//	}
	Xorm, err = xorm.NewEngine("mysql", db_def)
	if err != nil {
		err = errors.New(fmt.Sprintf("xorm.NewEngine(mysql, %s), error: %s", db_def, err))
		return
	}

	if maxidle > 0 {
		Xorm.SetMaxIdleConns(maxidle)
	}
	if maxconn > 0 {
		Xorm.SetMaxOpenConns(maxconn)
	}
	return
}
