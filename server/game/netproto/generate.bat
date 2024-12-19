:: 单个生成 protoc --go_out=../pb msgresponse.proto
:: 批量生成 protoc --go_out=. *.proto

protoc --go_out=. *.proto

pause