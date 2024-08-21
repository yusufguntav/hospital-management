package cmd

import (
	"github.com/yusufguntav/hospital-management/pkg/cache"
	c "github.com/yusufguntav/hospital-management/pkg/config"
	"github.com/yusufguntav/hospital-management/pkg/database"
	"github.com/yusufguntav/hospital-management/pkg/server"
	"github.com/yusufguntav/hospital-management/pkg/utils"
)

func StartApp() {
	config := c.InitConfig()
	utils.LoadEnvs()
	database.InitDB(config.Database)
	cache.InitRedis(config.Redis)
	server.LaunchHttpServer(config.App, config.Allows)
}
