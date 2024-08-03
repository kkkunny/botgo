package dto

// EventType 事件类型
type EventType string

// WSPayload websocket 消息结构
type WSPayload struct {
	WSPayloadBase
	Data       interface{} `json:"d,omitempty"`
	RawMessage []byte      `json:"-"` // 原始的 message 数据
}

// WSPayloadBase 基础消息结构，排除了 data
type WSPayloadBase struct {
	OPCode OPCode    `json:"op"`
	Seq    uint32    `json:"s,omitempty"`
	Type   EventType `json:"t,omitempty"`
}

// 以下为发送到 websocket 的 data

// WSIdentityData 鉴权数据
type WSIdentityData struct {
	Token      string   `json:"token"`
	Intents    Intent   `json:"intents"`
	Shard      []uint32 `json:"shard"` // array of two integers (shard_id, num_shards)
	Properties struct {
		Os      string `json:"$os,omitempty"`
		Browser string `json:"$browser,omitempty"`
		Device  string `json:"$device,omitempty"`
	} `json:"properties,omitempty"`
}

// WSResumeData 重连数据
type WSResumeData struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       uint32 `json:"seq"`
}

// 以下为会收到的事件data

// WSHelloData hello 返回
type WSHelloData struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// WSReadyData ready，鉴权后返回
type WSReadyData struct {
	Version   int    `json:"version"`
	SessionID string `json:"session_id"`
	User      struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Bot      bool   `json:"bot"`
	} `json:"user"`
	Shard []uint32 `json:"shard"`
}

// WSGuildData 频道 payload
type WSGuildData Guild

// WSGuildMemberData 频道成员 payload
type WSGuildMemberData Member

// WSChannelData 子频道 payload
type WSChannelData Channel

// WSMessageData 消息 payload
type WSMessageData Message

// WSATMessageData only at 机器人的消息 payload
type WSATMessageData Message

// WSDirectMessageData 私信消息 payload
type WSDirectMessageData Message

// WSMessageDeleteData 消息 payload
type WSMessageDeleteData MessageDelete

// WSPublicMessageDeleteData 公域机器人的消息删除 payload
type WSPublicMessageDeleteData MessageDelete

// WSDirectMessageDeleteData 私信消息 payload
type WSDirectMessageDeleteData MessageDelete

// WSAudioData 音频机器人的音频流事件
type WSAudioData AudioAction

// WSMessageReactionData 表情表态事件
type WSMessageReactionData MessageReaction

// WSMessageAuditData 消息审核事件
type WSMessageAuditData MessageAudit

// WSThreadData 主题事件
type WSThreadData Thread

// WSPostData 帖子事件
type WSPostData Post

// WSReplyData 帖子回复事件
type WSReplyData Reply

// WSForumAuditData 帖子审核事件
type WSForumAuditData ForumAuditResult

// WSInteractionData 互动事件
type WSInteractionData Interaction

// WSC2CMessageData 单聊消息 payload
type WSC2CMessageData struct {
	// 消息ID
	ID string `json:"id"`
	// 发送者
	Author struct {
		// 用户 openid
		UserOpenID string `json:"user_openid"`
	} `json:"author"`
	// 文本消息类型
	Content string `json:"content"`
	// 消息生产时间
	Timestamp Timestamp `json:"timestamp"`
	// 富媒体文件附件
	Attachments []struct {
		// 文件类型，"image/jpeg","image/png","image/gif"，"file"，"video/mp4"，"voice"
		ContentType string `json:"content_type"`
		// 文件名称
		Filename string `json:"filename"`
		// 图片高度
		Height int64 `json:"height"`
		// 图片宽度
		Width int64 `json:"width"`
		// 文件大小
		Size int64 `json:"size"`
		// 文件链接
		URL string `json:"url"`
	} `json:"attachments,omitempty"`
}

// WSGroupATMessageData 群聊 only at 机器人的消息 payload
type WSGroupATMessageData struct {
	// 平台方消息 ID，可以用于被动消息发送
	ID string `json:"id"`
	// 发送者
	Author struct {
		// 用户在本群的 member_openid
		MemberOpenID string `json:"member_openid"`
	} `json:"author"`
	// 消息内容
	Content string `json:"content"`
	// 消息生产时间
	Timestamp Timestamp `json:"timestamp"`
	// 群聊的 openid
	GroupOpenID string `json:"group_openid"`
	// 富媒体文件附件
	Attachments []struct {
		// 文件类型，"image/jpeg","image/png","image/gif"，"file"，"video/mp4"，"voice"
		ContentType string `json:"content_type"`
		// 文件名称
		Filename string `json:"filename"`
		// 图片高度
		Height int64 `json:"height"`
		// 图片宽度
		Width int64 `json:"width"`
		// 文件大小
		Size int64 `json:"size"`
		// 文件链接
		URL string `json:"url"`
	} `json:"attachments,omitempty"`
}
