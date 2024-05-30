package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *Handler) *gin.Engine {
	r := gin.Default()

	// PostgreSQL routes
	r.POST("/start-postgres-replication", handler.StartPostgresReplication)
	r.POST("/stop-postgres-replication", handler.StopPostgresReplication)
	r.GET("/postgres-status", handler.GetPostgresStatus)

	// MySQL routes
	r.POST("/start-mysql-replication", handler.StartMySQLReplication)
	r.POST("/stop-mysql-replication", handler.StopMySQLReplication)
	r.GET("/mysql-status", handler.GetMySQLStatus)

	// MSSQL routes
	r.POST("/start-mssql-replication", handler.StartMSSQLReplication)
	r.POST("/stop-mssql-replication", handler.StopMSSQLReplication)
	r.GET("/mssql-status", handler.GetMSSQLStatus)

	return r
}
