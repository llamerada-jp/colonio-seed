module github.com/colonio/colonio-seed/src/seed

go 1.13

require (
	github.com/colonio/colonio-seed/src/proto v0.0.0
	github.com/gobwas/httphead v0.0.0-20180130184737-2c6c146eadee // indirect
	github.com/gobwas/pool v0.2.0 // indirect
	github.com/gobwas/ws v1.0.2
	github.com/golang/protobuf v1.3.2
	github.com/mailru/easygo v0.0.0-20190618140210-3c14a0dc985f
	golang.org/x/sys v0.0.0-20191029155521-f43be2a4598c // indirect
)

replace github.com/colonio/colonio-seed/src/proto => ../proto
