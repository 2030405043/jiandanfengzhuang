package dao

import (
	"github.com/cihub/seelog"
	"message/DBconf"
	"message/entity"
)

func AddMessage(entity entity.MseeageEntity) {
	DB := DBconf.GetDB()
	_, err := DB.Exec("INSERT INTO message(username, message) VALUES(?, ?);", entity.UserName, entity.Mseeage)

	if err != nil {
		seelog.Error("INSERT INTO message", err.Error())
		return
	}

//	lastId, err := res.LastInsertId()
//	fmt.Println(lastId)
}
