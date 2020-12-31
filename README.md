# Source to image s2i - extended build

This repo is an example on HOW-TO use `s2i` to:

* build a `GOLANG` application using a builder image e.g. `golang:1.15`
* running the built `GOLANG` application using a runtime image e.g. `registry.access.redhat.com/ubi8/ubi-minimal`

First pull the builder and the runtime images:

```
docker pull golang:1.15
docker pull registry.access.redhat.com/ubi8/ubi-minimal
```

Then run `s2i`:

```
s2i build . golang:1.15 goserver --runtime-image registry.access.redhat.com/ubi8/ubi-minimal --runtime-artifact /go/bin/goserver
```

Finally start the application and test it:

```
docker run --rm -d -p 8080:8080 goserver
curl http://localhost:8080/
```

> NOTE: we are using a builder image which is not an `s2i` ready image; it doesn't contain any `s2i` script nor has the `io.openshift.s2i.scripts-url` label set

## References:

* https://github.com/openshift/source-to-image/blob/master/docs/runtime_image.md
