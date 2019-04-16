set GOOS=windows
packr
go build -ldflags="-s -w"
packr clean
del /f TestCompre.exe
D:\upx\upx.exe -9 -oTestCompre.exe "%~dp0Test.exe"
PAUSE