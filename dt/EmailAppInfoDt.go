package dt

type EmailAppInfoDt struct {
	/**
	 * 唯一标识
	 */
	UniqueId string `json:"unique_id"`

	/**
	 * 推送接收邮箱
	 */
	NotifyEmail string `json:"notify_email"`

	/**
	 * 是否验证
	 */
	Verified bool `json:"verified"`

	/**
	 * 创建时间
	 */
	Ctime string `json:"ctime"`
}
