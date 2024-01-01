package main

import "fmt"

type IState interface {
	Enter()
	Exit()
}

type HTTPConnectingState struct {
	stateMachine StateMachine
}

func (s *HTTPConnectingState) Enter() {
	fmt.Println("==== State: HTTPConnectingState ====")
	fmt.Println("Connecting to NATS-server...")
	fmt.Println("Connecting to database...")
}
func (s *HTTPConnectingState) Exit() {
	fmt.Println("Exiting state HTTPConnectingState...")
}

type HTTPListeningState struct {
	stateMachine StateMachine
}

func (s *HTTPListeningState) Enter() {
	fmt.Println("==== State: HTTPListeningState ====")
	fmt.Println("Starting to listen at port 127.0.0.1...")
}
func (s *HTTPListeningState) Exit() {
	fmt.Println("Exiting state HTTPListeningState...")
}

type HTTPShutdownState struct {
	stateMachine StateMachine
}

func (s *HTTPShutdownState) Enter() {
	fmt.Println("==== State: HTTPShutdownState ====")
	fmt.Println("Stopped to listen at port 127.0.0.1...")
}
func (s *HTTPShutdownState) Exit() {
	fmt.Println("Exiting state HTTPShutdownState...")
}

type StateMachine struct {
	CurrentState IState
}

func (s *StateMachine) ListenAtPort() {
	s.SwitchState(&HTTPListeningState{})
}
func (s *StateMachine) ShutdownServer() {
	s.SwitchState(&HTTPShutdownState{})
}
func (s *StateMachine) InitServer() {
	s.SwitchState(&HTTPConnectingState{})
}
func (s *StateMachine) SwitchState(state IState) {
	if s.CurrentState != nil {
		s.CurrentState.Exit()
	}

	s.CurrentState = state
	state.Enter()
}

func main() {
	stateMachine := StateMachine{}

	stateMachine.InitServer()

	stateMachine.ListenAtPort()

	stateMachine.ShutdownServer()
}
