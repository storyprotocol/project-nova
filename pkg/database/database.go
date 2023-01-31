package database

import (
	backoff "github.com/cenkalti/backoff/v4"
	_ "github.com/lib/pq" // postgres driver
	"github.com/pkg/errors"
	"github.com/project-nova/backend/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

const (
	DefaultMaxRetry = 10
)

// NewGormDB connects to postgres database and returns gorm.DB
func NewGormDB(connStr string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Use default exponential backoff: https://github.com/cenkalti/backoff/blob/v4/exponential.go#L75
	retryErr := backoff.Retry(func() error {
		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
			Logger: glogger.Default.LogMode(glogger.Silent),
		})
		if err != nil {
			logger.Errorf("Failed to connect to db (%s): %v. Retrying...", connStr, err)
			return err
		}

		return nil
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), DefaultMaxRetry))

	if retryErr != nil {
		return nil, errors.Wrapf(retryErr, "failed to connect to db (%s)", connStr)
	}

	return db, nil
}
