package main

import (
	"bufio"
	"net"
	"time"
)

//func main() {
//
//	// Run a simple http server
//	mux := http.NewServeMux()
//
//	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//		writer.WriteHeader(200)
//		writer.Write([]byte("OK\n"))
//	})
//
//	port, ok := os.LookupEnv("PORT")
//	if !ok {
//		port = "8080"
//	}
//
//	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
//
//	if err != nil {
//		panic(err)
//	}
//}

func main() {

	// Run a simple tcp server
	server, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	for {
		// only accept one connection at a time
		conn, err := server.Accept()

		if err != nil {
			panic(err)
		}

		time.Sleep(10 * time.Second)

		//r := bufio.NewReader(conn)
		w := bufio.NewWriter(conn)

		w.WriteString("OK\n")
		w.Flush()
		conn.Close()
	}
}
