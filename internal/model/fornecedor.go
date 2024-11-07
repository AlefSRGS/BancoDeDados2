package model

import "gorm.io/gorm"

type Fornecedor struct {
    gorm.Model
    ID            uint   `gorm:"primaryKey" json:"id"`
    Nome          string `gorm:"type:varchar(100)" json:"nome"`
    EstadoOrigem  string `gorm:"type:char(2);check:estado_origem IN ('AC', 'AL', 'AP', 'AM', 'BA', 'CE', 'DF', 'ES', 'GO', 'MA', 'MT', 'MS', 'MG', 'PA', 'PB', 'PR', 'PE', 'PI', 'RJ', 'RN', 'RS', 'RO', 'RR', 'SC', 'SP', 'SE', 'TO')" json:"estado_origem"`
}

func (Fornecedor) TableName() string {
    return "fornecedor" 
}
