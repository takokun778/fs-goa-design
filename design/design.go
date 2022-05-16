package design

import . "goa.design/goa/v3/dsl"

// OPEN API のタイトル的なやつら
var _ = API("Hello User API", func() {
	Title("Hello User Service")
	Version("0.0.1")
	Description("Hello User Servise（APIの説明）")
	Server("hello", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

// サービスの単位（DDDにおけるドメインモデル単位くらいのニュアンスでもいいかと）
var _ = Service("user", func() {
	Description("User Serviceドメインモデル単位くらいのニュアンス")
	Method("hello", func() {
		Payload(HelloRequestBody)
		Result(HelloResponseBody)
		HTTP(func() {
			POST("/hello")
		})
	})
})

var HelloRequestBody = Type("HelloRequestBody", func() {
	Attribute("name", String, "user name")
})

var HelloResponseBody = Type("HelloResponseBody", func() {
	Attribute("greet", String, "greet result")
})
