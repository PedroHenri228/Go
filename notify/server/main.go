package main

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true // Permitir conexões de qualquer origem
    },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("Erro ao fazer upgrade para WebSocket: %v", err)
        return
    }
    defer ws.Close()

    log.Println("Conexão WebSocket estabelecida com sucesso!")

    for {
        _, msg, err := ws.ReadMessage()
        if err != nil {
            log.Printf("Erro ao ler mensagem: %v", err)
            break
        }
        log.Printf("Mensagem recebida: %s", msg)
        // Envia a mensagem de volta para todos os clientes conectados
        err = ws.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            log.Printf("Erro ao enviar mensagem: %v", err)
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleConnections)
    log.Println("Servidor iniciado na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
