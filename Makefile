PREFIX=$(DESTDIR)/usr/local
RUNTIME_GOPATH=$(GOPATH):`pwd`
VERSION=`git tag | tail -n 1`
GOOS=`go env GOOS`
GOARCH=`go env GOARCH`

kinesis-cat:	main.go src/kinesis_cat/kinesis-cat.go dependencies
	GOPATH=$(RUNTIME_GOPATH) go build -o kinesis-cat main.go

dependencies:
	go get github.com/crowdmob/goamz/kinesis
	go get code.google.com/p/go-uuid/uuid
	go get github.com/vaughan0/go-ini

install: kinesis-cat
	install -m 755 kinesis-cat $(PREFIX)/bin/

clean:
	rm -f kinesis-cat *.tar.gz

test:
	GOPATH=$(RUNTIME_GOPATH) go test src/**/*_test.go

package: clean kinesis-cat
	tar zcf kinesis-cat-$(VERSION)-${GOOS}-$(GOARCH).tar.gz ./kinesis-cat
