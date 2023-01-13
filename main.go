package main


import (
	"database/sql"
	"fmt"
	"formative-15/database"
	"github.com/gin-gonic/gin"
	"formative-15/controllers"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"os"
)

const(
	host = "localhost"
	port = 5432
	user = "mirza"
	password = "12Dwiana!"
	dbname = "practice"
)

var (
	DB *sql.DB
	err error
)

func main()  {

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed to load config")
	} else {
		fmt.Println("success loaded config")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("PGHOST"), 
	os.Getenv("PGPORT"), 
	os.Getenv("PGUSER"), 
	os.Getenv("PGPASSWORD"), 
	os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB connection failed")
		panic(err)
	} else {
		fmt.Println("DB connection success")
	}

	defer DB.Close()

	database.DbMigrate(DB)


	// router gin
	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run(":"+os.Getenv("PGPORT"))

}

