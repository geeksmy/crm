package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Procurement", func() {
	Description("采购服务")

	Error("internal_server_error", ErrorResult)
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/procurement")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("Get", func() {
		Description("获取单个采购")
		Meta("swagger:summary", "获取单个采购")

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

		Result(Procurement, func() {
			View("default")
		})

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("List", func() {
		Description("获取采购列表")
		Meta("swagger:summary", "获取采购列表")

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
			Field(1, "items", ArrayOf(Procurement, func() {
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
		Description("更新采购")
		Meta("swagger:summary", "更新采购")

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
			Attribute("code", String, "采购编码", func() {
				Example("123asd123123asd")
			})
			Attribute("money", Int, "采购金额", func() {
				Example(123)
			})
			Attribute("is_sales_return", Boolean, "采购还是退货", func() {
				Enum(false, true)
				Example(false)
			})
			Attribute("note", String, "备注", func() {
				Example("备注")
			})
			Attribute("head_id", String, "负责人ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})

			Required("token", "id")
		})

		Result(Procurement, func() {
			View("default")
		})

		HTTP(func() {
			PUT("/update")
			Response(StatusOK)
		})
	})

	Method("Create", func() {
		Description("创建采购")
		Meta("swagger:summary", "创建采购")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("supplier_id", String, "供应商ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("code", String, "采购编码", func() {
				Example("123asd123123asd")
			})
			Attribute("money", Int, "采购金额", func() {
				Example(123)
			})
			Attribute("is_sales_return", Boolean, "采购还是退货", func() {
				Enum(false, true)
				Example(false)
			})
			Attribute("note", String, "备注", func() {
				Example("备注")
			})
			Attribute("head_id", String, "负责人ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("founder_id", String, "创建人ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})

			Required("token", "supplier_id", "code", "money", "is_sales_return", "note", "head_id", "founder_id")
		})

		Result(Procurement, func() {
			View("default")
		})

		HTTP(func() {
			POST("/create")
			Response(StatusOK)
		})
	})

	Method("Delete", func() {
		Description("删除采购")
		Meta("swagger:summary", "删除采购")

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
