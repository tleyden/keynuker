env GOOS=linux GOARCH=amd64 go build -o exec main.go && \
zip action.zip exec && \
wsk action delete monitor-activations && \
wsk action create monitor-activations --timeout 300000 --web true --docker tleyden5iwx/openwhisk-dockerskeleton action.zip