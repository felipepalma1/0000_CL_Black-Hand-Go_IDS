package Exercise

import (
	"fmt"
	"net"
)

func E001() {

	for i := 1; i <= 10000; i++ {
		address := fmt.Sprintf("vsrv.nino.tres.cl:%d", i)
		connection, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		err = connection.Close()
		if err != nil {
			return
		}
		fmt.Printf("%d Open\n", i)
	}
}
