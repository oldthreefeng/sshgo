# build.sh v3.0.2
COMMIT_SHA1=$(git rev-parse --short HEAD || echo "0.0.0")
BUILD_TIME=$(date "+%F %T")
go build -ldflags "-X github.com/oldhtreefeng/sshgo/cmd.Version=$1 -X github.com/oldhtreefeng/sshgo/cmd.Build=${COMMIT_SHA1} -X 'github.com/oldhtreefeng/sshgo/cmd.BuildTime=${BUILD_TIME}'" sshgo.go
