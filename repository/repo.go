package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/go-redis/redis/v8"
	pb "github.com/nightsilvertech/clockwerk/protocs/api/v1"
	_interface "github.com/nightsilvertech/clockwerk/repository/interface"
)

type storage struct {
	cache *redis.Client
}

const (
	// redisSchedulerData scheduler:dade9eeb-115e-488d-b9f1-f0f0675e73d9:0463ef5d-b676-40d5-b6c2-89e79e60a25f / scheduler:id:reference_id
	redisSchedulerData = `scheduler:%s:%s`
	// redisSchedulersData scheduler keys for select all data with this keys
	redisSchedulersData = `scheduler:*`
	// redisSchedulerRetryAttempts scheduler-retry-attempt:dade9eeb-115e-488d-b9f1-f0f0675e73d9:0463ef5d-b676-40d5-b6c2-89e79e60a25f / scheduler-retry-attempt:id:reference_id
	redisSchedulerRetryAttempts = `scheduler-retry-attempt:%s:%s`
	// redisSchedulerRetryAttemptsUsed scheduler-retry-attempt-used:dade9eeb-115e-488d-b9f1-f0f0675e73d9:0463ef5d-b676-40d5-b6c2-89e79e60a25f / scheduler-retry-attempt-used:id:reference_id
	redisSchedulerRetryAttemptsUsed = `scheduler-retry-attempt-used:%s:%s`
	// expiration never expired
	expiration = 0
	// defaultRedisDB default redis database
	defaultRedisDB = 0
)

func (s *storage) removeCredential(scheduler *pb.Scheduler) {
	scheduler.Username = ""
	scheduler.Password = ""
}

func (s *storage) RemRetryAttempts(id string, refID string) error {
	ctx := context.Background()
	defer ctx.Done()
	var key = fmt.Sprintf(redisSchedulerRetryAttempts, id, refID)
	return s.cache.Del(ctx, key).Err()
}

func (s *storage) SetRetryAttempts(id string, refID string, retryAttempts int32) error {
	ctx := context.Background()
	defer ctx.Done()
	var key = fmt.Sprintf(redisSchedulerRetryAttempts, id, refID)
	return s.cache.Set(ctx, key, retryAttempts, expiration).Err()
}

func (s *storage) GetRetryAttempts(id string, refID string) (int32, error) {
	ctx := context.Background()
	defer ctx.Done()
	var key = fmt.Sprintf(redisSchedulerRetryAttempts, id, refID)
	val, err := strconv.ParseInt(s.cache.Get(ctx, key).Val(), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

func (s *storage) RemRetryAttemptsUsed(id string, refID string) error {
	ctx := context.Background()
	defer ctx.Done()
	var key = fmt.Sprintf(redisSchedulerRetryAttemptsUsed, id, refID)
	return s.cache.Del(ctx, key).Err()
}

func (s *storage) SetRetryAttemptsUsed(id string, refID string, retryAttemptsUsed int32) error {
	ctx := context.Background()
	defer ctx.Done()
	var key = fmt.Sprintf(redisSchedulerRetryAttemptsUsed, id, refID)
	return s.cache.Set(ctx, key, retryAttemptsUsed, expiration).Err()
}

func (s *storage) GetRetryAttemptsUsed(id string, refID string) (int32, error) {
	ctx := context.Background()
	defer ctx.Done()
	var key = fmt.Sprintf(redisSchedulerRetryAttemptsUsed, id, refID)
	val, err := strconv.ParseInt(s.cache.Get(ctx, key).Val(), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

func (s *storage) Set(scheduler *pb.Scheduler) error {
	s.removeCredential(scheduler)
	ctx := context.Background()
	defer ctx.Done()
	var key = fmt.Sprintf(redisSchedulerData, scheduler.Id, scheduler.ReferenceId)
	bytes, err := json.Marshal(scheduler)
	if err != nil {
		return err
	}
	return s.cache.Set(ctx, key, string(bytes), expiration).Err()
}

func (s *storage) Rem(id string, refID string) error {
	ctx := context.Background()
	defer ctx.Done()
	var key = fmt.Sprintf(redisSchedulerData, id, refID)
	return s.cache.Del(ctx, key).Err()
}

func (s *storage) Get(id string, refID string) (res *pb.Scheduler, err error) {
	ctx := context.Background()
	defer ctx.Done()
	var scheduler pb.Scheduler
	var key = fmt.Sprintf(redisSchedulerData, id, refID)
	err = json.Unmarshal([]byte(s.cache.Get(ctx, key).Val()), &scheduler)
	if err != nil {
		return res, err
	}
	return &scheduler, nil
}

func (s *storage) All() (res *pb.Schedulers, err error) {
	ctx := context.Background()
	defer ctx.Done()
	var schedulers pb.Schedulers
	keys, err := s.cache.Keys(ctx, redisSchedulersData).Result()
	if err != nil {
		return res, err
	}
	for _, key := range keys {
		var scheduler pb.Scheduler
		err = json.Unmarshal([]byte(s.cache.Get(ctx, key).Val()), &scheduler)
		if err != nil {
			return res, err
		}
		schedulers.Schedulers = append(schedulers.Schedulers, &scheduler)
	}
	sort.Slice(schedulers.Schedulers, func(i, j int) bool {
		return schedulers.Schedulers[i].CreatedAt > schedulers.Schedulers[j].CreatedAt
	})
	return &schedulers, nil
}

func NewRepo(host, port, pass string) _interface.Storage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		DB:       defaultRedisDB,
		Password: pass,
	})
	return &storage{
		cache: rdb,
	}
}
