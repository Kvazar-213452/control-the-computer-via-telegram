# Telegram PC Control Bot on Windows

This program allows you to control a computer through a Telegram bot.

## Getting Started

### Prerequisites

- Install the repository.

### Build

Compile the program using the following command:

```cmd
git clone https://github.com/Kvazar-213452/control-the-computer-via-telegram
go build -ldflags="-H windowsgui"
```

And run it on the PC you want to control.
Before that, you need to insert the bot token into the .env file in the TOKEN field.
