package repository

import "github.com/vinicius457/BancoDeDados2/internal/model"


type VendaRepository interface {
	Create(Venda model.Venda) error
    GetByID(id uint) (model.Venda, error)
    Update(Venda model.Venda) error
    Delete(id uint) error
    GetAll() ([]model.Venda, error)
}