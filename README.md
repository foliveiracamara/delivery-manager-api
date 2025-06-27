# Delivery Manager API

API para gerenciamento de pacotes e cotações de frete, desenvolvida em Go com arquitetura limpa.

## 🚀 Funcionalidades

- ✅ **Criação de Pacotes**: Cadastro de pacotes com produto, peso e destino
- ✅ **Cotação de Fretes**: Obtenção de cotações de múltiplas transportadoras
- ✅ **Contratação de Transportadora**: Seleção e contratação de transportadora
- ✅ **Atualização de Status**: Controle do ciclo de vida do pacote
- ✅ **Validações de Negócio**: Regras que garantem integridade dos dados
- ✅ **Documentação Swagger**: API documentada e testável

## 🏗️ Arquitetura

O projeto segue os princípios da **Clean Architecture** com as seguintes camadas:

```
├── cmd/                 # Ponto de entrada da aplicação
├── internal/
│   ├── api/             # Camada de apresentação (HTTP)
│   ├── application/     # Casos de uso
│   ├── domain/          # Entidades e regras de negócio
│   ├── infrastructure/  # Implementações externas (DB, APIs)
│   └── service/         # Serviços de domínio
├── docs/                # Documentação Swagger
└── test.http            # Exemplos de requisições
```

## 📋 Pré-requisitos

