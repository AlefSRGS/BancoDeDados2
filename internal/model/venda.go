package model

import (
	"time"

	"gorm.io/gorm"
)

type Venda struct {
    gorm.Model
    ID         uint   `gorm:"primaryKey" json:"id"`
    IDCliente  uint   `gorm:"index" json:"id_cliente"`
    IDPrato    uint   `gorm:"index" json:"id_prato"`
    Quantidade int    `json:"quantidade"`
    Dia        time.Time `json:"dia"`  
    Hora       time.Time `json:"hora"`
    Valor      float64 `gorm:"type:numeric(10,2)" json:"valor"`
}
