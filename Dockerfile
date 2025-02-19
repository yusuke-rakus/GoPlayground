FROM golang:1.22rc1-bullseye

ENV PATH="$PATH:$GOROOT/bin:$GOPATH/bin"
ENV PATH="$PATH:$(go env GOPATH)/bin"

COPY . .

RUN apt-get update \
	&& apt-get install -y \
	&& go install github.com/air-verse/air@latest \
	&& go install -v golang.org/x/tools/gopls@latest \
	&& go install -v github.com/cweill/gotests/gotests@v1.6.0 \
	&& go install -v github.com/fatih/gomodifytags@v1.16.0 \
	&& go install -v github.com/josharian/impl@v1.1.0 \
	&& go install -v github.com/haya14busa/goplay/cmd/goplay@v1.0.0 \
	&& go install -v github.com/go-delve/delve/cmd/dlv@latest \
	&& go install -v honnef.co/go/tools/cmd/staticcheck@latest \
	&& go install -v golang.org/x/tools/cmd/goimports@latest

ENV CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64
