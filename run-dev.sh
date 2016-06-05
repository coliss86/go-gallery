#!/bin/sh

if [ -z $GOPATH ];
then
  echo "Variable \$GOPATH is not set." >&2
  exit 1
fi

godep go build && ./go-gallery conf/gallery.conf
