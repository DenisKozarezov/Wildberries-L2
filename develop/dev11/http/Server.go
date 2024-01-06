package main

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
