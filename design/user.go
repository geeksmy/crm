package design

import . "goa.design/goa/v3/dsl"

var _ = Service("User", func() {
	Description("微服务")

	Error("internal_server_error", ErrorResult)
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/user")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("LoginByUsername", func() {
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

		Result(func() {
			Attribute("errcode", Int, "错误码", func() {
				Minimum(0)
				Maximum(999999)
				Example(0)
			})
			Attribute("errmsg", String, "错误消息", func() {
				Example("")
			})
			Attribute("data", Session)
			Required("errcode", "errmsg")
		})

		HTTP(func() {
			POST("/login_by_username")
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
			Attribute("old_password", String, func() {
				MinLength(1)
				Example("123abc")
				MaxLength(128)
			})
			Attribute("new_password", String, func() {
				MinLength(6)
				Example("abc123")
				MaxLength(128)
			})
			Required("token", "old_password", "new_password")
		})

		Result(func() {
			Attribute("errcode", Int, "错误码", func() {
				Minimum(0)
				Maximum(999999)
				Example(0)
			})
			Attribute("errmsg", String, "错误消息", func() {
				Example("")
			})
			Attribute("data", SuccessResult)
			Required("errcode", "errmsg")
		})

		HTTP(func() {
			POST("/update_password")
			Response(StatusOK)
		})
	})

	Method("GetCaptchaImage", func() {
		Description("获取图形验证码")
		Meta("swagger:summary", "获取图形验证码")

		Result(func() {
			Attribute("errcode", Int, "错误码", func() {
				Minimum(0)
				Maximum(999999)
				Example(0)
			})
			Attribute("errmsg", String, "错误消息", func() {
				Example("")
			})
			Attribute("data", Captcha)
			Required("errcode", "errmsg")
		})
		HTTP(func() {
			POST("/get_captcha_image")
			Response(StatusOK)
		})
	})

})
