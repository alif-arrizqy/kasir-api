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

	// Bind environment variables explicitly
	viper.BindEnv("HOST")
	viper.BindEnv("PORT")
	viper.BindEnv("DBCONN") // Railway expects all uppercase

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	// load config with fallback to DBCONN (all uppercase)
	dbConn := viper.GetString("DBConn")
	if dbConn == "" {
		dbConn = viper.GetString("DBCONN") // Try uppercase version
	}
	if dbConn == "" {
		dbConn = os.Getenv("DBCONN") // Direct fallback
	}

	config := config.Config{
		Host:   viper.GetString("HOST"),
		Port:   viper.GetString("PORT"),
		DBConn: dbConn,
	}

	// Validate config
	if config.DBConn == "" {
		log.Fatal("Database connection string (DBCONN/DBConn) is empty. Check your environment variables")
	}
	
	// Debug: log host dari DBConn saja (tanpa password) untuk memastikan env terbaca
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