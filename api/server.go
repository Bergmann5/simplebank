package api

import (
	"fmt"

	db "github.com/Bergmann5/simplebank/db/sqlc"
	"github.com/Bergmann5/simplebank/token"
	"github.com/Bergmann5/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves http requestsfor our banking service.
type Server struct {
	store      *db.Store
	route      *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

//Newserver creates a new HTTP server and setup routing
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// User route
	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// // Account routes
	authRoutes.POST("/account", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)

	// // Transfer routes
	authRoutes.POST("/transfer", server.createTransfer)
	authRoutes.GET("/transfer/:id", server.getTransfer)
	authRoutes.GET("/transfers", server.listTransfers)

	server.route = router
	return server, nil
}

// func (server *Server) setupRouter() {

// 	router := gin.Default()
// 	// User route
// 	router.POST("/user", server.createUser)
// 	router.POST("/user/login", server.loginUser)

// 	// Account routes
// 	router.POST("/account", server.createAccount)
// 	router.GET("/accounts/:id", server.getAccount)
// 	router.GET("/accounts", server.listAccounts)

// 	// Entry routes
// 	router.POST("/entry", server.createEntry)
// 	router.GET("/entry/:id", server.getEntry)
// 	router.GET("/entries", server.listEntries)

// 	// Transfer routes
// 	router.POST("/transfer", server.createTransfer)
// 	router.GET("/transfer/:id", server.getTransfer)
// 	router.GET("/transfers", server.listTransfers)

// 	server.route = router
// }

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.route.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
