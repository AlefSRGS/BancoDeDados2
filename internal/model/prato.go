package model

import "gorm.io/gorm"

type Prato struct {
    ID            uint    `gorm:"primaryKey" json:"id"`
    Nome          string  `json:"nome"`
    Descricao     string  `json:"descricao"`
    Valor         float64 `gorm:"type:numeric(10,2)" json:"valor"`
    Disponibilidade bool   `gorm:"default:true" json:"disponibilidade"`
}
