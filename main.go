package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

type DemoResponse struct {
	Data string `json:"data"`
	Ts   string `json:"ts"`
}

var (
	addr = flag.String("addr", ":8080", "TCP address to listen to")
)

func main() {
	flag.Parse()

	h := requestHandlerJson

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")
	ctx.SetContentType("text/plain; charset=utf8")
}

func requestHandlerJson(ctx *fasthttp.RequestCtx) {
	output := DemoResponse{Data: "hello world", Ts: time.Now().String()}
	bodyBuffer := new(bytes.Buffer)
	err := json.NewEncoder(bodyBuffer).Encode(output)
	if err != nil {
		ctx.SetStatusCode(500)
		return
	}

	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetStatusCode(200)
	ctx.Response.SetBodyString(bodyBuffer.String())
	//ctx.Response.SetBodyString("{\"data\":\"hello\"}")
}
