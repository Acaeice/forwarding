package conf

import (
	"github.com/spf13/viper"
)

// GetNSQ 获取NSQ配置
func GetNSQ() string {
	return viper.GetString("nsq")
}
