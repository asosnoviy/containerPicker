package server

import (
	"containerPicker/picker"
	"log"
	"net/http"
)

type Server struct {
	port       string
	picker     *picker.Picker
	httpserver *http.ServeMux
}

func NewServer(port string) *Server {
	mux := http.NewServeMux()
	return &Server{port: ":" + port, picker: picker.NewPicker(), httpserver: mux}
}

func (s *Server) Start() {

	s.httpserver.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { s.call(w, r) })

	err := http.ListenAndServe(s.port, s.httpserver)
	if err != nil {
		log.Fatalln(err)
	}

}
