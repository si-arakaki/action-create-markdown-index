FROM mcr.microsoft.com/devcontainers/base:noble

RUN apt-get update && apt-get upgrade -y

RUN apt install -y curl

ARG GOLANG_VERSION=1.24.0
RUN curl -L https://go.dev/dl/go${GOLANG_VERSION}.linux-arm64.tar.gz | tar -C /usr/local -zx

USER vscode

ENV PATH=${PATH}:/usr/local/go/bin
ENV PATH=${PATH}:/home/vscode/go/bin

RUN go install golang.org/x/tools/gopls@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
