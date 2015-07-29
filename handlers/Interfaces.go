package handler

import "net/http"

type Entity interface {
	CreateDevice(w http.ResponseWriter, r *http.Request)
	GetDevices(w http.ResponseWriter, r *http.Request)
}

func NewEntity() (entity Entity, e error) {
	entity = &DeviceHandller{}
	return
}
