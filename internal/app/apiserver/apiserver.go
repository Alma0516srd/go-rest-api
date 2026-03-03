package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Apiserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *Apiserver {
	return &Apiserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Apiserver) Start() error {
	err := s.configureLogger()
	if err != nil {
		return err
	}
	s.configRouter()
	s.logger.Info("Starting API server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Apiserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
func (s *Apiserver) configRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *Apiserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World")
	}
}
