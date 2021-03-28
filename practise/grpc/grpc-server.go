package main

import (
	"context"

	"google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion7

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
