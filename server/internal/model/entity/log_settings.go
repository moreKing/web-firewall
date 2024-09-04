// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// LogSettings is the golang structure for table log_settings.
type LogSettings struct {
	Id            int64  `json:"id"            orm:"id"             ` // 主键
	Name          string `json:"name"          orm:"name"           ` // 接口名称
	Loginname     string `json:"loginname"     orm:"loginname"      ` // 登陆名
	Username      string `json:"username"      orm:"username"       ` // 用户名
	ClientIp      string `json:"clientIp"      orm:"client_ip"      ` // 登陆IP
	UserId        int64  `json:"userId"        orm:"user_id"        ` // 用户id
	Success       bool   `json:"success"       orm:"success"        ` // 操作成功为true，当响应码与code==0 时才为成功
	DepartmentId  int64  `json:"departmentId"  orm:"department_id"  ` // 部门id
	RequestMethod string `json:"requestMethod" orm:"request_method" ` // 请求方式
	RequestPath   string `json:"requestPath"   orm:"request_path"   ` //
	RequestBody   string `json:"requestBody"   orm:"request_body"   ` // 请求内容
	ResponseCode  int    `json:"responseCode"  orm:"response_code"  ` // 相应码
	ResponseError string `json:"responseError" orm:"response_error" ` // 错误内容
	ResponseBody  string `json:"responseBody"  orm:"response_body"  ` // 相应内容
	CreatedAt     int64  `json:"createdAt"     orm:"created_at"     ` //
}
