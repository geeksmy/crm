package design

import . "goa.design/goa/v3/dsl"

var Procurement = ResultType("Procurement", func() {
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Example("519151ca-6250-4eec-8016-1e14a68dc448")
		})
		Attribute("supplier", Supplier, "供应商")
		Attribute("code", String, "采购编码", func() {
			Example("123asd123123asd")
		})
		Attribute("money", Int, "采购金额", func() {
			Example(123)
		})
		Attribute("is_sales_return", Boolean, "采购还是退货", func() {
			Example(false)
		})
		Attribute("note", String, "产品备注", func() {
			Example("备注")
		})
		Attribute("head", Head, "负责人")
		Attribute("founder", Founder, "创建人")

		Required("id", "supplier", "code", "money", "is_sales_return", "note", "head", "founder")
	})

	View("default", func() {
		Attribute("id")
		Attribute("supplier")
		Attribute("code")
		Attribute("money")
		Attribute("is_sales_return")
		Attribute("note")
		Attribute("head")
		Attribute("founder")
	})
})
