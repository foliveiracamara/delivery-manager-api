package service

// import (
// 	"github.com/foliveiracamara/delivery-manager-api/internal/domain"
// 	"github.com/foliveiracamara/delivery-manager-api/internal/domain/vo"
// )

// // ShippingDomainService representa o serviço de domínio para cotação de fretes
// type ShippingDomainService struct {
// 	carriers []*domain.Carrier
// }

// // NewShippingDomainService cria uma nova instância do serviço de domínio
// func NewShippingDomainService() *ShippingDomainService {
// 	return &ShippingDomainService{
// 		carriers: domain.GetAvailableCarriers(),
// 	}
// }

// // GetQuotes calcula as cotações disponíveis para um envio
// func (s *ShippingDomainService) GetQuotes(request vo.ShippingRequest) []vo.ShippingQuote {
// 	if !request.IsValid() {
// 		return []vo.ShippingQuote{}
// 	}

// 	var quotes []vo.ShippingQuote
// 	region := request.GetRegion()

// 	for _, carrier := range s.carriers {
// 		if carrier.IsAvailableForRegion(region) {
// 			price, days, _ := carrier.CalculateShipping(region, request.WeightKg)
// 			quotes = append(quotes, vo.NewShippingQuote(
// 				carrier.Name,
// 				carrier.ID,
// 				price,
// 				days,
// 			))
// 		}
// 	}

// 	return quotes
// }

// // GetBestQuote retorna a melhor cotação (menor preço)
// func (s *ShippingDomainService) GetBestQuote(request vo.ShippingRequest) (*vo.ShippingQuote, bool) {
// 	quotes := s.GetQuotes(request)
// 	if len(quotes) == 0 {
// 		return nil, false
// 	}

// 	bestQuote := quotes[0]
// 	for _, quote := range quotes[1:] {
// 		if quote.EstimatedPrice < bestQuote.EstimatedPrice {
// 			bestQuote = quote
// 		}
// 	}

// 	return &bestQuote, true
// }

// // GetCarrierByID retorna uma transportadora pelo ID
// func (s *ShippingDomainService) GetCarrierByID(carrierID string) (*domain.Carrier, bool) {
// 	for _, carrier := range s.carriers {
// 		if carrier.ID == carrierID {
// 			return carrier, true
// 		}
// 	}
// 	return nil, false
// }
