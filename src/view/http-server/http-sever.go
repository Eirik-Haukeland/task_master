package http_server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
)

const keyServerAddr = "serverAddr"

func HttpServer() {
	ctx := context.Background()

	router := Router()

	server := &http.Server{
		Addr:    ":3334",
		Handler: router,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server one closed\n")
	} else if err != nil {
		fmt.Printf("error starting for server one: %s\n", err)
		os.Exit(1)
	}

}
