GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o  public/linux/ifsc main.go
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o  public/win/ifsc.exe main.go
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o  public/darwin/ifsc main.go
