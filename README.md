### Kubernetes Convenience

Kubernetes Convenience (`k8sc`) is an application that attempts to eliminate the verbose usage of some `kubectl` commands 
and introduces convenience commands to ease the process of interacting with Kubernetes clusters. This project is
being developed by being directly inspired by the amazing `kubectx` and `kubens` CLI tools ([GitHub](https://github.com/ahmetb/kubectx))
written by [Ahmet Alp Balkan](https://github.com/ahmetb), and it's really a way for me to practice Go.
For now, `k8sc` can be seen as exactly the same as `kubectx`, but in the future `kubens` and more features will be 
incorporated.

PS: The package is not released on any package manager at the moment, however, managers like `apt` and `homebrew` will
surely be added.