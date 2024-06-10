package models

// Usuario representa la tabla Usuario
type Usuario struct {
	UID         string `gorm:"primaryKey;size:20;not null;unique"`
	DisplayName string
	Nombre      string `gorm:"not null"`
	Rut         string `gorm:"not null;unique"`
}

// Billetera representa la tabla Billetera
type Billetera struct {
	ID        uint     `gorm:"primaryKey"`
	UsuarioID string   `gorm:"size:20;unique"`
	Monedas   []Moneda `gorm:"many2many:billetera_monedas"`
	Usuario   Usuario  `gorm:"foreignKey:UsuarioID"`
}

// Moneda representa la tabla Moneda
type Moneda struct {
	ID       uint   `gorm:"primaryKey"`
	Symbol   string `gorm:"unique"`
	URL      string
	Cantidad int // Cantidad de esta moneda en la billetera
}
