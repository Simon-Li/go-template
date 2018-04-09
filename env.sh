echo "Exporting GOPATH"
export GOPATH=$(pwd)/go
export GOBIN=$(pwd)/go/bin
export PATH=.:$GOPATH/bin:$PATH

