package skelmi

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"log/slog"
)

type Server struct {
	port   int32
	logger *slog.Logger
	echo   *echo.Echo
}

func NewServer(port int32, logger *slog.Logger) *Server {
	eApp := echo.New()
	eApp.HideBanner = true
	eApp.Validator = NewDefaultValidator()
	return &Server{
		port:   port,
		logger: logger,
		echo:   eApp,
	}
}

// RegisterBlock registers a logical block's methods
// Each block provides its handle functions via GetHandlers method
// Each block provides list of middlewares required for it's work
func (s *Server) RegisterBlock(block Block) {
	group := s.echo.Group(block.GetPath())
	for _, m := range block.GetMiddlewares() {
		group.Use(m.Relay)
	}
	for _, h := range block.GetHandlers() {
		group.Add(h.Method, h.Path, h.Handle, h.Middlewares...)
	}
}

// Use adds app global middleware
func (s *Server) Use(f echo.MiddlewareFunc) {
	s.echo.Use(f)
}

// SetValidator sets the validator for the app
// Default value is instance of DefaultValidator
func (s *Server) SetValidator(validator echo.Validator) {
	s.echo.Validator = validator
}

// Start starts HTTP server on giving port
func (s *Server) Start() {
	err := s.echo.Start(fmt.Sprintf("0.0.0.0:%d", s.port))
	log.Fatal(err)
}
