package api

import (
	db "github.com/Bergmann5/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves http requestsfor our banking service.
type Server struct {
	store *db.Store
	route *gin.Engine
}

//Newserver creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// Account routes
	router.POST("/account", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	// Entry routes
	router.POST("/entry", server.createEntry)
	router.GET("/entry/:id", server.getEntry)
	router.GET("/entries", server.listEntries)

	// Transfer routes
	router.POST("/transfer", server.createTransfer)
	router.GET("/transfer/:id", server.getTransfer)
	router.GET("/transfers", server.listTransfers)

	server.route = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.route.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
