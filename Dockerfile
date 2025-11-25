FROM golang:latest AS build
WORKDIR /src
COPY src /src
RUN go mod init github.com/wthys/advent-of-code-2025 && go mod tidy && go install
RUN go test ./... && go build -o /out/aoc2025

FROM scratch AS bin
COPY --from=build /out/aoc2025 /
