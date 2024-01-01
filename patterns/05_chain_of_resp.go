package main

import "fmt"

type IHandler interface {
	SetNextHandler(Handler IHandler)
	Handle(request *Request)
}

const (
	Guest         = "Guest"
	Administrator = "Administrator"
	None          = "None"
)

type Request struct {
	Login, Password string
	SertificateHash string
	UserStatus      string
}

type AuthorizedUserHandler struct {
	next IHandler
}

func (h *AuthorizedUserHandler) CheckUser(login string, password string) (bool, string) {
	if login == "guest" && password == "guest" {
		return true, Guest
	}

	if login == "admin" && password == "admin" {
		return true, Administrator
	}

	return false, None
}
func (h *AuthorizedUserHandler) Handle(request *Request) {
	// Check if user exists in database...
	isAuthorized, status := h.CheckUser(request.Login, request.Password)
	if isAuthorized {
		request.UserStatus = status
		fmt.Println("User is authorized.")

		if h.next != nil {
			h.next.Handle(request)
		}
	} else {
		fmt.Println("User's credentials are not correct. Authorization failed!")
	}
}
func (h *AuthorizedUserHandler) SetNextHandler(Handler IHandler) {
	h.next = Handler
}

type SertifiedUserHandler struct {
	next IHandler
}

func (h *SertifiedUserHandler) Handle(request *Request) {
	if request.SertificateHash == "myHash32" {
		fmt.Println("User is sertified.")

		if h.next != nil {
			h.next.Handle(request)
		}
	} else {
		fmt.Println("User's sertification hash is not correct!")
	}
}
func (h *SertifiedUserHandler) SetNextHandler(Handler IHandler) {
	h.next = Handler
}

type AdministratorUserHandler struct {
	next IHandler
}

func (h *AdministratorUserHandler) Handle(request *Request) {
	if request.UserStatus == Administrator {
		fmt.Println("UserStatus: Administrator.")

		if h.next != nil {
			h.next.Handle(request)
		}
	} else {
		fmt.Println("This web-resource is restricted. Only administrators are allowed!")
	}
}
func (h *AdministratorUserHandler) SetNextHandler(Handler IHandler) {
	h.next = Handler
}

func main() {
	authorizedHandler := &AuthorizedUserHandler{}
	sertifiedHandler := &SertifiedUserHandler{}
	adminHandler := &AdministratorUserHandler{}

	authorizedHandler.SetNextHandler(sertifiedHandler)
	sertifiedHandler.SetNextHandler(adminHandler)

	request1 := &Request{
		Login:           "guest",
		Password:        "guest",
		SertificateHash: "myHash32",
	}
	authorizedHandler.Handle(request1)

	fmt.Println()

	request2 := &Request{
		Login:           "admin",
		Password:        "admin",
		SertificateHash: "myHash32",
	}
	authorizedHandler.Handle(request2)

	fmt.Println()

	request3 := &Request{
		Login:           "admin",
		Password:        "admin",
		SertificateHash: "fakeHash32",
	}
	authorizedHandler.Handle(request3)
}
