# 🧾 Gateway de Pagamentos com Antifraude - GoLang

Este projeto é um **gateway de pagamentos simples e funcional** escrito em **Go (Golang)**, com verificação antifraude integrada. Ele fornece uma API REST que processa pagamentos e aplica regras de segurança básicas para identificar possíveis fraudes.

---

## 🚀 Funcionalidades

✅ API HTTP para recebimento de requisições de pagamento  
✅ Processamento e validação de dados de pagamento  
✅ Módulo antifraude com regras simples de segurança  
✅ Arquitetura modular e escalável usando `internal/`  
✅ Projeto pronto para produção ou extensão

---

## 📦 Instalação e Execução

### 1. Clone o repositório


git clone https://github.com/seu-usuario/gateway.git
cd gateway

2. Inicie o módulo Go
bash
go mod init gateway
go mod tidy

3. Execute a aplicação
bash
go run ./cmd/main.go

A aplicação estará disponível em:
📍 http://localhost:8080/payment

🔐 Regras Antifraude
O sistema considera um pagamento como fraude se:

💸 O valor for maior que R$ 10.000,00

💳 O número do cartão começar com "0000"

🌐 IP ou dispositivo do cliente estiverem ausentes

Essas regras estão em:
internal/fraud/service/antifraud.go

📡 API - Endpoint de Pagamento
POST /payment\

📤 Exemplo de JSON de entrada:\
{\
  "id": "1",\
  "user_id": "user123",\
  "amount": 500,\
  "currency": "BRL",\
  "card_number": "4111111111111111",\
  "card_holder": "João Silva",\
  "card_expiry": "12/26",\
  "cvv": "123",\
  "ip_address": "192.168.0.1",\
  "device_id": "device-001"\
}

📥 Respostas possíveis:\
Código	Significado\
200	Pagamento aceito\
400	Erro de validação JSON\
403	Pagamento identificado como fraude\
405	Método HTTP não permitido\


👨‍💻 Autor
Yann Kistenmacker
