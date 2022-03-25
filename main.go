// main.go
// A basic server

// main.go file belongs to the main package
package main

// the net/http and os packages are imported into the file
import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mkstyl3/linkedin_jobs/models"
	"github.com/mkstyl3/linkedin_jobs/routes"

	"github.com/mkstyl3/linkedin_jobs/utils"
	"github.com/rs/zerolog/log"
)

func InitWebServer() {
	sPort := os.Getenv("SERVPORT")
	utils.LoadTemplates("templates/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":"+sPort, nil)
}

func ConfigureLog() {
	log.Logger = log.With().Caller().Logger()
}

func LoadEnvVariables() error {
	return godotenv.Load()
}

func main() {
	if err := LoadEnvVariables(); err != nil {
		log.Fatal().Msg("Error loading .env file.")
		return
	}
	ConfigureLog()
	models.InitDB()
	// models.Db.AutoMigrate(&models.Job{})

	InitWebServer()

	// defer func() {
	// 	if error := recover(); error != nil {
	// 		fmt.Println("Error:", error)
	// 	}
	// }()
}
