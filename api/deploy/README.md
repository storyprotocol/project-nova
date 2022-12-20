# Overview

This contains k8s deployment files to the EKS cluster. For more details of kustomize concepts:
https://kubernetes-sigs.github.io/kustomize/api-reference/glossary/

# `base/`

This includes the [base](https://kubernetes-sigs.github.io/kustomize/api-reference/glossary/#base) deployment
configurations. This is shared acriss all environments.

# `overlays/dev`

This includes custom patches for `dev` environment.

# `overlays/staging`

This includes custom patches for `staging` environment.

# `overlays/prod`

This includes custom patches for `prod` environment.
