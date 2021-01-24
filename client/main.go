package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	// certPool := x509.NewCertPool()
	// certBytes, err := ioutil.ReadFile("./server/server.crt")
	// if err != nil {
	// 	log.Fatal("Failed to read server.crt")
	// }
	// certPool.AppendCertsFromPEM(certBytes)
	config := tls.Config{
		// RootCAs: certPool,
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "localhost:443", &config)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := rpc.NewClient(conn)

	var result Result

	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square", err)
	}
	log.Printf("%d^2=%d", result.Num, result.Ans)
	// asyncCall := client.Go("Cal.Square", 12, &result, nil)

	// <-asyncCall.Done
	// log.Printf("%d^2=%d", result.Num, result.Ans)
}
