package service

import (
	"message/dao"
	"message/entity"
)

func AddMessage(message entity.MseeageEntity) {

	dao.AddMessage(message)
}
