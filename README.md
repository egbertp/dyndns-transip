# dyndns-transip

[![](https://travis-ci.org/egbertp/dyndns-transip.svg?branch=master)](https://travis-ci.org/github/egbertp/dyndns-transip)


Is a small little executable that will update a domain record of your choice that is hosted by [TransIP](https://www.transip.nl/). Run it via cron to keep your DNS record up to date.

This project is heavily inspired by [`go-transip-dyndns`](https://github.com/jlentink/go-transip-dyndns) made by Jason Lentink. The inner working is similar. I used GoMods, updated to GoLang 1.15 and created binaries for more platforms and architectures (amongst others OpenBSD and FreeBSD).

# Getting started 

## Configure TransIP

1. Create a access key for the API. `(Control panel > My Account > API)`
2. Add a label and press create.
3. Save Key to a file. `e.g. transip.key`

## Download & install

Download the latest version compatible with your OS and architecture

One-liner for `Linux` on `amd64`
```sh
$ curl -s https://api.github.com/repos/egbertp/dyndns-transip/releases/latest |  jq -r '.assets[] | select(.browser_download_url | contains("linux_amd64")) | .browser_download_url' | xargs -n 1 curl -O -sSL
```

Move the binary to `/usr/local/bin`
```sh
$ chmod +x dyndns-transip_linux_amd64
$ sudo mv dyndns-transip_linux_amd64 /usr/local/bin/.
$ sudo ln -s /usr/local/bin/dyndns-transip_linux_amd64 /usr/local/bin/dyndns-transip
```

## Example config file

name: `dyndns-transip.yaml`

place the config file at `/etc/dyndns-transip.yaml` or in the directory where you execute the command.

```yaml
username: "transip-username"
private-key: "/path-to/private.key"

verbose: false

domain: "yourdomain.nl"
domain-entry: "subdomain"
domain-ttl: 60
```

## Create DNS record

```sh
$ dyndns-transip create
```

For more info: `dyndns-transip help`

## Configure cron

```sh
$ crontab -e
```

Add to crontab
```
0 * * * * /usr/local/bin/go-transip-dyndns
```

## No association with Transip

This tool has been created to compensate for the fact that my ISP does not provide me with a static IP address. There is no association with Transip.

## Built with

* [Cobra](https://github.com/spf13/cobra) - Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.
* [Viper](https://github.com/spf13/viper) - Viper is a complete configuration solution for Go applications including 12-Factor apps.
* [mitchellh/gox](github.com/mitchellh/gox) - Gox - Simple Go Cross Compilation

A big thanks to all authors and contributors of these libraries and tools!

## References

* [One Liner to Download the Latest Release from Github Repo](https://gist.github.com/steinwaywhw/a4cd19cda655b8249d908261a62687f8)
* [ipify API](https://www.ipify.org/)