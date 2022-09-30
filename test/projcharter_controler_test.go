package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"projectcharter/helper"
	"projectcharter/middleware"
	"projectcharter/models/domain"
	"projectcharter/routes"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = setupTestDB()
)

func setupTestDB() *gorm.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	db.AutoMigrate(&domain.ProjectCharter{})
	return db
}

func CloseTestDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	helper.PanicIfError(err)
	dbSQL.Close()
}

func TruncateTable(db *gorm.DB, tableName string) {
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", tableName))
}

func SetupRouter() *gin.Engine {
	err := godotenv.Load()
	helper.PanicIfError(err)
	/**
	@description Setup Database Connection
	*/

	/**
	@description Init Router
	*/
	router := gin.Default()
	/**
	@description Setup Mode Application
	*/
	if os.Getenv("GO_ENV") != "production" && os.Getenv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if os.Getenv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	/**
	@description Setup Middleware
	*/

	/**
	@description Init All Route
	*/
	routes.NewProjectCharterRoutes(db, router)
	router.Use(middleware.ErrorHandler())
	router.Use(cors.Default())

	return router
}

func TestCreatePreojectCharterSuccess(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)
	TruncateTable(db, "user")

	router := SetupRouter()

	requestBody := strings.NewReader(`
		"name": "Test",
		"description": "First name",
		"updatedby": "test 1",
		"deletedby": "test 1",
	`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/pcharter", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	//router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusCreated, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])
	assert.Equal(t, "Test", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "First name", responseBody["data"].(map[string]interface{})["description"])
	assert.Equal(t, "test 1", responseBody["data"].(map[string]interface{})["updatedby"])
	assert.Equal(t, "test 1", responseBody["data"].(map[string]interface{})["deletedby"])
	fmt.Println()
}

func TestCreateUserFailed(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)

	router := SetupRouter()

	requestBody := strings.NewReader(`
		"name": "Test",
		"description": "First name",
		"updatedby": "test 1",
		"deletedby": "test 1",
	`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/user", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusBadRequest)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
	assert.Equal(t, "Genre", responseBody["errors"].([]interface{})[0].(map[string]interface{})["key"])
	assert.Equal(t, "Error : failed on the required tag", responseBody["errors"].([]interface{})[0].(map[string]interface{})["message"])
}
