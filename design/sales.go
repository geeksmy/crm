package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Sales", func() {
	Description("销售服务")

	Error("internal_server_error", ErrorResult)
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/sales")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("Get", func() {
		Description("获取单个销售")
		Meta("swagger:summary", "获取单个销售")

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

		Result(Sales, func() {
			View("default")
		})

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("List", func() {
		Description("获取销售列表")
		Meta("swagger:summary", "获取销售列表")

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
			Field(1, "items", ArrayOf(Sales, func() {
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
		Description("更新销售")
		Meta("swagger:summary", "更新销售")

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
			Attribute("name", String, "销售单名", func() {
				Example("xx销售")
			})
			Attribute("customer_id", String, "客户ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("money", Int, "销售金额", func() {
				Example(123)
			})
			Attribute("consignment_date", String, "销售日期", func() {
				Example("2021-01-01")
			})
			Attribute("is_sales_return", Boolean, "销售还是退货", func() {
				Enum(false, true)
				Example(false)
			})
			Attribute("note", String, "备注", func() {
				Example("备注")
			})
			Attribute("head_id", String, "负责人ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})

			Required("token", "id", "name", "customer_id", "money", "is_sales_return", "note", "head_id")
		})

		Result(Sales, func() {
			View("default")
		})

		HTTP(func() {
			PUT("/update")
			Response(StatusOK)
		})
	})

	Method("Create", func() {
		Description("创建销售")
		Meta("swagger:summary", "创建销售")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("name", String, "销售单名", func() {
				Example("xx销售")
			})
			Attribute("code", String, "销售编码", func() {
				Example("123asd123123asd")
			})
			Attribute("customer_id", String, "客户ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("money", Int, "销售金额", func() {
				Example(123)
			})
			Attribute("consignment_date", String, "销售日期", func() {
				Example("2021-01-01")
			})
			Attribute("is_sales_return", Boolean, "销售还是退货", func() {
				Enum(false, true)
				Example(false)
			})
			Attribute("note", String, "备注", func() {
				Example("备注")
			})
			Attribute("head_id", String, "负责人", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("founder_id", String, "创建人", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})

			Required("token", "name", "code", "customer_id", "money", "consignment_date", "is_sales_return", "note", "head_id", "founder_id")
		})

		Result(Sales, func() {
			View("default")
		})

		HTTP(func() {
			POST("/create")
			Response(StatusOK)
		})
	})

	Method("Delete", func() {
		Description("删除销售")
		Meta("swagger:summary", "删除销售")

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
