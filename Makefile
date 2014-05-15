TARGETS = auth
test: packages
		bash -c 'go test -v ./... | tee go-test.out; RET_CODE=$${PIPESTATUS[0]}; (go2xunit -input go-test.out -output xunit.xml) 2> /dev/null || echo -n; exit $${RET_CODE}'

packages:
	go get code.google.com/p/go.tools/cmd/cover
	go get github.com/axw/gocov/gocov
	go get gopkg.in/matm/v1/gocov-html

cover: packages
	rm -rf *.out
	rm -rf cover.json
	touch cover.json
	@for t in $(TARGETS); \
	do \
		gocov test $$t/ -v >> cover.json; \
	done;

	gocov-html cover.json > cover.html
