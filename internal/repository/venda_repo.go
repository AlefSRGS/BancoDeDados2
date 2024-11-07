package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/vinicius457/BancoDeDados2/internal/model"
)

// Interface VendaRepoInterface define os métodos para o CRUD de Venda
type VendaRepoInterface interface {
	Create(ctx context.Context, venda model.Venda) error
	Read(ctx context.Context, id int) (model.Venda, error)
	Update(ctx context.Context, venda model.Venda) error
	Delete(ctx context.Context, id int) error
}

// Estrutura VendaRepo que implementa VendaRepoInterface
type VendaRepo struct {
	DB *sql.DB
}

// Construtor para VendaRepo, retorna a interface VendaRepoInterface
func NewVendaRepo(db *sql.DB) VendaRepoInterface {
	return &VendaRepo{DB: db}
}

// Create - insere uma nova venda no banco de dados
func (repo *VendaRepo) Create(ctx context.Context, venda model.Venda) error {
	query := "INSERT INTO venda (id_cliente, id_prato, quantidade, dia, hora, valor) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := repo.DB.ExecContext(ctx, query, venda.IDCliente, venda.IDPrato, venda.Quantidade, venda.Dia, venda.Hora, venda.Valor)
	return err
}

// Read - busca uma venda por ID
func (repo *VendaRepo) Read(ctx context.Context, id int) (model.Venda, error) {
	query := "SELECT id_cliente, id_prato, quantidade, dia, hora, valor FROM venda WHERE id = $1"
	row := repo.DB.QueryRowContext(ctx, query, id)

	var venda model.Venda
	if err := row.Scan(&venda.IDCliente, &venda.IDPrato, &venda.Quantidade, &venda.Dia, &venda.Hora, &venda.Valor); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return venda, nil // Não encontrou, retorna venda vazia
		}
		return venda, err
	}
	return venda, nil
}

// Update - atualiza uma venda por ID
func (repo *VendaRepo) Update(ctx context.Context, venda model.Venda) error {
	query := "UPDATE venda SET id_cliente = $1, id_prato = $2, quantidade = $3, dia = $4, hora = $5, valor = $6 WHERE id = $7"
	_, err := repo.DB.ExecContext(ctx, query, venda.IDCliente, venda.IDPrato, venda.Quantidade, venda.Dia, venda.Hora, venda.Valor, venda.ID)
	return err
}

// Delete - exclui uma venda por ID
func (repo *VendaRepo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM venda WHERE id = $1"
	_, err := repo.DB.ExecContext(ctx, query, id)
	return err
}