- **Go 1.21+** ([Download](https://golang.org/dl/))

## 🛠️ Instalação

1. **Clone o repositório:**
```bash
git clone <repository-url>
cd delivery-manager-api
```

2. **Instale as dependências:**
```bash
go mod download
```

3. **Instale o Swag (para documentação):**
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## 🚀 Executando o Projeto

### **Opção 1: Execução Direta**
```bash
go run cmd/main.go server start
```

### **Opção 2: Build e Execução**
```bash
go build -o main.exe cmd/main.go
./main.exe server start
```

### **Opção 3: Usando Make (se disponível)**
```bash
make run
```

A API estará disponível em: **http://localhost:5000**

## 📚 Documentação da API

### **Swagger UI**
Após iniciar a aplicação, acesse: **http://localhost:5000/swagger/index.html**

### **Endpoints Principais**

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/health` | Health check da API |
| `POST` | `/package/` | Criar novo pacote |
| `GET` | `/package/{id}` | Buscar pacote por ID |
| `POST` | `/package/{id}/quote` | Obter cotações de frete |
| `POST` | `/package/hire-carrier` | Contratar transportadora |
| `PUT` | `/package/status` | Atualizar status do pacote |

## 🧪 Testes

### **Executar Todos os Testes**
```bash
go test ./internal/... -v
```

### **Executar Testes com Cobertura**
```bash
go test ./internal/... -v -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

## 📝 Exemplos de Uso

### **1. Criar um Pacote**
```bash
curl -X POST http://localhost:5000/package/ \
  -H "Content-Type: application/json" \
  -d '{
    "produto": "Smartphone Samsung Galaxy S23",
    "peso_kg": 0.25,
    "estado_destino": "SP"
  }'
```

### **2. Obter Cotações de Frete**
```bash
curl -X POST http://localhost:5000/package/{package-id}/quote
```

### **3. Contratar Transportadora**
```bash
curl -X POST http://localhost:5000/package/hire-carrier \
  -H "Content-Type: application/json" \
  -d '{
    "package_id": "{package-id}",
    "carrier_id": "nebulix"
  }'
```

### **4. Atualizar Status**
```bash
curl -X PUT http://localhost:5000/package/status \
  -H "Content-Type: application/json" \
  -d '{
    "package_id": "{package-id}",
    "status": "enviado"
  }'
```

## 🔒 Validações de Negócio

### **1. Validações de Criação de Pacote**
- **Produto**: Obrigatório, mínimo 2 caracteres, máximo 100 caracteres
- **Peso**: Obrigatório, maior que 0kg, máximo 1000kg
- **Estado de Destino**: Obrigatório, exatamente 2 caracteres alfabéticos
- **Região de Destino**: Deve ser uma região válida (sul, sudeste, centro-oeste, nordeste, norte)

### **2. Validações de Status**
- **Status Válidos**: Apenas `criado`, `esperando_coleta`, `coletado`, `enviado`, `entregue`, `extraviado`
- **Status que Requerem Transportadora**: Os seguintes status só podem ser aplicados a pacotes com transportadora contratada:
  - `esperando_coleta`
  - `coletado`
  - `enviado`
  - `entregue`
  - `extraviado`

### **3. Validações de Transportadora**
- **Pacote Único**: Um pacote não pode ter mais de uma transportadora
- **Região de Atendimento**: A transportadora deve atender a região do pacote
- **Transportadora Existente**: A transportadora deve existir no sistema

### **4. Validações de Cotação**
- **Peso Mínimo**: Para pacotes muito leves, o preço mínimo é o preço por kg da região
- **Região Válida**: Apenas transportadoras que atendem a região são consideradas
- **Ordenação**: Cotações são ordenadas por prazo de entrega (mais rápido primeiro)

### **Exemplos de Validação**

#### **❌ Falha - Status sem Transportadora**
```bash
curl -X PUT http://localhost:5000/package/status \
  -H "Content-Type: application/json" \
  -d '{
    "package_id": "123",
    "status": "enviado"
  }'
# Resposta: 400 - "Package cannot be marked as 'enviado' without a carrier assigned"
```

#### **❌ Falha - Estado Inválido**
```bash
curl -X POST http://localhost:5000/package/ \
  -H "Content-Type: application/json" \
  -d '{
    "produto": "Smartphone",
    "peso_kg": 0.25,
    "estado_destino": "XX"
  }'
# Resposta: 400 - "Invalid state: XX"
```

#### **❌ Falha - Transportadora Duplicada**
```bash
# Tentar contratar segunda transportadora
curl -X POST http://localhost:5000/package/hire-carrier \
  -H "Content-Type: application/json" \
  -d '{
    "package_id": "123",
    "carrier_id": "rotafacil"
  }'
# Resposta: 409 - "Package already has a carrier"
```

#### **✅ Sucesso - Fluxo Completo**
```bash
# 1. Criar pacote
# 2. Contratar transportadora
# 3. Atualizar status (funciona)
```

## 🏢 Transportadoras Disponíveis

| ID | Nome | Regiões Atendidas |
|----|------|-------------------|
| `nebulix` | Nebulix Logística | Sul, Sudeste |
| `rotafacil` | RotaFácil Transportes | Sul, Sudeste, Centro-Oeste, Nordeste |
| `moventra` | Moventra Express | Centro-Oeste, Nordeste |

## 📊 Status dos Pacotes

| Status | Descrição | Requer Transportadora |
|--------|-----------|----------------------|
| `criado` | Pacote criado | ❌ |
| `esperando_coleta` | Aguardando coleta | ✅ |
| `coletado` | Pacote coletado | ✅ |
| `enviado` | Pacote enviado | ✅ |
| `entregue` | Pacote entregue | ✅ |
| `extraviado` | Pacote extraviado | ✅ |

## 🛠️ Desenvolvimento

### **Gerar Documentação Swagger**
```bash
swag init -g cmd/main.go
```

### **Comandos Make Disponíveis**
```bash
make run          # Executar aplicação
make doc          # Gerar documentação Swagger
make test         # Executar testes
make test-coverage # Executar testes com cobertura
```

## 🚀 Próximos Passos

### **📊 Persistência de Dados**
- [ ] **PostgreSQL com Docker**: Migrar de armazenamento em memória para banco de dados PostgreSQL
- [ ] **Connection Pool**: Configurar pool de conexões
- [ ] **Docker Compose**: Criar ambiente completo com PostgreSQL

### **🚨 Gestão de Pacotes Extraviados**
- [ ] **Workflow de Investigação**: Implementar fluxo completo para pacotes extraviados
- [ ] **Notificações**: Sistema de notificação para cliente e transportadora
- [ ] **Processo de Indenização**: Automatizar cálculo e processamento de indenizações
- [ ] **Relatórios**: Gerar relatórios de extravio por transportadora/região
- [ ] **Status de Investigação**: Adicionar status como "em_investigacao", "indenizado"

### **📈 Monitoramento e Observabilidade**
- [ ] **Logs Estruturados**: Implementar logging estruturado com níveis
- [ ] **Métricas**: Adicionar métricas de performance e negócio
- [ ] **Alertas**: Sistema de alertas para falhas e métricas críticas

### **🔧 Melhorias Técnicas**
- [ ] **Cache**: Implementar cache para cotações e dados de transportadoras
- [ ] **Autenticação**: Implementar sistema de autenticação JWT
- [ ] **Abstração**: Aumentar orientação de serviços por interfaces e não implementações
