package model

import "gorm.io/gorm"

type Fornecedor struct {
    ID            uint   `gorm:"primaryKey" json:"id"`
    Nome          string `json:"nome"`
    EstadoOrigem  string `gorm:"type:char(2);check:estado_origem IN ('AC', 'AL', 'AP', 'AM', 'BA', 'CE', 'DF', 'ES', 'GO', 'MA', 'MT', 'MS', 'MG', 'PA', 'PB', 'PR', 'PE', 'PI', 'RJ', 'RN', 'RS', 'RO', 'RR', 'SC', 'SP', 'SE', 'TO')" json:"estado_origem"`
}
