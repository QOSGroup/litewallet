#!/bin/bash

rm -f  src/main/resources/liblitewallet.so
rm -f  src/main/resources/liblitewallet.dll

go build -buildmode=c-shared -o litewallet.so  litewallet.go
go build -buildmode=c-shared -o litewallet.dll litewallet.go

mv litewallet.so src/main/resources/liblitewallet.so
mv litewallet.dll src/main/resources/liblitewallet.dll