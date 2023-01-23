package api

// db: utilização do module 'teste' definido no go.mod + o caminho das pastas 'db' e 'sqlc'
import (
	db "teste/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Mesma tipagem definida no store.go
type Server struct {
	store  *db.ExecuteStore
	router *gin.Engine
}

// Função que recebe o Store e o Router
// Default: refere-se à Engine do Gin
func InstanceServer(store *db.ExecuteStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// utilização dos endpoints criados no 'products.go'
	router.POST("/product", server.createProduct)
	router.PUT("/product", server.updateProduct)
	router.DELETE("/product/:id", server.deleteProduct)
	router.GET("/product/:id", server.getProduct)
	router.GET("/products", server.getProducts)

	server.router = router
	return server
}

// Função de Start
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has one error:": err.Error()}
}
