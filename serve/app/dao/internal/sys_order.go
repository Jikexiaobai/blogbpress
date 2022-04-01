// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"time"

	"fiber/app/model"
)

// SysOrderDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type SysOrderDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns sysOrderColumns
}

// SysOrderColumns defines and stores column names for table sys_order.
type sysOrderColumns struct {
	OrderId           string //                                                                                                                             
    OrderNum          string // 订单编号                                                                                                                    
    UserId            string // 下单人id                                                                                                                    
    AuthorId          string // 作者id                                                                                                                      
    PayMethod         string // 支付方式：1支付宝，2微信，3余额                                                                                             
    OrderMoney        string // 订单金额                                                                                                                    
    DistrictMoney     string // 优惠金额                                                                                                                    
    AuthorMoney       string // 作者收益                                                                                                                    
    ServiceMoney      string // 服务费                                                                                                                      
    PaymentMoney      string // 支付金额                                                                                                                    
    OrderPoint        string // 订单积分                                                                                                                    
    OrderType         string // 订单类型: 1 充值，2打赏充电，3内容购买，4查看话题隐藏内容  ,5加入付费圈子,6购买付费课程,7查看付费答案，8开通vip，9认证付费  
    OrderMode         string // 订单方式：1虚拟物品，2实体物品                                                                                              
    ShippingMoney     string // 运输费用                                                                                                                    
    ShippingAddress   string // 收获地址                                                                                                                    
    ShippingCompName  string // 快递公司                                                                                                                    
    ShippingPhone     string // 收货人联系方式                                                                                                              
    ShippingName      string // 收货人姓名                                                                                                                  
    ShippingSn        string // 快递单号                                                                                                                    
    ShippingTime      string // 发货时间                                                                                                                    
    DetailId          string //                                                                                                                             
    DetailModule      string //                                                                                                                             
    Status            string // 订单状态 1未支付，2已支付                                                                                                   
    Invoice           string // 发票抬头                                                                                                                    
    PayTime           string // 支付时间                                                                                                                    
    ReceiveTime       string // 收货时间                                                                                                                    
    CreateTime        string // 创建时间                                                                                                                    
    UpdateTime        string // 最后修改时间
}

