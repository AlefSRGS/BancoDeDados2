package repository

import "github.com/vinicius457/BancoDeDados2/internal/model"

type UsosRepository interface {
    Create(uso model.Usos) error                 
    GetByID(idPrato, idIngrediente uint) (model.Usos, error) 
    Update(uso model.Usos) error                 
    Delete(idPrato, idIngrediente uint) error   
    GetAll() ([]model.Usos, error)               
}
