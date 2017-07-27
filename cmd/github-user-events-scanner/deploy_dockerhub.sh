# go1.8.3 darwin/amd64
goversion=$(go version)

if [[ $goversion == *"go1.7"* ]]; then
  echo "Must use go 1.8 or later due to https://groups.google.com/d/msg/golang-nuts/9SaVxumSc-Y/rNAI8R7_BAAJ"
  exit 1
fi

env GOOS=linux GOARCH=amd64 go build -o action main.go && \
docker build -t tleyden5iwx/github-user-events-scanner . && \
docker push tleyden5iwx/github-user-events-scanner
