// util_test.go
package util

import (
	"api/log"
	"fmt"
	"testing"
	 
)

func Test_Util(t *testing.T) {
 CIDR="1.2.0.0/8"
	if  er := ValidateIp("1.2.4.198"); er != nil {
		log.Error(er)
		t.Fail()
	} else {
		fmt.Printf("This Is valid.")
	}

}
