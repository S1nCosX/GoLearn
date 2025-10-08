package server

type Server interface {
	TurnOn()
	TurnOff()
	run()
}
