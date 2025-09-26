package redis_service

import (
	"context"
	"time"
)

// StartRedisHealthCheck run goroutine ping Redis
func (r *RedisService) StartRedisHealthCheck(ctx context.Context, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// Ping Redis
				err := r.RedisClient.Ping(ctx).Err()
				if err != nil {
					r.IsAlive.Store(false)
					//log.Println("Redis service is not alive:", err.Error())
				} else {
					r.IsAlive.Store(true)
					//log.Println("Redis service is alive:")
				}
			}
		}
	}()
}
