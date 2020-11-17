package util

const (
	SUCCESS         = 0 //成功
	SUCCESS_Partial = 1 //部分成功（批量操作）

	//公共错误
	Fail                   = 10000 //操作失败
	LackParameter          = 10001 //缺少必要的参数
	InvaildParameterFormat = 10002 //参数格式有误
	InvaildParameter       = 10004 //参数验证失败
	NotFound               = 10005 //记录不存在
	BadRequest             = 10006 //记录不存在

	//用户模块
	UsernameOrPwdError   = 40100 //用户名或密码不正确
	LoginOvertime        = 40101 //登录过期，请重新登录
	UsernameOrPwdEmpty   = 40102 //用户名或密码为空
	NoPower              = 40103 //没有访问权限
	PwdUnIdentical       = 40104 //两次密码不一致
	OldPwdError          = 40105 //旧密码错误
	UserNameExists       = 40106 //用户名已经存在
	UserNotExists        = 40107 //用户不存在
	UserPermissionDend   = 40108 //没有权限
	UserPermissionExpiry = 40109 //授权过期

	DeviceBinded         = 40211 //设备重启中
	DeviceRebootFrequent = 40211 //5分
	DeviceChangeStoreBad = 40213 //淘宝系设备不能切换门店

	//商品操作模块
	ProductStatesError        = 40301 //产品状态错误，禁止该操作
	ProductErrorSingleItemNot = 40302 //关联商品N未上架，不可上架该组合商品
	CanNotShelfThisBrand      = 40303 //所选门店不能上架XX品牌
	PrdIdNotBelongOrDelete    = 40304 //商品不属于该门店或以删除
	TooManyPrdUrls            = 40305 //批量操作太多数据了
	MustWebUrl                = 40306 //必须是http(s):开头的链接

	StoreNickExist        = 40401 //门店名已存在
	StoreHasNoBrand       = 40402 //总店还没有创建品牌
	StoreHasNoAccessToken = 40403 //门店无第三方授权码

	BrandNameExist  = 40501 //品牌名称已存在
	BrandHasProduct = 40502 //品牌已有商品，禁止禁用

	ActvStartTimeOut = 40601 //活动开始时间超过当前时间！
	ActvTimeRepet    = 40602 //活动时间重叠
	ActvEndTimeOut   = 40603 //活动结束时间超过当前时间！

	CouponCodeEmpty   = 40701 //优惠券码不能为空
	CouponCodeFileErr = 40702 //优惠券码文件格式错误

	StoreManagerExist = 40801 //店长已存在
	StaffAccountExist = 40802 //员工账号已存在
	HasNoCostAct      = 40803 //无可创建的付费账号

	//服务器错误
	ServerError = 50000
	DbError     = 50002 //请求超时
	TimeOut     = 50003 //请求超时
)

const (
	SUCCESSMsg = "成功"

	//公共错误
	FailMsg                   = "操作失败"
	LackParameterMsg          = "缺少必要的参数"
	InvaildParameterFormatMsg = "参数格式有误"
	InvaildParameterMsg       = "参数验证失败"
	NotFoundMsg               = "记录不存在"
	BadRequestMsg             = "请求失败"

	//用户模块
	UsernameOrPwdErrorMsg   = "用户名或密码不正确"
	LoginOvertimeMsg        = "登录过期，请重新登录"
	UsernameOrPwdEmptyMsg   = "用户名或密码为空"
	NoPowerMsg              = "没有访问权限"
	PwdUnIdenticalMsg       = "两次密码不一致"
	OldPwdErrorMsg          = "旧密码错误"
	UserNameExistsMsg       = "用户名已存在"
	UserNotExistsMsg        = "User is not exists"
	UserPermissionDendMsg   = "没有操作权限"
	UserPermissionExpiryMsg = "授权过期"

	DeviceBindedMsg         = "设备已经在别的商家绑定了"
	DeviceRebootFrequentMsg = "5分钟之内不能重复操作"
	DeviceChangeStoreBadMsg = "淘宝系设备不能切换门店"

	TooManyPrdUrlsMsg = "太多记录"

	StoreNickExistMsg  = "门店名称已存在"
	StoreHasNoBrandMsg = "总店还没有创建品牌"

	BrandNameExistMsg        = "品牌名已存在"
	BrandHasProductMsg       = "品牌下已有商品"
	StoreHasNoAccessTokenMsg = "门店无第三方授权码"
	MustWebUrlMsg            = "必须是http(s):开头的链接"

	ActvStartTimeOutMsg = "活动开始时间必须大于当前时间"
	ActvTimeRepetMsg    = "活动时间重叠"
	ActvEndTimeOutMsg   = "活动结束时间必须大于当前时间"

	CouponCodeEmptyMsg   = "优惠券码为空"
	CouponCodeFileErrMsg = "券码文件格式错误"

	StoreManagerExistMsg = "店长已存在，一个门店只能创建一位店长"
	StaffAccountExistMsg = "员工账号已存在"
	HasNoCostActMsg      = "无可创建的付费账号"

	//服务器错误
	ServerErrorMsg = "服务器出错"
	TimeOutMsg     = "请求超时"
	DbErrorMsg     = "Db operation failed"
)

const (
	//OnlineSeconds 超时
	OnlineSeconds = 3 * 60
)
