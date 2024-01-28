package main

// import (
// 	"fmt"
// 	"time"
// )

// type Message struct {
// 	From    string
// 	Payload string // geralmente é um slice de bytes
// }

// type Server struct {
// 	msgch  chan Message
// 	quitch chan struct{} // tem menos bytes que bool e int, teoricamente 0 bytes
// }

// func (s *Server) startAndListen() {

// running:
// 	for {
// 		select {
// 		// vai bloquear aqui até alguém mandar uma msg pro channel
// 		case msg := <-s.msgch:
// 			fmt.Printf("mensagem recebida de: %s payload %s\n", msg.From, msg.Payload)

// 		case <-s.quitch:
// 			fmt.Println("O servidor está fazendo um graceful shutdown")
// 			// lógica para o graceful shutdown -- vai depender mt da aplicação
// 			// talvez precise atualizer o DB, talvez garantir que todas as requests terminam

// 			break running
// 		default:
// 		}
// 	}

// 	fmt.Println("O servidor foi desligado.")
// }

// func sendMessageToServer(msgch chan Message, payload string) {
// 	msg := Message{
// 		From:    "Victor",
// 		Payload: payload,
// 	}
// 	fmt.Println("Enviando mensagem")

// 	msgch <- msg
// }

// func gracefulQuitServer(quitch chan struct{}) {
// 	close(quitch)
// }

// func main() {
// 	s := &Server{
// 		msgch:  make(chan Message),
// 		quitch: make(chan struct{}),
// 	}

// 	go s.startAndListen()

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		go sendMessageToServer(s.msgch, "Olá mundo!")
// 	}()

// 	go func() {
// 		time.Sleep(4 * time.Second)
// 		gracefulQuitServer(s.quitch)
// 	}()

// 	// select {}
// }
