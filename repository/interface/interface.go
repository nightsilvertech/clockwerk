package _interface

import (
	pb "github.com/nightsilvertech/clockwerk/protocs/api/v1"
)

type Storage interface {
	Set(scheduler *pb.Scheduler) error
	Rem(id string, refID string) error
	Get(id string, refID string) (*pb.Scheduler, error)
	All() (*pb.Schedulers, error)
	SetRetryAttempts(id string, refID string, retryAttempts int32) error
	GetRetryAttempts(id string, refID string) (int32, error)
	SetRetryAttemptsUsed(id string, refID string, retryAttemptsUsed int32) error
	GetRetryAttemptsUsed(id string, refID string) (int32, error)
	RemRetryAttempts(id string, refID string) error
	RemRetryAttemptsUsed(id string, refID string) error
}
