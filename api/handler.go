package api

import (
	"net/http"

	"github.com/cqllum/db2ch/replication"

	"github.com/gin-gonic/gin"
)

type Handler struct {
    ReplicationService      *replication.ReplicationService
    MySQLReplicationService *replication.MySQLReplicationService
    MSSQLReplicationService *replication.MSSQLReplicationService
}

func NewHandler(repService *replication.ReplicationService, mySQLRepService *replication.MySQLReplicationService, mssqlRepService *replication.MSSQLReplicationService) *Handler {
    return &Handler{
        ReplicationService: repService,
        MySQLReplicationService: mySQLRepService,
        MSSQLReplicationService: mssqlRepService,
    }
}

func (h *Handler) StartPostgresReplication(c *gin.Context) {
    slotName := c.Query("slotName")
    if slotName == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "slotName is required"})
        return
    }
    err := h.ReplicationService.StartReplication(slotName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "Postgres replication started"})
}

func (h *Handler) StopPostgresReplication(c *gin.Context) {
    h.ReplicationService.StopReplication()
    c.JSON(http.StatusOK, gin.H{"status": "Postgres replication stopped"})
}

func (h *Handler) GetPostgresStatus(c *gin.Context) {
    status := h.ReplicationService.GetStatus()
    c.JSON(http.StatusOK, gin.H{"status": status})
}

func (h *Handler) StartMySQLReplication(c *gin.Context) {
    slotName := c.Query("slotName")
    if slotName == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "slotName is required"})
        return
    }
    err := h.MySQLReplicationService.StartReplication(slotName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "MySQL replication started"})
}

func (h *Handler) StopMySQLReplication(c *gin.Context) {
    h.MySQLReplicationService.StopReplication()
    c.JSON(http.StatusOK, gin.H{"status": "MySQL replication stopped"})
}

func (h *Handler) GetMySQLStatus(c *gin.Context) {
    status := h.MySQLReplicationService.GetStatus()
    c.JSON(http.StatusOK, gin.H{"status": status})
}

func (h *Handler) StartMSSQLReplication(c *gin.Context) {
    slotName := c.Query("slotName")
    if slotName == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "slotName is required"})
        return
    }
    err := h.MSSQLReplicationService.StartReplication(slotName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "MSSQL replication started"})
}

func (h *Handler) StopMSSQLReplication(c *gin.Context) {
    h.MSSQLReplicationService.StopReplication()
    c.JSON(http.StatusOK, gin.H{"status": "MSSQL replication stopped"})
}

func (h *Handler) GetMSSQLStatus(c *gin.Context) {
    status := h.MSSQLReplicationService.GetStatus()
    c.JSON(http.StatusOK, gin.H{"status": status})
}
