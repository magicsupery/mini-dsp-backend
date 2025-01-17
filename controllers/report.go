package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"mini-dsp-backend/utils"
)

func GetHourlyReport(c *gin.Context) {
	campaignID := c.Query("campaign_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if campaignID == "" || startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required params"})
		return
	}

	sqlStr := fmt.Sprintf(`
        SELECT
            stat_hour,
            campaign_id,
            SUM(impressions) AS impressions,
            SUM(clicks) AS clicks,
            SUM(installs) AS installs,
            SUM(pay_count) AS pay_count,
            SUM(pay_amount) AS pay_amount
        FROM ad_stats_hourly
        WHERE campaign_id = %s
          AND stat_hour >= '%s'
          AND stat_hour < '%s'
        GROUP BY stat_hour, campaign_id
        ORDER BY stat_hour ASC
    `, campaignID, startDate, endDate)

	rows, err := utils.DorisDB.Query(sqlStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	type HourlyStats struct {
		StatHour    time.Time `json:"stat_hour"`
		CampaignId  int64     `json:"campaign_id"`
		Impressions int64     `json:"impressions"`
		Clicks      int64     `json:"clicks"`
		Installs    int64     `json:"installs"`
		PayCount    int64     `json:"pay_count"`
		PayAmount   float64   `json:"pay_amount"`
	}
	var results []HourlyStats

	for rows.Next() {
		var s HourlyStats
		if err := rows.Scan(
			&s.StatHour,
			&s.CampaignId,
			&s.Impressions,
			&s.Clicks,
			&s.Installs,
			&s.PayCount,
			&s.PayAmount,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, s)
	}
	c.JSON(http.StatusOK, results)
}
