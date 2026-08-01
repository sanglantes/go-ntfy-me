# go-ntfy-me
An extendable ntfy.sh client in Go.

## Building

```powershell
git clone https://github.com/sanglantes/go-ntfy-me
cd go-ntfy-me
go build -ldflags="-s -w -H windowsgui" cmd\main.go -o gnm.exe
```

This will build go-ntfy-me as a console-less application.

## Configuration

go-ntfy-me will look for configuration options specified in the `.env` file. The `.env` file must be located in the same directory as the executable.

### Options
- `NTFY_ENDPOINT`: The ntfy endpoint URI together with GET parameters. The endpoint must emit JSON. Example: `https://ntfy.sh/my_topic/json`
- `NTFY_AUTH_TOKEN`: A bearer token to access private topics. Optional.
- `ADD_TO_START_UP`: Add go-ntfy-me to the Windows start up folder. Accepted values are `true`/`false`.
- `CONNECTION_RETRY_TIME`: If a connection is lost, retry in specified amount of seconds. If polling is selected, this will decide the interval for which the ntfy topic is polled.

## Writing your own actions and plugins

