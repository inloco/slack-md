# slack-md

[![GoDoc](https://godoc.org/github.com/inloco/slack-md?status.svg)](https://godoc.org/github.com/inloco/slack-md)

Golang Slack Markdown struct serializer based on reflection.

### The struct you give us 

```go
type MyUserAction struct {
	URL string

	User            *MyUser
	NotSoUsefulInfo string `slack_md:"-"`
}

type MyUser struct {
	UserID      string
	SomeUserKey string `slack_md:"obfuscate"`
}
```

### How it will show up on your Slack channel


