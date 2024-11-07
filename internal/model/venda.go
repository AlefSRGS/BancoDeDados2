package model

import (
	"time"

	"gorm.io/gorm"
)

type Venda struct {
    gorm.Model
    IDCliente   uint      `gorm:"not null;foreignKey:IDCliente" json:"id_cliente"` 
    Cliente     Cliente   `gorm:"foreignKey:IDCliente" json:"cliente"`
    
    IDPrato     uint      `gorm:"not null;foreignKey:IDPrato" json:"id_prato"`
    Prato       Prato     `gorm:"foreignKey:IDPrato" json:"prato"` 
    
    Quantidade  int       `json:"quantidade"`
    Dia         time.Time `gorm:"type:date" json:"dia"`
    Hora        time.Time `gorm:"type:time" json:"hora"`
    Valor       float64   `gorm:"type:numeric(10,2)" json:"valor"`
}

func (Venda) TableName() string {
    return "venda"
}