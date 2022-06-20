package main

import (
	"log"
	"net"
)

/*
type staticHandler struct {
	http.Handler
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("../www")))
	http.ListenAndServe(":5000", nil)
}
*/

func main() {
	listener, e := net.Listen("tcp", ":3000")
	if e != nil {
		log.Fatalf("fail to bind address")
		return
	}
	defer listener.Close()
	for {
		con, e := listener.Accept()
		if e != nil {
			log.Println("listener error")
			continue
		} else {
			log.Println("con success")
		}
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
