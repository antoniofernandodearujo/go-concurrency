package service

import (
	"sync"
	"testing"
)

// Testa o caso de um único usuário comprando um ingresso.
func TestSingleTicketPurchase(t *testing.T) {
	ticketService := NewTicketService(1)

	// Usando WaitGroup para garantir que o serviço inicie antes de qualquer requisição
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticketService.Start()
	}()
	wg.Wait() // Espera até que o serviço tenha iniciado

	// Faz a requisição de compra de um ingresso
	ticketService.RequestPurchase(1)
	ticketService.Close()

	// Verifica se os ingressos restantes são 0
	if ticketService.availableTickets != 0 {
		t.Errorf("Esperado 0 ingressos disponíveis, mas restaram %d", ticketService.availableTickets)
	} else {
		t.Log("Passou: TestSingleTicketPurchase")
	}
}

// Testa o caso de múltiplos usuários comprando ingressos até que acabem.
func TestMultipleTicketPurchases(t *testing.T) {
	ticketService := NewTicketService(3)

	// Usando WaitGroup para garantir que o serviço inicie antes de qualquer requisição
	started := make(chan bool)
	go func() {
		ticketService.Start()
		close(started)
	}()

	<-started // Espera o serviço iniciar

	// Faz as requisições de compra de ingressos
	for i := 1; i <= 3; i++ {
		ticketService.RequestPurchase(i)
	}
	ticketService.Close()

	// Verifica se os ingressos restantes são 0
	if ticketService.availableTickets != 0 {
		t.Errorf("Esperado 0 ingressos disponíveis, mas restaram %d", ticketService.availableTickets)
	} else {
		t.Log("Passou: TestMultipleTicketPurchases")
	}
}

// Testa a concorrência com múltiplos usuários tentando comprar ingressos simultaneamente.
func TestConcurrentPurchases(t *testing.T) {
	ticketService := NewTicketService(10)

	// Usando WaitGroup para garantir que o serviço inicie antes de qualquer requisição
	started := make(chan bool)
	go func() {
		ticketService.Start()
		close(started)
	}()

	<-started // Espera o serviço iniciar

	// Usando WaitGroup para aguardar todas as goroutines
	var wg sync.WaitGroup
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()
			ticketService.RequestPurchase(userID)
		}(i)
	}
	wg.Wait() // Aguarda todas as goroutines terminarem
	ticketService.Close()

	// Verifica se todos os ingressos disponíveis foram comprados
	if ticketService.availableTickets != 0 {
		t.Errorf("Esperado 0 ingressos disponíveis, mas restaram %d", ticketService.availableTickets)
	} else {
		t.Log("Passou: TestConcurrentPurchases")
	}
}

// Testa a integridade dos dados para evitar que o mesmo ingresso seja comprado duas vezes.
func TestNoDuplicatePurchases(t *testing.T) {
	ticketService := NewTicketService(5)

	// Usando WaitGroup para garantir que o serviço inicie antes de qualquer requisição
	started := make(chan bool)
	go func() {
		ticketService.Start()
		close(started)
	}()

	<-started // Espera o serviço iniciar

	// Usando WaitGroup para aguardar todas as goroutines
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()
			ticketService.RequestPurchase(userID)
		}(i)
	}
	wg.Wait() // Aguarda todas as goroutines terminarem
	ticketService.Close()

	// Verifica quantos ingressos foram comprados
	purchasedTickets := 0
	for _, ticket := range ticketService.tickets {
		if ticket.Purchased {
			purchasedTickets++
		}
	}

	// Verifica se não houve compras duplicadas e que o total de ingressos comprados é correto
	if purchasedTickets != 5 {
		t.Errorf("Esperado 5 ingressos comprados, mas %d foram comprados", purchasedTickets)
	} else {
		t.Log("Passou: TestNoDuplicatePurchases")
	}
}
