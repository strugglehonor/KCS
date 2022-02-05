package redis

var RxdCli *RedisCli

// only for test
func InitEnv() error {
	url := "10.130.9.98"
	passwd := "xxxxxx"

	txOption := NewDefaultOption()
	txOption.URL = url
	txOption.Password = passwd
	txOption.DB = 0

	RxdCli = NewRedisCli(txOption)
	if RxdCli != nil {
		return nil
	}

	return nil
}
