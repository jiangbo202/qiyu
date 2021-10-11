/**
 * @Author: jiangbo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2021/06/19 3:54 下午
 */

package pkg

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp:=viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("blog_service/configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
