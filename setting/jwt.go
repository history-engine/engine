package setting

import "github.com/spf13/viper"

var (
	JwtKey    = "he-jwt-key"
	JwtSecret = []byte("ub9V6ntqsC57Uzp2")
)

func loadJwt() {
	v := viper.Sub("jwt")
	JwtKey = "he" + v.GetString("jwt_key")
	JwtSecret = []byte(v.GetString("jwt_secret"))
}
