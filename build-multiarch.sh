GOOS=windows GOARCH=amd64 go build -o bin/win_amd64/arm-xxd arm-xxd.go 
GOOS=darwin GOARCH=arm64 go build -o bin/mac_arm/arm-xxd arm-xxd.go 
GOOS=linux GOARCH=arm64 go build -o bin/linux_arm/arm-xxd arm-xxd.go