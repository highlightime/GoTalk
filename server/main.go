package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, e := net.Listen("tcp", ":3000")
	if e != nil {
		log.Fatalf("fail to bind address")
		return
	}
	defer listener.Close()
	var temp string
	for {
		con, e := listener.Accept()
		if e != nil {
			log.Println("listener error")
			continue
		} else {
			log.Println("con success")
		}
		go func() {
			for {
				fmt.Scan(&temp)
				_, e := con.Write([]byte(temp))
				if e != nil {
					log.Println("write error")
					return
				}
			}
		}()
		go func() {
			buf := make([]byte, 1000)
			for {
				cnt, e := con.Read(buf)
				if e != nil {
					log.Println("read error")
					return
				}
				if cnt > 0 {
					data := buf[:cnt]
					log.Println(string(data))
				}
			}
		}()
	}
}
