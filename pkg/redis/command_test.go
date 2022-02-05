package redis

import (
	"testing"
)

func TestRedisCli_Append(t *testing.T) {
	if err := InitEnv(); err != nil {
		panic(err)
	}

	RxdCli.Del("TestRedisCli_Append")
	defer RxdCli.Del("TestRedisCli_Append")

	_, err := RxdCli.Set("TestRedisCli_Append", "test")
	if err != nil {
		t.Error(err.Error())
	}
	result, _ := RxdCli.Get("TestRedisCli_Append")
	if result != "test" {
		t.Error(result)
	}
	result, _ = RxdCli.Get("TestRedisCli_Append_Nil")
	if result != "" {
		t.Error(result)
	}
}
