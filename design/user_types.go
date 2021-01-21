package design

import . "goa.design/goa/v3/dsl"

var Superior = ResultType("Superior", func() {
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Example("519151ca-6250-4eec-8016-1e14a68dc448")
		})
		Attribute("name", String, "姓名", func() {
			Example("张三")
		})
		Required("id", "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
})

var Founder = ResultType("Founder", func() {
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Example("519151ca-6250-4eec-8016-1e14a68dc448")
		})
		Attribute("name", String, "姓名", func() {
			Example("张三")
		})
		Required("id", "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
})

var Head = ResultType("Head", func() {
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Example("519151ca-6250-4eec-8016-1e14a68dc448")
		})
		Attribute("name", String, "姓名", func() {
			Example("张三")
		})
		Required("id", "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
})

var User = ResultType("User", func() {
	Description("用户")
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Example("519151ca-6250-4eec-8016-1e14a68dc448")
		})
		Attribute("username", String, "用户名", func() {
			Example("用户名")
		})
		Attribute("name", String, "姓名", func() {
			Example("张三")
		})
		Attribute("mobile", String, "手机号", func() {
			Example("1808001010")
		})
		Attribute("email", String, "邮箱", func() {
			Example("adb@adb.com")
		})
		Attribute("jobs", Int, "职位", func() {
			Example(1)
			Enum(1, 2, 3)
			Description("1 - 推销员，2 - 经理，3 - 管理员")
		})
		Attribute("is_admin", Boolean, "是否是管理员", func() {
			Example(false)
		})
		Attribute("superior", Superior, "直属上级")
		Attribute("group", Group, "所属组")

		Required("id", "username", "name", "mobile", "email", "jobs", "is_admin", "superior", "group")
	})

	View("default", func() {
		Attribute("id")
		Attribute("username")
		Attribute("mobile")
		Attribute("name")
		Attribute("email")
		Attribute("jobs")
		Attribute("is_admin")
		Attribute("superior")
		Attribute("group")
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

var Captcha = ResultType("Captcha", func() {
	Description("图形验证码")
	ContentType("application/json")

	Attributes(func() {
		Attribute("image", String, "图片base64", func() {
		})
		Attribute("captchaId", String, "验证码ID", func() {
		})
		Required("image", "captchaId")
	})

	View("default", func() {
		Attribute("image")
		Attribute("captchaId")
	})
})
