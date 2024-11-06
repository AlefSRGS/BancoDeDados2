package model

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
    gorm.Model
    ID         uint      `gorm:"primaryKey" json:"id"`
    Nome       string    `gorm:"type:varchar(100)" json:"nome"`
    Sexo       string    `gorm:"type:char(1);check:sexo IN ('m', 'f', 'o')" json:"sexo"`
    Idade      int       `json:"idade"`
    Nascimento time.Time `gorm:"type:date" json:"nascimento"` 
    Pontos     int       `gorm:"default:0" json:"pontos"`    
}

func (Cliente) TableName() string {
    return "cliente"
}
