// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"github.com/gofiber/fiber/v2"
	uuid "github.com/google/uuid"
	"time"
)

type ServiceRoute interface {
	SetRoutes(route *fiber.App)
}

type Option func(srv *Server)
type Handler = fiber.Handler
type ErrorHandler func(err error) error

func Service(svc ServiceRoute) Option {
	return func(srv *Server) {
		if srv.srvHTTP != nil {
			svc.SetRoutes(srv.Fiber())
		}
	}
}

func MyService(svc *httpMyService) Option {
	return func(srv *Server) {
		if srv.srvHTTP != nil {
			srv.httpMyService = svc
			svc.SetRoutes(srv.Fiber())
		}
	}
}

func SetFiberCfg(cfg fiber.Config) Option {
	return func(srv *Server) {
		srv.config = cfg
		srv.config.DisableStartupMessage = true
	}
}

func SetReadBufferSize(size int) Option {
	return func(srv *Server) {
		srv.config.ReadBufferSize = size
	}
}

func MaxBodySize(max int) Option {
	return func(srv *Server) {
		srv.config.BodyLimit = max
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(srv *Server) {
		srv.config.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(srv *Server) {
		srv.config.WriteTimeout = timeout
	}
}

func WithRequestID(headerName string) Option {
	return func(srv *Server) {
		srv.headerHandlers[headerName] = func(value string) Header {
			if value == "" {
				value = uuid.New().String()
			}
			return Header{

				LogKey:        "requestID",
				LogValue:      value,
				ResponseKey:   headerName,
				ResponseValue: value,
				SpanKey:       "requestID",
				SpanValue:     value,
			}
		}
	}
}

func WithHeader(headerName string, handler HeaderHandler) Option {
	return func(srv *Server) {
		srv.headerHandlers[headerName] = handler
	}
}

func Use(args ...interface{}) Option {
	return func(srv *Server) {
		if srv.srvHTTP != nil {
			srv.srvHTTP.Use(args...)
		}
	}
}
