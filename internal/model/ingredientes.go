package model

import (
	"time"

	"gorm.io/gorm"
)

type Ingrediente struct {
    ID             uint   `gorm:"primaryKey" json:"id"`
    Nome           string `json:"nome"`
    DataFabricacao time.Time `json:"data_fabricacao"` 
    DataValidade   time.Time `json:"data_validade"`   
    Quantidade     int    `json:"quantidade"`
    Observacao     string `json:"observacao"`
}
