package util

import (
	"crypto/rand"
	"fmt"
	"net"
	"strings"

	"github.com/saurabh2013/CRPCDemo/config"
	"github.com/saurabh2013/CRPCDemo/log"
)

/*
Genarate New psedulo Guid id as unique identifier.
*/
func Guid() (uuid string, err error) {
	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		log.Error(err, nil, "Error while GUID creation")
		return
	}
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}

/*
Validate Ip Address.
	*Should be a valid ip address
	*Should be whith in pre apecified range.
*/
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
		return fmt.Errorf("Error in parsing CIDR ip range Err: %s", err)
	}

	if !cidrnet.Contains(ipaddr) {
		return fmt.Errorf("IP '%v' is not in CIDR range '%v'", ip, cidrnet)
	}
	return nil
}
