package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("User", func() {
	Description("用户服务")

	Error("internal_server_error", ErrorResult)
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/user")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("Get", func() {
		Description("获取单个用户")
		Meta("swagger:summary", "获取单个用户")

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

			Required("token", "id")
		})

		Result(User, func() {
			View("default")
		})

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("List", func() {
		Description("获取用户列表")
		Meta("swagger:summary", "获取用户列表")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("cursor", Int, "cursor of page", func() {
				Example(0)
			})
			Attribute("limit", Int, "limit of items", func() {
				Example(20)
			})

			Required("token")
		})

		Result(func() {
			Field(1, "items", ArrayOf(User, func() {
				View("default")
			}))
			Field(2, "nextCursor", Int, "下一页游标", func() {
				Example(100)
			})
			Field(3, "total", Int, "总记录数")
			Required("items", "nextCursor", "total")
		})

		HTTP(func() {
			GET("")
			Params(func() {
				Param("cursor")
				Param("limit")
			})
			Response(StatusOK)
		})
	})

	Method("Update", func() {
		Description("更新用户")
		Meta("swagger:summary", "更新用户")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("id", String, "用户ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("name", String, "姓名", func() {
				Example("张三")
			})
			Attribute("mobile", String, "手机号", func() {
				Format(FormatRegexp)
				Pattern(`^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$`)
				MaxLength(11)
				MinLength(11)
				Example("1808001010")
			})
			Attribute("email", String, "邮箱", func() {
				Format(FormatEmail)
				Example("adb@adb.com")
			})
			Attribute("jobs", Int, "职位", func() {
				Example(1)
				Enum(1, 2, 3)
				Description("1 - 推销员，2 - 经理，3 - 管理员")
			})
			Attribute("superior_id", String, "直属上级ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("group_id", String, "所属组", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})

			Required("token", "id")
		})

		Result(User, func() {
			View("default")
		})

		HTTP(func() {
			PUT("/update")
			Response(StatusOK)
		})
	})

	Method("Create", func() {
		Description("创建用户")
		Meta("swagger:summary", "创建用户")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("username", String, "用户名", func() {
				Pattern(`(\d+.*[a-zA-Z]+)|([a-zA-Z]+.*\d+)`)
				Example("abc")
			})
			Attribute("password", String, "用户密码", func() {
				Example("abc")
			})
			Attribute("name", String, "姓名", func() {
				Example("张三")
			})
			Attribute("mobile", String, "手机号", func() {
				Format(FormatRegexp)
				Pattern(`^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$`)
				MaxLength(11)
				MinLength(11)
				Example("1808001010")
			})
			Attribute("email", String, "邮箱", func() {
				Format(FormatEmail)
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
			Attribute("superior_id", String, "直属上级ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("group_id", String, "所属组", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})

			Required("token", "username", "password", "name", "mobile", "email", "jobs", "is_admin", "superior_id", "group_id")
		})

		Result(User, func() {
			View("default")
		})

		HTTP(func() {
			POST("/create")
			Response(StatusOK)
		})
	})

	Method("Delete", func() {
		Description("删除用户")
		Meta("swagger:summary", "删除用户")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("ids", ArrayOf(String, func() {
				Example(ExampleUUID)
			}), func() {
				MaxLength(100)
			})

			Required("token", "ids")
		})

		Result(SuccessResult)

		HTTP(func() {
			DELETE("/delete")
			Response(StatusOK)
		})
	})
})
