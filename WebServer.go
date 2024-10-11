package web

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type WebServer struct {
	lock       sync.Mutex
	httpServer http.Server
	endpoints  map[string]func(*http.Request) ServerResponse
}

func NewWebServer(addr string) WebServer {
	return WebServer{
		lock: sync.Mutex{},
		httpServer: http.Server{
			Addr: addr,
		},
		endpoints: make(map[string]func(*http.Request) ServerResponse),
	}
}

func (server *WebServer) Start() {
	server.httpServer.Handler = server
	log.Fatal(server.httpServer.ListenAndServe())
}

func (server *WebServer) RegisterGet(endpoint string, lambda func(request *http.Request) ServerResponse) {
	endpoint = "GET " + strings.TrimSpace(endpoint)
	server.endpoints[endpoint] = lambda
	http.Handle(endpoint, server)
}

func (server *WebServer) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	urlPath := request.Method + " " + request.URL.String()
	endpoint := server.endpoints[urlPath]
	if endpoint == nil {
		fmt.Println("Unrecognized " + request.Method + " request at endpoint " + request.URL.String())
		return
	}
	response := server.endpoints[urlPath](request)
	for key, value := range response.header {
		responseWriter.Header().Add(key, value)
	}
	responseWriter.WriteHeader(response.statusCode)
	responseWriter.Write(response.body)
}
