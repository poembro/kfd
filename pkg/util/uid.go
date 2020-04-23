package util

import (
    "database/sql"
    //"kfd/pkg/logger"
    "kfd/pkg/util/uid"
)

var ( 
    DeviceIdUid      *uid.Uid
)

const (
    DeviceIdBusinessId = "device_id" // 设备id
)

func InitUID(db *sql.DB) {
    var err error
    DeviceIdUid, err = uid.NewUid(db, DeviceIdBusinessId, 5)
    if err != nil {
        //logger.Sugar.Error(err)
        panic(err)
    }
}
