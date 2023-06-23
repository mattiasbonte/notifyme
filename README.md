# Notify Me

> Send notification to yourself directly on the command line

### Installation

- `git clone <repository>`
- `go mod tidy`
- `go install`

- Add `.config/notifyme/config.toml` file

```toml config.toml
[telegram]
  auth_token = "your telegram auth token"
  chat_id = "your telegram chat id"
}
```

### Help

> `go run main.go --help`

### Examples

```bash
sleep(10) && notifyme -m "sleeping done" -t "beep"

./my_long_running_import_script.sh && notifyme -m "import complete" -t "alert"
```
