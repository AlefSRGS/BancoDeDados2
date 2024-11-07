package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/vinicius457/BancoDeDados2/internal/handler"
	"github.com/vinicius457/BancoDeDados2/internal/repository"
	"github.com/vinicius457/BancoDeDados2/config"
)

func initRoutes(router *gin.Engine) {
	// Obtém a conexão com o banco de dados usando config.GetDbConnection
	gormDB := config.GetDbConnection()

	// Extrai o *sql.DB de *gorm.DB
	db, err := gormDB.DB()
	if err != nil {
		panic("Failed to get *sql.DB from *gorm.DB")
	}

	// Inicializa os handlers
	handler.InitializingHandlers()

	// Configura a rota de teste
	router.GET("/ping", handler.Ping)

	// Configuração das rotas da API para as entidades
	SetupAPIRoutes(router, db)
}

// SetupAPIRoutes - Configura as rotas principais da API para as entidades
func SetupAPIRoutes(router *gin.Engine, db *sql.DB) {
	// Inicializando os repositórios e handlers
	clienteRepo := repository.NewClienteRepo(db)
	clienteHandler := handler.NewClienteHandler(clienteRepo)

	fornecedorRepo := repository.NewFornecedorRepo(db)
	fornecedorHandler := handler.NewFornecedorHandler(fornecedorRepo)

	ingredienteRepo := repository.NewIngredienteRepo(db)
	ingredienteHandler := handler.NewIngredienteHandler(ingredienteRepo)

	pratoRepo := repository.NewPratoRepo(db)
	pratoHandler := handler.NewPratoHandler(pratoRepo)

	usosRepo := repository.NewUsosRepo(db)
	usosHandler := handler.NewUsosHandler(usosRepo)

	vendaRepo := repository.NewVendaRepo(db)
	vendaHandler := handler.NewVendaHandler(vendaRepo)

	funcoesRepo := repository.NewFuncoesRepo(db)
	funcoesHandler := handler.NewFuncoesHandler(funcoesRepo)

	// Definindo as rotas para cada entidade
	api := router.Group("/api")
	{
		api.POST("/clientes", clienteHandler.CreateCliente)
		api.GET("/clientes/:id", clienteHandler.GetCliente)
		api.PUT("/clientes/:id", clienteHandler.UpdateCliente)
		api.DELETE("/clientes/:id", clienteHandler.DeleteCliente)

		api.POST("/fornecedores", fornecedorHandler.CreateFornecedor)
		api.GET("/fornecedores/:id", fornecedorHandler.GetFornecedor)
		api.PUT("/fornecedores/:id", fornecedorHandler.UpdateFornecedor)
		api.DELETE("/fornecedores/:id", fornecedorHandler.DeleteFornecedor)

		api.POST("/ingredientes", ingredienteHandler.CreateIngrediente)
		api.GET("/ingredientes/:id", ingredienteHandler.GetIngrediente)
		api.PUT("/ingredientes/:id", ingredienteHandler.UpdateIngrediente)
		api.DELETE("/ingredientes/:id", ingredienteHandler.DeleteIngrediente)

		api.POST("/pratos", pratoHandler.CreatePrato)
		api.GET("/pratos/:id", pratoHandler.GetPrato)
		api.PUT("/pratos/:id", pratoHandler.UpdatePrato)
		api.DELETE("/pratos/:id", pratoHandler.DeletePrato)

		api.POST("/usos", usosHandler.CreateUsos)
		api.GET("/usos/:id_prato/:id_ingrediente", usosHandler.GetUsos)
		api.DELETE("/usos/:id_prato/:id_ingrediente", usosHandler.DeleteUsos)

		api.POST("/vendas", vendaHandler.CreateVenda)
		api.GET("/vendas/:id", vendaHandler.GetVenda)
		api.PUT("/vendas/:id", vendaHandler.UpdateVenda)
		api.DELETE("/vendas/:id", vendaHandler.DeleteVenda)

		// Rota para aplicar reajuste
		api.POST("/reajuste", funcoesHandler.Reajuste)
		
		// Rota para sorteio de cliente
		api.POST("/sorteio_cliente", funcoesHandler.SorteioCliente)
		
		// Rota para estatísticas de vendas
		api.GET("/estatisticas_vendas", funcoesHandler.EstatisticasVendas)
	}
}
