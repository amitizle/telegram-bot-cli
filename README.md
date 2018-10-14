# Telegram bot CLI

A super simple CLI to send on-demand messages on behalf of a Telegram bot - **not as a regular user**.

## Config / flags

All flags can be replaced by environment variables with the prefix `TELEGRAM_`.
So `--token` can be replaced by defining `TELEGRAM_TOKEN`.

### Example

```bash
#!/usr/bin/env bash

export TELEGRAM_TOKEN="SECRET_TOKEN"
export TELEGRAM_CHATID="000001111"

/opt/telegram-bot-cli message text "Server time: $(date)"
```

## Test

TODO
