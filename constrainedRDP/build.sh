#!/usr/bin/env bash

echo "building linux binary"
go build -o ../dist/constdp/constdp
chmod +x ../dist/constdp/constdp
echo "..........done............."

echo "building windows binary"
GOOS=windows go build -o ../dist/constdp/constdp.exe
echo "..........done............."

\cp -rf ./resource/cfg.toml ../dist/constdp/
node docs.js
