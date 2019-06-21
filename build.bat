SET CGO_LDFLAGS=-Wl,--kill-at
SET GOOS=windows
SET GOARCH=386
go build -ldflags "-s -w" -buildmode=c-shared -o app.dll