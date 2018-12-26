package controller

import (
	"github.com/cihub/seelog"
	"message/entity"
	"message/service"
	"net/http"
)

func AddMessage(w http.ResponseWriter, r *http.Request) {



	seelog.Debug("==================")


	username := r.PostFormValue("username")
	message := r.PostFormValue("message")

	//mseeageEntity := new(entity.MseeageEntity)
	mseeageEntity := entity.MseeageEntity{UserName: username, Mseeage: message}

	service.AddMessage(mseeageEntity)

	w.Write([]byte("success"))
}
