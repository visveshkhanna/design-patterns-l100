package main

import "fmt"

type AuthHandler struct{ BaseHandler }

type LogHandler struct{ BaseHandler }

type DataHandler struct{ BaseHandler }

func (h *AuthHandler) Handle(request string) {
	if request == "auth" || request == "process" {
		fmt.Println("AuthHandler: authenticated")
		h.callNext("process")
		return
	}
	fmt.Println("AuthHandler: skipping")
	h.callNext(request)
}

func (h *LogHandler) Handle(request string) {
	fmt.Println("LogHandler: logging request ->", request)
	h.callNext(request)
}

func (h *DataHandler) Handle(request string) {
	if request == "process" {
		fmt.Println("DataHandler: processing data")
		return
	}
	fmt.Println("DataHandler: nothing to do")
	h.callNext(request)
}
