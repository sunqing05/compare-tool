#!/bin/sh

cur_dir=$PWD

build() {
    bin="$1"
    o="$4"

    mkdir -p "$bin/data"

    echo "building..."

    CGO_ENABLED=0 GOOS=$3 GOARCH=$2 go build -o "$bin/$o"
    
    # 复制配置文件
    src="$cur_dir/conf.yaml"
    cp -f $src "$bin/conf.yaml"
}

build $1 $2 $3 $4
