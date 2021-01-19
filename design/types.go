package design

import . "goa.design/goa/v3/dsl"

var (
	ExampleJwt  = "eyJhbGciOiJIUz..."
	ExampleUUID = "91cc3eb9-ddc0-4cf7-a62b-c85df1a9166f"
)

var SuccessResult = ResultType("SuccessResult", func() {
	Description("成功信息")
	ContentType("application/json")
	TypeName("SuccessResult")

	Attributes(func() {
		Attribute("ok", Boolean, "success", func() {
			Example(true)
		})
		Required("ok")
	})

	View("default", func() {
		Attribute("ok")
	})
})

var User = ResultType("User", func() {
	Description("用户")
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID")
		Attribute("username", String, "用户名")
		Attribute("mobile", String, "手机号")

		Required("id", "username",  "mobile")
	})

	View("default", func() {
		Attribute("id")
		Attribute("username")
		Attribute("mobile")
	})
})

var Credentials = Type("Credentials", func() {
	Field(1, "token", String, "JWT token", func() {
		Example(ExampleJwt)
	})
	Field(7, "expires_in", Int, "有效时长（秒）：生成之后x秒内有效", func() {
		Example(25200)
	})
	Required("token", "expires_in")
})

var Session = ResultType("Session", func() {
	Description("会话")
	ContentType("application/json")
	Attributes(func() {
		Field(1, "user", User)
		Field(2, "credentials", Credentials)
		Required("user", "credentials")
	})

	View("default", func() {
		Attribute("user")
		Attribute("credentials")
	})
})

var Captcha = Type("Captcha", func() {
	Attribute("image", String, "图片base64", func() {
	})
	Attribute("captchaId", String, "验证码ID", func() {
	})
	Required("image", "captchaId")
})
