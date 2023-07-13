package manager

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var onceLoadDB sync.Once

type InfraManager interface {
	GetDB() *sql.DB
}

type infraManager struct {
	db *sql.DB
}

var onceLoadDB sync.Once

func (i *infraManager) initDb() {
	psqlconn := fmt.Sprintf("user=%s host=%s password=%s dbname=%s sslmode=disable", i.cfg.User, i.cfg.Host, i.cfg.Password, i.cfg.Name)
	onceLoadDB.Do(func() {
		db, err := sql.Open(i.cfg.Driver, psqlconn)
		if err != nil {
			log.Fatal("Cannot start app, error when connect to DB", err.Error())
		}

		im.db = db
	})
	return im.db
}

func NewInfraManager() InfraManager {
	return &infraManager{}
}
