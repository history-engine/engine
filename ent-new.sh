#!/bin/bash

if [ ! -n "$1" ];then
    echo "please input name, example: User"
    exit 0
fi

go run -mod=mod entgo.io/ent/cmd/ent new --target ./ent/schema $1
