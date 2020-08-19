package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCountTagFail   = NewError(20010002, "统计标签失败")
	ErrorCreateTagFail  = NewError(20010003, "新增标签失败")
	ErrorUpdateTagFail  = NewError(20010004, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010005, "删除标签失败")
)
