package repository

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"SkipAdsV2/entities"
	"context"
	"log"
	"strconv"
	"time"
)

func (r *RepoMySQL) SeedPackage() {
	gameIds := []string{"1", "2", "3", "4", "5"}

	packages := []httpmodel.CreatePackageRequest{
		{
			PackageID:    "PKG001",
			Name:         "1",
			Quantity:     10,
			Type:         entities.PackageTypePurchase,
			ExpiresAfter: 30,
			Games:        gameIds,
		},
		{
			PackageID:    "PKG002",
			Name:         "2",
			Quantity:     20,
			Type:         entities.PackageTypePurchase,
			ExpiresAfter: 45,
			Games:        gameIds,
		},
		{
			PackageID:    "PKG003",
			Name:         "3",
			Quantity:     30,
			Type:         entities.PackageTypePurchase,
			ExpiresAfter: 60,
			Games:        gameIds,
		},
		{
			PackageID:    "PKG004",
			Name:         "4",
			Quantity:     1,
			Type:         entities.PackageTypeExchange,
			ExpiresAfter: 60,
			Games:        gameIds,
		},
		{
			PackageID:    "PKG005",
			Name:         "5",
			Quantity:     3,
			Type:         entities.PackageTypeExchange,
			ExpiresAfter: 60,
			Games:        gameIds,
		},
	}

	for _, pkg := range packages {
		p, g := pkg.ConvertToPackageAndPackageGames()
		r.CreatePackage(context.Background(), p, g)
	}
}

func (r *RepoMySQL) SeedSkipAds() {
	quantityMap := map[int]uint32{
		1: 10,
		2: 20,
		3: 30,
		4: 1,
		5: 3,
	}

	typeMap := map[int]entities.EventAddSkipAdsType{
		1: entities.EventAddSkipAdsPurchase,
		2: entities.EventAddSkipAdsPurchase,
		3: entities.EventAddSkipAdsPurchase,
		4: entities.EventAddSkipAdsExchange,
		5: entities.EventAddSkipAdsExchange,
	}
	pkg1 := "PKG001"
	pkg2 := "PKG002"
	pkg3 := "PKG003"
	pkg4 := "PKG004"
	pkg5 := "PKG005"
	packageMap := map[int]*string{
		1: &pkg1,
		2: &pkg2,
		3: &pkg3,
		4: &pkg4,
		5: &pkg5,
	}
	expireMap := map[int]int{
		1: 30,
		2: 45,
		3: 60,
		4: 60,
		5: 60,
	}
	// i: user ID
	// j: package ID
	// z: nums of event per user
	var batchSize = 1000
	var events []entities.EventAddSkipAds

	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 5; j++ {
			for z := 1; z <= 50; z++ {
				s := strconv.Itoa(i)
				event := entities.EventAddSkipAds{
					UserID:        s,
					PackageID:     packageMap[j],
					SourceEventID: "123",
					Quantity:      quantityMap[j],
					QuantityUsed:  0,
					Type:          typeMap[j],
					ExpiresAt:     time.Now().Add(time.Duration(expireMap[j]) * 24 * time.Hour),
				}
				event.SetPriority()
				events = append(events, event)

				// flush when batch >= batchSize
				if len(events) >= batchSize {
					if err := r.db.Create(&events).Error; err != nil {
						panic(err)
					}
					events = events[:0] // reset slice
					log.Println("current userID: ", i)
				}

			}
		}
	}

	// insert last batch
	if len(events) > 0 {
		if err := r.db.Create(&events).Error; err != nil {
			panic(err)
		}
	}
}
