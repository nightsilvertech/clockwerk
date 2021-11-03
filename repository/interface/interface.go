package _interface

import (
	pb "github.com/nightsilvertech/clockwerk/protocs/api/v1"
)

type Storage interface {
	Set(scheduler *pb.Scheduler) error
	Rem(id string, entryID int32) error
	Get(id string, entryID int32) (*pb.Scheduler, error)
	All() (*pb.Schedulers, error)
}
