package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBProvider is an abstraction over a database connection used for DI.
type DBProvider interface {
	DB() *gorm.DB
	Close() error
}

// GormProvider implements DBProvider using a *gorm.DB.
type GormProvider struct {
	db *gorm.DB
}

func (p *GormProvider) DB() *gorm.DB { return p.db }

func (p *GormProvider) Close() error {
	if p.db == nil {
		return nil
	}
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// NewGormProvider creates a DBProvider from a DSN.
func NewGormProvider(dsn string) (DBProvider, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// configure underlying sql.DB
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		if err := sqlDB.Ping(); err != nil {
			log.Printf("database ping failed: %v", err)
		}
	}
	return &GormProvider{db: db}, nil
}

// NewGormProviderFromEnv creates a DBProvider using env vars or GOOSE_DBSTRING.
func NewGormProviderFromEnv() (DBProvider, error) {
	dsn := os.Getenv("GOOSE_DBSTRING")
	if dsn == "" {
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		name := os.Getenv("DB_NAME")
		if port == "" {
			port = "5432"
		}
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, name)
	}
	return NewGormProvider(dsn)
}

// Convenience wrappers keeping previous API but returning a provider for DI.
func ConnectProviderFromEnv() (DBProvider, error) { return NewGormProviderFromEnv() }
