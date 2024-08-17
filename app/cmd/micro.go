package cmd

import (
	"github.com/yusufguntav/hospital-management/pkg/config"
	"github.com/yusufguntav/hospital-management/pkg/database"
	"github.com/yusufguntav/hospital-management/pkg/server"
	"github.com/yusufguntav/hospital-management/pkg/utils"
)

func StartApp() {
	config := config.InitConfig()
	utils.LoadEnvs()
	database.InitDB(config.Database)
	server.LaunchHttpServer(config.App, config.Allows)

}
