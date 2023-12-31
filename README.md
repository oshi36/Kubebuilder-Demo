# Kubebuilder Demo
This repository contains Kubebuilder demo which I presented at [Kubernetes Community Days Bengaluru 2023](https://community.cncf.io/events/details/cncf-kcd-bengaluru-presents-kubernetes-community-days-bengaluru-2023-in-person/). You can find the hands-on lab for this [here](https://cloudyuga.guru/hands_on_lab/kubebuilder-intro).

[Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) is a framework for building Kubernetes APIs using [Custom Resource Definitions (CRDs)](https://kubernetes.io/docs/tasks/access-kubernetes-api/extend-api-custom-resource-definitions).


## Description
This demo consists of creating a custom resource and custom controller, which creates a pod based on some spec.
![demores_ctr_kubebuilder](https://github.com/oshi36/Kubebuilder-Demo/assets/47573417/f08d5206-6c94-46dd-b605-a9540557f3b1)

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).
Also install [Go](https://go.dev/dl/) and [Kubebuilder](https://book.kubebuilder.io/quick-start.html) CLI.
### Scaffolding of Project and Bootstrapping operator
![kubebuilder_workflow](https://github.com/oshi36/Kubebuilder-Demo/assets/47573417/3d339b64-a502-4fb1-923c-7fb6593df564)
- Create an empty directory and initialize it : `kubebuilder init --domain=demo.kcd.io --repo project`
- Create API for CRDs , custom reosurce and controller : `kubebuilder create api --group demo --version v1 --kind DemoResource`
- Define the custom resources and implement the custom controller logic.


 
### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/project:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/project:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

