module main

go 1.20

replace main/ystruct => ./ystruct

require (
	gopkg.in/yaml.v2 v2.4.0
	main/ystruct v0.0.0-00010101000000-000000000000
)

require (
	github.com/kr/pretty v0.3.1 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)
