package sqlite

import (
	"RestApiGolang/internal/config"
	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/jinzhu/gorm"
	"time"

	log "RestApiGolang/internal/logger"
	_ "github.com/mattn/go-sqlite3" // init db_sqlite3 driver
)

var Db *gorm.DB

func chooseDBDriver(name, openStr string) goose.DBDriver {
	d := goose.DBDriver{
		Name:    name,
		OpenStr: openStr,
		Import:  "github.com/mattn/go-sqlite3",
		Dialect: &goose.Sqlite3Dialect{},
	}

	return d
}

func Setup(c *config.Config) error {
	const op = "internal.database.sqlite.Setup"
	conf := c

	migrateConf := &goose.DBConf{
		MigrationsDir: conf.MigrationsPath,
		Env:           "production",
		Driver:        chooseDBDriver(conf.DatabaseName, conf.DatabasePath),
	}

	latest, err := goose.GetMostRecentDBVersion(migrateConf.MigrationsDir)
	if err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return err
	}

	i := 0
	for {
		Db, err = gorm.Open(conf.DatabaseName, conf.DatabasePath)
		if err == nil {
			break
		}
		if err != nil && i >= 10 {
			log.Errorf("%s: %s", op, err.Error())
			return err
		}
		i++
		log.Warn("waiting for database to be up...")
		time.Sleep(5 * time.Second)
	}
	Db.LogMode(false)
	Db.SetLogger(log.Logger)
	Db.DB().SetMaxOpenConns(1)
	if err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return err
	}
	err = goose.RunMigrationsOnDb(migrateConf, migrateConf.MigrationsDir, latest, Db.DB())
	if err != nil {
		log.Errorf("%s: %s", op, err.Error())
		return err
	}

	return nil
}
