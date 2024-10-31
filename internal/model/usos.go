package model

import "gorm.io/gorm"

type Usos struct {
    IDPrato      uint `gorm:"primaryKey;column:id_prato" json:"id_prato"`
    IDIngrediente uint `gorm:"primaryKey;column:id_ingrediente" json:"id_ingrediente"`
}
