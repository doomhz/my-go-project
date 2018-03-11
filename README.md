My Go Project
=============

For testing purposes.

The code works only when I move the files `main.go` and `db.go` in the same directory `my-lib`
and change the package name to `main` for both `db.go` and `main.go`.

Install
-------

Access docker dev environment with `make docker`.

Install `dep` with `go get -u github.com/golang/dep/cmd/dep`.

Go to `my-lib` and install dependencies with `dep`:

    cd src/my-lib/
    dep ensure --update
    cd /go

Run the app with `go run src/my-lib/main.go`.

The following dependency error should pop-up:

    root@c348fbca321e:/go# go run src/my-lib/main.go
    src/my-lib/mySubLib/db.go:6:2: cannot find package "_/go/src/my-lib/vendor/github.com/go-sql-driver/mysql" in any of:
            /usr/local/go/src/_/go/src/my-lib/vendor/github.com/go-sql-driver/mysql (from $GOROOT)
            /go/src/_/go/src/my-lib/vendor/github.com/go-sql-driver/mysql (from $GOPATH)

Any idea how to fix this? Any help would be much appreciated. Thanks in advance.
