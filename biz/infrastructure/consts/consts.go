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
	DeleteStatus = -1
)

// mongo关键字
const (
	SET = "$set"
)

// http
const (
	GlmUrl                    = "https://open.bigmodel.cn/api/paas/v4/chat/completions"
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
	DefaultPrompt    = `请作为高精度日程解析引擎，严格遵循以下处理规则，

输入：
- 当前时间：{current_time} (由系统提供的YYYY-MM-DD HH:mm:ss格式本地时间，默认时区为东八区)
- 用户输入："{user_input}" (用户原始日程描述)

处理规则：
1. 标题生成：
  - 提取核心事件主体，去除时间状语/修饰词
  - 保留关键动词+名词结构（如"提交报告"）
  - 长度限制：中文≤15字，英文≤30字符

2. 时间解析：
  - 绝对时间处理：
    - 识别"YYYY-MM-DD"格式直接采用
    - 转换"X月X日"为当年对应日期
  - 相对时间计算：
    - 基于current_time计算相对日期
    - 支持：今天/明天/后天/大后天/周X/下周X/下下周X/大下周X
    - 支持：x天后/x个工作日后（排除周末）
  - 默认值：当无法识别时间时，使用current_time的日期

我将给你一句日常的对话，请分析这段话，并且从中分析出日程，如果有多个任务则需要分析出多个日程。
输出要求如下:
1. 输出的格式是：一个json格式的数组，每个元素包含一个json对象，json对象格式如下
    - {"schedules":[{"origin":"原话", "title":"日程标题", "description":"日程描述","ddl":"截止时间"}]}
    - 只返回一个json数据，不要包含其他任何文字，不要返回markdown，不要用反引号包裹json，直接以{开头。
2. 字段有：
    - origin：分析出这个任务的原话；
    - title：这个日程的简要标题；
    - description：这个日程的详细描述；
    - ddl：这个任务的截止时间，如果分析不出来就用 0001-01-01 00:00:00 这个时间零值。
        - 注意区分会议等活动和普通任务截止时间，若日程为会议等活动请以它们的开始时间作为该项的值

当前时间为：%s，请分析这句话：%s`
)
