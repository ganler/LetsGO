This guy finally starts to learn GO.

## Why GO

This is a response to issue#1 [link](https://github.com/ganler/LetsGO/issues/1). 

Refer to this [article](https://productcoalition.com/reasons-why-golang-is-better-than-other-programming-languages-4714082bb1b1).

> * Fast-paving compilation and execution
> * Do away with the need of working with different subsets of languages for one project.
> * A boost to code readability and documentation
> * Offering a thoroughly consistent language
> * Facilitating easy versioning of the program
> * Allowing developing with multiple languages
> * Allowing easier maintenance of dependencies

Any personal reasons? Yes, of course. I am just too bored that I want to learn something different and all. Thanks, @PragmaTwice.

## Resources

### Setup

* [Installation](https://golang.org/doc/install): For Windows users, it is recommeded to install via a `.msi` file.
    * for mac: `brew install go`
* Add `$GOROOT/bin` into your `$PATH` if you want to make go compiler directly executable. (ignore this step if you are using a package manager / `.msi` on Windows)
* Set `$GOPATH` to somewhere you want to put all your go files.
* `go get -v github.com/nsf/gocode` to set up your `$GOPATH` (by installing a package).
* You can also manually set up `$GOPATH`:
    1. `mkdir src && mkdir bin && mkdir pkg`
    2. Put your projects under the `src` folder.

> You may not want to touch `GOROOT` where the go compiler packages locate. Change it if you need different versions of GOs.

> More about `GOPATH`: **It is OK to have more than one `$GOPATH`.**
> If you have multiple `GOPATH` entries, only the first one will be used for downloading source codes(i.e., `go get`), but all you entries can be searched for source codes. 

### Proxy

```shell
go env -w GO111MODULE=on
go env -w GOPROXY="https://goproxy.io,direct"

# Set environment variable allow bypassing the proxy for selected modules (optional)
go env -w GOPRIVATE="*.corp.example.com"

# Set environment variable allow bypassing the proxy for specified organizations (optional)
go env -w GOPRIVATE="example.com/org_name"
```

See [https://goproxy.io/](https://goproxy.io/) for details.

### Install this repo

```
### Uniform approach
go get -v github.com/ganler/LetsGO

### For Windows PowerShell users:
cd $Env:GOPATH\src
mkdir -p github.com && cd github.com
mkdir -p ganler && cd ganler
git clone https://github.com/ganler/LetsGO.git

### For non-Windows users:
cd $GOPATH/src
mkdir -p github.com && cd github.com
mkdir -p ganler && cd ganler
git clone https://github.com/ganler/LetsGO.git
```

### Basic CMDs

```shell
go version
go env  # see the env vars.
go get -v XXX # The `v` flag is for looking at details of downloading.
```

### Tutorial

* [YOUTUBE](https://www.youtube.com/watch?v=YS4e4q9oBaU)
* [Go by Example](https://gobyexample.com/)
* [PKGs](https://golang.org/pkg/)

### Docs


### Standards

[Conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```
* `$description`: should be imperitive. (e.g., `add` instead of `added`)
* `$type`: fix | feat | build | refactor | test | BREAKINGCHANGE
* `$type!`: the `!` means something breaking occured.
* Examples:
```
fix: correct dead lock bug by variable lk
-------------------------------------------------
feat(basic): add code example about printing`
-------------------------------------------------
fix: correct minor typos in codes

see the issue for details

Reviewed-by: XZT
Refs #233
```
