package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"

	"github.com/exfly/cslab/Code/Go/go-ipc/api"
	"github.com/exfly/cslab/Code/Go/go-ipc/conf"
)

func main() {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)

	task := new(api.Task)
	// Publish the receivers methods
	err := rpc.Register(task)
	if err != nil {
		log.Fatal("Format of service Task isn't correct. ", err)
	}
	// Register a HTTP handler
	rpc.HandleHTTP()

	addr := conf.RPCSvrListener
	listener, e := net.Listen(addr.NetWork, addr.Address)
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on %v", addr)
	go func() {
		// Start accept incoming HTTP connections
		err = http.Serve(listener, nil)
		if err != nil {
			log.Fatal("Error serving: ", err)
		}
	}()

	<-stopCh
	log.Print("stop the server")
	if addr.NetWork == "unix" {
		if err = os.Remove(addr.Address); err != nil {
			log.Printf("remove socket file err: %v", err)
		} else {
			log.Printf("deleted socket file: %v", addr.Address)
		}
	}
}
