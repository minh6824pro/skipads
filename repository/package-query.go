package repository

import (
	"SkipAdsV2/entities"
	"SkipAdsV2/errorcode"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (r *RepoMySQL) GetPurchasePackageByID(ctx context.Context, packageID *string) (entities.Package, error) {
	var pkg entities.Package

	err := r.db.WithContext(ctx).First(&pkg, "id= ? and type = ?", packageID, entities.EventAddSkipAdsPurchase).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errorcode.ErrorService{
				InternalError: errors.New("package doesn't exists"),
				ErrorCode:     errorcode.ErrInvalidRequest,
			}
		}
		return pkg, err
	}
	return pkg, nil
}

func (r *RepoMySQL) GetExchangePackageByID(ctx context.Context, packageID *string) (entities.Package, error) {
	var pkg entities.Package

	err := r.db.WithContext(ctx).First(&pkg, "id= ? and type = ?", packageID, entities.EventAddSkipAdsExchange).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = &errorcode.ErrorService{
				InternalError: errors.New("package doesn't exists"),
				ErrorCode:     errorcode.ErrInvalidRequest,
			}
		}
		return pkg, err
	}
	return pkg, nil
}
