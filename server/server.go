package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gautamb02/social-api/api/packages/users"
	"github.com/gautamb02/social-api/api/rest"
	"github.com/gautamb02/social-api/configreader"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	Config *configreader.Config
}

var (
	sqlDBS = map[string]*sql.DB{
		DB_SOCIALAPI: nil,
	}
	logger *log.Logger
)

func initLogger(appname string) (*os.File, error) {
	logFile, err := os.OpenFile(fmt.Sprintf("%s.log", appname), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %v", err)
	}
	logger = log.New(logFile, "[SOCIAL-API] ", log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(logFile) // Set default logger to write to file
	return logFile, nil
}

func setUpMySQLConnector(dbConfig configreader.MySQLConnectionConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %v", err)
	}

	db.SetMaxOpenConns(dbConfig.MaxOpenConnections)
	db.SetMaxIdleConns(dbConfig.MaxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(dbConfig.MaxConnectionLifetime) * time.Second)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	logger.Println("Successfully connected to MySQL:", dbConfig.Database)
	return db, nil
}

func (s *Server) Setup() {
	cfg := s.Config

	logFile, err := initLogger(cfg.AppName)
	if err != nil {
		log.Fatal("Failed to initialize logger: ", err)
	}
	defer logFile.Close()

	logger.Println("🔧 Setting up MySQL connection...")
	sqlDBS[DB_SOCIALAPI], err = setUpMySQLConnector(cfg.DB.MySQl.SocialAPIDB)
	if err != nil {
		logger.Fatalf("Error connecting to %s database: %v", cfg.DB.MySQl.SocialAPIDB.Database, err)
	}
}

func (s *Server) Start() {
	logger.Println("🚀 Starting the server...")

	handlers := []rest.IHTTPHandlerProvider{
		users.NewUserModule(sqlDBS[DB_SOCIALAPI]),
	}

	router := rest.NewGRouter()
	router.Use(middleware.Logger)

	router.Route("/social", func(r rest.Router) {
		r.Register(handlers)
	})

	httpServer := &http.Server{
		Addr:    "0.0.0.0:1512",
		Handler: router,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
