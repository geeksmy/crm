package design

import . "goa.design/goa/v3/dsl"

var Group = ResultType("Group", func() {
	ContentType("application/json")

	Attributes(func() {
		Attribute("id", String, "ID", func() {
			Example("519151ca-6250-4eec-8016-1e14a68dc448")
		})
		Attribute("name", String, "组名", func() {
			Example("普通用户")
		})
		Required("id", "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
})
