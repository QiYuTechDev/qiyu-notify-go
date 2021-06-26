package dt

type ChatwootWebHookDt struct {
	/**
	 * 事件名称
	 */
	Event string `json:"event"`

	/**
	 * 消息ID
	 */
	Id string `json:"id"`

	/**
	 * 消息内容
	 */
	Content string `json:"content"`

	/**
	 * 发送时间
	 */
	CreatedAt string `json:"created_at"`

	/**
	 * 消息类型
	 */
	MessageType string `json:"message_type"`

	/**
	 * 内容类型
	 */
	ContentType string `json:"content_type"`

	/**
	 * 内容属性
	 */
	ContentAttributes interface{} `json:"content_attributes"`

	/**
	 * 来源ID
	 */
	SourceId string `json:"source_id"`

	/**
	 * 发送者
	 */
	Sender ChatwootSenderDt `json:"sender"`

	/**
	 * 联系人
	 */
	Contact ChatwootContactDt `json:"contact"`

	/**
	 * 会话信息
	 */
	Conversation ChatwootConversationDt `json:"conversation"`

	/**
	 * 账号信息
	 */
	Account ChatwootAccountDt `json:"account"`
}
