package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/bitcoin-trading-automation/internal/config"
	"github.com/bitcoin-trading-automation/internal/redis-server/models"
)

type RedisServerHandler struct {
	RedisRepository models.RedisRepository
	Config          config.Config
}

func NewRedisServerHandler(cfg config.Config) (*RedisServerHandler, error) {
	redisRepository, err := models.NewRedisRepository(cfg)
	if err != nil {
		return nil, err
	}

	return &RedisServerHandler{
		RedisRepository: *redisRepository,
		Config:          cfg,
	}, nil
}

func (h *RedisServerHandler) GetRedisHandler(ctx *gin.Context) {
	index := ctx.Param("index")
	key := ctx.Param("key")

	redisIndex, err := strconv.Atoi(index)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}

	if redisIndex < 0 || redisIndex >= h.Config.Redis.IndexCount {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Index is out of range"})
		return
	}

	if key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	redis := h.RedisRepository.GetRedis(redisIndex)
	value, err := redis.Get(key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"value": value})
}

func (h *RedisServerHandler) PostRedisHandler(ctx *gin.Context) {
	index := ctx.Param("index")
	key := ctx.Param("key")

	redisIndex, err := strconv.Atoi(index)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}

	if redisIndex < 0 || redisIndex >= h.Config.Redis.IndexCount {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Index is out of range"})
		return
	}

	if key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	var requestBody struct {
		Value any           `json:"value"`
		TTL   time.Duration `json:"ttl"`
	}

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redis := h.RedisRepository.GetRedis(redisIndex)
	err = redis.Set(key, requestBody.Value, requestBody.TTL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *RedisServerHandler) DeleteRedisHandler(ctx *gin.Context) {
	index := ctx.Param("index")
	key := ctx.Param("key")

	redisIndex, err := strconv.Atoi(index)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}

	if redisIndex < 0 || redisIndex >= h.Config.Redis.IndexCount {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Index is out of range"})
		return
	}

	if key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	redis := h.RedisRepository.GetRedis(redisIndex)
	err = redis.Del(key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
