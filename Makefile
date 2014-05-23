TARGETS = auth commands
test: packages
		@for t in $(TARGETS); \
		do \
			cd $$t && go test -v && cd ..; \
		done;

packages:
	go get code.google.com/p/go.tools/cmd/cover
	go get github.com/axw/gocov/gocov
	go get gopkg.in/matm/v1/gocov-html
	go get github.com/mattn/goveralls

cover: packages
	rm -rf *.out
	rm -rf cover*
	touch cover.json
	@for t in $(TARGETS); \
	do \
		gocov test github.com/Shaked/getpocket/$$t/ -v >> cover_$$t.json; \
		gocov-html cover_$$t.json >> cover_$$t.html; \
		goveralls -gocovdata=cover_$$t.json -service=travis-ci; \
	done;



