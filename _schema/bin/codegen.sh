#!/usr/bin/env bash

_clean() {
  find ../api/go-sdk -name "*.go" -exec rm -f {} \;
  find ../api/web-sdk -name "*.ts" -exec rm -f {} \;
  mkdir -p ../api/web-sdk/src
}


_protobuf_gen() {
  docker run --rm \
    --volume "${PWD}":/workspace/schema \
    --volume "${PWD}/../api":/workspace/api \
    --workdir /workspace/schema \
    bufbuild/buf generate --verbose
}

_transform_string() {
  echo "${1}" \
  |  sed -r 's|/|_|g' \
  | awk 'BEGIN{FS=OFS="_"} {for(i=1;i<=NF;i++){$i=toupper(substr($i,1,1)) substr($i,2)}} 1' \
  | sed 's/_Pb//g' \
  | sed 's/_//g'
}

_export_module() {
  while read -r module; do
   if [[ "${module}" =~ ^(.*/src/)(.*)(\.ts|_pb\.ts)$ ]]; then
     target="${BASH_REMATCH[1]}"
     module_name=${BASH_REMATCH[2]}
     export_as=$( _transform_string "${module_name}" )
     echo "export * as ${export_as} from \"./${module_name}\";" >> "${target}"index.ts
   fi
  done
}


_export_ts_index() {
  find ../api/web-sdk -name "*.ts" | _export_module
}

_clean
_protobuf_gen
_export_ts_index