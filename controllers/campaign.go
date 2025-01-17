package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mini-dsp-backend/services"
)

type CampaignController struct {
	CampaignService services.CampaignService
}

func NewCampaignController(service services.CampaignService) *CampaignController {
	return &CampaignController{CampaignService: service}
}

func (cc *CampaignController) CreateCampaign(c *gin.Context) {
	var req struct {
		AdvertiserID int64   `json:"advertiser_id"`
		Name         string  `json:"name"`
		Budget       float64 `json:"budget"`
		BidType      string  `json:"bid_type"`
		BidAmount    float64 `json:"bid_amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	campaign, err := cc.CampaignService.CreateCampaign(req.AdvertiserID, req.Name, req.Budget, req.BidAmount, req.BidType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

func (cc *CampaignController) ListCampaigns(c *gin.Context) {
	campaigns, err := cc.CampaignService.ListCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaigns)
}

func (cc *CampaignController) GetCampaign(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	campaign, err := cc.CampaignService.GetCampaignByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

func (cc *CampaignController) UpdateCampaign(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req struct {
		Name      string  `json:"name"`
		Budget    float64 `json:"budget"`
		BidType   string  `json:"bid_type"`
		BidAmount float64 `json:"bid_amount"`
		Status    int     `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	campaign, err := cc.CampaignService.GetCampaignByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	campaign.Name = req.Name
	campaign.Budget = req.Budget
	campaign.BidType = req.BidType
	campaign.BidAmount = req.BidAmount
	campaign.Status = req.Status

	updatedCampaign, err := cc.CampaignService.UpdateCampaign(campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCampaign)
}

func (cc *CampaignController) DeleteCampaign(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = cc.CampaignService.DeleteCampaign(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "campaign deleted"})
}
