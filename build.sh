GOOS=linux GOARCH=amd64 go build -o pkg/go-data-client
echo "go build: OK"

img="hub.caishuo.com/go-data-client:latest"
docker build $docker_opts -t $img .
echo "docker build: OK"

docker push $img
echo "docker push: OK"
