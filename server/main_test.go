package main

import (
	"context"
	"reflect"
	"testing"

	hello_world "github.com/cwyu57/grpc-go-hello-world/hello-world"
)

func Test_server_SayHello(t *testing.T) {
	type fields struct {
		UnimplementedGreeterServer hello_world.UnimplementedGreeterServer
	}
	type args struct {
		ctx context.Context
		in  *hello_world.HelloRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hello_world.HelloResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "World",
			fields: fields{},
			args: args{
				ctx: context.TODO(),
				in:  &hello_world.HelloRequest{Name: "World"},
			},
			want: &hello_world.HelloResponse{
				Message: "Hello World",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedGreeterServer: tt.fields.UnimplementedGreeterServer,
			}
			got, err := s.SayHello(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
