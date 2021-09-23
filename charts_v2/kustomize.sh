#!/bin/sh
cat <&0 > all.yaml
# you can also use "kustomize build ." if you have it installed.
exec kubectl kustomize && rm all.yaml