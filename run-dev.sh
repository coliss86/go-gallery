#!/bin/sh

if [ -z $GOPATH ];
then
  echo "Variable \$GOPATH is not set."
  DIR="$(cd "$(dirname "$0")" && pwd -P)"
  GOPATH="$(cd "$DIR/../../../../" && pwd -P)"
  export GOPATH
  echo "Setting it to '$GOPATH'"
  if [ ! -e "$GOPATH/src" ];
  then
    echo "Can't determine automatically the \$GOPATH value" >&2
    exit 1
  fi
fi

godep go build && ./go-gallery conf/gallery.conf
