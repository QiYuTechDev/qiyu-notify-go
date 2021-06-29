package dt

type TplAppUpdateDt struct {
	/**
	 * 模版应用名称
	 */
	AppName string `json:"app_name"`

	/**
	 * 微信应用UUID
	 */
	WxApp string `json:"wx_app"`

	/**
	 * 钉钉应用UUID
	 */
	DdApp string `json:"dd_app"`

	/**
	 * 邮箱应用UUID
	 */
	EmailApp string `json:"email_app"`

	/**
	 * 模版
	 */
	TplCode string `json:"tpl_code"`

	/**
	 * 用备注
	 */
	Comment string `json:"comment"`

	/**
	 * 唯一标识
	 */
	UniqueId string `json:"unique_id"`
}
