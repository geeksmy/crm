package design

import . "goa.design/goa/v3/dsl"

var Product = ResultType("Product", func() {
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID", func() {
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
		Attribute("founder", Founder, "创建人")
		Required("id", "name", "unit", "cost_price", "market_price", "note", "image", "code", "size", "type", "is_shelves", "founder")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("unit")
		Attribute("cost_price")
		Attribute("market_price")
		Attribute("note")
		Attribute("image")
		Attribute("code")
		Attribute("size")
		Attribute("type")
		Attribute("is_shelves")
		Attribute("founder")
	})
})
