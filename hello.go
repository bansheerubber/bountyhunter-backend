package main

import "fmt"
import "github.com/valyala/fasthttp"

func requestHandler(context *fasthttp.RequestCtx) {
	switch string(context.Path()) {
		case "/hello":
			fmt.Fprintf(context, "hey motherfucker\n\n")

		default:
			context.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func main() {
	fasthttp.ListenAndServe(":8090", requestHandler)
}