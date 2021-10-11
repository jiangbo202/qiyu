/**
 * @Author: jiangbo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2021/06/19 4:03 下午
 */

package global

import (
	"jiangbo.com/blog_service/pkg"
	"jiangbo.com/blog_service/pkg/logger"
)

var (
	ServerSetting   *pkg.ServerSettingS
	AppSetting      *pkg.AppSettingS
	DatabaseSetting *pkg.DatabaseSettingS

	Logger          *logger.Logger
	JWTSetting          *pkg.JWTSetting
)
