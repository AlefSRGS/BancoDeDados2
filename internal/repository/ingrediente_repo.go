package repository

import "github.com/vinicius457/BancoDeDados2/internal/model"


type IngredienteRepository interface {
	Create(ingrediente model.Ingrediente) error
    GetByID(id uint) (model.Ingrediente, error)
    Update(ingrediente model.Ingrediente) error
    Delete(id uint) error
    GetAll() ([]model.Ingrediente, error)
}