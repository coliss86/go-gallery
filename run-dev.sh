#!/bin/sh

if [ -z $GOPATH ];
then
  echo "Variable \$GOPATH is not set."
  DIR=$(dirname \"$(readlink -f \"$0\")\")
  GOPATH=$(readlink -f $DIR/../../../../)
  export GOPATH
  echo "Setting it to '$GOPATH'"
  if [ ! -e "$GOPATH/src" ];
  then
    echo "Can't determine automatically the \$GOPATH value" >&2
    exit 1
  fi
fi

godep go build && ./go-gallery conf/gallery.conf
