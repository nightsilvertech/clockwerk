package _interface

import (
	pb "gitlab.com/nbdgocean6/clockwerk/protocs/api/v1"
)

type Clockwerk interface {
	pb.ClockwerkServer
}
