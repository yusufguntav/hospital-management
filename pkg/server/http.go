package server

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Depado/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yusufguntav/hospital-management/app/api/routes"
	"github.com/yusufguntav/hospital-management/pkg/config"
	"github.com/yusufguntav/hospital-management/pkg/database"
	"github.com/yusufguntav/hospital-management/pkg/domains/hospital"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func LaunchHttpServer(appc config.App, allows config.Allows) {
	log.Println("Starting HTTP Server...")
	gin.SetMode(gin.ReleaseMode)

	app := gin.New()
	app.Use(gin.LoggerWithFormatter(func(log gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s\"\n",
			log.TimeStamp.Format("2006-01-02 15:04:05"),
			log.ClientIP,
			log.Method,
			log.Path,
			log.Request.Proto,
			log.StatusCode,
			log.Latency,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(otelgin.Middleware(appc.Name))
	app.Use(cors.New(cors.Config{
		AllowMethods:     allows.Methods,
		AllowHeaders:     allows.Headers,
		AllowOrigins:     allows.Origins,
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	p := ginprom.New(
		ginprom.Engine(app),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
		ginprom.Ignore("/swagger/*any"),
	)
	app.Use(p.Instrument())

	db := database.DBClient()
	api := app.Group("/api/v1")

	hospital_repo := hospital.NewHospitalRepository(db)
	hospital_service := hospital.NewHospitalService(hospital_repo)
	routes.HospitalRoutes(api.Group("/hospital"), hospital_service)

	fmt.Println("Server is running on port " + appc.Port)
	app.Run(net.JoinHostPort(appc.Host, appc.Port))
}
