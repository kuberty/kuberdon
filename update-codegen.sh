#/bin/sh
#Attribution to @trstringer [https://medium.com/@trstringer/create-kubernetes-controllers-for-core-and-custom-resources-62fc35ad64a3]('s post)

ROOT_PACKAGE="github.com/kuberty/kuberdon"
CUSTOM_RESOURCE_NAME="registry"
CUSTOM_RESOURCE_VERSION="v1beta1"


mkdir $GOPATH/src/$ROOT_PACKAGE -p
cp -rf . $GOPATH/src/$ROOT_PACKAGE/

go get -u k8s.io/code-generator/...


cd $GOPATH/src/k8s.io/code-generator

./generate-groups.sh all "$ROOT_PACKAGE/pkg/client" "$ROOT_PACKAGE/pkg/apis" "$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"
