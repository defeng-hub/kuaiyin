package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BusDataCleanGetPageReq struct {
	dto.Pagination     `search:"-"`
    TaskName string `form:"taskName"  search:"type:contains;column:task_name;table:bus_data_clean" comment:"任务名称"`
    Status string `form:"status"  search:"type:exact;column:status;table:bus_data_clean" comment:"状态"`
    BusDataCleanOrder
}

type BusDataCleanOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:bus_data_clean"`
    TaskName string `form:"taskNameOrder"  search:"type:order;column:task_name;table:bus_data_clean"`
    FileName string `form:"fileNameOrder"  search:"type:order;column:file_name;table:bus_data_clean"`
    FilePath string `form:"filePathOrder"  search:"type:order;column:file_path;table:bus_data_clean"`
    RstPath string `form:"rstPathOrder"  search:"type:order;column:rst_path;table:bus_data_clean"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:bus_data_clean"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bus_data_clean"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bus_data_clean"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:bus_data_clean"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:bus_data_clean"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:bus_data_clean"`
    
}

func (m *BusDataCleanGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BusDataCleanInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    TaskName string `json:"taskName" comment:"任务名称"`
    FileName string `json:"fileName" comment:"源文件名称"`
    FilePath string `json:"filePath" comment:"源文件地址"`
    RstPath string `json:"rstPath" comment:"目标文件地址"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *BusDataCleanInsertReq) Generate(model *models.BusDataClean)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.TaskName = s.TaskName
    model.FileName = s.FileName
    model.FilePath = s.FilePath
    model.RstPath = s.RstPath
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *BusDataCleanInsertReq) GetId() interface{} {
	return s.Id
}

type BusDataCleanUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    TaskName string `json:"taskName" comment:"任务名称"`
    FileName string `json:"fileName" comment:"源文件名称"`
    FilePath string `json:"filePath" comment:"源文件地址"`
    RstPath string `json:"rstPath" comment:"目标文件地址"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *BusDataCleanUpdateReq) Generate(model *models.BusDataClean)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.TaskName = s.TaskName
    model.FileName = s.FileName
    model.FilePath = s.FilePath
    model.RstPath = s.RstPath
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *BusDataCleanUpdateReq) GetId() interface{} {
	return s.Id
}

// BusDataCleanGetReq 功能获取请求参数
type BusDataCleanGetReq struct {
     Id int `uri:"id"`
}
func (s *BusDataCleanGetReq) GetId() interface{} {
	return s.Id
}

// BusDataCleanDeleteReq 功能删除请求参数
type BusDataCleanDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BusDataCleanDeleteReq) GetId() interface{} {
	return s.Ids
}
