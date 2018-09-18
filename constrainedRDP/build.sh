#!/usr/bin/env bash
distLinux=../dist/constdp/linux
distWin=../dist/constdp/windows
distMac=../dist/constdp/mac

echo "building linux binary"
GOOS=linux GOARCH=amd64 go build -o ${distLinux}/constdp
GOOS=linux GOARCH=386   go build -o ${distLinux}/constdp_32bit

chmod +x ${distLinux}/constdp
chmod +x ${distLinux}/constdp_32bit
echo "....................done.................."

echo "building windows binary"
GOOS=windows GOARCH=amd64 go build -o ${distWin}/constdp.exe
GOOS=windows GOARCH=386   go build -o ${distWin}/constdp_32bit.exe
echo "...................done.................."

echo "building mac binary"
GOOS=darwin GOARCH=amd64 go build -o ${distMac}/constdp
GOOS=darwin GOARCH=386   go build -o ${distMac}/constdp_32bit
echo "...................done.................."

for dist in ${distMac}  ${distWin} ${distLinux}
do
    \cp -rf ./resource/config.toml ${dist}/
done

node docs.js
