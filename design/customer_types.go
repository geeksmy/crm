package design

import . "goa.design/goa/v3/dsl"

var Customer = ResultType("Customer", func() {
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Example("519151ca-6250-4eec-8016-1e14a68dc448")
		})
		Attribute("name", String, "客户姓名", func() {
			Example("张三")
		})
		Attribute("src", Int, "客户来源", func() {
			Example(1)
		})
		Attribute("mobile", String, "客户手机", func() {
			Example("18000001234")
		})
		Attribute("url", String, "客户网址", func() {
			Example("http://www.baidu.com")
		})
		Attribute("email", String, "客户邮箱", func() {
			Example("adb@adb.com")
		})
		Attribute("industry", Int, "客户行业", func() {
			Example(1)
		})
		Attribute("level", Int, "客户等级", func() {
			Example(1)
		})
		Attribute("note", String, "备注", func() {
			Example("备注")
		})
		Attribute("address", String, "客户地址", func() {
			Example("咸阳")
		})
		Required("id", "name", "src", "mobile", "url", "email", "industry", "level", "note", "address")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("src")
		Attribute("mobile")
		Attribute("url")
		Attribute("email")
		Attribute("industry")
		Attribute("level")
		Attribute("note")
		Attribute("address")
	})
})
