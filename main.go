package main

import (
	"fmt"
	"kasir-api/config"
	"kasir-api/database"
	"kasir-api/routes"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	// load env
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	// load config
	config := config.Config{
		Host:   viper.GetString("HOST"),
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DBConn"),
	}

	// Manual override if viper fails to map automatically (common in some envs)
	if config.DBConn == "" {
		// Try to bind explicitly case-sensitive
		viper.BindEnv("DBConn")
		config.DBConn = viper.GetString("DBConn")
	}

	// Validate config
	if config.DBConn == "" {
		log.Fatal("Database connection string (DBConn) is empty. Check your .env file or environment variables.")
	}
	// Debug: log host dari DBConn saja (tanpa password) untuk memastikan .env terbaca
	if idx := strings.Index(config.DBConn, "@"); idx != -1 {
		log.Println("DB target:", config.DBConn[idx+1:])
	} else {
		log.Println("DBConn set (host tidak terdeteksi)")
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Setup routes (dependency injection dilakukan di dalam SetupRoutes)
	routes.SetupRoutes(db)

	addr := config.Host + ":" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}
