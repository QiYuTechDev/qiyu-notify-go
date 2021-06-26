package dt

type WxAppInfoDt struct {
	/**
	 * 唯一标识
	 */
	UniqueId string `json:"unique_id"`

	/**
	 * 应用名称
	 */
	AppName string `json:"app_name"`

	/**
	 * 企业微信 -> 公司ID
	 */
	CorpId string `json:"corp_id"`

	/**
	 * 企业微信 -> 微信令牌
	 */
	WxToken string `json:"wx_token"`

	/**
	 * 企业微信 -> AES密钥
	 */
	AesKey string `json:"aes_key"`

	/**
	 * 企业微信 -> AgentID
	 */
	AgentId int64 `json:"agent_id"`

	/**
	 * 企业微信 -> Secret
	 */
	Secret string `json:"secret"`

	/**
	 * 已验证
	 */
	Verified bool `json:"verified"`

	/**
	 * 验证时间
	 */
	VerifyTime string `json:"verify_time"`

	/**
	 * 创建时间
	 */
	Ctime string `json:"ctime"`
}
