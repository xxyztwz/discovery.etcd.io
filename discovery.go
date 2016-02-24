package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"github.com/coreos/go-systemd/activation"

	_ "github.com/coreos/discovery.etcd.io/http"
)

var addr = flag.String("addr", "", "web service address")

func main() {
	var ETCD_COON = os.Getenv("ETCD_COON")
	var BASE_URL = os.Getenv("BASE_URL")
	if( ETCD_COON == "" || BASE_URL == "" ){
		panic("Need envronment ETCD_COON and BASE_URL")
	}
	
	log.SetFlags(0)
	flag.Parse()

	if *addr != "" {
		http.ListenAndServe(*addr, nil)
	}

	listeners, err := activation.Listeners(true)
	if err != nil {
		panic(err)
	}

	if len(listeners) != 1 {
		panic("Unexpected number of socket activation fds")
	}

	http.Serve(listeners[0], nil)
}
