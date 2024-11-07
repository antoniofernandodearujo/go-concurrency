package service

import (
	"log"
	"math/rand"
	"sync"
	"tickets-concurrency/internal/models"
	"time"

	"github.com/fatih/color"
)

var (
	infoColor    = color.New(color.FgBlue).SprintFunc()
	successColor = color.New(color.FgGreen).SprintFunc()
	errorColor   = color.New(color.FgRed).SprintFunc()
)

// TicketService estrutura do serviço de tickets.
type TicketService struct {
	tickets          []models.Ticket
	ticketsChan      chan int
	mu               sync.Mutex
	totalTickets     int
	availableTickets int
}

// NewTicketService inicializa o serviço com um número de ingressos.
func NewTicketService(totalTickets int) *TicketService {
	tickets := make([]models.Ticket, totalTickets)
	for i := 0; i < totalTickets; i++ {
		tickets[i] = models.Ticket{ID: i + 1, Purchased: false}
	}
	return &TicketService{
		tickets:          tickets,
		ticketsChan:      make(chan int),
		totalTickets:     totalTickets,
		availableTickets: totalTickets,
	}
}

// PurchaseTicket processa um pedido de compra de ingresso.
func (s *TicketService) PurchaseTicket(userID int) {
	// Simulando o tempo de processamento do servidor
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)+100)) // tempo aleatório de 100ms a 300ms

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.availableTickets <= 0 {
		log.Printf("[%s] %s User %d: No tickets available.\n", time.Now().Format(time.RFC3339), errorColor("❌"), userID)
		return
	}

	// Processa a compra do próximo ingresso disponível.
	for i := 0; i < s.totalTickets; i++ {
		if !s.tickets[i].Purchased {
			s.tickets[i].Purchased = true
			s.availableTickets--
			log.Printf("[%s] %s User %d purchased ticket %d. Tickets remaining: %d\n",
				time.Now().Format(time.RFC3339), successColor("✔️"), userID, s.tickets[i].ID, s.availableTickets)
			break
		}
	}
}

// Start inicia a concorrência e lida com múltiplos usuários tentando comprar ingressos.
func (s *TicketService) Start() {
	log.Printf("[%s] %s Ticket service started.\n", time.Now().Format(time.RFC3339), infoColor("ℹ️"))
	for userID := range s.ticketsChan {
		log.Printf("[%s] %s User %d: Request received.\n", time.Now().Format(time.RFC3339), infoColor("⏳"), userID)
		go s.PurchaseTicket(userID)
	}
}

// RequestPurchase adiciona uma requisição na fila de compras.
func (s *TicketService) RequestPurchase(userID int) {
	log.Printf("[%s] %s User %d: Requesting purchase.\n", time.Now().Format(time.RFC3339), infoColor("🔄"), userID)
	s.ticketsChan <- userID
}

// Close encerra o canal de requisições.
func (s *TicketService) Close() {
	close(s.ticketsChan)
	log.Printf("[%s] %s Ticket service closed.\n", time.Now().Format(time.RFC3339), infoColor("🔒"))
}
