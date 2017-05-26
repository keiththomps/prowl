# Prowl CLI

A simple CLI for sending iPhone push notifications via [prowlapp.com](prowlapp.com).

## Installation

If you have go installed you can use

```shell
$ go get -u github.com/keiththomps/prowl
```

If not you can grab [the latest release](https://github.com/keiththomps/prowl/releases/latest) and place it in your $PATH.

## Configuration

The CLI is designed to only be used with a single API key currently and you can configure that by using the `prowl config api` command:

```shell
$ prowl config api [YOUR API KEY]
```

The key will be stored in `~/.prowl.json` and utilized in your API requests.

## Usage

After you've added your API key to the configuration then you can send yourself a notification using the `prowl send` command:

```shell
# invoked via `prowl send [EVENT NAME] [DESCRIPTION]`
$ prowl send "Pizza Ready" "You should go eat pizza!"
```

There are more flags for `prowl send` so see `prowl send --help` for more details.
