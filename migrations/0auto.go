package migrations

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 15/08/20 15.00
**/

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/mmuflih/envgo/conf"
)

type Migration struct {
	db *gorm.DB
	c  conf.Config
}

func NewMigration(db *gorm.DB, c conf.Config) *Migration {
	return &Migration{db, c}
}

func (m Migration) Run() {
	env := m.c.GetString("env")
	if strings.Contains(env, "prod") {
		env = "production"
	}
	if strings.Contains(env, "stag") {
		env = "staging"
	}
	if strings.Contains(env, "dev") {
		env = "development"
	}
	log.Println("=>", "Running on", env, "environment")
	if env == "development" {
		log.Println("=>", "Disabled Auto Migrate on", env, "environment")
		return
	}
	log.Println("=>", "Auto Migrate")
	out, err := exec.Command("sql-migrate", "up", "--env="+env).Output()
	if err != nil {
		log.Println("<>", err, "Auto Migrate")
		return
	}
	fmt.Printf("%s", out)
}
