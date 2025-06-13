module github.com/datauniverse-lab/earth-asd

go 1.23

replace (
	git.datau.co.kr/benz/benz-common => ../../../git.datau.co.kr/benz/benz-common
	git.datau.co.kr/ferrari/ferrari-common => ../../../git.datau.co.kr/ferrari/ferrari-common
	github.com/datauniverse-lab/earth-common => ../earth-common
	github.com/datauniverse-lab/tesla-common => ../tesla-common
)

require (
	git.datau.co.kr/benz/benz-common v0.0.0-00010101000000-000000000000
	git.datau.co.kr/ferrari/ferrari-common v0.0.0-00010101000000-000000000000
	github.com/datauniverse-lab/earth-common v0.0.0-00010101000000-000000000000
	github.com/datauniverse-lab/tesla-common v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.6.0
	github.com/utahta/go-cronowriter v1.2.0
	google.golang.org/grpc v1.49.0
)

require (
	github.com/beevik/guid v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lestrrat-go/strftime v0.0.0-20180220091553-9948d03c6207 // indirect
	github.com/pkg/errors v0.8.1-0.20180311214515-816c9085562c // indirect
	github.com/santhosh-tekuri/jsonschema/v5 v5.1.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/yannh/kubeconform v0.6.1 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
