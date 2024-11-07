package model

import "gorm.io/gorm"

type Prato struct {
    gorm.Model
    Nome          string    `gorm:"type:varchar(100)" json:"nome"`
    Descricao     string    `gorm:"type:text" json:"descricao"`
    Valor         float64   `gorm:"type:numeric(10,2)" json:"valor"`
    Disponibilidade bool    `gorm:"default:true" json:"disponibilidade"`
}

func (Prato) TableName() string {
    return "prato"
}