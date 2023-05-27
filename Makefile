clean:
	rm -rf build

binaries: clean
	env GOOS=linux GOARCH=amd64 go build -o build/gptsh-linux-amd64 .
	env GOOS=darwin GOARCH=arm64 go build -o build/gptsh-macos-arm64 .
	env GOOS=windows GOARCH=arm64 go build -o build/gptsh-windows-amd64 .