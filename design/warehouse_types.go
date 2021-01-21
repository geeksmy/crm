package design

import . "goa.design/goa/v3/dsl"

var Warehouse = ResultType("Warehouse", func() {
	ContentType("application/json")

	Attributes(func() {
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
			Example(1)
		})
		Attribute("founder", Founder, "创建人")
		Required("id", "name", "code", "address", "type", "founder")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("code")
		Attribute("address")
		Attribute("type")
		Attribute("founder")
	})
})
