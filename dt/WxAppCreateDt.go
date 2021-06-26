package dt

type WxAppCreateDt struct {
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
}
