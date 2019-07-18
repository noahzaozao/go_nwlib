module github.com/noahzaozao/go_nwlib

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.8
	github.com/micro/go-config v1.1.0
	github.com/micro/go-micro v1.7.0 // indirect
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	google.golang.org/grpc v1.21.1
)

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f
