package main

import (
	"fmt"
	"kasir-api/config"
	"kasir-api/database"
	"kasir-api/middlewares"
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
	viper.BindEnv("API_KEY")

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

	// load config with fallback to API_KEY (all uppercase)
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("API_KEY") // Direct fallback
	}
	if apiKey == "" {
		apiKey = viper.GetString("APIKey") // Try uppercase version
	}

	config := config.Config{
		Host:   viper.GetString("HOST"),
		Port:   viper.GetString("PORT"),
		DBConn: dbConn,
		APIKey: apiKey,
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Middleware setup
	apiKeyMiddleware := middlewares.APIKey(config.APIKey)
	
	// Setup routes (dependency injection dilakukan di dalam SetupRoutes)
	routes.SetupRoutes(db, apiKeyMiddleware)

	addr := config.Host + ":" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}