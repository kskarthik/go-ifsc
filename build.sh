GOOS=linux GOARCH=amd64 go build -o  public/ifsc main.go
GOOS=windows GOARCH=amd64 go build -o  public/ifsc.exe main.go
GOOS=darwin GOARCH=amd64 go build -o  public/ifsc-darwin main.go
