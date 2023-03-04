package main

import (
	"containerPicker/cmd"
	"fmt"

	"github.com/kardianos/service"
)

const serviceName = "Container  picker service"
const serviceDescription = "Web service for quick conteiner picking"

type program struct{}

func (p program) Start(s service.Service) error {
	fmt.Println(s.String() + " started")
	go p.run()
	return nil
}

func (p program) Stop(s service.Service) error {
	fmt.Println(s.String() + " stopped")
	return nil
}

func (p program) run() {

	cmd.Execute()

	// Port := "5555"
	// myserver := server.NewServer(Port)
	// myserver.Start()

}

func main() {
	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Cannot create the service: " + err.Error())
	}
	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}
