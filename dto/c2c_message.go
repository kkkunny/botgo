package dto

// C2CMessageType 人人、群聊消息类型
type C2CMessageType uint8

const (
	// C2CMessageTypeText 文本
	C2CMessageTypeText C2CMessageType = 0
	// C2CMessageTypeMarkdown markdown
	C2CMessageTypeMarkdown C2CMessageType = 2
	// C2CMessageTypeArk ark
	C2CMessageTypeArk C2CMessageType = 3
	// C2CMessageTypeEmbed embed
	C2CMessageTypeEmbed C2CMessageType = 4
	// C2CMessageTypeMedia media
	C2CMessageTypeMedia C2CMessageType = 7
)

// C2CMessageToCreate 创建人人、群聊消息的结构体定义
type C2CMessageToCreate struct {
	// 消息类型
	MessageType C2CMessageType `json:"msg_type"`
	// 文本内容
	Content string `json:"content,omitempty"`
	// Markdown对象
	Markdown *MarkdownMessage `json:"markdown,omitempty"`
	// Keyboard对象
	Keyboard *KeyboardMessage `json:"keyboard,omitempty"`
	// Ark对象
	Ark *ArkMessage `json:"ark,omitempty"`
	// 富媒体单聊的file_info
	Media string `json:"media,omitempty"`
	// 【暂未支持】消息引用
	MessageReference interface{} `json:"message_reference,omitempty"`
	// 前置收到的事件 ID，用于发送被动消息，支持事件："INTERACTION_CREATE"、"C2C_MSG_RECEIVE"、"FRIEND_ADD"
	EventID EventType `json:"event_id,omitempty"`
	// 前置收到的用户发送过来的消息 ID，用于发送被动（回复）消息
	MsgID string `json:"msg_id,omitempty"`
	// 回复消息的序号，与 msg_id 联合使用，避免相同消息id回复重复发送，不填默认是1。相同的 msg_id + msg_seq 重复发送会失败。
	MsgSeq *int64 `json:"msg_seq,omitempty"`
}

// MarkdownMessage Markdown消息结构体
type MarkdownMessage struct {
	// markdown 模版id，申请模版后获得
	CustomTemplateID string `json:"custom_template_id,omitempty"`
	// 原生 markdown 文本内容
	Content string `json:"content,omitempty"`
	// 模版内变量与填充值的kv映射
	Params []MarkdownParam `json:"params,omitempty"`
}

// MarkdownParam Markdown参数结构体
type MarkdownParam struct {
	// key
	Key string `json:"key"`
	// value
	Values string `json:"values"`
}

// KeyboardMessage Keyboard消息结构体
type KeyboardMessage struct {
	Rows []KeyboardRow `json:"rows"`
}

// KeyboardRow Keyboard row 结构体
type KeyboardRow struct {
	// 内部可以是KeyboardButton，也可以是KeyboardRow
	Buttons []interface{} `json:"buttons"`
}

// KeyboardButton Keyboard button 结构体
type KeyboardButton struct {
	// 按钮ID：在一个keyboard消息内设置唯一
	ID         string                `json:"id,omitempty"`
	RenderData KeyboardRenderData    `json:"render_data"`
	Action     KeyboardMessageAction `json:"action"`
}

// KeyboardRenderData Keyboard render data 结构体
type KeyboardRenderData struct {
	// 按钮样式：0 灰色线框，1 蓝色线框
	Style int64 `json:"style"`
	// 按钮上的文字
	Label string `json:"label"`
	// 点击后按钮的上文字
	VisitedLabel string `json:"visited_label"`
}

