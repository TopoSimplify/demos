#!/usr/bin/env bash
dpDIR=$PWD
tmpDIR=../dist/constdp/tmp
mkdir -p ${tmpDIR}

distLinux=../dist/constdp/linux
distWin=../dist/constdp/windows
distMac=../dist/constdp/mac


#goos exec ext
function build_archive () {
    goos=$1
    exec=$2
    ext=$3
    arch_postfix=$4

    echo "building $goos binary"
    GOOS=${goos} GOARCH=amd64 go build -o ${tmpDIR}/${exec}${ext}
    GOOS=${goos} GOARCH=386   go build -o ${tmpDIR}/${exec}_32bit${ext}
    chmod +x ${tmpDIR}/${exec}${ext}
    chmod +x ${tmpDIR}/${exec}_32bit${ext}

    \cp -rf ./resource/config.toml ${tmpDIR}/

    cd ${tmpDIR}
        zip -r constdp_${arch_postfix}.zip *
        mv constdp_${arch_postfix}.zip ../
        rm -rf *
    cd ${dpDIR}
    echo ".................done - $goos.................."
}

build_archive "linux" "constdp" "" "linux"
build_archive "windows" "constdp" ".exe" "windows"
build_archive "darwin" "constdp" "" "mac"

node docs.js
