# Crypto Bot
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/mbtamuli/crypto?label=go%20module)](https://github.com/mbtamuli/crypto/tags)

## Binance Bot
The [binance-bot](./bots/binance-bot) that fetches data from the Binance API and sends updates via Telegram to users registered in the database.

### Setup
1. Create a bot using Telegram's bot for bot creation :robot: - [BotFather](https://core.telegram.org/bots#6-botfather)
2. A GitHub release is created using [GoReleaser](https://github.com/mbtamuli/crypto/blob/main/.github/workflows/release.yml) to publish the latest version of the Go binary.
3. The Go binary is run using [GitHub Actions](https://github.com/mbtamuli/crypto/actions/workflows/run.yml).

## Telegram CLI
The [telegram-users-cli](./bots/telegram-users-cli) is used to register users in a database with their Telegram Chat ID.

## Roadmap

- [x] Store Telegram user's Chat ID in a database
- [ ] Fetch data from Binance
- [ ] Send updates via Telegram