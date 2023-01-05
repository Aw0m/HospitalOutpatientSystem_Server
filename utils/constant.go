package utils

var MySigningKey = []byte("dev123")

const (
	Duration   = 24 * 30
	TimeLayout = "2006-01-02"

	OpenIDURL = "https://api.weixin.qq.com/sns/jscode2session?appid=wxbfb49af0d2946439&secret=edec5769dee37f9c82369f4ab66b8126&js_code=%s&grant_type=authorization_code"
)
