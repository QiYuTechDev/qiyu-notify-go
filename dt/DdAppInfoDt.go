package dt

type DdAppInfoDt struct {
	/**
	 * 唯一标识
	 */
	UniqueId string `json:"unique_id"`

	/**
	 * 应用名称
	 */
	AppName string `json:"app_name"`

	/**
	 * 钉钉访问令牌
	 */
	DdAk string `json:"dd_ak"`

	/**
	 * 钉钉加密密钥
	 */
	DdSk string `json:"dd_sk"`

	/**
	 * 已验证
	 */
	Verified bool `json:"verified"`

	/**
	 * 创建时间
	 */
	Ctime string `json:"ctime"`
}
