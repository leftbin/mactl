# mactl

cli to manage macbooks.

## install

```shell
brew install leftbin/tap/mactl
```

## usage

### bootstrap

a new macbook can be fully bootstrapped with this one command

```shell
mactl bootstrap run
```

if you would like to see what comes included as part of bootstrap, you can look at the check list

```shell
mactl bootstrap checklist
```

### bundles

you can add group of related apps, called bundles

```shell
mactl bundle list
mactl bundle add --name hotkey
```

### apps

you can add a single app

```shell
mactl app list
mactl app add --name docker
```

### zshrc

zshrc is generated as part of bootstrap. it can be generated outside bootstrap as well.

```shell
mactl zshrc generate
```

### git

git can be configured to clone, pull and push to git remote repositories with correct commit author details.

```shell
mactl git init --username <git-username> --email <leftbin-email> --workspace <git-username> --host <github.com>
```

> note: after running the above command, you would have to paste the public key, which is copied to your clipboard by
> mactl in the url shown on the console

## release

release process is currently manual. releasing a new version requires a binary to be uploaded
to [gs://afs-planton-pos-uc1-ext-file-repo](https://console.cloud.google.com/storage/browser/afs-planton-pos-uc1-ext-file-repo/tool/mactl/download)
location on GCS.

this can be accomplished by running the below make command after logging into gcp using gcloud.

```shell
make release v=<version>
```
