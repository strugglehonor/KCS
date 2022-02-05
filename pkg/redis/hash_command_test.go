package redis

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRedisCli_HSet(t *testing.T) {
	if err := InitEnv(); err != nil {
		panic(err)
	}

	Convey("HSet", t, func() {
		if err := InitEnv(); err != nil {
			panic(err)
		}

		RxdCli.HSet("TestRedisCli_HSet", "test", "testvalue")

		RxdCli.HDel("TestRedisCli_HSet", "test")
		rst, _ := RxdCli.HGet("TestRedisCli_HSet", "test")
		So(rst, ShouldEqual, "")

		param := map[string]interface{}{
			"test1": 1000,
			"test2": "value2",
		}
		_, _ = RxdCli.HMSet("TestRedisCli_HSet", param)

		mapRst, _ := RxdCli.HGetAll("TestRedisCli_HSet")
		So(2, ShouldEqual, len(mapRst))
		So("1000", ShouldEqual, mapRst["test1"])
		So("value2", ShouldEqual, mapRst["test2"])

		bret, _ := RxdCli.HExists("TestRedisCli_HSet", "test1")
		So(bret, ShouldEqual, true)
		RxdCli.HDel("TestRedisCli_HSet", "test1")
		bret, _ = RxdCli.HExists("TestRedisCli_HSet", "test1")
		So(bret, ShouldEqual, false)
		RxdCli.Del("TestRedisCli_HSet")
	})
}

func TestRedisCli_HIncrBy(t *testing.T) {
	if err := InitEnv(); err != nil {
		panic(err)
	}

	Convey("TestRedisCli_HIncrBy", t, func() {
		RxdCli.Del("TestRedisCli_HIncrBy")
		RxdCli.HIncrBy("TestRedisCli_HIncrBy", "test", 1)
		rst, _ := RxdCli.HGet("TestRedisCli_HIncrBy", "test")

		So(rst, ShouldEqual, "1")
		RxdCli.HIncrBy("TestRedisCli_HIncrBy", "test", 1)
		rst, err := RxdCli.HGet("TestRedisCli_HIncrBy", "test")
		So(rst, ShouldEqual, "2")

		RxdCli.HIncrByFloat("TestRedisCli_HIncrBy", "test2", 1.2)
		rst, err = RxdCli.HGet("TestRedisCli_HIncrBy", "test2")
		So(rst, ShouldEqual, "1.2")

		slist, _ := RxdCli.HKeys("TestRedisCli_HIncrBy")

		So(slist[0], ShouldEqual, "test")
		So(slist[1], ShouldEqual, "test2")

		l, _ := RxdCli.HLen("TestRedisCli_HIncrBy")
		So(l, ShouldEqual, 2)

		slist, _ = RxdCli.HVals("TestRedisCli_HIncrBy")
		So(slist[0], ShouldEqual, "2")
		So(slist[1], ShouldEqual, "1.2")

		mapResult, _ := RxdCli.HMGet("TestRedisCli_HIncrBy", []string{"test", "test2"})
		So(mapResult["test"], ShouldEqual, "2")
		So(mapResult["test2"], ShouldEqual, "1.2")

		RxdCli.HSetNX("TestRedisCli_HIncrBy", "test3", "value3")
		RxdCli.Del("TestRedisCli_HIncrBy")

		value := map[string]interface{}{
			"user_relate": 1,
			"user_id":     1234,
			"machine":     "qwertyuiop",
			"test1":       "12312424534234234",
			"test2":       "jdksjfksdlfjwieurisdjfslkdjfslkdfjsdfsdfwerwe",
			"test3":       "jjxkdksieusdfhjaosdlfkjfwiefjkjfskdfwiejfdlsdjfsl",
			"test4":       "jjxkdksieusdfhjaosdlfkjfwiefjkjfskdfwiejfdlsdjfsl",
		}
		RxdCli.HMSet("TestHashBenchmark", value)
		rsp, err := RxdCli.HGet("TestHashBenchmark", "user_ixd")
		t.Log(rsp, " ", err)

		rsp, err = RxdCli.HGet("test:hashfile:11", "11ac04643f2b065b24f331f86a804aeb")
		t.Log(rsp, " ", err)

		result, e := RxdCli.HGetAll("heheda")
		t.Log(result, " ", e)
	})
}
