package dt

type DdMsgDt struct {
	/**
	 * 来源IP
	 */
	SrcIp string `json:"src_ip"`

	/**
	 * 消息类型
	 */
	MsgType string `json:"msg_type"`

	/**
	 * 消息内容
	 */
	Content string `json:"content"`

	/**
	 * 发送时间
	 */
	Ctime string `json:"ctime"`
}
