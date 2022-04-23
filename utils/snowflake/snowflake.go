package snowflake

import (
	"association/global"
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	// 格式化 1月2号下午3时4分5秒  2006年
	st, err = time.Parse("2006-01-02 15:03:04", startTime)
	if err != nil {
		global.ASS_LOG.Error("Time initialization failed ")
		return
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		global.ASS_LOG.Error("Snowflake ID node creation failed ")
		return
	}
	return
}

// GenID 获取生成的id
func GenID() int64 {
	return node.Generate().Int64()
}
