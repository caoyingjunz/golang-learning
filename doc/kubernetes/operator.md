# Operator

### Mac 安装 operator-sdk

```
curl -LO https://github.com/operator-framework/operator-sdk/releases/download/v1.4.2/operator-sdk_darwin_amd64
chmod +x operator-sdk_darwin_amd64
sudo cp operator-sdk_darwin_amd64 /usr/local/go/bin/operator-sdk
```

### 初始化项目
* [Quickstart](https://sdk.operatorframework.io/docs/building-operators/golang/quickstart/)
* [Example](http://www.dockone.io/article/8733)
* [Fulltutorial](https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/)

```
mkdir podset-operator
cd podset-operator
operator-sdk init --domain example.com --repo github.com/example/podset-operator
```

### 创建 api

```
operator-sdk create api --group cache --version v1alpha1 --kind PodSet --resource --controller
```

### 自定义控制器代码
TODO

### 使用内置的 Makefile build and push image
```
export USERNAME=<quay-namespace>
export OPERATOR_IMG="quay.io/$USERNAME/memcached-operator:v0.0.1"
make docker-build docker-push IMG=$OPERATOR_IMG
```

### 部署 CRD
```
 kubectl apply -f config/samples/cache_v1alpha1_podset.yaml
```

