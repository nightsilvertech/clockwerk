package _interface

import (
	pb "github.com/nightsilvertech/clockwerk/protocs/api/v1"
)

type Clockwerk interface {
	pb.ClockwerkServer
}
