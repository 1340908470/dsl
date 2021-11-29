// Package model.message 是以发送者的角度，向用户发送消息
package model

// Message 消息
type Message struct {
	ChatId  string      `json:"chat_id,omitempty"`
	MsgType string      `json:"msg_type,omitempty"`
	RootId  string      `json:"root_id,omitempty"`
	Card    MessageCard `json:"card,omitempty"`
}

// MessageCard 消息卡片
type MessageCard struct {
	Config   MessageCardConfig `json:"config,omitempty"`
	Header   MessageCardHeader `json:"header,omitempty"`
	Elements []MessageElement  `json:"elements,omitempty"`
}

// MessageCardConfig 描述消息卡片的功能属性
type MessageCardConfig struct {
	EnableForward bool `json:"enable_forward,omitempty"` //是否允许卡片被转发
	UpdateMulti   bool `json:"update_multi,omitempty"`   //更新卡片的内容是否对所有收到这张卡片的人员可见，默认为false，即仅操作用户可见卡片的更新内容。
}

// MessageCardHeader 用于配置卡片的标题内容
type MessageCardHeader struct {
	Template string                 `json:"template,omitempty"`
	Title    MessageCardHeaderTitle `json:"title,omitempty"`
}

type MessageCardHeaderTitle struct {
	Tag     string `json:"tag,omitempty"`
	Content string `json:"content,omitempty"`
}

// MessageElement 用于定义卡片的正文内容
// 消息卡片的正文内容由 MessageElement 堆砌而成
// 	- 内容模块
//  - 图片模块
//  - 交互模块
type MessageElement struct {
	// 根据Tag选择是哪一种模块
	Tag string `json:"tag,omitempty"`

	// 内容模块 Tag = "div"
	Text   MessageContentText    `json:"text,omitempty"` // Text 和 Fields 二选一
	Fields []MessageElementField `json:"fields,omitempty"`
	//Extra  MessageElementExtra   `json:"extra,omitempty"` // 可选，附加的元素展示在文本内容右侧

	// 交互模块 Tag = "action"
	Actions []MessageElementAction `json:"actions,omitempty"`
	// TODO：还有图片模块、交互模块没写
}

// MessageElementField 内容模块的多文本展示
type MessageElementField struct {
	IsShort bool               `json:"is_short,omitempty"` // 是否并排布局
	Text    MessageContentText `json:"text,omitempty"`
}

// MessageElementExtra 附加元素，图片、按钮等元素是 n选1 的关系
type MessageElementExtra struct {
	// 图片元素
	MessageContentImage

	// 按钮元素
	MessageContentButton

	// 日期选择
	MessageContentDatePicker

	// TODO：还有 selectMenu、overflow、datePicker 几种元素没写
}

// MessageContentText 文本元素
type MessageContentText struct {
	Tag     string `json:"tag,omitempty"`     // 文本的内容类型，可以是 "plain_text" 或 "lark_md"
	Content string `json:"content,omitempty"` // 文本内容
	Lines   int    `json:"lines,omitempty"`   // 可选参数，用于设置内容显示行数
}

// MessageContentImage 图片元素
type MessageContentImage struct {
	Tag     string             `json:"tag,omitempty"` // 固定: Tag = "img"
	ImgKey  string             `json:"img_key,omitempty"`
	Alt     MessageContentText `json:"alt,omitempty"`     // 图片 hover 说明
	Preview bool               `json:"preview,omitempty"` // 可选参数，点击后是否放大图片，默认为 true
}

// MessageElementAction 交互模块交互组件，可以是 button 或 datePicker 或 overflow 或 selectMenu
type MessageElementAction struct {
	// 按钮
	MessageContentButton

	// 日期选择
	//MessageContentDatePicker
}

// MessageContentButton 按钮元素，可用于内容块的 extra 字段和交互块的 actions 字段
// TODO: 目前还不能使用多端跳转链接
type MessageContentButton struct {
	Tag  string             `json:"tag,omitempty"`  // 固定: Tag = "button"
	Text MessageContentText `json:"text,omitempty"` // 指定按钮的文本
	Url  string             `json:"url,omitempty"`  // 可选参数，跳转链接，与 multi_url 互斥
	//MultiUrl MessageContentUrl         `json:"multi_url,omitempty"` // 可选参数，多端跳转链接
	Type  string                    `json:"type,omitempty"`  // 可选参数，按钮样式，默认为 "default" 可以是 "default" 或 "primary" 或 "danger"
	Value MessageContentButtonValue `json:"value,omitempty"` // 用户在点击按钮后，向 Process 中添加的键值
}

type MessageContentUrl struct {
}

// MessageContentButtonValue 点击按钮后返回一对键值
type MessageContentButtonValue struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// MessageContentDatePicker 时间选择组件
type MessageContentDatePicker struct {
	Tag string `json:"tag,omitempty"` // 三选一: Tag = "date_picker" ｜ "picker_time" ｜ "picker_datetime" 【日期｜时间｜日期+时间】

	InitialDate     string `json:"initial_date,omitempty"`     // 可选参数，日期模式下的初始值，格式："yyyy-MM-dd"
	InitialTime     string `json:"initial_time,omitempty"`     // 可选参数，时间模式下的初始值，格式："HH:mm"
	InitialDatetime string `json:"initial_datetime,omitempty"` // 可选参数，日期时间模式的初始值，	格式："yyyy-MM-dd HH:mm"

	Value MessageContentDatePickerValue `json:"value,omitempty"` // 用户选择后，回传的参数
}

type MessageContentDatePickerValue struct {
}
