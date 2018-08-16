#!/bin/bash

VERSION="$(grep "version.*=" main.go | cut -d "\"" -f2 | cut -d "\"" -f1)"

go build
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -ldflags "-H windowsgui" -o ImgViewer.exe

tar cvzf img_viewer_v${VERSION}_linux_amd64.tgz img_viewer
7z a img_viewer_v${VERSION}_win_amd64.zip ImgViewer.exe

mv img_viewer /home/$USER/dev/go/bin/
mv ImgViewer.exe /home/$USER/comas/share/