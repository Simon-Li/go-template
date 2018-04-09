# Go project template

A template contains various utilities for Golang projet developments, it gives unit test, environment setting, logging utility frameworks out of the box by default so that Go developers should be able to jump into development without hassle of these utilities setup.

## Build

1. If using Mac OS, install [`brew`](https://brew.sh/)
2. Install `Golang` via `brew install go`
3. Install `direnv` via `brew install direnv`
4. Add the following to your shell profile config file (i.e. `.zshrc`)

   ```
   eval "$(direnv hook zsh)"
   ```
5. source your shell profile config file
   ```
   source ~/.zsh
   ```

## Unit test packages

- github.com/onsi/ginkgo
- github.com/onsi/gomega

## Environment setting

- Install direnv (more info: [Homebrew for OSX](http://brewformulas.org/Direnv))
- Create ".envrc" file and add the following scripts
  ```
  export GOPATH=$PWD/go
  export GOBIN=$PWD/go/bin
  export PATH=.:$GOPATH/bin:$PATH
  ```
