package dt

type DdAppCreateDt struct {
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
}
