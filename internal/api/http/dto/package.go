package dto

// PackageRequest representa a requisição para criar um novo pacote
// @Description Dados necessários para criar um novo pacote
type PackageRequest struct {
	Product       string  `json:"produto" validate:"required,min=2,max=100" example:"Camisa tamanho G"`
	WeightKg      float64 `json:"peso_kg" validate:"required,gt=0,lte=1000" example:"0.6"`
	EstadoDestino string  `json:"estado_destino" validate:"required,len=2,alpha" example:"PR"`
}

// ShippingsQuoteRequest representa a requisição para obter cotações de frete
// @Description Dados necessários para obter cotações de frete
type ShippingsQuoteRequest struct {
	PackageID string `json:"package_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
}

// HireCarrierRequest representa a requisição para contratar uma transportadora
// @Description Dados necessários para contratar uma transportadora
type HireCarrierRequest struct {
	PackageID string `json:"package_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	CarrierID string `json:"carrier_id" validate:"required" example:"nebulix"`
}

// UpdateStatusRequest representa a requisição para atualizar o status de um pacote
// @Description Dados necessários para atualizar o status de um pacote
type UpdateStatusRequest struct {
	PackageID string `json:"package_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Status    string `json:"status" validate:"required" example:"enviado"`
}

// End Requests

// PackageResponse representa a resposta de um pacote
// @Description Resposta com os dados de um pacote
type PackageResponse struct {
	ID            string                 `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Product       string                 `json:"produto" example:"Camisa tamanho G"`
	WeightKg      float64                `json:"peso_kg" example:"0.6"`
	EstadoDestino string                 `json:"estado_destino" example:"PR"`
	RegiaoDestino string                 `json:"regiao_destino" example:"sul"`
	Status        string                 `json:"status" example:"criado"`
	Shipping      *ShippingQuoteResponse `json:"entrega,omitempty"`
}

// ShippingQuoteResponse representa uma cotação de frete
// @Description Dados de uma cotação de frete
type ShippingQuoteResponse struct {
	Transportadora    string  `json:"transportadora" example:"Nebulix Logística"`
	PrecoEstimado     float64 `json:"preco_estimado" example:"42.50"`
	PrazoEstimadoDias int     `json:"prazo_estimado_dias" example:"4"`
	TransportadoraID  string  `json:"transportadora_id" example:"nebulix"`
}

// End Responses

// HealthCheckResponse representa a resposta de saúde da API
// @Description Resposta simples de status da API
type HealthCheckResponse struct {
	Message string `json:"message" example:"OK!"`
}

// SuccessResponse representa uma resposta de sucesso genérica
type SuccessResponse struct {
	Message string `json:"message" example:"Operation completed successfully"`
}

// CreatePackageResponse representa a resposta de criação de pacote
type CreatePackageResponse struct {
	Message string `json:"message" example:"Package created successfully"`
	ID      string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
}
