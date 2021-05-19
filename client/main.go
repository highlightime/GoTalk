package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	con, error := net.Dial("tcp", ":3000")
	var temp string

	if nil != error {
		log.Printf("tcp dial failed")
	} else {
		for {
			go func() {
				buf := make([]byte, 1000)
				for {
					cnt, e := con.Read(buf)
					if e != nil {
						log.Println(e)
						return
					}
					if cnt > 0 {
						data := buf[:cnt]
						log.Println(string(data))
					}
				}
			}()

			fmt.Scan(&temp)
			if temp == "end" {
				con.Close()
				log.Println("end")
				return
			}
			_, e := con.Write([]byte(temp))
			if e != nil {
				log.Printf("write error")
			}
		}
	}
}
