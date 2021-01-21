package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Inventory", func() {
	Description("库存服务")

	Error("internal_server_error", ErrorResult)
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/inventory")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("Get", func() {
		Description("获取单个库存")
		Meta("swagger:summary", "获取单个库存")

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

		Result(Inventory, func() {
			View("default")
		})

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("List", func() {
		Description("获取库存列表")
		Meta("swagger:summary", "获取库存列表")

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
			Field(1, "items", ArrayOf(Inventory, func() {
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
		Description("更新库存")
		Meta("swagger:summary", "更新库存")

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
			Attribute("product_id", String, "产品ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("number", Int, "产品数量", func() {
				Example(123)
			})
			Attribute("code", String, "库存单号", func() {
				Example("123qwe123qwe")
			})
			Attribute("warehouse_id", String, "仓库ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("type", Int, "库存类型", func() {
				Enum(1, 2, 3)
				Example(1)
			})
			Attribute("inventory_date", String, "库存时间", func() {
				Example("20210101")
			})
			Attribute("in_and_out", Int, "入库还是出库", func() {
				Enum(1, 2)
				Example(1)
			})
			Attribute("note", String, "备注", func() {
				Example("备注")
			})
			Attribute("head_id", String, "负责人ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})

			Required("token", "id", "product_id", "number", "code", "warehouse_id", "type", "inventory_date", "in_and_out", "note", "head_id")
		})

		Result(Inventory, func() {
			View("default")
		})

		HTTP(func() {
			PUT("/update")
			Response(StatusOK)
		})
	})

	Method("Create", func() {
		Description("创建库存")
		Meta("swagger:summary", "创建库存")

		Security(JWTAuth, func() {
			Scope("role:user")
		})

		Payload(func() {
			Token("token", String, func() {
				Description("JWT used for authentication")
				Example(ExampleJwt)
			})
			Attribute("product_id", String, "产品ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("number", Int, "产品数量", func() {
				Example(123)
			})
			Attribute("code", String, "库存单号", func() {
				Example("123qwe123qwe")
			})
			Attribute("warehouse_id", String, "仓库ID", func() {
				Example("519151ca-6250-4eec-8016-1e14a68dc448")
			})
			Attribute("type", Int, "库存类型", func() {
				Enum(1, 2, 3)
				Example(1)
			})
			Attribute("inventory_date", String, "库存时间", func() {
				Example("20210101")
			})
			Attribute("in_and_out", Int, "入库还是出库", func() {
				Enum(1, 2)
				Example(1)
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

			Required("token", "product_id", "number", "code", "warehouse_id", "type", "inventory_date", "in_and_out", "note", "head_id", "founder_id")
		})

		Result(Inventory, func() {
			View("default")
		})

		HTTP(func() {
			POST("/create")
			Response(StatusOK)
		})
	})

	Method("Delete", func() {
		Description("删除库存")
		Meta("swagger:summary", "删除库存")

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
