package redis_swipe_repo

import (
	"context"
	"fmt"

	"github.com/kbiits/dealls-take-home-test/domain/repository"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
)

type RedisSwipeRepo struct {
	rdb *redis.Client
}

func NewRedisSwipeRepo(
	rdb *redis.Client,
) repository.SwipeCacheRepository {
	return &RedisSwipeRepo{
		rdb: rdb,
	}
}

func (r *RedisSwipeRepo) GetNonSwiped(ctx context.Context, date, swiperProfileID string, profileIDs []string) ([]string, error) {
	keyBF := buildBloomFilterCacheKey(swiperProfileID, date)

	profileIDsAny := lo.Map(profileIDs, func(id string, _ int) interface{} {
		return id
	})

	results := r.rdb.BFMExists(ctx, keyBF, profileIDsAny...)
	if results.Err() != nil {
		return nil, results.Err()
	}

	nonSwipedProfiles := make([]string, 0)
	for i, result := range results.Val() {
		if !result {
			nonSwipedProfiles = append(nonSwipedProfiles, profileIDs[i])
		}
	}

	return nonSwipedProfiles, nil
}

// bloom filter is used to store swiped profiles
func (r *RedisSwipeRepo) MarkAsSwiped(ctx context.Context, date, swiperProfileID, profileID string) (bool, error) {
	keyBF := buildBloomFilterCacheKey(swiperProfileID, date)
	boolRes := r.rdb.BFAdd(ctx, keyBF, profileID)
	if boolRes.Err() != nil {
		return false, boolRes.Err()
	}

	return boolRes.Val(), nil
}

func (r *RedisSwipeRepo) ClearYesterdayBloomFilter(ctx context.Context, date string) error {
	keysYesterday := r.rdb.Keys(ctx, "swiped_by:*:"+date).Val()
	if len(keysYesterday) == 0 {
		return nil
	}

	_, err := r.rdb.Del(ctx, keysYesterday...).Result()
	if err != nil {
		return err
	}

	return nil
}

func buildBloomFilterCacheKey(userID, date string) string {
	return fmt.Sprintf("swiped_by:%s:%s", userID, date)
}
