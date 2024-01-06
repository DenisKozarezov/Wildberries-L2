package main

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"services"
	"strings"

	"github.com/go-playground/validator"
)

type ServerStatus int8

type ServerConfig struct {
	IP       string `validate:"required"`
	Port     string `validate:"required"`
	Protocol string `validate:"required"`
}

const (
	SERV_Listening ServerStatus = 0
	SERV_Shutdown  ServerStatus = 1
)

type Server struct {
	Status ServerStatus
}

func (s *Server) setupServices() *http.ServeMux {
	mux := http.NewServeMux()
	services.SetupCalenderServiceHandlers(mux)

	log.Println("ServeMux handlers are setup.")

	return mux
}
func (s *Server) StartListening(config *ServerConfig) error {
	if len(strings.TrimSpace(config.IP)) == 0 {
		log.Panicln("Invalid network address for listeting...")
	}

	s.Status = SERV_Listening
	mux := s.setupServices()

	address := strings.Join([]string{config.IP, config.Port}, ":")

	log.Printf("Server is starting to listen at address %s", address)

	if err := http.ListenAndServe(address, mux); err != nil {
		return err
	}

	return nil
}

func (s *Server) StopListening(ip string, port int) {
	s.Status = SERV_Shutdown
	log.Println("Server is stopped.")
}

func main() {
	file, err := os.ReadFile("C:\\dev\\Wildberries-L2\\develop\\dev11\\http\\config.json")

	if err != nil {
		log.Fatalln("Could not read a server config:", err)
	}

	config := &ServerConfig{}
	if err = json.Unmarshal(file, &config); err != nil {
		log.Fatalln("Could not parse json:", err)
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, fieldErr := range errs {
			fmt.Printf("Field %s: %s\n", fieldErr.Field(), fieldErr.Tag())
		}
	}

	server := &Server{}
	if err = server.StartListening(config); err != nil {
		log.Fatalln("Server cannot start to listen:", err)
	}
}
