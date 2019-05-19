package io

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/alittlebrighter/go-intro/calculator"
)

type server struct {
	input  chan Msg
	output <-chan Msg

	inFlight map[uint]func(Msg)
}

func (s *server) StartServer(serveAt string, input chan Msg, output <-chan Msg) {
	s.input = input
	s.output = output

	s.inFlight = make(map[uint]func(Msg))

	go s.handleOutput()

	http.HandleFunc("/calculate", s.calculate)

	log.Println("serving calculator at", serveAt+"/calculate")
	log.Fatal(http.ListenAndServe(serveAt, nil))
}

func (s *server) calculate(w http.ResponseWriter, r *http.Request) {
	msg := NewMsg()
	var err error

	if msg.Operation, err = calculator.ParseCalcOperation(r.FormValue("op")); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, val := range strings.Split(r.FormValue("params"), ",") {
		intVal, err := strconv.Atoi(val)
		if err == nil {
			msg.Params = append(msg.Params, intVal)
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	s.inFlight[msg.Id()] = func(msg Msg) {
		body, _ := json.Marshal(msg)
		w.Write(body)
		w.WriteHeader(http.StatusOK)
		wg.Done()
	}

	s.input <- msg

	wg.Wait()
}

func (s *server) handleOutput() {
	for msg := range s.output {
		if cb, ok := s.inFlight[msg.Id()]; ok {
			cb(msg)
		} else {
			log.Println("Could not find callback for Message", msg.Id())
		}
	}
}

type uniqueIdContainer struct {
	id uint
}

func (uid uniqueIdContainer) Id() uint {
	return uid.id
}

type Msg struct {
	uniqueIdContainer
	Params    []int
	Operation calculator.CalcOperation
	Result    int
}

func NewMsg() Msg {
	return Msg{
		uniqueIdContainer: uniqueIdContainer{appNonce.New()},
		Params:            []int{},
	}
}

type nonce struct {
	val uint
}

func (n *nonce) New() uint {
	n.val++
	return n.val
}

var appNonce = new(nonce)
