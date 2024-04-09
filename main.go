package main

import (
	"os"

	"github.com/nncoding/go-basic/config"
	"github.com/nncoding/go-basic/modules/servers"
	"github.com/nncoding/go-basic/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnection(cfg.Db())
	// defer คือทำงานก่อนที่ func จะ return กลับไป
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
