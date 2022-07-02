package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"cs101/pkg/telementry/logging"
	"cs101/pkg/telementry/logging/zap"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Config struct {
	Env     string
	Version string

	DB struct {
		Mysql struct {
			Addr string
			User string
			Pass string
			Name string
		}
	}

	Proxy string

	Tracing struct {
		Addr        string
		ServiceName string
	}

	MessageQueue struct {
		Kafka struct {
			GroupID string
			Brokers []string
			Topic   string
		}
	}

	Cache struct {
		Redis struct {
			Addr string
			Pass string
		}
	}
}

func (s *Server) init() {
	s.initLogger()
	s.initHTTP()
	checkErr := func(name string, err error) {
		if err != nil {
			s.logger.Fatalf("init %s failed: %+v", name, err)
		}
	}
	checkErr("database", s.initDB())

}

func (s *Server) initHTTP() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.GET("/health", func(c *gin.Context) { c.Writer.WriteHeader(200) })
	engine.GET("/info", func(c *gin.Context) {
		var response = struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		}{
			Name:   "cs101",
			Status: "ready",
		}
		s.logger.WithField("response", response).Info("/info")
		c.JSON(200, response)
	})

	s.http = &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
}

func (s *Server) initLogger() {
	c := zap.LocalConfig()
	if s.cfg.Env != "local" {
		c = zap.ReleaseConfig()
	}
	logger := zap.New(c)
	logging.SetDefaultLogger(logger)
	s.logger = logging.FromContext(context.Background())
}

func (s *Server) initDB() error {
	var (
		cfg   = s.cfg.DB.Mysql
		dbLog = fmt.Sprintf("addr=%s user=%s pass=%s db_name=%s", cfg.Addr, cfg.User, cfg.Pass, cfg.Name)
	)

	s.logger.Infof("Connecting to db %s", dbLog)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tpc(%s)/%s?parseTime=true&charset=utf8",
		cfg.User, cfg.Pass, cfg.Addr, cfg.Name))

	if err != nil {
		return errors.Wrapf(err, "open db: %s", dbLog)
	}

	if pingErr := db.Ping(); pingErr != nil {
		return errors.Wrapf(pingErr, "ping db: %s", dbLog)
	}

	s.db = db
	return nil
}
