module github.com/deislabs/smi-sdk-go

go 1.12

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gogo/protobuf v1.2.0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/kubernetes-client/go v0.0.0-20190625181339-cd8e39e789c7
	github.com/stretchr/testify v1.3.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190720062849-3043179095b6
	k8s.io/apimachinery v0.0.0-20190719140911-bfcf53abc9f8
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20190712204705-3dccf664f023 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190718062839-c8a0b81cb10e
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190717022731-0bb8574e0887
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190717023132-0c47f9da0001
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190717022600-77f3a1fe56bb
)
