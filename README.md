# go-ntfy-me
An extendable ntfy.sh client in Go.

## Building

```powershell
git clone https://github.com/sanglantes/go-ntfy-me
cd go-ntfy-me
go build -ldflags="-s -w -H windowsgui" -o gnm.exe .\cmd```

This will build go-ntfy-me as a console-less application.

## Configuration

go-ntfy-me will look for configuration options specified in the `.env` file. The `.env` file must be located in the same directory as the executable.

### Options
- `NTFY_ENDPOINT`: The ntfy endpoint URI together with GET parameters. The endpoint must emit JSON. Example: `https://ntfy.sh/my_topic/json`
- `NTFY_AUTH_TOKEN`: A bearer token to access private topics. Optional.
- `ADD_TO_START_UP`: Add go-ntfy-me to the Windows start up folder. Accepted values are `true`/`false`.
- `CONNECTION_RETRY_TIME`: If a connection is lost, retry in specified amount of seconds. If polling is selected, this will decide the interval for which the ntfy topic is polled.

## Extending the client

go-ntfy-me can be extended with compile-time plugins. Custom "actions" are added to the `actions/custom` package and then registered in `actions/init.go`.

Begin by constructing a `Registration` struct for your plugin.

```go
var ExportDefaultAction = action.Registration{
	Name:       "default",
	Action:     defaultAct,
	Priority:   1,
	IsBlocking: false,
}
```

- `Name`: Unique action name. By default, all registered actions are run on the event `ntfy.msg`, but single events can be called with `r.Run(name string, e event.Event)`.
- `Action`: The function that will be run when `ntfy.msg` is published.
- `Priority`: The priority of the action decides when it fires. A higher priority means it fires before other actions with lower priorities.
- `IsBlocking`: If set to `true`, make `Action` finish before proceeding with other actions.

An example of an `Action`:
```go
func defaultAct(e event.Event, eb event.EventBus) {
	switch msgx := e.Data.(type) {
	case ntfy.NtfyMessage:
		printMsg(msgx)

	case []ntfy.NtfyMessage:
		for _, msg := range msgx {
			printMsg(msg)
		}
	}
}

func printMsg(msg ntfy.NtfyMessage) {
	if msg.Event == "message" {
		fmt.Println(msg.Message)
	}
}
```

At last, register the `action.Registration` struct to `actions/init.go` with `r.Register(custom.ExportDefaultAction)` in the `Install` function.

Note that `ntfy.msg` may send data that is either `[]ntfy.NtfyMessage` or `ntfy.NtfyMessage`.

The `ntfy.NtfyMessage` struct looks like this:
```go
type NtfyMessage struct {
	ID         string   `json:"id"`
	Time       int      `json:"time"`
	Expires    int      `json:"expires"`
	Event      string   `json:"event"`
	Topic      string   `json:"topic"`
	Priority   int      `json:"priority"`
	Tags       []string `json:"tags"`
	Click      string   `json:"click"`
	Attachment struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Size    int    `json:"size"`
		Expires int    `json:"expires"`
		URL     string `json:"url"`
	} `json:"attachment"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
```