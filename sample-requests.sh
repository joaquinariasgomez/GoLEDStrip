#!/bin/bash

function sample-get() {
    curl -X GET localhost:8888/action | jq
}

function sample-post() {
    echo $1
    body=`cat sample-post-action.json`
    curl -X POST localhost:8888/action -d "$body"
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