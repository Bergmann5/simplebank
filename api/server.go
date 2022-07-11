package api

import (
	db "github.com/Bergmann5/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
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

	router.POST("/account", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

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
