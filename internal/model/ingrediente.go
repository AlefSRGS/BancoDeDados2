package model

import (
	"time"

	"gorm.io/gorm"
)

type Ingrediente struct {
    gorm.Model
    ID             uint      `gorm:"primaryKey" json:"id"`
    Nome           string    `gorm:"type:varchar(100)" json:"nome"`
    DataFabricacao time.Time `gorm:"type:date" json:"data_fabricacao"` 
    DataValidade   time.Time `gorm:"type:date" json:"data_validade"`   
    Quantidade     int       `json:"quantidade"`
    Observacao     string    `gorm:"type:text" json:"observacao"`
}

func (Ingrediente) TableName() string {
    return "ingrediente"
}
