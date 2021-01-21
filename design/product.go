package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Product", func() {
	Description("产品服务")

	Error("internal_server_error", ErrorResult)
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/product")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("Get", func() {
		Description("获取单个产品")
		Meta("swagger:summary", "获取单个产品")

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

		Result(Product, func() {
			View("default")
		})

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("List", func() {
		Description("获取产品列表")
		Meta("swagger:summary", "获取产品列表")

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
			Field(1, "items", ArrayOf(Product, func() {
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
		Description("更新产品")
		Meta("swagger:summary", "更新产品")

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
			Attribute("name", String, "产品名", func() {
				Example("灌装辣椒")
			})
			Attribute("unit", Int, "产品单位", func() {
				Example(1)
				Enum(1, 2, 3, 4, 5, 6, 7, 8)
			})
			Attribute("cost_price", Int, "成本价", func() {
				Example(123)
			})
			Attribute("market_price", Int, "市场价", func() {
				Example(123)
			})
			Attribute("note", String, "产品备注", func() {
				Example("备注")
			})
			Attribute("image", String, "产品图片", func() {
				Example("/images/123.jpg")
			})
			Attribute("size", String, "产品规格", func() {
				Example("瓶")
			})
			Attribute("type", Int, "产品类型", func() {
				Example(1)
				Enum(1, 2, 3, 4, 5, 6, 7, 8)
			})
			Attribute("is_shelves", Boolean, "是否上架", func() {
				Example(false)
			})

			Required("token", "id", "name", "unit", "cost_price", "market_price", "note", "image", "size", "type", "is_shelves")
		})

		Result(Product, func() {
			View("default")
		})

		HTTP(func() {
			PUT("/update")
			Response(StatusOK)
		})
	})

	Method("Create", func() {
		Description("创建产品")
		Meta("swagger:summary", "创建产品")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("name", String, "产品名", func() {
				Example("灌装辣椒")
			})
			Attribute("unit", Int, "产品单位", func() {
				Example(1)
				Enum(1, 2, 3, 4, 5, 6, 7, 8)
			})
			Attribute("cost_price", Int, "成本价", func() {
				Example(123)
			})
			Attribute("market_price", Int, "市场价", func() {
				Example(123)
			})
			Attribute("note", String, "产品备注", func() {
				Example("备注")
			})
			Attribute("image", String, "产品图片", func() {
				Example("/images/123.jpg")
			})
			Attribute("code", String, "产品编码", func() {
				Example("123asd123123asd")
			})
			Attribute("size", String, "产品规格", func() {
				Example("瓶")
			})
			Attribute("type", Int, "产品类型", func() {
				Example(1)
				Enum(1, 2, 3, 4, 5, 6, 7, 8)
			})
			Attribute("is_shelves", Boolean, "是否上架", func() {
				Example(false)
			})
			Attribute("founder_id", String, "创建人ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})

			Required("token", "name", "unit", "cost_price", "market_price", "note", "image", "code", "size", "type", "is_shelves", "founder_id")
		})

		Result(Product, func() {
			View("default")
		})

		HTTP(func() {
			POST("/create")
			Response(StatusOK)
		})
	})

	Method("Delete", func() {
		Description("删除产品")
		Meta("swagger:summary", "删除产品")

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
