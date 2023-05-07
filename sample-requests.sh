#!/bin/bash

function sample-get() {
    curl -X GET localhost:8888/action | jq
}

function sample-post() {
    echo $1
    if [[ "$1" == 1 ]]; then
      body=`cat post-office-lights.json`
      curl -X POST localhost:8888/action -d "$body"
    else
      body=`cat post-decrease-brightness.json`
      curl -X POST localhost:8888/action -d "$body"
    fi
}

while getopts 'p:gh' opt; do
  case "$opt" in
    p)
      arg="$OPTARG"
      echo "Processing option 'POST' with '${OPTARG}' argument"
      sample-post ${OPTARG}
      ;;
    g)
      echo "Processing option 'GET'"
      sample-get
      ;;
    ?|h)
      echo "Usage: $(basename $0) [-p arg](POST) [-g](GET)"
      exit 1
      ;;
  esac
done
shift "$(($OPTIND -1))"