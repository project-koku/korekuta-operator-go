# Build the manager binary
FROM gcr.io/gcp-runtimes/go1-builder:1.13 as builder

WORKDIR /workspace
COPY . .
COPY .git .git
# Build
RUN GIT_COMMIT=$(git rev-list -1 HEAD) && \
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
/usr/local/go/bin/go build -ldflags "-X controllers.GitCommit=$GIT_COMMIT" -mod vendor -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

# For terminal access, use this image:
# FROM gcr.io/distroless/base:debug-nonroot

WORKDIR /
COPY --from=builder /workspace/manager .
USER nonroot:nonroot

ENTRYPOINT ["/manager"]
