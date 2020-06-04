package db

import (
    "database/sql" 
    _ "github.com/go-sql-driver/mysql"
)

var DBCli *sql.DB

func InitMysqlDB(cnf string) {
    var err error
    //返回一个连接池对象，不是单个连接
    DBCli, err = sql.Open("mysql", cnf)
    DBCli.SetMaxOpenConns(20) //设置最大打开的连接数，默认值为0表示不限制。
    DBCli.SetMaxIdleConns(10) //设置闲置的连接数。

    // DBCli.ping() 
    if err != nil {
        panic(err)
    }
}
