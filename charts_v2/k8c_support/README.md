# Approach using upstream charts

## Basics
1. We use a single parent chart called k8c_support.
1. We add any chart that we want to migrate in the `dependencies` section in the `chart.yaml` file under `k8c_support` directory.
1. We add overrider values of parent chart in `values.yaml` in `k8c_support` directory by adding a top level key matching the chart name in dependencies section. e.g. we have `minio` as name dependecies section as well as in values.yaml

## Advanced
There are things which are kkp specific in existing charts which cannot be simply added by overriding values.yaml file. For such things, we have two approaches.

### 1. using kustomize as post-renderer to helm.

> post-renderer can only be used with helm v3.1+

To use post-renderer, 
1. Create a directory for each chart that needs customization.
1. place a `kustomization.yaml` file in that directory and then any other files used by kustomize to achieve desired effects.
1. run command like this in each chart directory to see the updated output.
helm template minio ../ --values user-supplied-values.yaml --post-renderer ../kustomize.sh
1. The `kustomize.sh` is a simple wrapper to pass helm output stream to kustomize

e.g. In case of minio, you can see that kustomization

1. adds a backup container to existing minio deployment
2. add a volume to existin deployment spec
3. can also create a new secret based on secret created as part of Helm output. (Commented out for now)


### 2. Using extra yaml files (when you want to refer things from values.yaml)
* Sometimes, you don't want to add a static patch to helm generated content. You want to add dynamic content based on values provided in values.yaml.
* Unfortunately, kustomize does not support this directly (supported partially via `replacements` but it feels too verbose). 
* So other option could be add our own custom yaml files which add the required resources. 
* This option can work only if we are creating completely new resource.
* But advantage of this option is, since we are writing helm template, we can take advantage of templating constructs like if/else and other go templating functions.

e.g. In case of minio, you can see that a file `secret-external.yaml` was added in templates folder of parent chart. It refers to `minio.accessKey` and `minio.secretKey` variables to generate dynamic content as per values provided by end user.