## 🛒 Ticket Purchase Service: Concorrência em Go

Este projeto tem como objetivo demonstrar o uso da concorrência em Go, simulando um serviço de compra de ingressos para um evento, onde múltiplos usuários tentam comprar ingressos simultaneamente. Ao utilizar goroutines e mutexes, o projeto garante que o sistema seja eficiente e seguro, mesmo com múltiplas requisições concorrentes.

### 🚀 O Projeto
Este projeto simula um serviço de ingressos em um evento. Ele permite que múltiplos usuários tentem comprar ingressos ao mesmo tempo, e gerencia o estado dos ingressos para garantir que a compra seja realizada de forma correta. O foco principal aqui é a concorrência em Go e como gerenciar acesso a recursos compartilhados de forma segura e eficiente.

### 🔑 Conceitos-chave do Projeto:
- Concorrência: O sistema permite que múltiplos usuários tentem comprar ingressos simultaneamente, usando goroutines para simular essas requisições.
- Mutex: Usado para garantir que apenas um processo consiga acessar o recurso compartilhado (o ingresso) por vez.
- Canal (Channel): Utilizado para comunicar os usuários ao serviço de tickets de maneira eficiente e sem bloqueios.
- Sincronização: Usando sync.WaitGroup para controlar a execução das goroutines e garantir que todas as compras sejam processadas antes do término do serviço.

### 🔧 Tecnologias Usadas:
- Go: Linguagem de programação para backend e manipulação de concorrência.
- Goroutines e Channels: Para gerenciar concorrência e comunicação entre processos.
- Mutexes: Para garantir o acesso exclusivo aos recursos compartilhados (ingressos).

---
  
### 🛠️ Como Executar o Projeto
Siga os passos abaixo para rodar o projeto localmente em sua máquina:

1️⃣ Clone o Repositório
```bash
  git clone https://github.com/seu-usuario/ticket-purchase-service.git
  cd ticket-purchase-service
```

2️⃣ Compilar e Executar
Este projeto não requer nenhuma dependência externa, então basta compilar e rodar o código Go.

``` bash
go run main.go
```

3️⃣ Testar o Serviço

Você pode rodar os testes para garantir que o serviço de compra de ingressos esteja funcionando corretamente. O Go já vem com uma ferramenta integrada de testes, então basta executar:

```bash
go test -v
```

4️⃣ Verificando a Execução

Ao rodar o serviço, você verá logs de cada requisição de compra de ingresso processada, assim como o estado atual do serviço.

Exemplo de saída no terminal:

```bash
ℹ️ Ticket service started.
🔄 User 1: Requesting purchase.
✔️ User 1 purchased ticket 1. Tickets remaining: 2
...
🔒 Ticket service closed.
```

### 🔍 Testes Implementados
O projeto inclui testes que validam o comportamento do sistema de compra de ingressos:

- TestSingleTicketPurchase: Verifica se o serviço permite a compra de um único ingresso.
- TestMultipleTicketPurchases: Verifica se o serviço permite a compra de múltiplos ingressos, respeitando o número de ingressos disponíveis.
- TestConcurrentPurchases: Testa a concorrência, permitindo que múltiplos usuários tentem comprar ingressos ao mesmo tempo.
- TestNoDuplicatePurchases: Garante que não haja compras duplicadas de ingressos.

### 🎯 Objetivo do Projeto
O foco principal deste projeto é a demonstração do uso de concorrência e sincronização no Go. Embora o exemplo seja simples, os conceitos de goroutines, mutexes e canal são fundamentais em sistemas concorrentes que exigem alta performance e segurança.

📚 Aprendizado:
- Como usar goroutines para executar funções simultaneamente.
- Como sincronizar tarefas concorrentes usando sync.WaitGroup e sync.Mutex.
- Como evitar problemas de race condition e garantir que o estado do sistema seja consistente.

#### 📈 Post no LinkedIn
Sinta-se à vontade para conferir o meu post sobre este projeto no LinkedIn onde compartilho mais detalhes sobre como a concorrência foi implementada e como ela pode ser aplicada em sistemas de alto desempenho.

Link do Post: 

### 🤝 Contribuições
Contribuições são bem-vindas! Se você quiser melhorar este projeto ou adicionar novas funcionalidades, basta criar um pull request. Vamos manter a concorrência sempre em alta! 🚀

### 🔒 Licença
Este projeto está licenciado sob a MIT License - veja o LICENSE.md para mais detalhes.

🔧 Autor:
Antônio Fernando
