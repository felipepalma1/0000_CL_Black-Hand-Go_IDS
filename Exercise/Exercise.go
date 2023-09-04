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

func E002() {

	for i := 1; i <= 10000; i++ {
		go func(j int) {

			address := fmt.Sprintf("vsrv.nino.tres.cl:%d", j)
			connection, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			err = connection.Close()
			if err != nil {
				return
			}
			fmt.Printf("%d Open\n", j)
		}(i)
	}
}
