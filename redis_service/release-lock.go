package redis_service

import (
	"context"
	"fmt"
)

const releaseScript = `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end
`

// ReleaseLock deletes the lock only if the token matches
func (r *RedisService) ReleaseLock(ctx context.Context, userID uint32, token string) error {
	key := fmt.Sprintf(UserSkipAdsLockKeyPattern, userID)
	res, err := r.RedisClient.Eval(ctx, releaseScript, []string{key}, token).Result()
	if err != nil {
		return err
	}
	if res.(int64) == 0 {
		// not released because token mismatch or key expired
	}
	return nil
}
