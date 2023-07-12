package manager

import (
	"database/sql"
	"fmt"
	"sync"

	"enigmacamp.com/final-project/team-4/track-prosto/config"
	_ "github.com/lib/pq"
)

type InfraManager interface {
	DbConn() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg config.Config
}

var onceLoadDB sync.Once

func (i *infraManager) initDb() {
	psqlconn := fmt.Sprintf("user=%s host=%s password=%s dbname=%s sslmode=disable", i.cfg.User, i.cfg.Host, i.cfg.Password, i.cfg.Name)
	onceLoadDB.Do(func() {
		db, err := sql.Open(i.cfg.Driver, psqlconn)
		if err != nil {
			panic(err)
		}
		i.db = db
	})
	fmt.Println("DB Connected")
}
func (i *infraManager) DbConn() *sql.DB {
	return i.db
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.initDb()
	return &infra
}
