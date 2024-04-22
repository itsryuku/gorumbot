# gorumbot
[rumble](https://rumble.com) livestream viewbot written in go!

## How does it work?
This viewbot doesn't depend on headless browsers or proxies. Instead, we obtain valid viewer IDs from Rumble and use them to increase livestream views.
It continues to run, sending those viewers every 60 seconds until you stop it.

## How to install golang

Here's are some resources to learn about golang and how to install it on windows
- [GeeksforGeeks - How to Install Go on Windows](https://www.geeksforgeeks.org/how-to-install-go-on-windows/)
- [YouTube - Install GO on Windows 11 in 2 minutes](https://www.youtube.com/watch?v=EPpZbwAr4k8)

## Installation

Once you have installed go, run the following command to get the repo:
```sh
go install -v github.com/itsryuku/gorumbot@latest
```

## Usage

```sh
gorumbot -h # this will display help

gorumbot -u <livestream url> -b <number of bots>
```

## Preview

![Preview](https://raw.githubusercontent.com/itsryuku/gorumbot/main/assets/preview.gif)

If you need anything, feel free to email me at hunter@ryukudz.com or add me on discord: ryuku_
