package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(start string, machineID int64) (err error) {
	var startTime time.Time
	startTime, err = time.Parse("2006-01-02", start)
	if err != nil {
		return
	}
	snowflake.Epoch = startTime.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

func Generate() int64 {
	return node.Generate().Int64()
}
