package setting

var Redis = struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}{
	Addr:     "127.0.0.1",
	Password: "",
	DB:       0,
	PoolSize: 100,
}

func loadRedis() {

}
