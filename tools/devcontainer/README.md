# Set-up Development Environment

**Table of Contents**
- [1. Prerequisite](#1-prerequisite)
- [2. Start podman machine](#2-start-podman-machine)
- [3. Start Dev Containers](#3-start-dev-containers)
- [etc](#etc)


## 1. Prerequisite

### podman
https://podman.io/docs/installation
```bash
# On mac
brew install podman
```

### 'Dev Containers' vscode extension
1. Install 'Dev Containers' in extension.
2. In `'Containers:docker path'` of 'vscode Settings', Edit `'docker'` to `'podman'`


## 2. Start podman machine
Start QEMU based virtual machine because Windows & Mac can't run container directly.
```bash
podman machine init test \
  --cpus 2 \
  --memory 4096 \
  --disk-size 32 \
  --image-path stable

podman machine start test
```


## 3. Start Dev Containers

### Create Dev Container
1. In `"Dev Containers: Open Folder in Containers..."` of Command Palette.
2. Select the `'devcontainer'` directory which includes `.devcontainer.json` & `Dockerfile`.

=> A `'work'` directory under the 'devcontainer' of Host is `mounted` into a container.


### Attach to Running Dev Container

1. In `"Dev Containers: Attach to Running Container..."` of Command Palette.
2. Select `'dongle-dev'`



<br><br>
## Etc

### Verified Environments
- Apple M1 ventura 13.5