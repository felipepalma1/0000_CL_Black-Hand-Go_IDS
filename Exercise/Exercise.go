package Exercise

import (
	"fmt"
	"net"
)

func E001() {
	_, err := net.Dial("tcp", "vsrv.nino.tres.cl:3000")
	if err == nil {
		fmt.Println("Connection Successful")
	}
}
