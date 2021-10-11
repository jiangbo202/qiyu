/**
 * @Author: jiangbo
 * @Description:
 * @File:  module_code
 * @Version: 1.0.0
 * @Date: 2021/06/19 7:13 下午
 */

package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorRetrieveTagFail = NewError(20010001, "获取标签失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")
	ErrorUploadFileFail = NewError(20030001, "上传文件失败")

	ErrorGetArticleListFail = NewError(20010011, "获取文章列表失败")
	ErrorCreateArticleFail  = NewError(20010012, "创建文章失败")
	ErrorCountArticleFail   = NewError(20010015, "统计文章失败")
)

