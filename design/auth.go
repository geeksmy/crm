package design

import . "goa.design/goa/v3/dsl"

var _ = Service("Auth", func() {
	Description("鉴权服务")

	Error("internal_server_error", ErrorResult)
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/auth")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("Login", func() {
		Description("使用账号密码登录")
		Meta("swagger:summary", "使用账号密码登录")

		Payload(func() {
			Attribute("username", String, "用户名", func() {
				Example("user")
				MinLength(1)
				MaxLength(128)
			})
			Attribute("password", String, "密码", func() {
				Example("password")
				MinLength(1)
				MaxLength(128)
			})
			Attribute("humanCode", String, "图形验证码", func() {
				Default("")
				MinLength(4)
				MaxLength(8)
			})
			Attribute("captchaId", String, "图形验证码ID", func() {
				Default("")
				MinLength(1)
				MaxLength(128)
			})
			Required("username", "password")
		})

		Result(Session, func() {
			View("default")
		})

		HTTP(func() {
			POST("/login")
			Response(StatusOK)
		})
	})

	Method("UpdatePassword", func() {
		Description("修改登录密码")
		Meta("swagger:summary", "修改登录密码")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("id", String, func() {
				Example(ExampleUUID)
			})
			Attribute("old_password", String, func() {
				MinLength(1)
				Example("123abc")
				MaxLength(128)
			})
			Attribute("new_password", String, func() {
				MinLength(8)
				Example("abc123")
				MaxLength(128)
			})
			Required("token", "id", "old_password", "new_password")
		})

		Result(SuccessResult, func() {
			View("default")
		})

		HTTP(func() {
			POST("/update_password")
			Response(StatusOK)
		})
	})

	Method("CaptchaImage", func() {
		Description("获取图形验证码")
		Meta("swagger:summary", "获取图形验证码")

		Result(Captcha, func() {
			View("default")
		})

		HTTP(func() {
			GET("/captcha_image")
			Response(StatusOK)
		})
	})

})
