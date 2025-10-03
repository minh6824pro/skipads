package httpmodel

import (
	"SkipAdsV2/entities"
	"SkipAdsV2/errorcode"
	"errors"
)

type CreatePackageRequest struct {
	PackageID    string               `json:"package_id" binding:"required"`
	Name         string               `json:"name" binding:"required"`
	Quantity     uint32               `json:"quantity" binding:"required"`
	Type         entities.PackageType `json:"type" binding:"required"`
	ExpiresAfter uint32               `json:"expires_after" binding:"required"`
	Games        []string             `json:"games" binding:"required"`
}

func (r *CreatePackageRequest) Validate() error {
	if r.Type != entities.PackageTypePurchase && r.Type != entities.PackageTypeExchange {
		return &errorcode.ErrorService{
			InternalError: errors.New("package type must be purchase or exchange"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	return nil
}

func (r *CreatePackageRequest) ConvertToPackageAndPackageGames() (*entities.Package, []*entities.PackageGame) {
	pkg := &entities.Package{
		ID:           r.PackageID,
		Name:         r.Name,
		Quantity:     r.Quantity,
		Type:         r.Type,
		ExpiresAfter: r.ExpiresAfter,
	}
	var games []*entities.PackageGame
	for _, game := range r.Games {
		games = append(games, &entities.PackageGame{
			PackageID: r.PackageID,
			AppID:     game,
		})
	}
	return pkg, games
}

func (r *CreatePackageRequest) ConvertToPackageGames() []*entities.PackageGame {
	var games []*entities.PackageGame
	for _, game := range r.Games {
		games = append(games, &entities.PackageGame{
			PackageID: r.PackageID,
			AppID:     game,
		})
	}
	return games
}
