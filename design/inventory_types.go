package design

import . "goa.design/goa/v3/dsl"

var Inventory = ResultType("Inventory", func() {
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Example("519151ca-6250-4eec-8016-1e14a68dc448")
		})
		Attribute("product", Product, "产品")
		Attribute("number", Int, "产品数量", func() {
			Example(123)
		})
		Attribute("code", String, "库存单号", func() {
			Example("123qwe123qwe")
		})
		Attribute("warehouse", Warehouse, "仓库")
		Attribute("type", Int, "库存类型", func() {
			Example(1)
		})
		Attribute("inventory_date", String, "库存时间", func() {
			Example("20210101")
		})
		Attribute("in_and_out", Int, "入库还是出库", func() {
			Example(1)
		})
		Attribute("note", String, "备注", func() {
			Example("备注")
		})
		Attribute("head", Head, "负责人")
		Attribute("founder", Founder, "创建人")
		Required("id", "product", "number", "code", "warehouse", "type", "inventory_date", "in_and_out", "note", "head", "founder")
	})

	View("default", func() {
		Attribute("id")
		Attribute("product")
		Attribute("number")
		Attribute("code")
		Attribute("warehouse")
		Attribute("type")
		Attribute("inventory_date")
		Attribute("in_and_out")
		Attribute("note")
		Attribute("head")
		Attribute("founder")
	})
})
