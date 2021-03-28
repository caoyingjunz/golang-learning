package main

import (
	"context"

	"google.golang.org/grpc"
)

// https://github.com/grpc/grpc-go/tree/master/examples

const _ = grpc.SupportPackageIsVersion7

// 客户端
type KubezClientInterface interface {
	GetName(ctx context.Context, opts ...grpc.CallOption) (*string, error)
}

type KubezClient struct {
	cc grpc.ClientConnInterface
}

func NewKubezClient(cc grpc.ClientConnInterface) KubezClientInterface {
	return &KubezClient{cc}
}

func (k *KubezClient) GetName(ctx context.Context, opts ...grpc.CallOption) (*string, error) {
	var rest *string
	if err := k.cc.Invoke(ctx, "/rpc.Kubez/GetName", rest, opts...); err != nil {
		return nil, err
	}

	return rest, nil
}

type KubezServerInterface interface {
	GetName(ctx context.Context)
}
