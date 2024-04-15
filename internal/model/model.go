package model

type Customer struct {
	CustomerId string `gorm:"primaryKey" json:"customerId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"emailAddress"`
	Phone      string `json:"phoneNumber"`
	Address    string `json:"address"`
}

type Products struct {
	ProductId string  `gorm:"primaryKey" json:"productId"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	VendorId  string  `json:"vendorId"`
}

type Services struct {
	ServiceId string  `gorm:"primaryKey" json:"serviceId"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
}

type Vendors struct {
	VendorId string  `gorm:"primaryKey" json:"vendorId"`
	Name     string  `json:"name"`
	Contact  float32 `json:"contact"`
	Phone    string  `json:"phoneNumber"`
	Email    string  `json:"emailAddress"`
	Address  string  `json:"address"`
}
