package dt

type TplAppInfoDt struct {
	/**
	 * 模版应用名称
	 */
	AppName string `json:"app_name"`

	/**
	 * 唯一标识
	 */
	UniqueId string `json:"unique_id"`

	/**
	 * 微信应用信息
	 */
	WxApp WxAppInfoDt `json:"wx_app"`

	/**
	 * 邮箱应用信息
	 */
	EmailApp EmailAppInfoDt `json:"email_app"`

	/**
	 * 钉钉应用信息
	 */
	DdApp DdAppInfoDt `json:"dd_app"`

	/**
	 * 模版
	 */
	TplCode string `json:"tpl_code"`

	/**
	 * 备注
	 */
	Comment string `json:"comment"`

	/**
	 * 创建时间
	 */
	Ctime string `json:"ctime"`
}
