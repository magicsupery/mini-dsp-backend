package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mini-dsp-backend/services"
)

type AdvertiserController struct {
	AdvertiserService services.AdvertiserService
}

// NewAdvertiserController 创建一个新的 AdvertiserController 实例
func NewAdvertiserController(service services.AdvertiserService) *AdvertiserController {
	return &AdvertiserController{AdvertiserService: service}
}

// CreateAdvertiser 创建一个新的广告主
func (ac *AdvertiserController) CreateAdvertiser(c *gin.Context) {
	var req struct {
		Name    string `json:"name" binding:"required"`
		Contact string `json:"contact" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	advertiser, err := ac.AdvertiserService.CreateAdvertiser(req.Name, req.Contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, advertiser)
}

// ListAdvertisers 列出所有广告主
func (ac *AdvertiserController) ListAdvertisers(c *gin.Context) {
	advertisers, err := ac.AdvertiserService.ListAdvertisers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, advertisers)
}

// GetAdvertiser 获取特定广告主的详情
func (ac *AdvertiserController) GetAdvertiser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid advertiser ID"})
		return
	}

	advertiser, err := ac.AdvertiserService.GetAdvertiserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Advertiser not found"})
		return
	}

	c.JSON(http.StatusOK, advertiser)
}

// UpdateAdvertiser 更新广告主的信息
func (ac *AdvertiserController) UpdateAdvertiser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid advertiser ID"})
		return
	}

	var req struct {
		Name    string `json:"name" binding:"required"`
		Contact string `json:"contact" binding:"required"`
		Status  int    `json:"status" binding:"required"` // 0: 禁用, 1: 可用
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取现有广告主
	advertiser, err := ac.AdvertiserService.GetAdvertiserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Advertiser not found"})
		return
	}

	// 更新字段
	advertiser.Name = req.Name
	advertiser.Contact = req.Contact
	advertiser.Status = req.Status

	updatedAdvertiser, err := ac.AdvertiserService.UpdateAdvertiser(advertiser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAdvertiser)
}

// DeleteAdvertiser 删除一个广告主
func (ac *AdvertiserController) DeleteAdvertiser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid advertiser ID"})
		return
	}

	err = ac.AdvertiserService.DeleteAdvertiser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Advertiser deleted successfully"})
}
