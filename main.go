package main

import (
	"fmt"
	syslog "log"
	"net/http"

	"github.com/saurabh2013/CRPCDemo/config"
	"github.com/saurabh2013/CRPCDemo/consts"
	"github.com/saurabh2013/CRPCDemo/handlers"
	"github.com/saurabh2013/CRPCDemo/log"
)

func main() {

	config.LoadConfig()
	mux := http.NewServeMux()
	//Create device Handler
	dh := handler.NewDeviceHandler()
	mux.HandleFunc(consts.CREATE_DEVICE, dh.CreateDevice)
	mux.HandleFunc(consts.DEVICE_LIST, dh.GetDevices)
	log.Info("Listening at port:", consts.PORT)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", consts.PORT), mux); err != nil {
		syslog.Printf("Listen and serve %s exits with error: %s", consts.PORT, err.Error())

	}
}
