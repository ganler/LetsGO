This guy finally starts to learn GO.

## Resources

### Setup

* [Installation](https://golang.org/doc/install): For Windows users, it is recommeded to install via a `.msi` file.
    * for mac: `brew install go`
* Add `$GOROOT/bin` into your `$PATH` if you want to make go compiler directly executable. (ignore this step if you are using a package manager / `.msi` on Windows)
* Set `$GOPATH` to somewhere you want to put all your go files.
* `go get github.com/nsf/gocode --verbose` to set up your `$GOPATH` (by installing a package).
* You can also manually set up `$GOPATH`:
    1. `mkdir src && mkdir bin && mkdir pkg`
    2. Put your projects under the `src` folder.

> You may not want to touch `GOROOT` where the go compiler packages locate. Change it if you need different versions of GOs.

> More about `GOPATH`: **It is OK to have more than one `$GOPATH`.**
> If you have multiple `GOPATH` entries, only the first one will be used for downloading source codes(i.e., `go get`), but all you entries can be searched for source codes. 

### Install this repo

```
### For Windows PowerShell users:
cd $Env:GOPATH\src
git clone https://github.com/ganler/LetsGO.git

### For non-Windows users:
cd $GOPATH/src
git clone https://github.com/ganler/LetsGO.git
```

### Basic CMDs

```shell
go version
go env  # see the env vars.
go get XXX --verbose # Tha `verbose` flag is for looking at details of downloading.
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
