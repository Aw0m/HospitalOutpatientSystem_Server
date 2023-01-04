package utils

var MySigningKey = []byte("dev123")

const (
	Duration   = 24 * 30
	TimeLayout = "2006-01-02"

	OpenIDURL = "https://api.weixin.qq.com/sns/jscode2session?appid=wxf02df1739164ba92&secret=35ca848d9e2f3ba0b88ff9eedc4b2eb2&js_code=%s&grant_type=authorization_code"
)
