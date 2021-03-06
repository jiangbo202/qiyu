/**
 * @Author: jiangbo
 * @Description:
 * @File:  module_code
 * @Version: 1.0.0
 * @Date: 2021/06/19 7:13 下午
 */

package errcode

var (
	ErrorGetTagListFail  = NewError(20010001, "获取标签列表失败")
	ErrorRetrieveTagFail = NewError(20010002, "获取标签失败")
	ErrorCreateTagFail   = NewError(20010003, "创建标签失败")
	ErrorUpdateTagFail   = NewError(20010004, "更新标签失败")
	ErrorDeleteTagFail   = NewError(20010005, "删除标签失败")
	ErrorCountTagFail    = NewError(20010006, "统计标签失败")
	ErrorUploadFileFail  = NewError(20030001, "上传文件失败")

	ErrorGetArticleListFail  = NewError(20010011, "获取文章列表失败")
	ErrorRetrieveArticleFail = NewError(20010012, "获取文章失败")
	ErrorCreateArticleFail   = NewError(20010013, "创建文章失败")
	ErrorUpdateArticleFail   = NewError(20010014, "更新文章失败")
	ErrorDeleteArticleFail   = NewError(20010015, "删除文章失败")
	ErrorCountArticleFail    = NewError(20010016, "统计文章失败")

	ErrorGetCategoryListFail  = NewError(20010021, "获取类别列表失败")
	ErrorRetrieveCategoryFail = NewError(20010022, "获取类别失败")
	ErrorCreateCategoryFail   = NewError(20010023, "创建类别失败")
	ErrorUpdateCategoryFail   = NewError(20010024, "更新类别失败")
	ErrorDeleteCategoryFail   = NewError(20010025, "删除类别失败")
	ErrorCountCategoryFail    = NewError(20010026, "统计类别失败")
)
