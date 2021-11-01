package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/nbdgocean6/clockwerk/executors"
	executorhttp "gitlab.com/nbdgocean6/clockwerk/executors/http"
	pb "gitlab.com/nbdgocean6/clockwerk/protocs/api/v1"
	_interfacerepo "gitlab.com/nbdgocean6/clockwerk/repository/interface"
	_interface "gitlab.com/nbdgocean6/clockwerk/service/interface"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
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
		_, err := c.DeleteScheduler(nil, &pb.SelectScheduler{Id: scheduler.Id, EntryId: scheduler.EntryId})
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

func (c clockwerk) GetSchedulers(_ context.Context, _ *emptypb.Empty) (res *pb.Schedulers, err error) {
	res, err = c.repo.All()
	log.Println("success to get schedulers totals", len(res.Schedulers))
	return
}

func (c clockwerk) AddScheduler(_ context.Context, scheduler *pb.Scheduler) (res *pb.Scheduler, err error) {
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

func (c clockwerk) DeleteScheduler(_ context.Context, selectScheduler *pb.SelectScheduler) (res *emptypb.Empty, err error) {
	c.crn.Remove(cron.EntryID(selectScheduler.EntryId))
	log.Println("success to delete scheduler")
	return res, c.repo.Rem(selectScheduler.Id, selectScheduler.EntryId)
}

func (c clockwerk) ToggleScheduler(_ context.Context, toggle *pb.SelectToggle) (res *emptypb.Empty, err error) {
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
	return nil, nil
}

func NewClockwerk(crn *cron.Cron, repo _interfacerepo.Storage) _interface.Clockwerk {
	crn.Start()
	return &clockwerk{
		crn:  crn,
		repo: repo,
	}
}
