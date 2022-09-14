package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("port", 3090, "Port")
	host = flag.String("host", "localhost", "Host")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	// channel is used to pass random info to unlock the func
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()
	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done
}

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
