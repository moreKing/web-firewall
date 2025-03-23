package model

import (
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gmeta"
)

type PasswordComplex struct {
	Length      int `json:"length" v:"min:6|max:200"`           //密码长度
	Validity    int `json:"validity" v:"in:1,30,60,90,180,365"` //密码有效期：1-不限；具体天数。本地认证密码过期时间,仅支持 1 30 60 90 180 365 选择一个值
	Expire      int `json:"expire" v:"in:0,1"`                  //密码过期处理：0-禁止登录；1-仅提醒'
	DifferTimes int `json:"differTimes" v:"min:1|max:30"`       //不能与前多少次密码一致 1-30
	Complex     int `json:"complex" v:"in:1,2,3,4"`             //本地密码复杂度 1. 不限制  2.至少包含数字、字母 3.至少包含字母、数字、特殊字符 4.至少包含大小写字母、数字、特殊字符
}

type SystemConfPasswordComplex struct {
	gmeta.Meta `orm:"table:system_conf"`
	*entity.SystemConf
	Config PasswordComplex
}

// AuthenticateConf 认证配置
type AuthenticateConf struct {
	TotpOffset    int    `json:"totpOffset" v:"min:1|max:30"`    // totp动态码有效期
	TotpIssuer    string `json:"totpIssuer"`                     // totp标识
	MessageOffset int    `json:"messageOffset" v:"min:1|max:30"` // 短信验证码有效期
	EmailOffset   int    `json:"emailOffset" v:"min:1|max:30"`   // 邮件验证码有效期
}

type SystemConfAuthenticate struct {
	gmeta.Meta `orm:"table:system_conf"`
	*entity.SystemConf
	Config AuthenticateConf
}

type WebSession struct {
	Timeout int `json:"timeout" v:"min:10|max:1440"` // web 登录超时时间
}

type SystemConfWebSession struct {
	gmeta.Meta `orm:"table:system_conf"`
	*entity.SystemConf
	Config WebSession
}

// Message 系统短信认证配置
type Message struct {
	State      int    `json:"state" v:"in:0,1,2#状态 0禁用 1 内置 2 python 中的一个"`
	URL        string `json:"url" v:"required-if:state,1|url"`
	Method     string `json:"method" v:"required-if:state,1"`
	EncType    string `json:"encType" v:"required-if:state,1"`
	Parameters []struct {
		Key   string `json:"key" v:"required#请求参数变量名不能为空"`
		Value any    `json:"value" v:"required#请求参数值不能为空"`
	} `json:"parameters"`
	Content string `json:"content" v:"required-if:state,1"`
}

type SystemConfMessage struct {
	gmeta.Meta `orm:"table:system_conf"`
	*entity.SystemConf
	Config Message
}

// Email 系统邮箱认证配置
type Email struct {
	Enable   bool   `json:"enable"`
	SMTP     string `json:"smtp" v:"required-if:enable,true"` // smtp地址
	Port     int    `json:"port" `
	Email    string `json:"email" v:"required-if:enable,true|email"`
	Account  string `json:"account" v:"required-if:enable,true"` // 发送邮件时显示的名称
	Protocol int    `json:"protocol" v:"required-if:enable,true"`
	Password string `json:"password"`
}

type SystemConfEmail struct {
	gmeta.Meta `orm:"table:system_conf"`
	*entity.SystemConf
	Config Email
}

// FileServer 文件服务器 sftp/sftp
type FileServer struct {
	Protocol int    `json:"protocol" v:"in:0,1,2"` // 0 禁用 1.ftp 2 sftp
	Addr     string `json:"addr" v:"required-unless:protocol,0"`
	Port     int    `json:"port" `
	WorkDir  string `json:"workDir" v:"required-unless:protocol,0"`
	Account  string `json:"account" v:"required-unless:protocol,0"`

	Auth      int    `json:"auth" v:"required-unless:protocol,0"` // 0 密码 1.rsa 2 ed25519
	Password  string `json:"password,omitempty"  `
	Cipher    string `json:"cipher,omitempty" `    // 密钥密码 仅支持rsa
	SecretKey string `json:"secretKey,omitempty" ` // 密钥文件 支持rsa  ed25519
}

type SystemConfFileServer struct {
	gmeta.Meta `orm:"table:system_conf"`
	*entity.SystemConf
	Config FileServer
}

// AccountExceptionRule 异常账号规则
type AccountExceptionRule struct {
	Zombie     int `json:"zombie" v:"min:0|max:365"`     // 多少天没登录 判断是僵尸账号 0为跳过验证，最少30
	PwdTimeout int `json:"pwdTimeout" v:"min:0|max:365"` // 密码多久没有改 算超时 0为跳过验证，最少30
	PwdWeak    struct {
		Length int `json:"length" v:"min:5"` // 密码长度 最少5
		Digit  int `json:"digit" v:"min:0"`  // 数字数量
		Upper  int `json:"upper" v:"min:0"`  // 大写字母数量
		Lower  int `json:"lower" v:"min:0"`  // 小写字数
		Other  int `json:"other"  v:"min:0"` // 其他字符数
	} `json:"pwdWeak"` // 弱密码规则
}

type SystemConfAccountExceptionRule struct {
	gmeta.Meta `orm:"table:system_conf"`
	*entity.SystemConf
	Config AccountExceptionRule
}

type SystemStatus struct {
	CPUTotal          int     `json:"cpuTotal"`
	CpuLogicalTotal   int     `json:"cpuLogicalTotal"`
	CpuPercent        float64 `json:"cpuPercent"`
	MemoryTotal       uint64  `json:"memoryTotal"`
	MemoryUsed        uint64  `json:"memoryUsed"`
	MemoryUsedPercent float64 `json:"memoryUsedPercent"`
	Partitions        any     `json:"partitions"`
}

// 第三方堡垒机同步
type AppSync struct {
	Enable   bool   `json:"enable"`
	Company  string `json:"company" v:"required-if:enable,true"` // 1 齐治
	IP       string `json:"ip" v:"required-if:enable,true"`
	Port     int    `json:"port" v:"required-if:enable,true"`
	Account  string `json:"account" v:"required-if:enable,true"`
	Password string `json:"password" v:"required-if:enable,true"`
	//PushPwd    bool   `json:"pushPwd"`    // 改密后推送密码
	//Auto       bool   `json:"auto"`       // 自动执行
	//Processing bool   `json:"processing"` // 同步处理创建 删除资产/账号  仅同步资产信息，创建资产时同步账号信息，后续不再同步账号信息，保障唯一性
	//Interval   int    `json:"interval"`   // 同步间隔 单位：分钟 范围：10-1440 至少一天同步一次
}

type SystemConfAppSync struct {
	gmeta.Meta `orm:"table:system_conf"`
	*entity.SystemConf
	Config AppSync
}
