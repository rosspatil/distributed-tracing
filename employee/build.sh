GOOS=linux go build employee -ldflags "-s -w"
upx employee
docker build . --build-arg BUILDKIT_INLINE_CACHE=1   -t employee
docker run --rm -p 8090:8090 -p 8091:8091 -p 8092:8092 --security-opt=seccomp:unconfined --name=employee employee:latest