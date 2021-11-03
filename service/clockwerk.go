package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/nightsilvertech/clockwerk/gvar"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"

	"github.com/nightsilvertech/clockwerk/executors"
	executorhttp "github.com/nightsilvertech/clockwerk/executors/http"
	pb "github.com/nightsilvertech/clockwerk/protocs/api/v1"
	_interfacerepo "github.com/nightsilvertech/clockwerk/repository/interface"
	_interface "github.com/nightsilvertech/clockwerk/service/interface"
	"github.com/robfig/cron/v3"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type clockwerk struct {
	repo _interfacerepo.Storage
	crn  *cron.Cron
}

func (c clockwerk) shellExecutor(scheduler *pb.Scheduler) (res string, err error) {
	log.Println("shell executor on development")
	return res, nil
}

func (c clockwerk) httpExecutor(scheduler *pb.Scheduler) (res string, err error) {
	switch scheduler.Method {
	case executorhttp.MethodPost:
		return executorhttp.Post(scheduler)
	case executorhttp.MethodPut:
		return executorhttp.Put(scheduler)
	case executorhttp.MethodGet:
		return executorhttp.Get(scheduler)
	case executorhttp.MethodDelete:
		return executorhttp.Delete(scheduler)
	default:
		return res, errors.New("unknown http method")
	}
}

func (c clockwerk) persistenceCheck(scheduler *pb.Scheduler) {
	if !scheduler.Persist {
		c.crn.Remove(cron.EntryID(scheduler.EntryId))
		err := c.repo.Rem(scheduler.Id, scheduler.EntryId)
		if err != nil {
			log.Println(fmt.Sprintf("not persist failed to delete scheduler with id %s and entry id %d", scheduler.Id, scheduler.EntryId))
		} else {
			log.Println(fmt.Sprintf("not persist success to delete scheduler with id %s", scheduler.Id))
		}
	}
}

func (c clockwerk) execution(scheduler *pb.Scheduler) {
	if !scheduler.Disabled {
		switch scheduler.Executor {
		case executors.HTTP:
			res, err := c.httpExecutor(scheduler)
			if err != nil {
				log.Println(fmt.Sprintf("failed to run http executor with scheduler id %s error %+v", scheduler.Id, err))
			} else {
				c.persistenceCheck(scheduler)
				log.Println(fmt.Sprintf("resposen from http executor with scheduler id %s response %s", scheduler.Id, res))
			}
		case executors.SHELL:
			res, err := c.shellExecutor(scheduler)
			if err != nil {
				log.Println(fmt.Sprintf("failed to run shell executor with scheduler id %s error %+v", scheduler.Id, err))
			} else {
				c.persistenceCheck(scheduler)
				log.Println(fmt.Sprintf("resposen from shell executor with scheduler id %s response %s", scheduler.Id, res))
			}
		default:
			_, err := c.DeleteScheduler(nil, &pb.SelectScheduler{Id: scheduler.Id, EntryId: scheduler.EntryId})
			if err != nil {
				log.Println(fmt.Sprintf("failed to delete scheduler with id %s", scheduler.Id))
			}
		}
	}
}

func (c clockwerk) verifyBasicAuth(username, password string) error {
	key := fmt.Sprintf("%s_%s", gvar.HashKeyMap, username)
	hashedPassword, ok := gvar.SyncMapHashStorage.Load(key)
	if !ok {
		return errors.New("username not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword.(string)), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (c clockwerk) GetSchedulers(ctx context.Context, _ *emptypb.Empty) (res *pb.Schedulers, err error) {
	res, err = c.repo.All()
	log.Println("success to get schedulers totals", len(res.Schedulers))
	return
}

func (c clockwerk) AddScheduler(ctx context.Context, scheduler *pb.Scheduler) (res *pb.Scheduler, err error) {
	err = c.verifyBasicAuth(scheduler.Username, scheduler.Password)
	if err != nil {
		return res, err
	}
	entry, err := c.crn.AddFunc(scheduler.Spec, func() {
		c.execution(scheduler)
	})
	if err != nil {
		return res, err
	}
	scheduler.Id = uuid.NewV4().String()
	scheduler.EntryId = int32(entry)
	scheduler.CreatedAt = time.Now().Unix()
	log.Println("success to add scheduler")
	return scheduler, c.repo.Set(scheduler)
}

func (c clockwerk) DeleteScheduler(ctx context.Context, selectScheduler *pb.SelectScheduler) (res *emptypb.Empty, err error) {
	err = c.verifyBasicAuth(selectScheduler.Username, selectScheduler.Password)
	if err != nil {
		return res, err
	}
	c.crn.Remove(cron.EntryID(selectScheduler.EntryId))
	log.Println("success to delete scheduler")
	return res, c.repo.Rem(selectScheduler.Id, selectScheduler.EntryId)
}

func (c clockwerk) ToggleScheduler(ctx context.Context, toggle *pb.SelectToggle) (res *emptypb.Empty, err error) {
	err = c.verifyBasicAuth(toggle.Username, toggle.Password)
	if err != nil {
		return res, err
	}
	scheduler, err := c.repo.Get(toggle.Id, toggle.EntryId)
	if err != nil {
		return res, err
	}
	c.crn.Remove(cron.EntryID(toggle.EntryId))
	err = c.repo.Rem(toggle.Id, toggle.EntryId)
	if err != nil {
		return res, err
	}
	scheduler.Disabled = toggle.Disabled
	entry, err := c.crn.AddFunc(scheduler.Spec, func() {
		c.execution(scheduler)
	})
	if err != nil {
		return res, err
	}
	scheduler.EntryId = int32(entry)
	log.Println("success to toggle scheduler")
	return res, c.repo.Set(scheduler)
}

func (c clockwerk) Backup(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	schedulerData, err := c.repo.All()
	if err != nil {
		return nil, err
	}
	for _, scheduler := range schedulerData.Schedulers {
		// create new scheduler with new entry id
		entry, err := c.crn.AddFunc(scheduler.Spec, func() {
			c.execution(scheduler)
		})
		if err != nil {
			return nil, err
		}
		// remove data with old entry id
		err = c.repo.Rem(scheduler.Id, scheduler.EntryId)
		if err != nil {
			return nil, err
		}
		// insert it to redis
		scheduler.EntryId = int32(entry)
		err = c.repo.Set(scheduler)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func NewClockwerk(crn *cron.Cron, repo _interfacerepo.Storage) _interface.Clockwerk {
	crn.Start()
	return &clockwerk{
		crn:  crn,
		repo: repo,
	}
}
