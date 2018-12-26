package main

import (
	"fmt"
	"github.com/cihub/seelog"
	_ "message/route"
	"net/http"
)

func main() {
	fmt.Println("Hello,Message...")

	logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	e := http.ListenAndServe(":8000", nil)
	if e != nil {
		seelog.Error("server err,", e.Error())
		return
	}

}
