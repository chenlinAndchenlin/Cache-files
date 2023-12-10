package main

import "net/http"

type Server interface {
	Router(add string, Handfunc http.HandlerFunc)
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Router(add string, Handfunc http.HandlerFunc) {
	//TODO implement me
	panic("implement me")
}
