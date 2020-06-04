package db

import (
    "kfd/config"
)

func InitDB() {
    rdscnf := config.LogicConf.RedisIP
    InitRdsDB(rdscnf)

    mysqlcnf := config.LogicConf.MySQL
    InitMysqlDB(mysqlcnf)
}