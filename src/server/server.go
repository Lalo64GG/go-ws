package server

import (
	"log"
	"ws-go/src/websocket"
	"ws-go/src/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	host     string
	port     string
	httpAddr string
}

func New(host, port string) Server {
	gin.SetMode(gin.ReleaseMode)
	srv := Server{
		engine:   gin.New(),
		host:     host,
		port:     port,
		httpAddr: host + ":" + port,
	}
	srv.engine.Use(gin.Logger())
	srv.engine.Use(config.ConfigurationCors())

	// Ruta para WebSocket, donde cada cliente se conecta a una sala específica usando el query parameter ?room=<roomName>
	srv.engine.GET("/ws", websocket.WebSocketHandler)
	srv.engine.GET("/clients/", websocket.GetClients)

	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on " + s.host + ":" + s.port)
	return s.engine.Run(s.httpAddr)
}
