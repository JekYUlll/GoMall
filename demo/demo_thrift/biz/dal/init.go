package dal

import (
	"github.com/JekYUlll/GoMall/demo/demo_thrift/biz/dal/mysql"
	"github.com/JekYUlll/GoMall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
