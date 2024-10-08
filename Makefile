.PHONY: build run install-deps test-pdfgen release

install-deps:
	which modd || go install github.com/cortesi/modd/cmd/modd@latest
	go mod tidy
	go mod download
	go mod verify

build:
	rm -fr dist
	cp public/generate-report.html templates/layout.html.templ
	go build -ldflags "-s -w" -o dist/gooferr cmd/*.go 

run: build
	dist/gooferr

test-pdfgen:
	rm -f output.pdf
	/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --headless --disable-gpu --print-to-pdf http://localhost:3000/generate-report

release:
	rm -fr dist
	goreleaser --parallelism 1 --clean --rm-dist --skip-validate
	echo "Release process not implemented yet"