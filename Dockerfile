# builder image
FROM golang:1.17-alpine3.14 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go mod tidy && go build -o kubescape .

RUN /build/kubescape download framework nsa --output nsa.json


FROM golang:1.17-alpine3.14
COPY --from=builder /build/kubescape .
COPY --from=builder /build/nsa.json .
ENTRYPOINT [ "./kubescape" ]
CMD [ "scan", "framework", "nsa", "--use-from", "nsa.json", "--exclude-namespaces", "kube-system,kube-public" ]