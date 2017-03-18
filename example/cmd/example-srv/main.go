// Copyright 2015 Yik Yak, Inc. All rights reserved.
//

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/coyle/go-starter/example/controller"
	log "github.com/inconshreveable/log15"
)

var (
	apiPort         uint
	internalAPIPort uint
	debugPort       uint
)

func initializeFlags() {
	flag.UintVar(&apiPort, "apiPort", 8080, "api server port to use")
	flag.UintVar(&internalAPIPort, "internalAPIPort", 8081, "internal api server port to use")
	flag.UintVar(&debugPort, "debugPort", 8082, "debug server port to use")
}

func main() {
	mainLogger := log.New(log.Ctx{"main": "example/server"})
	initializeFlags()
	flag.Parse()

	validateFlags()
	startDebugServer(debugPort)

	startAPIServer(apiPort)

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChan
	mainLogger.Info("Server stopping", log.Ctx{"signal_received": sig})
}

func startAPIServer(apiPort uint) {
	go func() {
		mux := http.NewServeMux()

		mux.Handle("/", &controller.RootHandler{})
		mux.Handle("/data", &controller.GetData{})

		http.ListenAndServe(fmt.Sprintf(":%d", apiPort), mux)
	}()
}

func startInternalAPIServer(port int) {
	go func() {
		mux := http.NewServeMux()

		mux.Handle("/", &controller.RootHandler{})

		http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	}()
}

func startDebugServer(debugPort uint) {
	http.HandleFunc("/healthz", healthz)
	go func() { http.ListenAndServe(fmt.Sprintf(":%d", debugPort), nil) }()
}

func usage() {
	flag.Usage()
	os.Exit(2)
}

func validateFlags() {
	if apiPort == 0 {
		usage()
	}

	if debugPort == 0 {
		usage()
	}
}

// actually check something
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HEALTHY")
}
