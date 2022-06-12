package userconfig

import "github.com/spf13/viper"

const (
	prefixKey = "user_config."
)

func Get(key string) string {
	return viper.GetString(prefixKey + key)
}

func Set(key, value string) error {
	viper.Set(prefixKey+key, value)
	return viper.WriteConfig()
}
