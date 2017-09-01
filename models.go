package main

type Manufacturer struct{
	ManufacturerId string 'json:"manufacturerId"'
	//CartonsRecord []cartons 'json:"cartons"'
	Name string 'json:"name"'
}


type Units struct{

	UnitId string 'json:"unitId"'
	UnitCompositeKey string 'json:"unitCompositeKey"'
	ProductId string 'json:"productId"'
	ManufacturerId string 'json:"manufacturerId"'
	CartonId string '"json:cartonId"'
	TrackRecord string 'json:"trackRecord"'
	DateOfDepartureFromManufacturer string 'json:"dateOfDepartureFromManufacturer"'
	CurrentOwner string 'json:"currentOwner"'
	Recipient string 'json:"recipient"'
	Consumer string 'json:"consumer"'
	State string 'json:"state"'
	UnitStatus string 'json:"status"'

}

type Carton struct{
	CartonId string 'json:"cartonId"'
	CartonCompositeKey string 'json:"cartonCompositeKey"'
	UnitsInCarton []units 'json:"unitsInCarton"'
	DateofDeparture string 'json:"dateofDeparture"'
	ManufactureID string 'json:"manufactureID"'
	CurrentOwner string 'json:"currentOwner"'
	Recipient string 'json:"recipient"'
	CartonStatus string 'json:"cartonStatus"'

}

type Warehouse struct{
	WareHouseId string 'json:"wareHouseId"'
	

}


type Distributer struct{
	DistributerId string 'json:"distributerId"'
	
	
}

type Retailer struct{
	RetailerId string 'json:"retailerId"'
	
}