var (
	// SysOrder is globally public accessible object for table sys_order operations.
	SysOrder = &SysOrderDao{
		M:     g.DB("default").Model("sys_order").Safe(),
		DB:    g.DB("default"),
		Table: "sys_order",
		Columns: sysOrderColumns{
			OrderId:          "order_id",            
            OrderNum:         "order_num",           
            UserId:           "user_id",             
            AuthorId:         "author_id",           
            PayMethod:        "pay_method",          
            OrderMoney:       "order_money",         
            DistrictMoney:    "district_money",      
            AuthorMoney:      "author_money",        
            ServiceMoney:     "service_money",       
            PaymentMoney:     "payment_money",       
            OrderPoint:       "order_point",         
            OrderType:        "order_type",          
            OrderMode:        "order_mode",          
            ShippingMoney:    "shipping_money",      
            ShippingAddress:  "shipping_address",    
            ShippingCompName: "shipping_comp_name",  
            ShippingPhone:    "shipping_phone",      
            ShippingName:     "shipping_name",       
            ShippingSn:       "shipping_sn",         
            ShippingTime:     "shipping_time",       
            DetailId:         "detail_id",           
            DetailModule:     "detail_module",       
            Status:           "status",              
            Invoice:          "invoice",             
            PayTime:          "pay_time",            
            ReceiveTime:      "receive_time",        
            CreateTime:       "create_time",         
            UpdateTime:       "update_time",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *SysOrderDao) Ctx(ctx context.Context) *SysOrderDao {
	return &SysOrderDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *SysOrderDao) As(as string) *SysOrderDao {
	return &SysOrderDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *SysOrderDao) TX(tx *gdb.TX) *SysOrderDao {
	return &SysOrderDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *SysOrderDao) Master() *SysOrderDao {
	return &SysOrderDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *SysOrderDao) Slave() *SysOrderDao {
	return &SysOrderDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *SysOrderDao) Args(args ...interface{}) *SysOrderDao {
	return &SysOrderDao{M: d.M.Args(args ...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *SysOrderDao) LeftJoin(table ...string) *SysOrderDao {
	return &SysOrderDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *SysOrderDao) RightJoin(table ...string) *SysOrderDao {
	return &SysOrderDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *SysOrderDao) InnerJoin(table ...string) *SysOrderDao {
	return &SysOrderDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *SysOrderDao) Fields(fieldNamesOrMapStruct ...interface{}) *SysOrderDao {
	return &SysOrderDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *SysOrderDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *SysOrderDao {
	return &SysOrderDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *SysOrderDao) Option(option int) *SysOrderDao {
	return &SysOrderDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *SysOrderDao) OmitEmpty() *SysOrderDao {
	return &SysOrderDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *SysOrderDao) Filter() *SysOrderDao {
	return &SysOrderDao{M: d.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (d *SysOrderDao) Where(where interface{}, args ...interface{}) *SysOrderDao {
	return &SysOrderDao{M: d.M.Where(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *SysOrderDao) WherePri(where interface{}, args ...interface{}) *SysOrderDao {
	return &SysOrderDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *SysOrderDao) And(where interface{}, args ...interface{}) *SysOrderDao {
	return &SysOrderDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *SysOrderDao) Or(where interface{}, args ...interface{}) *SysOrderDao {
	return &SysOrderDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *SysOrderDao) Group(groupBy string) *SysOrderDao {
	return &SysOrderDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *SysOrderDao) Order(orderBy ...string) *SysOrderDao {
	return &SysOrderDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *SysOrderDao) Limit(limit ...int) *SysOrderDao {
	return &SysOrderDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *SysOrderDao) Offset(offset int) *SysOrderDao {
	return &SysOrderDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *SysOrderDao) Page(page, limit int) *SysOrderDao {
	return &SysOrderDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *SysOrderDao) Batch(batch int) *SysOrderDao {
	return &SysOrderDao{M: d.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (d *SysOrderDao) Cache(duration time.Duration, name ...string) *SysOrderDao {
	return &SysOrderDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *SysOrderDao) Data(data ...interface{}) *SysOrderDao {
	return &SysOrderDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.SysOrder.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *SysOrderDao) All(where ...interface{}) ([]*model.SysOrder, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.SysOrder
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.SysOrder.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *SysOrderDao) One(where ...interface{}) (*model.SysOrder, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.SysOrder
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *SysOrderDao) FindOne(where ...interface{}) (*model.SysOrder, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.SysOrder
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *SysOrderDao) FindAll(where ...interface{}) ([]*model.SysOrder, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.SysOrder
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Struct retrieves one record from table and converts it into given struct.
// The parameter <pointer> should be type of *struct/**struct. If type **struct is given,
// it can create the struct internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not nil.
//
// Eg:
// user := new(User)
// err  := dao.User.Where("id", 1).Struct(user)
//
// user := (*User)(nil)
// err  := dao.User.Where("id", 1).Struct(&user)
func (d *SysOrderDao) Struct(pointer interface{}, where ...interface{}) error {
	return d.M.Struct(pointer, where...)
}

// Structs retrieves records from table and converts them into given struct slice.
// The parameter <pointer> should be type of *[]struct/*[]*struct. It can create and fill the struct
// slice internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not empty.
//
// Eg:
// users := ([]User)(nil)
// err   := dao.User.Structs(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Structs(&users)
func (d *SysOrderDao) Structs(pointer interface{}, where ...interface{}) error {
	return d.M.Structs(pointer, where...)
}

// Scan automatically calls Struct or Structs function according to the type of parameter <pointer>.
// It calls function Struct if <pointer> is type of *struct/**struct.
// It calls function Structs if <pointer> is type of *[]struct/*[]*struct.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved and given pointer is not empty or nil.
//
// Eg:
// user  := new(User)
// err   := dao.User.Where("id", 1).Scan(user)
//
// user  := (*User)(nil)
// err   := dao.User.Where("id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err   := dao.User.Scan(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Scan(&users)
func (d *SysOrderDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *SysOrderDao) Chunk(limit int, callback func(entities []*model.SysOrder, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.SysOrder
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *SysOrderDao) LockUpdate() *SysOrderDao {
	return &SysOrderDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *SysOrderDao) LockShared() *SysOrderDao {
	return &SysOrderDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *SysOrderDao) Unscoped() *SysOrderDao {
	return &SysOrderDao{M: d.M.Unscoped()}
}