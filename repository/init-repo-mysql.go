package repository

import (
	"SkipAdsV2/config"
	"SkipAdsV2/entities"
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RepoMySQL struct {
	cfg config.Config
	db  *gorm.DB
}

var RepoMySQLInstance *RepoMySQL

func NewRepoMysql(cfg config.Config) (*RepoMySQL, error) {
	if RepoMySQLInstance != nil {
		return RepoMySQLInstance, nil
	}

	rp := &RepoMySQL{
		cfg: cfg,
		db:  nil,
	}
	err := rp.connection()
	if err != nil {
		return nil, err
	}
	return rp, nil
}

func (m *RepoMySQL) connection() error {
	ctx, cancel := context.WithTimeout(context.Background(), m.cfg.MySQL.TimeToConnect)
	defer cancel()

	database, err := gorm.Open(mysql.Open(m.cfg.MySQL.URI), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := database.DB()
	if err != nil {
		return err
	}
	// Check the connection
	err = sqlDB.PingContext(ctx)
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(int(m.cfg.MySQL.MaxIdleConnections))
	sqlDB.SetMaxOpenConns(int(m.cfg.MySQL.MaxConnections))
	sqlDB.SetConnMaxIdleTime(m.cfg.MySQL.TimeToConnect)
	sqlDB.SetConnMaxLifetime(m.cfg.MySQL.TimeToConnect)
	m.db = database
	return nil
}

func (m *RepoMySQL) InitTable() error {
	err := m.db.AutoMigrate(
		&entities.User{},
		&entities.Package{},
		&entities.EventAddSkipAds{},
		&entities.EventSubSkipAds{},
		&entities.EventAddSkipAdsArchive{},
	)
	if err != nil {
		return err
	}
	return nil
}
