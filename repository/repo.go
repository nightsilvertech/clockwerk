package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	pb "gitlab.com/nbdgocean6/clockwerk/protocs/api/v1"
	_interface "gitlab.com/nbdgocean6/clockwerk/repository/interface"
	"sort"
)

type storage struct {
	cache *redis.Client
}

const (
	// redisSchedulerData scheduler:dade9eeb-115e-488d-b9f1-f0f0675e73d9:1 / scheduler:id:entry_id
	redisSchedulerData = `scheduler:%s:%d`
	// redisSchedulersData scheduler keys for select all data with this keys
	redisSchedulersData = `scheduler:*`
	// expiration never expired
	expiration = 0
)

func (s *storage) Set(scheduler *pb.Scheduler) error {
	ctx := context.Background()
	defer ctx.Done()
	var key = fmt.Sprintf(redisSchedulerData, scheduler.Id, scheduler.EntryId)
	bytes, err := json.Marshal(scheduler)
	if err != nil {
		return err
	}
	return s.cache.Set(ctx, key, string(bytes), expiration).Err()
}

func (s *storage) Rem(id string, entryID int32) error {
	ctx := context.Background()
	defer ctx.Done()
	var key = fmt.Sprintf(redisSchedulerData, id, entryID)
	return s.cache.Del(ctx, key).Err()
}

func (s *storage) Get(id string, entryID int32) (res *pb.Scheduler, err error) {
	ctx := context.Background()
	defer ctx.Done()
	var scheduler pb.Scheduler
	var key = fmt.Sprintf(redisSchedulerData, id, entryID)
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
		Password: pass,
		DB:       0,
	})
	return &storage{
		cache: rdb,
	}
}
