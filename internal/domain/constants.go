package domain

type PackageStatus string

const (
	StatusCreated       PackageStatus = "criado"
	StatusWaitingPickup PackageStatus = "esperando_coleta"
	StatusCollected     PackageStatus = "coletado"
	StatusShipped       PackageStatus = "enviado"
	StatusDelivered     PackageStatus = "entregue"
	StatusLost          PackageStatus = "extraviado"
)

type DestinationRegion string

const (
	DestinationRegionMidwest   DestinationRegion = "centro-oeste"
	DestinationRegionNortheast DestinationRegion = "nordeste"
	DestinationRegionNorth     DestinationRegion = "norte"
	DestinationRegionSoutheast DestinationRegion = "sudeste"
	DestinationRegionSouth     DestinationRegion = "sul"
)

// StateToRegionMapping maps Brazilian states to regions
var StateToRegionMapping = map[string]DestinationRegion{
	// Sul
	"RS": DestinationRegionSouth, // Rio Grande do Sul
	"SC": DestinationRegionSouth, // Santa Catarina
	"PR": DestinationRegionSouth, // Paraná

	// Sudeste
	"SP": DestinationRegionSoutheast, // São Paulo
	"RJ": DestinationRegionSoutheast, // Rio de Janeiro
	"MG": DestinationRegionSoutheast, // Minas Gerais
	"ES": DestinationRegionSoutheast, // Espírito Santo

	// Centro-Oeste
	"GO": DestinationRegionMidwest, // Goiás
	"MT": DestinationRegionMidwest, // Mato Grosso
	"MS": DestinationRegionMidwest, // Mato Grosso do Sul
	"DF": DestinationRegionMidwest, // Distrito Federal

	// Nordeste
	"BA": DestinationRegionNortheast, // Bahia
	"SE": DestinationRegionNortheast, // Sergipe
	"AL": DestinationRegionNortheast, // Alagoas
	"PE": DestinationRegionNortheast, // Pernambuco
	"PB": DestinationRegionNortheast, // Paraíba
	"RN": DestinationRegionNortheast, // Rio Grande do Norte
	"CE": DestinationRegionNortheast, // Ceará
	"PI": DestinationRegionNortheast, // Piauí
	"MA": DestinationRegionNortheast, // Maranhão

	// Norte
	"TO": DestinationRegionNorth, // Tocantins
	"PA": DestinationRegionNorth, // Pará
	"AP": DestinationRegionNorth, // Amapá
	"RR": DestinationRegionNorth, // Roraima
	"AM": DestinationRegionNorth, // Amazonas
	"AC": DestinationRegionNorth, // Acre
	"RO": DestinationRegionNorth, // Rondônia
}

func GetRegionFromState(state string) (DestinationRegion, bool) {
	region, exists := StateToRegionMapping[state]
	return region, exists
}
