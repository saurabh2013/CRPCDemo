package models

import (
	"fmt"
	"strings"

	"github.com/saurabh2013/CRPCDemo/log"
)

var (
	Devices              []Device                       //according to putput schema
	GlobalDeviceListByIp = make(map[string]interface{}) //for internal use
)

type Device struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	IpAddress string `json:"ipaddress"`
}

func NewDevice() *Device {
	return &Device{}

}

//Function to add devices.
func (this *Device) CreateDevice() (map[string]interface{}, error) {

	//Validate Device with existing devices address
	//Create device code here.
	if device, er := validateDevice(this); er != nil {
		return device, er
	}
	if len(strings.TrimSpace(this.Id)) < 1 {
		return nil, fmt.Errorf("Found empty device id.")
	}
	d := Device{Id: this.Id, Name: this.Name, Desc: this.Desc, IpAddress: this.IpAddress}
	Devices = append(Devices, d)
	log.LogTofilef("Device Created Successfully... Devices: %v", d)
	return nil, nil

}

//This funcction will check few validation like.
//	Is Input Ip Address allocation to other device.
//  Get that Device and send Back to User
//
func validateDevice(d *Device) (map[string]interface{}, error) {
	if d == nil {
		return nil, fmt.Errorf("Found Emty device to validate.")
	}

	if len(GlobalDeviceListByIp) > 0 {
		if _device, k := GlobalDeviceListByIp[d.IpAddress]; k {
			return map[string]interface{}{"Existing_Device_With_This_IP": _device},
				fmt.Errorf("Provided Ip address is already allocatied to another device.")
		}
	}
	//You are here that means you did not find any device id with this ip
	addToGlobalList(d)
	return nil, nil
}

//Add Deviceto global list
//Device is new it adds device to global device list, Id by ip
func addToGlobalList(d *Device) {
	GlobalDeviceListByIp[d.IpAddress] = d
}

//Get the device list from the local stoted deivice list.
func (this *Device) GetDevices() (*[]Device, error) {
	if len(Devices) < 1 {
		return nil, fmt.Errorf("Currently no device is available to show. pleasse add some.")
	}
	return &Devices, nil
}

//Get the device information for provided ip address.
func (this *Device) GetDevicesByIp(ip string) (interface{}, error) {
	if len(Devices) < 1 {
		return nil, fmt.Errorf("Currently no device is available to show. pleasse add some.")
	}
	if _d, k := GlobalDeviceListByIp[strings.TrimSpace(ip)]; k {
		return _d, nil
	}
	return &Devices, fmt.Errorf("No Device found with this ip '%s'", ip)

}
