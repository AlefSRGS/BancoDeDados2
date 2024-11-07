package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/vinicius457/BancoDeDados2/internal/model"
)

// Interface UsosRepoInterface define os métodos para o CRUD de Usos
type UsosRepoInterface interface {
	Create(ctx context.Context, uso model.Usos) error
	Read(ctx context.Context, idPrato, idIngrediente int) (model.Usos, error)
	Delete(ctx context.Context, idPrato, idIngrediente int) error
}

// Estrutura UsosRepo que implementa UsosRepoInterface
type UsosRepo struct {
	DB *sql.DB
}

// Construtor para UsosRepo, retorna a interface UsosRepoInterface
func NewUsosRepo(db *sql.DB) UsosRepoInterface {
	return &UsosRepo{DB: db}
}

// Create - insere um novo uso no banco de dados
func (repo *UsosRepo) Create(ctx context.Context, uso model.Usos) error {
	query := "INSERT INTO usos (id_prato, id_ingrediente) VALUES ($1, $2)"
	_, err := repo.DB.ExecContext(ctx, query, uso.IDPrato, uso.IDIngrediente)
	return err
}

// Read - busca um uso por IDs de prato e ingrediente
func (repo *UsosRepo) Read(ctx context.Context, idPrato, idIngrediente int) (model.Usos, error) {
	query := "SELECT id_prato, id_ingrediente FROM usos WHERE id_prato = $1 AND id_ingrediente = $2"
	row := repo.DB.QueryRowContext(ctx, query, idPrato, idIngrediente)

	var uso model.Usos
	if err := row.Scan(&uso.IDPrato, &uso.IDIngrediente); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uso, nil // Não encontrou, retorna uso vazio
		}
		return uso, err
	}
	return uso, nil
}

// Delete - exclui um uso por IDs de prato e ingrediente
func (repo *UsosRepo) Delete(ctx context.Context, idPrato, idIngrediente int) error {
	query := "DELETE FROM usos WHERE id_prato = $1 AND id_ingrediente = $2"
	_, err := repo.DB.ExecContext(ctx, query, idPrato, idIngrediente)
	return err
}
