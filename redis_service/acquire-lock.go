package redis_service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// AcquireLock tries to set a lock with a unique token
func (r *RedisService) AcquireLock(ctx context.Context, userID uint32, expiration time.Duration) (bool, string, error) {
	key := fmt.Sprintf(UserSkipAdsLockKeyPattern, userID)
	token := uuid.NewString()
	ok, err := r.RedisClient.SetNX(ctx, key, token, expiration).Result()
	if err != nil {
		return false, "", err
	}
	if !ok {
		return false, "", nil
	}
	return true, token, nil
}
