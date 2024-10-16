package web

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Cors struct {
	AllowedOrigins   map[string]bool
	AllowedMethods   map[string]bool
	AllowedHeaders   map[string]bool
	AllowCredentials bool
	ExposeHeaders    map[string]bool
	MaxAge           int
}

func (cors *Cors) validadeOrigin(request *http.Request) error {
	origin := request.Header.Get("Origin")
	if origin == "" {
		return errors.New(ERROR_NOT_A_CORS_REQUEST)
	}
	_, wildcardOrigin := cors.AllowedOrigins["*"]
	_, originOk := cors.AllowedOrigins[origin]
	if !originOk && !wildcardOrigin {
		return errors.New(ERROR_ORIGIN_NOT_ALLOWED + " " + origin)
	}
	return nil
}

func (cors *Cors) validateMethod(request *http.Request) error {
	_, wildcardMethod := cors.AllowedMethods["*"]
	_, methodOk := cors.AllowedMethods[request.Method]
	if !methodOk && !wildcardMethod {
		return errors.New(ERROR_METHOD_NOT_ALLOWED + " " + request.Method)
	}
	return nil
}

func (cors *Cors) validateHeaders(request *http.Request) error {
	_, wildcardHeader := cors.AllowedHeaders["*"]
	if !wildcardHeader {
		for key := range request.Header {
			fmt.Println(key)
			_, headerOk := cors.AllowedHeaders[key]
			if !headerOk {
				return errors.New(ERROR_HEADER_NOT_ALLOWED + " " + key)
			}
		}
	}
	return nil
}

func (cors *Cors) validateRequest(responseWriter http.ResponseWriter, request *http.Request) bool {
	if cors.BuildOptions(responseWriter, request) {
		return true
	}
	if err := cors.validate(responseWriter, request); err != nil {
		return err.Error() != ERROR_NOT_A_CORS_REQUEST
	}
	if request.Method == http.MethodOptions {
		writeHeader(responseWriter, cors.ExposeHeaders, HEADER_ACCESS_CONTROL_EXPOSE_HEADERS)
	}
	if cors.AllowCredentials {
		responseWriter.Header().Add(HEADER_ACCESS_CONTROL_CREDENTIAL, "true")
	}
	responseWriter.Header().Add(HEADER_MAX_AGE, strconv.Itoa(cors.MaxAge))
	return false
}

func (cors *Cors) validate(responseWriter http.ResponseWriter, request *http.Request) error {
	if originErr := cors.validadeOrigin(request); originErr != nil {
		if originErr.Error() == ERROR_NOT_A_CORS_REQUEST {
			return originErr
		}
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte{})
		return originErr
	} else if methodErr := cors.validateMethod(request); methodErr != nil {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		responseWriter.Write([]byte{})
		return methodErr
	} else if corsErr := cors.validateHeaders(request); corsErr != nil {
		responseWriter.WriteHeader(http.StatusForbidden)
		responseWriter.Write([]byte{})
		return corsErr
	}
	return nil
}

func (cors *Cors) BuildOptions(responseWriter http.ResponseWriter, request *http.Request) bool {
	if !isValidOptionRequest(request) {
		return false
	}
	writeHeader(responseWriter, cors.AllowedMethods, HEADER_ACCESS_CONTROL_ALLOW_METHODS)
	writeHeader(responseWriter, cors.AllowedHeaders, HEADER_ACCESS_CONTROL_ALLOW_HEADERS)
	writeHeader(responseWriter, cors.AllowedOrigins, HEADER_ACCESS_CONTROL_ALLOW_ORIGIN)
	responseWriter.Header().Add(HEADER_ACCESS_CONTROL_MAX_AGE, strconv.Itoa(cors.MaxAge))
	if cors.AllowCredentials {
		responseWriter.Header().Add(HEADER_ACCESS_CONTROL_CREDENTIAL, "true")
	}
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte{})
	return true
}
