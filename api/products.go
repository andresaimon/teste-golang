package api

import (
	"net/http"

	db "teste/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Construção dos Endpoints:

// Create
type createProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Price int32  `json:"price" binding:"required"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	// validando o binding:"required"
	var req createProductRequest
	err := ctx.ShouldBindJSON(&req)

	// tratamento de erro
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateProductParams{
		Name:  req.Name,
		Price: req.Price,
	}

	product, err := server.store.CreateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusCreated, product)
}

// Get
// Parâmetro do request: URI
type getProductRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	product, err := server.store.GetProduct(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, product)
}

// Delete
type deleteProductRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteProduct(ctx *gin.Context) {
	var req deleteProductRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = server.store.DeleteProduct(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, true)
}

// Update
// Required: somente o ID
type updateProductRequest struct {
	ID    int32  `json:"id" binding:"required"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
}

func (server *Server) updateProduct(ctx *gin.Context) {
	var req updateProductRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// mesma tipagem definida no product.sql.go
	arg := db.UpdateProductParams{
		ID:    req.ID,
		Name:  req.Name,
		Price: req.Price,
	}

	product, err := server.store.UpdateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, product)
}

// Get Products
// não possui parâmetros de URL e JSON
func (server *Server) getProducts(ctx *gin.Context) {
	products, err := server.store.GetProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, products)
}
