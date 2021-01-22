package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Warehouse", func() {
	Description("仓库服务")

	Error("internal_server_error", ErrorResult)
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/warehouse")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("Get", func() {
		Description("获取单个仓库")
		Meta("swagger:summary", "获取单个仓库")

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

		Result(Warehouse, func() {
			View("default")
		})

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("List", func() {
		Description("获取仓库列表")
		Meta("swagger:summary", "获取仓库列表")

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
			Field(1, "items", ArrayOf(Warehouse, func() {
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
		Description("更新仓库")
		Meta("swagger:summary", "更新仓库")

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
			Attribute("name", String, "仓库名", func() {
				Example("1号仓库")
			})
			Attribute("code", String, "仓库编码", func() {
				Example("123awe12qwe")
			})
			Attribute("address", String, "仓库地址", func() {
				Example("地址")
			})
			Attribute("type", Int, "仓库状态", func() {
				Enum(1, 2, 3)
				Example(1)
			})

			Required("token", "id")
		})

		Result(Warehouse, func() {
			View("default")
		})

		HTTP(func() {
			PUT("/update")
			Response(StatusOK)
		})
	})

	Method("Create", func() {
		Description("创建仓库")
		Meta("swagger:summary", "创建仓库")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("name", String, "仓库名", func() {
				Example("1号仓库")
			})
			Attribute("code", String, "仓库编码", func() {
				Example("123awe12qwe")
			})
			Attribute("address", String, "仓库地址", func() {
				Example("地址")
			})
			Attribute("type", Int, "仓库状态", func() {
				Enum(1, 2, 3)
				Example(1)
			})
			Attribute("founder_id", String, "创建人ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})

			Required("token", "name", "code", "address", "type", "founder_id")
		})

		Result(Warehouse, func() {
			View("default")
		})

		HTTP(func() {
			POST("/create")
			Response(StatusOK)
		})
	})

	Method("Delete", func() {
		Description("删除仓库")
		Meta("swagger:summary", "删除仓库")

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
