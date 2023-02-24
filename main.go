package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"time"
)

type Server struct {
	Conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		Conns: make(map[*websocket.Conn]bool),
	}
}
func (s *Server) handlerWs(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())
	s.Conns[ws] = true
	s.readLoop(ws)
}
func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Println("read error", err)
			continue
		}
		msg := buf[:n]
		s.broadcast(msg)

	}
}
func (s *Server) broadcast(b []byte) {
	for ws := range s.Conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error ", err)
			}
		}(ws)
	}
}
func (s *Server) handleOrderBook(ws *websocket.Conn) {
	fmt.Println("new imcoming connection from client to order book", ws.RemoteAddr())
	for {
		payload := fmt.Sprintf("orderbook ----> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}
func main() {
	server := NewServer()
	http.Handle("/connect", websocket.Handler(server.handlerWs))
	http.Handle("/orderBook", websocket.Handler(server.handleOrderBook))
	http.ListenAndServe(":8000", nil)
}
