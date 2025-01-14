// 自动生成模板InterfaceTemplate
package interfacecase

import (
	"github.com/test-instructor/cheetah/server/global"
	"github.com/test-instructor/cheetah/server/model/interfacecase/customType"
	"github.com/test-instructor/cheetah/server/model/system"
	"gorm.io/datatypes"
)

type ApiType int

const (
	ApiTypeTemplate ApiType = 1 // 接口模板
	ApiTypeCase     ApiType = 2 // 测试用例
)

type transactionType string

const (
	transactionStart transactionType = "start"
	transactionEnd   transactionType = "end"
)

// ApiStep 结构体
// 如果含有time.Time 请自行import time包
type ApiStep struct {
	global.GVA_MODEL
	Name          string              `json:"name" form:"name" gorm:"column:name;comment:接口名称"`
	ApiType       ApiType             `json:"type" form:"type" gorm:"column:api_type;comment:接口名称"`
	RequestID     uint                `json:"-"`
	Request       *ApiRequest         `json:"request" form:"request"`
	TestCase      interface{}         `json:"testcase,omitempty" yaml:"testcase,omitempty" gorm:"-"` // *TestCasePath or *TestCase
	Transaction   *ApiStepTransaction `json:"transaction,omitempty" yaml:"transaction,omitempty"`
	Rendezvous    *ApiStepRendezvous  `json:"rendezvous,omitempty" yaml:"rendezvous,omitempty"`
	ThinkTime     *ApiStepThinkTime   `json:"think_time,omitempty" yaml:"think_time,omitempty"`
	ThinkTimeID   uint
	TransactionID uint
	RendezvousID  uint

	Variables       datatypes.JSONMap      `json:"variables" form:"variables" gorm:"column:variables;comment:;type:text"`
	Extract         datatypes.JSONMap      `json:"extract" form:"extract" gorm:"column:extract;comment:;type:text"`
	Validate        customType.TypeArgsMap `json:"validate" form:"validate" gorm:"column:validate;comment:;type:text"`
	ValidateNumber  uint                   `json:"validate_number" form:"validate_number"`
	ValidateJson    datatypes.JSON         `json:"validate_json" form:"validate_json" `
	ExtractJson     datatypes.JSON         `json:"extract_json" form:"extract_json"`
	VariablesJson   datatypes.JSON         `json:"variables_json" form:"variables_json"`
	Hooks           string                 `json:"hooks" form:"hooks" gorm:"column:hooks;"`
	SetupHooks      customType.TypeArgs    `json:"setup_hooks,omitempty" form:"setup_hooks,omitempty" gorm:"column:setup_hooks;type:text"`
	TeardownHooks   customType.TypeArgs    `json:"teardown_hooks,omitempty" form:"teardown_hooks,omitempty" gorm:"column:teardown_hooks;type:text"`
	ProjectID       uint                   `json:"-"`
	TTestCase       []ApiCaseStep          `json:"testCase" form:"testCase" gorm:"many2many:ApiCaseStepRelationship;"`
	Sort            uint                   `json:"sort" form:"sort" gorm:"column:sort;"`
	ExportHeader    datatypes.JSON         `json:"export_header" gorm:"column:export_header;comment:;type:text"`
	ExportParameter datatypes.JSON         `json:"export_parameter" gorm:"column:export_parameter;comment:;type:text"`

	Parent      uint           `json:"-"`
	Project     system.Project `json:"-"`
	ApiMenuID   uint           `json:"-"`
	ApiMenu     ApiMenu        `json:"-"`
	CreatedBy   system.SysUser `json:"-"`
	CreatedByID uint           `json:"-"`
	UpdateBy    system.SysUser `json:"-"`
	UpdateByID  uint           `json:"-"`
	DeleteBy    system.SysUser `json:"-"`
	DeleteByID  uint           `json:"-"`
}

type ApiStepTransaction struct {
	global.GVA_MODEL
	Name string          `json:"name" yaml:"name"`
	Type transactionType `json:"type" yaml:"type"`
}

type ApiStepRendezvous struct {
	global.GVA_MODEL
	Name    string  `json:"name" yaml:"name"`                           // required
	Percent float32 `json:"percent,omitempty" yaml:"percent,omitempty"` // default to 1(100%)
	Number  int64   `json:"number,omitempty" yaml:"number,omitempty"`
	Timeout int64   `json:"timeout,omitempty" yaml:"timeout,omitempty"` // milliseconds
}

type ApiStepThinkTime struct {
	global.GVA_MODEL
	Time float64 `json:"time" yaml:"time"`
}
