package hostclient

import (
	"context"

	"piblog/src/pb/host"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeleteBlogReq  = host.DeleteBlogReq
	DeleteBlogResp = host.DeleteBlogResp
	DeletePicResp  = host.DeletePicResp
	UploadBlogReq  = host.UploadBlogReq
	UploadBlogResp = host.UploadBlogResp
	UploadPicReq   = host.UploadPicReq
	UploadPicResp  = host.UploadPicResp

	Host interface {
		CreateBlog(ctx context.Context, in *UploadBlogReq, opts ...grpc.CallOption) (*UploadBlogResp, error)
		UpdateBlog(ctx context.Context, in *UploadBlogReq, opts ...grpc.CallOption) (*UploadBlogResp, error)
		DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogResp, error)
		UploadPic(ctx context.Context, in *UploadPicReq, opts ...grpc.CallOption) (*UploadPicResp, error)
		DeletePic(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeletePicResp, error)
	}

	defaultHost struct {
		cli zrpc.Client
	}
)

func NewHost(cli zrpc.Client) Host {
	return &defaultHost{
		cli: cli,
	}
}

func (m *defaultHost) CreateBlog(ctx context.Context, in *UploadBlogReq, opts ...grpc.CallOption) (*UploadBlogResp, error) {
	client := host.NewHostClient(m.cli.Conn())
	return client.CreateBlog(ctx, in, opts...)
}

func (m *defaultHost) UpdateBlog(ctx context.Context, in *UploadBlogReq, opts ...grpc.CallOption) (*UploadBlogResp, error) {
	client := host.NewHostClient(m.cli.Conn())
	return client.UpdateBlog(ctx, in, opts...)
}

func (m *defaultHost) DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogResp, error) {
	client := host.NewHostClient(m.cli.Conn())
	return client.DeleteBlog(ctx, in, opts...)
}

func (m *defaultHost) UploadPic(ctx context.Context, in *UploadPicReq, opts ...grpc.CallOption) (*UploadPicResp, error) {
	client := host.NewHostClient(m.cli.Conn())
	return client.UploadPic(ctx, in, opts...)
}

func (m *defaultHost) DeletePic(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeletePicResp, error) {
	client := host.NewHostClient(m.cli.Conn())
	return client.DeletePic(ctx, in, opts...)
}
