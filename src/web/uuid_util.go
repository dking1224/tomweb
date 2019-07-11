package web

import (
	"github.com/sony/sonyflake"
)

func GetUUID() uint64 {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	uuid, err := flake.NextID()
	if err != nil {
	}
	return uuid
}
