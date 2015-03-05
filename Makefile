TARGETS = auth commands commands\/modify
test: packages
		go test -v ./...

packages:
	! go get code.google.com/p/go.tools/cmd/cover
	go get golang.org/x/tools/cmd/cover
	go get github.com/axw/gocov/gocov
	go get gopkg.in/matm/v1/gocov-html
	go get github.com/modocache/gover
	go get github.com/mattn/goveralls
	
cover: packages
	rm -rf *.out
	rm -rf cover*
	touch cover.json
	@for t in $(TARGETS); \
	do \
		v=`echo $$t | sed 's/\//_/g'`; \
		gocov test github.com/Shaked/getpocket/$$t/ -v >> cover_$$v.json; \
		gocov-html cover_$$v.json >> cover_$$v.html; \
	done;

travis: packages
	rm -rf gover.coverprofile
	rm -rf profile.cov
	@for t in $(TARGETS); \
	do \
		go test -covermode=count -coverprofile=profile.cov github.com/Shaked/getpocket/$$t/; \
	done;
	$(HOME)/gopath/bin/gover 
	$(HOME)/gopath/bin/goveralls -repotoken 4aRDhgifgmEKSiBfuUvQa4whjauFFlkc2 -coverprofile=gover.coverprofile -service travis-ci
