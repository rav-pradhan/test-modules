module github.com/eldeal/test-modules/render

go 1.13

replace github.com/eldeal/test-modules/render => /Users/eleanor/Development/git/personal/test-modules/render

replace github.com/eldeal/test-modules/render/assets => /Users/eleanor/Development/git/personal/test-modules/render/assets

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/ONSdigital/dp-frontend-models v1.10.2
	github.com/ONSdigital/go-ns v0.0.0-20200902154605-290c8b5ba5eb
	github.com/ONSdigital/log.go v1.0.1
	github.com/c2h5oh/datasize v0.0.0-20200825124411-48ed595a09d2
	github.com/gosimple/slug v1.9.0
	github.com/jteeuwen/go-bindata v3.0.7+incompatible // indirect
	github.com/nicksnyder/go-i18n/v2 v2.1.2
	github.com/russross/blackfriday/v2 v2.1.0
	github.com/unrolled/render v1.0.3
	golang.org/x/text v0.3.5
)
