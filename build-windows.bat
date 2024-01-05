@echo off
set @GOOS=windows
set GOARCH=amd64
gogio -buildmode=exe -icon=appicon.png -arch=amd64 -target=windows -o BookmarkViewer.exe .
pause
