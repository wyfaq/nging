// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo/param"
	
	"time"
)

// Task 任务
type Task struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*Task
	namer   func(string) string
	connID  int
	
	Id            	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Uid           	uint    	`db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	GroupId       	uint    	`db:"group_id" bson:"group_id" comment:"分组ID" json:"group_id" xml:"group_id"`
	Name          	string  	`db:"name" bson:"name" comment:"任务名称" json:"name" xml:"name"`
	Type          	int     	`db:"type" bson:"type" comment:"任务类型" json:"type" xml:"type"`
	Description   	string  	`db:"description" bson:"description" comment:"任务描述" json:"description" xml:"description"`
	CronSpec      	string  	`db:"cron_spec" bson:"cron_spec" comment:"时间表达式" json:"cron_spec" xml:"cron_spec"`
	Concurrent    	uint    	`db:"concurrent" bson:"concurrent" comment:"是否支持多实例" json:"concurrent" xml:"concurrent"`
	Command       	string  	`db:"command" bson:"command" comment:"命令详情" json:"command" xml:"command"`
	WorkDirectory 	string  	`db:"work_directory" bson:"work_directory" comment:"工作目录" json:"work_directory" xml:"work_directory"`
	Env           	string  	`db:"env" bson:"env" comment:"环境变量(一行一个，格式为：var1=val1)" json:"env" xml:"env"`
	Disabled      	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	EnableNotify  	uint    	`db:"enable_notify" bson:"enable_notify" comment:"是否启用通知" json:"enable_notify" xml:"enable_notify"`
	NotifyEmail   	string  	`db:"notify_email" bson:"notify_email" comment:"通知人列表" json:"notify_email" xml:"notify_email"`
	Timeout       	uint64  	`db:"timeout" bson:"timeout" comment:"超时设置" json:"timeout" xml:"timeout"`
	ExecuteTimes  	uint    	`db:"execute_times" bson:"execute_times" comment:"累计执行次数" json:"execute_times" xml:"execute_times"`
	PrevTime      	uint    	`db:"prev_time" bson:"prev_time" comment:"上次执行时间" json:"prev_time" xml:"prev_time"`
	Created       	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated       	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	ClosedLog     	string  	`db:"closed_log" bson:"closed_log" comment:"是否(Y/N)关闭日志" json:"closed_log" xml:"closed_log"`
}

func (this *Task) Trans() *factory.Transaction {
	return this.trans
}

func (this *Task) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *Task) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *Task) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *Task) Objects() []*Task {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *Task) NewObjects() *[]*Task {
	this.objects = []*Task{}
	return &this.objects
}

func (this *Task) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *Task) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *Task) Short_() string {
	return "task"
}

func (this *Task) Struct_() string {
	return "Task"
}

func (this *Task) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *Task) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *Task) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *Task) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *Task) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Task) GroupBy(keyField string, inputRows ...[]*Task) map[string][]*Task {
	var rows []*Task
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*Task{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*Task{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *Task) KeyBy(keyField string, inputRows ...[]*Task) map[string]*Task {
	var rows []*Task
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*Task{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *Task) AsKV(keyField string, valueField string, inputRows ...[]*Task) map[string]interface{} {
	var rows []*Task
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *Task) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Task) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.ClosedLog) == 0 { this.ClosedLog = "N" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return
}

func (this *Task) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.ClosedLog) == 0 { this.ClosedLog = "N" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *Task) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *Task) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *Task) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	if val, ok := kvset["closed_log"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["closed_log"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *Task) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.ClosedLog) == 0 { this.ClosedLog = "N" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.ClosedLog) == 0 { this.ClosedLog = "N" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return 
}

func (this *Task) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *Task) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *Task) Reset() *Task {
	this.Id = 0
	this.Uid = 0
	this.GroupId = 0
	this.Name = ``
	this.Type = 0
	this.Description = ``
	this.CronSpec = ``
	this.Concurrent = 0
	this.Command = ``
	this.WorkDirectory = ``
	this.Env = ``
	this.Disabled = ``
	this.EnableNotify = 0
	this.NotifyEmail = ``
	this.Timeout = 0
	this.ExecuteTimes = 0
	this.PrevTime = 0
	this.Created = 0
	this.Updated = 0
	this.ClosedLog = ``
	return this
}

func (this *Task) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Uid"] = this.Uid
	r["GroupId"] = this.GroupId
	r["Name"] = this.Name
	r["Type"] = this.Type
	r["Description"] = this.Description
	r["CronSpec"] = this.CronSpec
	r["Concurrent"] = this.Concurrent
	r["Command"] = this.Command
	r["WorkDirectory"] = this.WorkDirectory
	r["Env"] = this.Env
	r["Disabled"] = this.Disabled
	r["EnableNotify"] = this.EnableNotify
	r["NotifyEmail"] = this.NotifyEmail
	r["Timeout"] = this.Timeout
	r["ExecuteTimes"] = this.ExecuteTimes
	r["PrevTime"] = this.PrevTime
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["ClosedLog"] = this.ClosedLog
	return r
}

func (this *Task) Set(key interface{}, value ...interface{}) factory.Model {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint(vv)
				case "Uid": this.Uid = param.AsUint(vv)
				case "GroupId": this.GroupId = param.AsUint(vv)
				case "Name": this.Name = param.AsString(vv)
				case "Type": this.Type = param.AsInt(vv)
				case "Description": this.Description = param.AsString(vv)
				case "CronSpec": this.CronSpec = param.AsString(vv)
				case "Concurrent": this.Concurrent = param.AsUint(vv)
				case "Command": this.Command = param.AsString(vv)
				case "WorkDirectory": this.WorkDirectory = param.AsString(vv)
				case "Env": this.Env = param.AsString(vv)
				case "Disabled": this.Disabled = param.AsString(vv)
				case "EnableNotify": this.EnableNotify = param.AsUint(vv)
				case "NotifyEmail": this.NotifyEmail = param.AsString(vv)
				case "Timeout": this.Timeout = param.AsUint64(vv)
				case "ExecuteTimes": this.ExecuteTimes = param.AsUint(vv)
				case "PrevTime": this.PrevTime = param.AsUint(vv)
				case "Created": this.Created = param.AsUint(vv)
				case "Updated": this.Updated = param.AsUint(vv)
				case "ClosedLog": this.ClosedLog = param.AsString(vv)
			}
	}
	return this
}

func (this *Task) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["uid"] = this.Uid
	r["group_id"] = this.GroupId
	r["name"] = this.Name
	r["type"] = this.Type
	r["description"] = this.Description
	r["cron_spec"] = this.CronSpec
	r["concurrent"] = this.Concurrent
	r["command"] = this.Command
	r["work_directory"] = this.WorkDirectory
	r["env"] = this.Env
	r["disabled"] = this.Disabled
	r["enable_notify"] = this.EnableNotify
	r["notify_email"] = this.NotifyEmail
	r["timeout"] = this.Timeout
	r["execute_times"] = this.ExecuteTimes
	r["prev_time"] = this.PrevTime
	r["created"] = this.Created
	r["updated"] = this.Updated
	r["closed_log"] = this.ClosedLog
	return r
}

func (this *Task) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *Task) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

