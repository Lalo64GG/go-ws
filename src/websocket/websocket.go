package websocket

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   int32
	Conn *websocket.Conn
}

type Room struct {
	clients   map[int32]*Client
	broadcast chan []byte
	lock      sync.Mutex
	count     int32
}

var rooms = make(map[string]*Room)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetClients(ctx *gin.Context) {
	roomName := ctx.Query("room")
	if roomName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Room name is required"})
		return
	}

	room := getRoom(roomName)
	clientCount := atomic.LoadInt32(&room.count)
	ctx.JSON(http.StatusOK, gin.H{"clientes": clientCount})
}

func getRoom(roomName string) *Room {
	room, exists := rooms[roomName]
	if !exists {
		room = &Room{
			clients:   make(map[int32]*Client),
			broadcast: make(chan []byte),
		}
		rooms[roomName] = room
		go room.handleMessages()
	}
	return room
}

func WebSocketHandler(c *gin.Context) {
	roomName := c.Query("room")
	if roomName == "" {
		c.JSON(400, gin.H{"error": "Room name is required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	room := getRoom(roomName)
	room.lock.Lock()
	clientID := atomic.AddInt32(&room.count, 1) // Generar un nuevo ID
	client := &Client{ID: clientID, Conn: conn}
	room.clients[clientID] = client
	room.lock.Unlock()

	
	message := []byte(`{"type": "connection", "id": ` + strconv.Itoa(int(clientID)) + `}`)
	room.broadcast <- message

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error al leer mensaje:", err)
			room.lock.Lock()
			delete(room.clients, clientID)
			room.lock.Unlock()
			break
		}

		room.broadcast <- message
	}
}

func (r *Room) handleMessages() {
	for {
		message := <-r.broadcast
		r.lock.Lock()
		for _, client := range r.clients {
			err := client.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Error al escribir mensaje:", err)
				client.Conn.Close()
				delete(r.clients, client.ID)
				atomic.AddInt32(&r.count, -1)
			}
		}
		r.lock.Unlock()
	}
}
