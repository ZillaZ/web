package web

const (
	HEADER_CONTENT_LENGTH                = "Content-Length"
	HEADER_CONTENT_TYPE                  = "Content-Type"
	HEADER_ACCESS_CONTROL_CREDENTIAL     = "Access-Control-Allow-Credentials"
	HEADER_ACCESS_CONTROL_ORIGIN         = "Access-Control-Allow-Origin"
	HEADER_MAX_AGE                       = "Access-Control-Max-Age"
	HEADER_ACCESS_CONTROL_METHOD         = "Access-Control-Request-Method"
	HEADER_ACCESS_CONTROL_HEADER         = "Access-Control-Request-Header"
	HEADER_ACCESS_CONTROL_ALLOW_METHODS  = "Access-Control-Allow-Methods"
	HEADER_ACCESS_CONTROL_ALLOW_HEADERS  = "Access-Control-Allow-Headers"
	HEADER_ACCESS_CONTROL_MAX_AGE        = "Access-Control-Max-Age"
	HEADER_ACCESS_CONTROL_ALLOW_ORIGIN   = "Access-Control-Allow-Origin"
	HEADER_ACCESS_CONTROL_EXPOSE_HEADERS = "Access-Control-Expose-Headers"
)

const (
	ERROR_METHOD_NOT_ALLOWED = "method not allowed"
	ERROR_ORIGIN_NOT_ALLOWED = "origin not allowed"
	ERROR_HEADER_NOT_ALLOWED = "header not allowed"
	ERROR_NOT_A_CORS_REQUEST = "not a cors request"
)
