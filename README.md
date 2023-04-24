# import-flyctl-error-2023-04

This codebase is a minimal repro of a bug when importing the flaps client:

```
$ go run .
# github.com/superfly/flyctl/api
/Users/tvd/go/pkg/mod/github.com/superfly/flyctl/api@v0.0.0-20220708073423-b6d7c3cf5161/client.go:38:21: undefined: logger.LoggerInterface
/Users/tvd/go/pkg/mod/github.com/superfly/flyctl/api@v0.0.0-20220708073423-b6d7c3cf5161/client.go:42:65: undefined: logger.LoggerInterface
/Users/tvd/go/pkg/mod/github.com/superfly/flyctl/api@v0.0.0-20220708073423-b6d7c3cf5161/http.go:15:34: undefined: logger.LoggerInterface
/Users/tvd/go/pkg/mod/github.com/superfly/flyctl/api@v0.0.0-20220708073423-b6d7c3cf5161/http.go:42:24: undefined: logger.LoggerInterface

$ echo $?
1
```
