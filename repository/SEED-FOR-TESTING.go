package repository

import (
	"SkipAdsV2/entities"
	"log"
	"time"
)

func (r *RepoMySQL) SeedUser() {
	for i := 10001; i <= 100000; i++ {
		user := entities.User{
			UserID: uint32(i),
			Name:   "Test",
		}
		r.db.Create(&user)
	}
}

func (r *RepoMySQL) SeedSkipAds() {
	quantityMap := map[int]uint32{
		1: 10,
		2: 20,
		3: 30,
		4: 10,
		5: 20,
	}

	typeMap := map[int]entities.EventAddSkipAdsType{
		1: entities.EventAddSkipAdsPurchase,
		2: entities.EventAddSkipAdsPurchase,
		3: entities.EventAddSkipAdsPurchase,
		4: entities.EventAddSkipAdsExchange,
		5: entities.EventAddSkipAdsExchange,
	}

	expireMap := map[int]int{
		1: 7,
		2: 30,
		3: 90,
		4: 7,
		5: 7,
	}
	// i: user ID
	// j: package ID
	// z: nums of event per user
	var batchSize = 1000
	var events []entities.EventAddSkipAds

	for i := 10001; i <= 100000; i++ {
		for j := 1; j <= 5; j++ {
			for z := 1; z <= 10; z++ {
				packageId := uint32(j)
				event := entities.EventAddSkipAds{
					UserID:        uint32(i),
					PackageID:     &packageId,
					SourceEventID: 123,
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
