package dt

type EmailVerifyArgs struct {
	/**
	 * 唯一标识
	 */
	UniqueId string `json:"unique_id"`

	/**
	 * 验证码
	 */
	Code string `json:"code"`
}
