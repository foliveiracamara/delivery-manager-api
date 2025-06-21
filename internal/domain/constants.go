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
	DestinationRegionMidwest   DestinationRegion = "midwest"   // Centro-Oeste
	DestinationRegionNortheast DestinationRegion = "northeast" // Nordeste
	DestinationRegionNorth     DestinationRegion = "north"     // Norte
	DestinationRegionSoutheast DestinationRegion = "southeast" // Sudeste
	DestinationRegionSouth     DestinationRegion = "south"     // Sul
)
