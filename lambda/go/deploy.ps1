## build go binary in docker
# alpine은 2번처럼 build해야 함
# 1. GOOS=linux go build main.go
# 2. env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/service service/main.go
# 3. env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o ./main ./main.go
##
# FIXME x aws configure 해결해야 함

$defaultContainerName = "demo"
$containerName = If ($args[0] -eq $null) { $defaultContainerName } Else { $args[0] }
$lambdaFuncName = "hello-lambda-golang"
$s3BucketName = "hello-lambda-golang"

docker exec $containerName sh -c "go env"

docker exec $containerName sh -c "env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o main main.go &&
zip function.zip main &&
aws lambda update-function-code --function-name $lambdaFuncName \
  --zip-file fileb://function.zip"
