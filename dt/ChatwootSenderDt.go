package dt

type ChatwootSenderDt struct {
	/**
	 * 发送者ID
	 */
	Id string `json:"id"`

	/**
	 * 发送者的名称
	 */
	Name string `json:"name"`

	/**
	 * 电子邮箱
	 */
	Email string `json:"email"`
}
