go get github.com/cloudwego/hertz/cmd/hz
go install github.com/cloudwego/hertz/cmd/hz

cd ..
# hz new --idl=./idl/api.thrift --handler_by_method -t=template=slim
hz update --idl=./idl/api.thrift --handler_by_method -t=template=slim