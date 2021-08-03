module github.com/kpango/ntokend

go 1.17

replace (
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/lint => golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/net => golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985
	golang.org/x/sync => golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys => golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c
	golang.org/x/text => golang.org/x/text v0.3.6
	golang.org/x/tools => golang.org/x/tools v0.1.5
	golang.org/x/xerrors => golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

require (
	github.com/AthenZ/athenz v1.10.28
	github.com/kpango/fastime v1.0.17
	github.com/kpango/glg v1.6.4
	github.com/pkg/errors v0.9.1
)

require (
	github.com/goccy/go-json v0.7.4 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
