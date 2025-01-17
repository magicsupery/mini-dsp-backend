package routers

import (
	"github.com/gin-gonic/gin"

	"mini-dsp-backend/controllers"
	"mini-dsp-backend/models"
	"mini-dsp-backend/repositories"
	"mini-dsp-backend/services"
	"mini-dsp-backend/utils"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 初始化数据库
	utils.InitDB()
	err := utils.DB.AutoMigrate(
		&models.User{},
		&models.Advertiser{},
		&models.Campaign{},
		&models.Creative{},
	)
	if err != nil {
		return nil
	}

	// 初始化 Repositories
	advRepo := repositories.NewAdvertiserRepo()
	campRepo := repositories.NewCampaignRepo()
	creatRepo := repositories.NewCreativeRepo()

	// 初始化 Services
	advertiserService := services.NewAdvertiserService(advRepo)
	campaignService := services.NewCampaignService(campRepo)
	creativeService := services.NewCreativeService(creatRepo)

	// 初始化 Controllers
	advCtrl := controllers.NewAdvertiserController(advertiserService)
	campCtrl := controllers.NewCampaignController(campaignService)
	creatCtrl := controllers.NewCreativeController(creativeService)

	// 路由绑定
	advertiserGroup := r.Group("/advertisers")
	{
		advertiserGroup.POST("", advCtrl.CreateAdvertiser)
		advertiserGroup.GET("", advCtrl.ListAdvertisers)
		advertiserGroup.GET("/:id", advCtrl.GetAdvertiser)
		advertiserGroup.PUT("/:id", advCtrl.UpdateAdvertiser)
		advertiserGroup.DELETE("/:id", advCtrl.DeleteAdvertiser)
	}

	campaignGroup := r.Group("/campaigns")
	{
		campaignGroup.POST("", campCtrl.CreateCampaign)
		campaignGroup.GET("", campCtrl.ListCampaigns)
		campaignGroup.GET("/:id", campCtrl.GetCampaign)
		campaignGroup.PUT("/:id", campCtrl.UpdateCampaign)
		campaignGroup.DELETE("/:id", campCtrl.DeleteCampaign)
	}

	creativeGroup := r.Group("/creatives")
	{
		creativeGroup.POST("", creatCtrl.CreateCreative)
		creativeGroup.GET("", creatCtrl.ListCreatives)
		creativeGroup.GET("/:id", creatCtrl.GetCreative)
		creativeGroup.PUT("/:id", creatCtrl.UpdateCreative)
		creativeGroup.DELETE("/:id", creatCtrl.DeleteCreative)
	}

	// 历史报表查询
	r.GET("/reports/hourly", controllers.GetHourlyReport)

	return r
}
