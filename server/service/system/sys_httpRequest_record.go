package system

import (
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system"
)

type OperationRecordService struct{}

// 添加记录,system.SysHttpRecord 表中
func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysHttpRecord) (err error) {
	err = global.NBUCTF_DB.Create(&sysOperationRecord).Error
	return err
}
