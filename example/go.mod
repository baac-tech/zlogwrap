module example

go 1.16

require (
	github.com/gofiber/fiber/v2 v2.15.0
	github.com/rs/zerolog v1.23.0
	ipanda.baac.tech/golib/zlogwrap v0.0.0-00010101000000-000000000000
)

replace ipanda.baac.tech/golib/zlogwrap => ../
