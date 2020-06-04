package util

import (
    "database/sql"
    //"kfd/pkg/logger"
    "kfd/pkg/util/uid"
)

var ( 
    DeviceIdUid      *uid.Uid
    BuildOrderSn     *uid.Uid //新增
)

const (
    DeviceIdBusinessId = "device_id" // 设备id
    OrderBusinessId = "order_sn" // 新增
)

func InitUID(db *sql.DB) {
    var err error
    DeviceIdUid, err = uid.NewUid(db, DeviceIdBusinessId, 5)
    if err != nil {
        //logger.Sugar.Error(err)
        panic(err)
    }

    // 新增
    BuildOrderSn, err = uid.NewUid(db, OrderBusinessId, 5)
    if err != nil {
        //logger.Sugar.Error(err)
        panic(err)
    }
}
