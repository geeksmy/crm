package design

import . "goa.design/goa/v3/dsl"

var Sales = ResultType("Sales", func() {
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Example("519151ca-6250-4eec-8016-1e14a68dc448")
		})
		Attribute("name", String, "销售单名", func() {
			Example("xx销售")
		})
		Attribute("code", String, "销售编码", func() {
			Example("123asd123123asd")
		})
		Attribute("customer", Customer, "客户")
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
		Attribute("head", Head, "负责人")
		Attribute("founder", Founder, "创建人")
		Required("id", "name", "code", "customer", "money", "consignment_date", "is_sales_return", "note", "head", "founder")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("code")
		Attribute("customer")
		Attribute("money")
		Attribute("consignment_date")
		Attribute("is_sales_return")
		Attribute("note")
		Attribute("head")
		Attribute("founder")
	})
})
