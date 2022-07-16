## Kubernetes Convenience

Kubernetes Convenience (`k8sc`) is an application that attempts to eliminate the verbose usage of some `kubectl` commands 
and introduces convenience commands to ease the process of interacting with Kubernetes clusters. This project is
being developed by being directly inspired by the amazing `kubectx` and `kubens` CLI tools ([GitHub](https://github.com/ahmetb/kubectx))
written by [Ahmet Alp Balkan](https://github.com/ahmetb), and it's really a way for me to practice Go.
For now, `k8sc` can be seen as exactly the same as `kubectx` and `kubens`, but in the future more features will be 
incorporated. To see how `k8sc` differs from the two tools, you can check out the [features](#Features) section.

PS: The package is not released on any package manager at the moment, however, managers like `apt` and `homebrew` will
surely be added. For now, to use the tool, you can install it manually using `go build` and then moving the binary to 
`/usr/local/bin` for UNIX based systems.

### Features

Print current context: `k8sc`

#### Contexts
* Change context: `cs` or `sc`
* Switch to the previous context: `cp` or `pc`

#### Namespaces
* Change namespace: `ns` or `sn`
* Switch to the previous namespace: `np` or `pn`

#### Pods
* Get pods: `gp` or `pg`
* Get pods running on a specific node: `gp -n <node-name>` or `gp --node <node-name>`
* Get wide output: `gp -w`

#### Generate completion scripts
To generate completion scripts for `bash`, `fish`, `powershell` and `zsh`, 
type `k8sc completion <shell-name>`