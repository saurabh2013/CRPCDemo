package util

import (
	"api/config"
	"api/log"
	"fmt"
	"net"
	"strings"

	"github.com/nu7hatch/gouuid"
)

  

//Genarate New Guid id as unique identifier.
func Guid() (string, error) {
	if u4, err := uuid.NewV4(); err == nil {
		return u4.String(), nil
	} else {
		log.Error(err, nil, "Error while GUID creation") //no context available
		return "", err
	}
}

//Validate Ip Address.
//	*Should be a valid ip address
//	*Should be whith in pre apecified range.
func ValidateIp(ip string) error {

	if len(strings.TrimSpace(ip)) < 1 {
		return fmt.Errorf("Found Empty Ip address.")
	}

	ipaddr := net.ParseIP(ip)
	if ipaddr.To4() == nil {
		return fmt.Errorf("'%v' is not an valid IPv4 address", ip)
	}
	_, cidrnet, err := net.ParseCIDR(config.CIDR)
	if err != nil {
		return fmt.Errorf("Error in parsing CIDR ip range Err: %s",err)
	}

	if !cidrnet.Contains(ipaddr) {
		return fmt.Errorf("IP '%v' is not in CIDR range '%v'", ip, cidrnet)
	}
	return nil
}
