module example

go 1.18

require (
	github.com/gofiber/fiber/v2 v2.34.0
	github.com/rs/zerolog v1.26.1
	ipanda.baac.tech/golib/zlogwrap v0.0.0-00010101000000-000000000000
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/klauspost/compress v1.15.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.37.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
)

replace ipanda.baac.tech/golib/zlogwrap => ../
