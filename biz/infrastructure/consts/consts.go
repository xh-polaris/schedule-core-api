package consts

var PageSize int64 = 10

// 数据库相关
const (
	ID           = "_id"
	UserID       = "user_id"
	Name         = "name"
	Status       = "status"
	Phone        = "phone"
	Top          = "top"
	Priority     = "priority"
	CreateTime   = "create_time"
	UpdateTime   = "update_time"
	DeleteTime   = "delete_time"
	DeleteStatus = 1
)

// mongo关键字
const (
	SET = "$set"
)

// http
const (
	GlmUrl                    = "https://open.bigmodel.cn/api/paas/v4/chat/completions"
	GlmModel                  = "glm-4-air"
	Post                      = "POST"
	PlatformSignInUrl         = "https://api.xhpolaris.com/platform/auth/sign_in"
	PlatformSetPasswordUrl    = "https://api.xhpolaris.com/platform/auth/set_password"
	PlatformSendVerifyCodeUrl = "https://api.xhpolaris.com/platform/auth/send_verify_code"
	ContentTypeJson           = "application/json"
	CharSetUTF8               = "UTF-8"
)

// 默认值
const (
	AppId            = 16
	DefaultStatus    = 0
	DefaultDone      = 0
	DefaultGroupName = "default"
	DefaultProgress  = -1
	DefaultPriority  = 0
	DefaultTop       = 0
	DefaultPrompt    = "你是一个日程管理大师，我将给你一句日常的对话，你将分析这段话，并且从中分析出日程，如果有多个任务则需要分析出多个日程。输出要求如下:输出的格式是{\"schedules\":[{\"origin\":\"原话\", \"title\":\"日程标题\", \"description\":\"日程描述\",\"ddl\":\"截至日期\"}]}，一个json格式的数组，每个元素包含一个json对象，字段有origin：分析出这个任务的原话，title：这个日程的简要标题，description：这个日程的详细描述，ddl：现在的时间是%s，由此推断这个任务的截止日期，如果分析不出来就用 0001-01-01 00:00:00 这个时间零值，请注意，请只返回一个json数据，不要包含其他任何文字，且不要返回markdown，不允许```包裹json，直接{开头。接下来请你分析这句话:%s"
)
