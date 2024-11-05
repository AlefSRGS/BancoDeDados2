package model

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
    gorm.Model
    ID         uint   `gorm:"primaryKey" json:"id"`
    Nome       string `json:"nome"`
    Sexo       string `gorm:"type:char(1);check:sexo IN ('m', 'f', 'o')" json:"sexo"`
    Idade      int    `json:"idade"`
    Nascimento time.Time `json:"data_nascimento"`
    Pontos     int    `gorm:"default:0" json:"pontos"`
}
