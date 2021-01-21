package design

import . "goa.design/goa/v3/dsl"

var (
	ExampleJwt  = "eyJhbGciOiJIUz..."
	ExampleUUID = "91cc3eb9-ddc0-4cf7-a62b-c85df1a9166f"
)

var SuccessResult = ResultType("SuccessResult", func() {
	Description("成功信息")
	ContentType("application/json")
	TypeName("SuccessResult")

	Attributes(func() {
		Attribute("ok", Boolean, "success", func() {
			Example(true)
		})
		Required("ok")
	})

	View("default", func() {
		Attribute("ok")
	})
})
