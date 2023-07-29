#!/bin/sh

protoc --proto_path=./proto\
    --go_out=./\
    --go_opt=Mmessages.proto=./protoc \
    ./proto/messages.proto
