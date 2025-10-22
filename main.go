package main

import (
	"context"
	"market/handle"
	"market/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phuslu/log"
)

func main() {
	log.DefaultLogger = log.Logger{
		Level:  log.InfoLevel,
		Caller: 1,
		Writer: &log.MultiEntryWriter{
			&log.ConsoleWriter{ColorOutput: true},
			// &log.FileWriter{Filename: "app.log", MaxSize: 100 << 20},
		},
	}
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Info().Msg("DSN Evnvironment variable not set, using default")
		dsn = "postgres://postgres:secretpassword@101.201.49.155:5432/mydatabase?sslmode=disable"
	}
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse DSN")
	}
	pgpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	queries := model.New(pgpool)
	router := handle.InitRouter(queries)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
