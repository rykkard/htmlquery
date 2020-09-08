# HTMLquery
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/rykkard/htmlquery)
[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/rykkard/htmlquery/issues)

## Install

- Installation
```bash
$ go get -v 'github.com/rykkard/htmlquery'
```

- Update
```bash
$ go get -v -u 'github.com/rykkard/htmlquery'
```

## Usage

```
$ hq -h
HTMLquery
A fairly simple html selector built around the goquery golang package.

Usage:
   hq [OPTIONS] <query> [<resources>]

Options:
   -u, --url                   enable URL mode
   -h, --help                  show help
```

## Examples

- Querying file resource with stdin
```
$ curl -s http://example.com |hq h1
<h1>Example Domain</h1>
```

- Querying url resources
```
$ hq -u h1 http://example.com
<h1>Example Domain</h1>
```

- Querying url resources with stdin
```
$ echo http://example.com |hq -u h1
<h1>Example Domain</h1>
```
