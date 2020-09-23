## build go binary in docker
# alpine은 2번처럼 build해야 함
# 1. GOOS=linux go build main.go
# 2. env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/service service/main.go
# 3. env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o ./main ./main.go
##
# FIXME x aws configure 해결해야 함

$ContainerName = $args[0]
$LambdaFuncName = "hello-lambda-golang"
$S3BucketName = "hello-lambda-golang"

docker exec $ContainerName sh -c "go env"

docker exec $ContainerName sh -c "env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o main main.go &&
zip function.zip main &&
aws lambda update-function-code --function-name $LambdaFuncName \
  --zip-file fileb://function.zip"
