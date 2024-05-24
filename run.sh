#!/bin/sh

R2AID=$(cat .R2AID) \
R2TOK=$(cat .R2TOK) \
R2SRV=$(cat .R2SRV) \
R2DEV=1 \
    go run cmd/R2/R2.go
