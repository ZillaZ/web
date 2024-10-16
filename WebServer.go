package web

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

var methodMap = map[string]bool{
	http.MethodConnect: true,
	http.MethodDelete:  true,
	http.MethodGet:     true,
	http.MethodHead:    true,
	http.MethodOptions: true,
	http.MethodPatch:   true,
	http.MethodPost:    true,
	http.MethodPut:     true,
	http.MethodTrace:   true,
}

type WebServer struct {
	lock       sync.Mutex
	httpServer http.Server
	endpoints  map[string]func(*http.Request) ServerResponse
	cors       Cors
}

func NewWebServer(addr string, cors Cors) WebServer {
	return WebServer{
		lock: sync.Mutex{},
		httpServer: http.Server{
			Addr: addr,
		},
		endpoints: make(map[string]func(*http.Request) ServerResponse),
		cors:      cors,
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
	alreadyWritten := server.cors.validateRequest(responseWriter, request)
	if alreadyWritten {
		return
	}
	urlPath := request.Method + " " + request.URL.String()
	endpoint := server.endpoints[urlPath]
	if endpoint == nil {
		fmt.Println("Unrecognized " + request.Method + " request at endpoint " + request.URL.String())
		return
	}
	response := server.endpoints[urlPath](request)
	for key, values := range response.header {
		for _, value := range values {
			responseWriter.Header().Add(key, value)
		}
	}
	responseWriter.WriteHeader(response.statusCode)
	responseWriter.Write(response.body)
}

func isValidOptionRequest(request *http.Request) bool {
	if request.Method != http.MethodOptions {
		return false
	}
	if method := request.Header.Get(HEADER_ACCESS_CONTROL_METHOD); !methodMap[method] {
		return false
	}
	if methodHeader := request.Header.Get(HEADER_ACCESS_CONTROL_HEADER); !isValidMethodHeader(methodHeader) {
		return false
	}
	return true
}

func isValidMethodHeader(header string) bool {
	return header != ""
}

func writeHeader(responseWriter http.ResponseWriter, values map[string]bool, header string) {
	for value := range values {
		responseWriter.Header().Add(header, value)
	}
}
