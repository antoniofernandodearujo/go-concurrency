package main

import (
	"fmt"
	"tickets-concurrency/internal/config"
	"tickets-concurrency/internal/service"
	"time"
)

func main() {
	// Inicializa o serviço de tickets
	ticketService := service.NewTicketService(config.TotalTickets)
	go ticketService.Start()

	// Simula múltiplos usuários tentando comprar ingressos
	for i := 1; i <= 20; i++ { // Exemplo com 20 tentativas
		go ticketService.RequestPurchase(i)
		time.Sleep(100 * time.Millisecond) // Pequeno intervalo entre as tentativas
	}

	// Dá tempo para todos os pedidos serem processados
	time.Sleep(2 * time.Second)

	// Encerra o canal
	ticketService.Close()
	fmt.Println("All ticket purchases processed.")
}
