package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Supplier", func() {
	Description("供应商服务")

	Error("internal_server_error", ErrorResult)
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/supplier")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("Get", func() {
		Description("获取单个供应商")
		Meta("swagger:summary", "获取单个供应商")

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

		Result(Supplier, func() {
			View("default")
		})

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("List", func() {
		Description("获取供应商列表")
		Meta("swagger:summary", "获取供应商列表")

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
			Field(1, "items", ArrayOf(Supplier, func() {
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
		Description("更新供应商")
		Meta("swagger:summary", "更新供应商")

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
			Attribute("name", String, "供应商名", func() {
				Example("xx公司")
			})
			Attribute("level", Int, "级别", func() {
				Enum(1, 2, 3, 4)
				Example(1)
			})
			Attribute("contact_name", String, "联系人姓名", func() {
				Example("张三")
			})
			Attribute("contact_phone", String, "联系人手机", func() {
				Format(FormatRegexp)
				Pattern(`^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$`)
				MaxLength(11)
				MinLength(11)
				Example("1808001010")
			})
			Attribute("contact_address", String, "联系人地址", func() {
				Example("咸阳")
			})
			Attribute("note", String, "产品备注", func() {
				Example("备注")
			})
			Attribute("head_id", String, "负责人ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})

			Required("token", "id")
		})

		Result(Supplier, func() {
			View("default")
		})

		HTTP(func() {
			PUT("/update")
			Response(StatusOK)
		})
	})

	Method("Create", func() {
		Description("创建供应商")
		Meta("swagger:summary", "创建供应商")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("name", String, "供应商名", func() {
				Example("xx公司")
			})
			Attribute("level", Int, "级别", func() {
				Enum(1, 2, 3, 4)
				Example(1)
			})
			Attribute("contact_name", String, "联系人姓名", func() {
				Example("张三")
			})
			Attribute("contact_phone", String, "联系人手机", func() {
				Format(FormatRegexp)
				Pattern(`^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$`)
				MaxLength(11)
				MinLength(11)
				Example("1808001010")
			})
			Attribute("contact_address", String, "联系人地址", func() {
				Example("咸阳")
			})
			Attribute("note", String, "产品备注", func() {
				Example("备注")
			})
			Attribute("head_id", String, "负责人ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("founder_id", String, "创建人ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})

			Required("token", "name", "level", "contact_name", "contact_phone", "contact_address", "note", "head_id", "founder_id")
		})

		Result(Supplier, func() {
			View("default")
		})

		HTTP(func() {
			POST("/create")
			Response(StatusOK)
		})
	})

	Method("Delete", func() {
		Description("删除供应商")
		Meta("swagger:summary", "删除供应商")

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