// KeyboardMessageAction Keyboard action结构体
type KeyboardMessageAction struct {
	// 设置 0 跳转按钮：http 或 小程序 客户端识别 scheme，设置 1 回调按钮：回调后台接口, data 传给后台，设置 2 指令按钮：自动在输入框插入 @bot data
	Type       int64 `json:"type"`
	Permission struct {
		// 0 指定用户可操作，1 仅管理者可操作，2 所有人可操作，3 指定身份组可操作（仅频道可用）
		Type int64 `json:"type"`
		// 有权限的用户 id 的列表
		SpecifyUserIDs []string `json:"specify_user_ids,omitempty"`
		// 有权限的身份组 id 的列表（仅频道可用）
		SpecifyRoleIDs []string `json:"specify_role_ids,omitempty"`
	} `json:"permission"`
	// 操作相关的数据
	Data string `json:"data"`
	// 指令按钮可用，指令是否带引用回复本消息，默认 false。支持版本 8983
	Reply bool `json:"reply,omitempty"`
	// 指令按钮可用，点击按钮后直接自动发送 data，默认 false。支持版本 8983
	Enter bool `json:"enter,omitempty"`
	// 本字段仅在指令按钮下有效，设置后后会忽略 action.enter 配置。
	// 设置为 1 时 ，点击按钮自动唤起启手Q选图器，其他值暂无效果。
	// （仅支持手机端版本 8983+ 的单聊场景，桌面端不支持）
	Anchor int64 `json:"anchor,omitempty"`
	// 【已弃用】可操作点击的次数，默认不限
	ClickLimit int64 `json:"click_limit,omitempty"`
	// 【已弃用】指令按钮可用，弹出子频道选择器，默认 false
	ATBotShowChannelList bool `json:"at_bot_show_channel_list,omitempty"`
	// 客户端不支持本action的时候，弹出的toast文案
	UnSupportTips string `json:"unsupport_tips"`
}

// ArkMessage Ark消息结构体
type ArkMessage struct {
	// 模版 id，管理端可获得或内邀申请获得
	// 以下默认可使用：
	// 23 链接+文本列表模板
	// 24 文本+缩略图模板
	// 37 大图模板
	TemplateID int64 `json:"template_id"`
	// 模版内变量与填充值的kv映射
	Params []ArkParam `json:"params"`
}

// ArkParam Ark参数结构体
type ArkParam struct {
	// key
	Key string `json:"key"`
	// value
	Value string `json:"value"`
}

// C2CMessageToReply 人人、群聊消息回复结构体定义
type C2CMessageToReply struct {
	// 消息唯一ID
	ID string `json:"id"`
	// 发送时间
	Timestamp Timestamp `json:"timestamp"`
}

// C2CFileType 人人、群聊富媒体类型
type C2CFileType uint8

const (
	// C2CFileTypeImage 图片png/jpg
	C2CFileTypeImage C2CFileType = iota + 1
	// C2CFileTypeVideo 视频mp4
	C2CFileTypeVideo
	// C2CFileTypeAudio 语音silk
	C2CFileTypeAudio
	// C2CFileTypeFile 文件（暂不开放）
	C2CFileTypeFile
)

// C2CFileToCreate 创建人人、群聊富媒体的结构体定义
type C2CFileToCreate struct {
	// 媒体类型
	FileType C2CFileType `json:"file_type"`
	// 需要发送媒体资源的url
	URL string `json:"url"`
	// 设置 true 会直接发送消息到目标端，且会占用主动消息频次
	SrvSendMsg bool `json:"srv_send_msg"`
	// 【暂未支持】
	FileData interface{} `json:"file_data,omitempty"`
}

// C2CFileToReply 人人、群聊富媒体回复结构体定义
type C2CFileToReply struct {
	// 文件 ID
	FileUUID string `json:"file_uuid"`
	// 文件信息，用于发消息接口的 media 字段使用
	FileInfo string `json:"file_info"`
	// 有效期，表示剩余多少秒到期，到期后 file_info 失效，当等于 0 时，表示可长期使用
	TTL int64 `json:"ttl"`
	// 发送消息的唯一ID，当srv_send_msg设置为true时返回
	ID string `json:"id,omitempty"`
}
