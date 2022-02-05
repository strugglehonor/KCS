package redis

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRedisCli_SAdd(t *testing.T) {
	if err := InitEnv(); err != nil {
		panic(err)
	}

	Convey("TestRedisCli_SAdd", t, func() {
		RxdCli.Del("TestRedisCli_SAdd")

		RxdCli.SAdd("TestRedisCli_SAdd", []string{"value1", "value2", "value3"})

		iCount, _ := RxdCli.SCard("TestRedisCli_SAdd")
		So(iCount, ShouldEqual, 3)

		bRet, _ := RxdCli.SIsMember("TestRedisCli_SAdd", "value1")
		So(bRet, ShouldEqual, true)

		bRet, _ = RxdCli.SIsMember("TestRedisCli_SAdd", "valuex")
		So(bRet, ShouldEqual, false)

		So(3, ShouldEqual, 3)
		//RxdCli.Del("TestRedisCli_SAdd")
	})
}
