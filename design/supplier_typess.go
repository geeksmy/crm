package design

import . "goa.design/goa/v3/dsl"

var Supplier = ResultType("Supplier", func() {
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Example("519151ca-6250-4eec-8016-1e14a68dc448")
		})
		Attribute("name", String, "供应商名", func() {
			Example("xx公司")
		})
		Attribute("level", Int, "级别", func() {
			Enum(1, 2, 3, 4)
			Example(1)
		})
		Attribute("contact_name", String, "联系人姓名", func() {
			Example("张三")
		})
		Attribute("contact_phone", String, "联系人手机", func() {
			Example("1808001010")
		})
		Attribute("contact_address", String, "联系人地址", func() {
			Example("咸阳")
		})
		Attribute("note", String, "产品备注", func() {
			Example("备注")
		})
		Attribute("head", Head, "负责人")
		Attribute("founder", Founder, "创建人")

		Required("id", "name", "level", "contact_name", "contact_phone", "contact_address", "note", "head", "founder")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("level")
		Attribute("contact_name")
		Attribute("contact_phone")
		Attribute("contact_address")
		Attribute("note")
		Attribute("head")
		Attribute("founder")
	})
})
