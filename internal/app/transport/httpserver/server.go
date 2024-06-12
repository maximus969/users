package httpserver

// HttpServer is a HTTP server for ports
type HttpServer struct {
	userService UserService
}

// NewHttpServer creates a new HTTP server for ports
func NewHttpServer(userService UserService) HttpServer {
	return HttpServer{
		userService: userService,
	}
}
