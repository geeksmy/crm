package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Customer", func() {
	Description("客户服务")

	Error("internal_server_error", ErrorResult)
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/customer")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("Get", func() {
		Description("获取单个客户")
		Meta("swagger:summary", "获取单个客户")

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

		Result(Customer, func() {
			View("default")
		})

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("List", func() {
		Description("获取客户列表")
		Meta("swagger:summary", "获取客户列表")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})

			Required("token")
		})

		Result(func() {
			Field(1, "items", ArrayOf(Customer, func() {
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
			Response(StatusOK)
		})
	})

	Method("Update", func() {
		Description("更新客户")
		Meta("swagger:summary", "更新客户")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("id", String, "ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("name", String, "客户姓名", func() {
				Example("张三")
			})
			Attribute("src", Int, "客户来源", func() {
				Example(1)
			})
			Attribute("mobile", String, "客户手机", func() {
				Format(FormatRegexp)
				Pattern(`^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$`)
				MaxLength(11)
				MinLength(11)
				Example("18000001234")
			})
			Attribute("url", String, "客户网址", func() {
				Format(FormatURI)
				Example("http://www.baidu.com")
			})
			Attribute("email", String, "客户邮箱", func() {
				Format(FormatEmail)
				Example("adb@adb.com")
			})
			Attribute("industry", Int, "客户行业", func() {
				Example(1)
			})
			Attribute("level", Int, "客户等级", func() {
				Example(1)
			})
			Attribute("note", String, "备注", func() {
				Example("备注")
			})
			Attribute("address", String, "客户地址", func() {
				Example("咸阳")
			})

			Required("token", "id", "name", "src", "mobile", "url", "email", "industry", "level", "note", "address")
		})

		Result(Customer, func() {
			View("default")
		})

		HTTP(func() {
			PUT("/update")
			Response(StatusOK)
		})
	})

	Method("Create", func() {
		Description("创建客户")
		Meta("swagger:summary", "创建客户")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("name", String, "客户姓名", func() {
				Example("张三")
			})
			Attribute("src", Int, "客户来源", func() {
				Example(1)
			})
			Attribute("mobile", String, "客户手机", func() {
				Format(FormatRegexp)
				Pattern(`^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$`)
				MaxLength(11)
				MinLength(11)
				Example("18000001234")
			})
			Attribute("url", String, "客户网址", func() {
				Format(FormatURI)
				Example("http://www.baidu.com")
			})
			Attribute("email", String, "客户邮箱", func() {
				Format(FormatEmail)
				Example("adb@adb.com")
			})
			Attribute("industry", Int, "客户行业", func() {
				Example(1)
			})
			Attribute("level", Int, "客户等级", func() {
				Example(1)
			})
			Attribute("note", String, "备注", func() {
				Example("备注")
			})
			Attribute("address", String, "客户地址", func() {
				Example("咸阳")
			})

			Required("token", "name", "src", "mobile", "url", "email", "industry", "level", "note", "address")
		})

		Result(Customer, func() {
			View("default")
		})

		HTTP(func() {
			POST("/create")
			Response(StatusOK)
		})
	})

	Method("Delete", func() {
		Description("删除客户")
		Meta("swagger:summary", "删除客户")

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
