# logit

**logit** is a CLI tool that consumes logs in JSON and feeds them into a logs storage like file, Amazon CloudWatch, or Slack.

Why it is important:

* Your application will contain only bussiness logic. Don't write your own logging library.
* Your application will not send secondary network requests.
* Your application will be faster. It's easy to make a JSON entry and feed it into stdout, and logit will take care of the rest.
* Your application will be more environment-agnostic. Here you run it on your machine, production runs on AWS, tomorrow you migrate on GCP. And in all cases, only the logit config is different, application is the same.

Supported handlers:

* console and files:
  * `json`
  * `logfmt`
  * `rolling`
  * `text`
  * `zalgo`
* databases:
  * `elastic`
  * `influxdb`
  * `mongodb`
  * `redis`
* messengers:
  * `discord`
  * `slack`
* log collectors:
  * `aws`
  * `fluentd`
  * `gcloud`
  * `graylog`
  * `sentry`
  * `syslog`
  * `loggly`
  * `logstash`

## Installation

The project is "go gettable":

```bash
go get github.com/life4/logit
```

The best practice for your production is to build a minimal binary with only handlers you need:

```bash
git clone https://github.com/life4/logit.git
cd logit
go build -tags 'h_clean,h_aws,h_gcloud' -o logit.bin .
```

Handlers `json`, `logfmt`, and `text` are always included. All other can be disabled by `h_clean` build tag and then individually enabled with `h_{HANDLER}` build tags.

## Usage

First of all, you need a config. Let's make the following `logit.toml`:

```toml
[[handler]]
format = "text"
```

And now, pipe some JSON into stdin of logit:

```bash
$ export text='{"animal":"walrus","level":"fatal","msg":"A huge walrus appears","time":"2020-12-29 15:09:21"}'
$ echo $text | logit
FATAL  [2020-12-29 15:09:21] A huge walrus appears    animal=walrus
```

Or just a plain text, why not:

```bash
$ echo 'hi' | logit
INFO   [2021-01-04 14:48:50] hi
```

Easy!

See [documentation](./docs/config.md) for more details.

## Producing JSON

There are some good libraries that you can use in your application to produce JSON logs:

* Go:
  * [onelog](https://github.com/francoispqt/onelog): fastest.
  * [logrus](github.com/sirupsen/logrus): most popular.
* Node.js:
  * [bunyan](https://github.com/trentm/node-bunyan)
* Python:
  * [structlog](https://github.com/hynek/structlog): powerful.
  * [python-json-logger](https://github.com/madzak/python-json-logger): formatter for the default [logging](https://docs.python.org/3/library/logging.html).
Rust:
  * [slog-json](https://github.com/slog-rs/json) for [slog-rs](https://github.com/slog-rs/slog).
