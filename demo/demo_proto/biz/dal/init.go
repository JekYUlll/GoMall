package dal

import (
	"github.com/JekYUlll/GoMall/demo/demo_proto/biz/dal/mysql"
	"github.com/JekYUlll/GoMall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
