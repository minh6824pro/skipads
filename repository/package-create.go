package repository

import (
	"SkipAdsV2/entities"
	"context"
)

func (r *RepoMySQL) CreatePackage(ctx context.Context, pkg *entities.Package, games []*entities.PackageGame) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(pkg).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Create(games).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	return err
}
