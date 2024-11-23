package node

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/me4hit/distributed-go-cache/cache"
)

type Handler struct {
	// Define the fields for Handler
	cacheManager *cache.LRUCache
	server       *gin.Engine
}

func (h *Handler) SetUpHandler() {
	h.server.GET("/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		value, found := h.cacheManager.Get(key)
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"message": "key not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
	})

	h.server.POST("/cache", func(c *gin.Context) {
		var json struct {
			Key   string `json:"key" binding:"required"`
			Value string `json:"value" binding:"required"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		h.cacheManager.Set(json.Key, json.Value)
		c.JSON(http.StatusOK, gin.H{"message": "cache set successfully"})
	})
	h.server.Run()
}

func (h *Handler) Run() {
	h.server.Run()
}

type Node struct {
	NodeID  string
	nodes   []string
	handler *Handler
}

func NewNode(nodeID string, cacheSize int) *Node {
	return &Node{
		NodeID:  nodeID,
		handler: &Handler{server: gin.Default(), cacheManager: cache.NewLRUCache(cacheSize)},
	}
}

func (n *Node) Start() {
	n.handler.SetUpHandler()
	n.handler.Run()
}
