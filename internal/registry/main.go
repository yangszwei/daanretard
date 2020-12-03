package registry

import (
	"daanretard/internal/infra/config"
	"daanretard/internal/infra/delivery"
	"daanretard/internal/infra/fbgraph"
	"daanretard/internal/infra/persistence"
	"log"
)

// Start initialize all parts of the application and start it
func Start() {
	cfg, err := config.Load("")
	if err != nil {
		log.Fatalln("Fail to start due to invalid configuration file.")
	}
	db, err := persistence.Open(cfg.DbDsn)
	if err != nil {
		log.Fatalln("Failed to connect to database.")
	}
	fb := fbgraph.New(cfg.FbPageID, cfg.FbAppID, cfg.FbAppSecret)
	service := newService(newRepo(db), fb)
	server := delivery.NewServer(cfg.Data)
	server.LoadTemplates()
	delivery.SetupSession(server, []byte(cfg.Secret))
	server.SetupRouterGroups()
	delivery.SetupUserAPI(server, service.user)
	if err := server.Run(cfg.Addr); err != nil {
		log.Fatalln(err)
	}
	server.Run(cfg.Addr)
}

// Init setup parts that are necessary for the app to start
func Init() {
	cfg, err := config.Load("")
	if err != nil {
		log.Fatalln("Fail to start due to invalid configuration file.")
	}
	db, err := persistence.Open(cfg.DbDsn)
	if err != nil {
		log.Fatalln("Failed to connect to database.")
	}
	autoMigrate(db)
}
