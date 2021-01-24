package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int, result *Result) error {
	result.Ans = num * num
	result.Num = num
	return nil
}

func main() {
	rpc.Register(new(Cal))
	cert, cErr := tls.LoadX509KeyPair("./server/server.crt", "./server/server.key")
	if cErr != nil {
		log.Fatal(cErr)
	}
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	listener, err := tls.Listen("tcp", ":443", &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Serving RPC server on port %d", 1234)
	// rpc.HandleHTTP()

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
	// log.Printf("Serviing RPC server on port %d", 1234)
	// if err := http.ListenAndServe(":1234", nil); err != nil {
	// 	log.Fatal("error serving", err)
	// }
}
