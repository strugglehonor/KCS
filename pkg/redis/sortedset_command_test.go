package redis

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRedisCli_ZAdd(t *testing.T) {
	if err := InitEnv(); err != nil {
		panic(err)
	}

	Convey("TestRedisCli_ZAdd", t, func() {
		RxdCli.Del("TestRedisCli_ZAdd")

		RxdCli.ZAdd("TestRedisCli_ZAdd", 1, "key1")
		RxdCli.ZAdd("TestRedisCli_ZAdd", 3, "key2")

		score, _ := RxdCli.ZScore("TestRedisCli_ZAdd", "key1")
		So(score, ShouldEqual, 1)

		RxdCli.ZIncrBy("TestRedisCli_ZAdd", 1, "key1")
		score, _ = RxdCli.ZScore("TestRedisCli_ZAdd", "key1")
		So(score, ShouldEqual, 2)

		rst, _ := RxdCli.ZRangeByScore("TestRedisCli_ZAdd", 0, -1)
		So(rst[0].Member, ShouldEqual, "key1")
		So(rst[0].Score, ShouldEqual, 2)

		r, _ := RxdCli.ZRank("TestRedisCli_ZAdd", "key1")
		So(r, ShouldEqual, 0)

		r, _ = RxdCli.ZRank("TestRedisCli_ZAdd", "keyx")
		So(r, ShouldEqual, -1)
	})
}
