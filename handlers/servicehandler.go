package handler

import (
	"api/log"
	"api/models"
	"api/util"
	"fmt"
	"net/http"
	"net/url"
)

type DeviceHandller struct{}

func NewDeviceHandler() *DeviceHandller {
	return &DeviceHandller{}
}

//Create devices
func (this *DeviceHandller) CreateDevice(w http.ResponseWriter, r *http.Request) {
	log.Info("Creating Device")
	d, e := parseGetReq(r)
	if e != nil {
		ProcessError(w, r, e)
		return
	}
	if er := validateCreateDevice(d); er != nil {
		ProcessError(w, r, er)
		return
	}

	if d, err := d.CreateDevice(); err != nil {
		if d != nil {
			output := map[string]interface{}{"Error": err.Error(), "Status": d}
			ProcessResponse(w, r, output)
		} else {
			ProcessError(w, r, err)
		}
	} else {
		ProcessResponse(w, r, "Device Created Successfully..")		
	}

	return
}

//Get Devices
func (this *DeviceHandller) GetDevices(w http.ResponseWriter, r *http.Request) {
	log.Info("Getting Devices")
 
	d := models.NewDevice()
	d, e := parseGetReq(r)
	if e != nil {
		ProcessError(w, r, e)
		return
	}
	if len(d.IpAddress) > 0 {
		if device, err := d.GetDevicesByIp(d.IpAddress); err != nil {
			ProcessError(w, r, err)
		} else {
			ProcessResponse(w, r, device)
		}
	} else {
		if devices, err := d.GetDevices(); err != nil {
			ProcessError(w, r, err)
		} else {
			ProcessResponse(w, r, devices)
		}
	}

}

//Parse Get Request to extract input params.
func parseGetReq(r *http.Request) (*models.Device, error) {

	d := models.NewDevice()
	if r.Method != "GET" {
		return nil, fmt.Errorf("Only Get Request is configured right now.")
	}

	u, err := url.Parse(r.RequestURI)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse url: %s", r.RequestURI)
	}
	//parse query
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse query in url: %s", r.RequestURI)
	}
	d.Name = q.Get("name")
	d.Desc = q.Get("desc")
	d.IpAddress = q.Get("ipaddress")
	//Device Id
	if g, err := util.Guid(); err != nil {
		return nil, err
	} else {
		d.Id = g
	}
	return d, nil
}

//Validate on demand
func validateCreateDevice(d *models.Device) error {

	if d == nil {
		return fmt.Errorf("Found Null Device.")
	}
	if len(d.Name) < 1 {
		return fmt.Errorf("Please specify 'name' as input param. ")
	}
	if len(d.Desc) < 1 {
		return fmt.Errorf("Please specify 'desc' as input param. ")
	}
	if len(d.IpAddress) > 0 {
		if er := util.ValidateIp(d.IpAddress); er != nil {
			return er
		}
	} else {
		return fmt.Errorf("Please specify 'ipaddress' as input param. ")
	}
	return nil
}
