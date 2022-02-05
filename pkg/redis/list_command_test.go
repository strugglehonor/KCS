package redis

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRedisCli_ListOption(t *testing.T) {
	if err := InitEnv(); err != nil {
		panic(err)
	}

	Convey("test", t, func() {
		RxdCli.Del("TestRedisCli_ListOption")
		iLen, _ := RxdCli.ListLen("TestRedisCli_ListOption")
		So(iLen, ShouldEqual, 0)

		RxdCli.ListLPush("TestRedisCli_ListOption", "value1")
		RxdCli.ListLPush("TestRedisCli_ListOption", "value2")
		RxdCli.ListLPush("TestRedisCli_ListOption", "value3")
		sRet, _ := RxdCli.ListIndex("TestRedisCli_ListOption", -1)
		So(sRet, ShouldEqual, "value1")

		sRet, _ = RxdCli.ListLPop("TestRedisCli_ListOption")
		So(sRet, ShouldEqual, "value3")

		RxdCli.Del("TestRedisCli_ListOption")

		dataStr := []string{"1", "2", "3"}
		_, err := RxdCli.ListRPush("testRpush", dataStr)
		So(err, ShouldBeNil)
		RxdCli.Del("testRpush")

		count := 0
		for {
			t.Log(RxdCli.ListBRPop("testBPOP", 3*time.Second))
			count += 1
			t.Log("tick ->", count)
			if count > 3 {
				break
			}
		}
	})
}
