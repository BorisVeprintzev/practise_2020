package apiserver

import (
	"apiserver/internal/store"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer - сам сервер
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	Store  *store.Store
}

// New - Создание нового сервера
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start - Запуск нового сервера. Если что то не так возвращает ошибку
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting APIserver")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// configureLogger - Присваиваем уровень логирования из конфига
// func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возвращаемых_результатов)
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	type request struct {
		name string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.Store = st

	return nil
}
