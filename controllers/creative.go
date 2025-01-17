package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mini-dsp-backend/services"
)

type CreativeController struct {
	CreativeService services.CreativeService
}

func NewCreativeController(service services.CreativeService) *CreativeController {
	return &CreativeController{CreativeService: service}
}

func (cc *CreativeController) CreateCreative(c *gin.Context) {
	var req struct {
		CampaignID     int64  `json:"campaign_id"`
		CreativeName   string `json:"creative_name"`
		CreativeType   string `json:"creative_type"`
		LandingPageUrl string `json:"landing_page_url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creative, err := cc.CreativeService.CreateCreative(req.CampaignID, req.CreativeName, req.CreativeType, req.LandingPageUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, creative)
}

func (cc *CreativeController) ListCreatives(c *gin.Context) {
	creatives, err := cc.CreativeService.ListCreatives()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, creatives)
}

func (cc *CreativeController) GetCreative(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	creative, err := cc.CreativeService.GetCreativeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Creative not found"})
		return
	}
	c.JSON(http.StatusOK, creative)
}

func (cc *CreativeController) UpdateCreative(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req struct {
		CreativeName   string `json:"creative_name"`
		CreativeType   string `json:"creative_type"`
		LandingPageUrl string `json:"landing_page_url"`
		Status         int    `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creative, err := cc.CreativeService.GetCreativeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Creative not found"})
		return
	}

	creative.CreativeName = req.CreativeName
	creative.CreativeType = req.CreativeType
	creative.LandingPageUrl = req.LandingPageUrl
	creative.Status = req.Status

	updatedCreative, err := cc.CreativeService.UpdateCreative(creative)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCreative)
}

func (cc *CreativeController) DeleteCreative(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = cc.CreativeService.DeleteCreative(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "creative deleted"})
}
