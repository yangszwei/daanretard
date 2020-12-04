package registry

import (
	"daanretard/internal/infra/delivery"
	"daanretard/internal/infra/fbgraph"
	"daanretard/internal/infra/persistence"
	"daanretard/internal/infra/validator"
	"log"
)

// Start initialize all parts of the application and start it
func Start() {
	v := validator.New()
	cfg, err := loadConfig("", v, false)
	if err != nil {
		log.Fatalln("Envfile not found.")
	}
	if err := cfg.Validate(); err != nil {
		log.Fatalln("Invalid config file: ", err.Error())
	}
	db, err := persistence.Open(cfg.DbDsn)
	if err != nil {
		log.Fatalln("Failed to connect to database.")
	}
	fb := fbgraph.New(cfg.FbPageID, cfg.FbGraphAppID, cfg.FbGraphAppSecret)
	service := newService(newRepo(db), fb)
	server := delivery.NewServer(cfg.DataPath)
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
	v := validator.New()
	cfg, err := loadConfig("", v, true)
	if err != nil {
		log.Fatalln("Envfile not found.")
	}
	if err := cfg.Validate(); err != nil {
		log.Fatalln("Invalid config file: ", err.Error())
	}
	db, err := persistence.Open(cfg.DbDsn)
	if err != nil {
		log.Fatalln("Failed to connect to database.")
	}
	autoMigrate(db)
}
