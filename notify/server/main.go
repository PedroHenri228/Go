package main

import (
    "log"
    "net/http"
    "sync"

    "github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)
var mutex = sync.Mutex{}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("Erro ao fazer upgrade para WebSocket: %v", err)
        return
    }
    defer ws.Close()

    mutex.Lock()
    clients[ws] = true
    mutex.Unlock()

    log.Println("Novo cliente conectado.")

    for {
        _, msg, err := ws.ReadMessage()
        if err != nil {
            log.Printf("Erro ao ler mensagem: %v", err)
            mutex.Lock()
            delete(clients, ws)
            mutex.Unlock()
            break
        }
        log.Printf("Mensagem recebida: %s", msg)
        broadcast <- msg
    }
}


func handleMessages() {
    for {

        msg := <-broadcast
        mutex.Lock()
        for client := range clients {
            err := client.WriteMessage(websocket.TextMessage, msg)
            if err != nil {
                log.Printf("Erro ao enviar mensagem: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
        mutex.Unlock()
    }
}

func main() {
    http.HandleFunc("/ws", handleConnections)

    fs := http.FileServer(http.Dir("../client"))

	http.Handle("/", fs)

    go handleMessages()


    log.Println("Servidor iniciado na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
