package main

import (
	"io"
	"net/http"
	"net/rpc"

	core "github.com/PrabodhaNirmani/vegetable-store/serverCore"
)

func main() {

	//create a new *VegetableStore object
	vegetables := core.NewVegetableStore()

	// register `vegetables` object with `rpc.DefaultServer`
	rpc.Register(vegetables)

	// register an HTTP handler for RPC communication on `http.DefaultServeMux` (default)
	// registers a handler on the `rpc.DefaultRPCPath` endpoint to respond to RPC messages
	// registers a handler on the `rpc.DefaultDebugPath` endpoint for debugging
	rpc.HandleHTTP()

	// test endpoint
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPS server live !!!")
	})

	// listen and serve default HTTP server
	http.ListenAndServe(":9000", nil)

}
