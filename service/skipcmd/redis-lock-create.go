package skipcmd

import (
	"context"
	"fmt"
	"time"
)

func (cmd *Command) CreateRedisLock(ctx context.Context, userID uint32) (string, error) {

	timeDelay := 100 * time.Millisecond
	maxRetry := 50
	expired := 2 * time.Minute

	for timeRetry := 0; timeRetry <= maxRetry; timeRetry++ {
		locked, token, err := cmd.redis.AcquireLock(ctx, userID, expired)
		if err != nil {
			return "", err
		}
		if !locked {
			// sleep and retry
			time.Sleep(timeDelay)
		} else {
			return token, nil
		}
	}
	return "", fmt.Errorf("failed to acquire lock, maximum retry reached for userID %d", userID)
}
