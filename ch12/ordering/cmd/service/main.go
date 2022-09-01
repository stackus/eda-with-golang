package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"

	"eda-in-golang/internal/config"
	"eda-in-golang/internal/system"
	"eda-in-golang/internal/web"
	"eda-in-golang/ordering"
	"eda-in-golang/ordering/migrations"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("ordering exitted abnormally: %s\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}
	s := system.NewSystem(cfg)
	err = s.InitDB()
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(s.DB())
	err = s.MigrateDB(migrations.FS)
	if err != nil {
		return err
	}
	err = s.InitJS()
	if err != nil {
		return err
	}
	s.InitLogger()
	s.InitMux()
	s.InitRpc()
	s.InitWaiter()

	// Mount general web resources
	s.Mux().Mount("/", http.FileServer(http.FS(web.WebUI)))

	err = ordering.Root(s.Waiter().Context(), s)
	if err != nil {
		return err
	}

	fmt.Println("started ordering service")
	defer fmt.Println("stopped ordering service")

	s.Waiter().Add(
		s.WaitForWeb,
		s.WaitForRPC,
		s.WaitForStream,
	)

	// go func() {
	// 	for {
	// 		var mem runtime.MemStats
	// 		runtime.ReadMemStats(&mem)
	// 		m.logger.Debug().Msgf("Alloc = %v  TotalAlloc = %v  Sys = %v  NumGC = %v", mem.Alloc/1024, mem.TotalAlloc/1024, mem.Sys/1024, mem.NumGC)
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }()

	return s.Waiter().Wait()
}
