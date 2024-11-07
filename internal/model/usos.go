package model

import "gorm.io/gorm"

type Usos struct {
    gorm.Model
    IDPrato      uint `gorm:"primaryKey;foreignKey:IDPrato" json:"id_prato"`
    IDIngrediente uint `gorm:"primaryKey;foreignKey:IDIngrediente" json:"id_ingrediente"`
    
    Prato        Prato       `gorm:"foreignKey:IDPrato;constraint:OnDelete:CASCADE;" json:"prato"`
    Ingrediente  Ingrediente `gorm:"foreignKey:IDIngrediente;constraint:OnDelete:CASCADE;" json:"ingrediente"`
}

