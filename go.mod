module github.com/noahzaozao/go_nwlib

go 1.12

require (
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.8
	github.com/micro/go-config v1.1.0
)

replace github.com/sourcegraph/go-diff => github.com/sourcegraph/go-diff v0.5.1

replace github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3 => golang.org/x/lint v0.0.0-20190409202823-5614ed5bae6fb75893070bdc0996a68765fdd275

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f
