package apiserver

import (
	"apiserver/internal/app/store"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// эта функция работы апишки
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err

	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting API server ")

	return http.ListenAndServe(s.config.BinAddr, s.router)
}

// эта функция для логгера
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

// функция для работы роутера

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.HandelHello())

}

func (s *APIServer) configureStore() error {
	st := store.New(&s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) HandelHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
