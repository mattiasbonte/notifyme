# Notify Me

> Send notification to yourself directly on the command line

### Installation

- Clone the repo
- `go mod tidy`
- `go install`

- Add `.env` file

```bash
TELEGRAM_AUTH_TOKEN=your_telegram_auth_token
TELEGRAM_CHAT_ID=your_telegram_chat_id
```

### Examples

```bash
sleep(10) && notifyme -m "sleeping done" -t "beep"

./my_long_running_import_script.sh && notifyme -m "import complete" -t "alert"
```
