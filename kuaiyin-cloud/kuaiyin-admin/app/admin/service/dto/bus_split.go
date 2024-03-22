package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BusSplitGetPageReq struct {
	dto.Pagination     `search:"-"`
    TaskName string `form:"taskName"  search:"type:contains;column:task_name;table:bus_split" comment:"任务名称"`
    Status string `form:"status"  search:"type:exact;column:status;table:bus_split" comment:"状态"`
    BusSplitOrder
}

type BusSplitOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:bus_split"`
    TaskName string `form:"taskNameOrder"  search:"type:order;column:task_name;table:bus_split"`
    FilePath string `form:"filePathOrder"  search:"type:order;column:file_path;table:bus_split"`
    FileSize string `form:"fileSizeOrder"  search:"type:order;column:file_size;table:bus_split"`
    RstPath string `form:"rstPathOrder"  search:"type:order;column:rst_path;table:bus_split"`
    OutputName string `form:"outputNameOrder"  search:"type:order;column:output_name;table:bus_split"`
    SplitRow string `form:"splitRowOrder"  search:"type:order;column:split_row;table:bus_split"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:bus_split"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bus_split"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bus_split"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:bus_split"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:bus_split"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:bus_split"`
    
}

func (m *BusSplitGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BusSplitInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    TaskName string `json:"taskName" comment:"任务名称"`
    FilePath string `json:"filePath" comment:"源文件地址"`
    FileSize string `json:"fileSize" comment:"文件大小"`
    RstPath string `json:"rstPath" comment:"目标文件地址"`
    OutputName string `json:"outputName" comment:"拆分后名称"`
    SplitRow string `json:"splitRow" comment:"拆分列号"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *BusSplitInsertReq) Generate(model *models.BusSplit)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.TaskName = s.TaskName
    model.FilePath = s.FilePath
    model.FileSize = s.FileSize
    model.RstPath = s.RstPath
    model.OutputName = s.OutputName
    model.SplitRow = s.SplitRow
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *BusSplitInsertReq) GetId() interface{} {
	return s.Id
}

type BusSplitUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    TaskName string `json:"taskName" comment:"任务名称"`
    FilePath string `json:"filePath" comment:"源文件地址"`
    FileSize string `json:"fileSize" comment:"文件大小"`
    RstPath string `json:"rstPath" comment:"目标文件地址"`
    OutputName string `json:"outputName" comment:"拆分后名称"`
    SplitRow string `json:"splitRow" comment:"拆分列号"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *BusSplitUpdateReq) Generate(model *models.BusSplit)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.TaskName = s.TaskName
    model.FilePath = s.FilePath
    model.FileSize = s.FileSize
    model.RstPath = s.RstPath
    model.OutputName = s.OutputName
    model.SplitRow = s.SplitRow
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *BusSplitUpdateReq) GetId() interface{} {
	return s.Id
}

// BusSplitGetReq 功能获取请求参数
type BusSplitGetReq struct {
     Id int `uri:"id"`
}
func (s *BusSplitGetReq) GetId() interface{} {
	return s.Id
}

// BusSplitDeleteReq 功能删除请求参数
type BusSplitDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BusSplitDeleteReq) GetId() interface{} {
	return s.Ids
}
