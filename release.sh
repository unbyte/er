GOOS=linux go build -ldflags "-w -s" -o release/er_linux cmd/er/main.go
GOOS=darwin go build -ldflags "-w -s" -o release/er_osx cmd/er/main.go
GOOS=windows go build -ldflags "-w -s" -o release/er_windows.exe cmd/er/main.go