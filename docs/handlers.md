# Handlers

## Text

* desc: colorful plain text, your best friend for local development.
* slug: `text`
* core: [logrus](https://github.com/sirupsen/logrus#formatters)
* fields:
  * `timestamp` (string, `"YYYY-MM-dd HH:mm:ss"`)
  * `sort` (bool, `true`)
  * `truncate_level` (bool, `false`)

## LogFmt

* desc: outputs logs in [logfmt](https://rollout.io/blog/logfmt-a-log-format-thats-easy-to-read-and-write/) format which is readable and at the same time [can be parsed](https://github.com/kr/logfmt). The best choice for CI or an environment without colors support.
* slug: `logfmt`
* core: [logrus](https://github.com/sirupsen/logrus#formatters)
* fields:
  * `timestamp` (string, `"YYYY-MM-dd HH:mm:ss"`)
  * `sort` (bool, `true`)

## JSON

* desc: outputs logs in JSON. It is helpful to transform plain text input into proper JSON logs, fill missed fields, or do other transformations on logs supported by logit.
* slug: `json`
* core: [logrus](https://github.com/sirupsen/logrus#formatters)
* fields:
  * `data_key` (string)
  * `timestamp` (string, `"YYYY-MM-dd HH:mm:ss"`)

## Rolling

* desc: rolling logging into a file (with file rotation).
* slug: `rolling`
* build tag: `h_clean,h_rolling`
* core: [lumberjack](github.com/natefinch/lumberjack)
* fields:
  * `file` (string)
  * `max_size` (int, `500`)
  * `max_age` (int, `28`)
  * `max_backups` (int, `3`)
  * `local_time` (bool, `false`)
  * `compress` (bool, `false`)

## AWS

* desc: save logs into [Amazon CloudWatch](https://aws.amazon.com/cloudwatch/).
* slug: `aws`
* build tag: `h_clean,h_aws`
* core: [logrus-cloudwatchlogs](https://github.com/kdar/logrus-cloudwatchlogs)
* fields:
  * `region` (string, `"us-east-1"`)
  * `group` (string)
  * `stream` (string)
  * `max_retries` (int, `-1`)
  * `ssl` (bool, `true`)

## Discord

* desc: send logs as messages into [Discord](https://discord.com/).
* slug: `discord`
* build tag: `h_clean,h_discord`
* core: [discordrus](github.com/kz/discordrus)
* fields:
  * `url` (string)
  * `username` (string, `"logit"`)
  * `author` (string)
  * `inline` (bool, `true`)
  * `timestamp` (string, `"YYYY-MM-dd HH:mm:ss"`)

## Elastic

* desc: save logs into an [ElasticSearch](https://www.elastic.co/elasticsearch/) index.
* slug: `elastic`
* build tag: `h_clean,h_elastic`
* core: [elogrus](https://github.com/sohlich/elogrus)
* fields:
  * `urls` (array of strings, `["http://localhost:9200"]`)
  * `host` (string, `"localhost"`)
  * `index` (string, `"logit"`)

## Fluentd

* desc: save logs into [Fluentd](https://www.fluentd.org/).
* slug: `fluentd`
* build tag: `h_clean,h_fluentd`
* core: [logrus_fluent](https://github.com/evalphobia/logrus_fluent)
* fields:
  * `host` (string, `"localhost"`)
  * `port` (int, `24224`)
  * `max_retry` (int, `0`)

## GCloud

* desc: send logs into [Google Cloud Logging](https://cloud.google.com/logging/) (ex StackDriver).
* slug: `gcloud`
* build tag: `h_clean,h_gcloud`
* core: [google-cloud-go](https://github.com/googleapis/google-cloud-go)
* fields:
  * `credentials` (string): path to credentials file.
  * `endpoint` (string)
  * `labels` (table string to string)
  * `log_name` (string)
  * `project_id` (string)

## Graylog

* desc: send logs into [Graylog](https://www.graylog.org/) over UDP.
* slug: `graylog`
* build tag: `h_clean,h_graylog`
* core: [logrus-graylog-hook](https://github.com/gemnasium/logrus-graylog-hook)
* fields:
  * `address` (string)

## InfluxDB

* desc: send logs into [InfluxDB](https://www.influxdata.com/).
* slug: `influxdb`
* build tag: `h_clean,h_influxdb`
* core: [logrus_influxdb](https://github.com/Abramovic/logrus_influxdb)
* fields:
  * `host` (string, `"localhost"`)
  * `port` (string, `6379`)
  * `username` (string)
  * `password` (string)
  * `database` (string, `"logit"`)
  * `precision` (string, `"ns"`)
  * `use_https` (bool, `false`)
  * `batch_count` (int, `0`)

## Loggly

* desc: send logs into [Loggly](https://www.loggly.com/).
* slug: `loggly`
* build tag: `h_clean,h_loggly`
* core: [logrusly](https://github.com/sebest/logrusly)
* fields:
  * `token` (string)
  * `host` (string)
  * `tags` (list of strings)

## Logstash

* desc: send logs into [Logstash](https://www.elastic.co/logstash).
* slug: `logstash`
* build tag: `h_clean,h_logstash`
* core: [logrus-logstash-hook](https://github.com/bshuster-repo/logrus-logstash-hook)
* fields:
  * `network` (string)
  * `address` (string)
  * `version` (string, `"1"`)
  * `type` (string, `"logit"`)

## MongoDB

* desc: save logs into a [MongoDB](https://www.mongodb.com/) collection.
* slug: `mongodb`
* build tag: `h_clean,h_mongodb`
* core: [mgorus](https://github.com/weekface/mgorus)
* fields:
  * `url` (string, `"localhost"`)
  * `db` (string)
  * `collection` (string)

## Redis

* desc: [RPUSH](https://redis.io/commands/rpush) logs into a [Redis](https://redis.io/) list.
* slug: `redis`
* build tag: `h_clean,h_redis`
* core: [logrus-redis-hook](https://github.com/rogierlommers/logrus-redis-hook)
* fields:
  * `host` (string, `"localhost"`)
  * `port` (int, `6379`)
  * `password` (string)
  * `key` (string, `"logit"`)
  * `format` (string)
  * `app` (string)
  * `source_host` (string)
  * `database` (int)
  * `ttl` (int, `"1h"`)

## Sentry

* desc: send logs into [Sentry](https://sentry.io/welcome/).
* slug: `sentry`
* build tag: `h_clean,h_sentry`
* core: [logrus_sentry](https://github.com/evalphobia/logrus_sentry)
* fields:
  * `dsn` (string)
  * `timeout` (string, `"20s"`)

## Slack

* desc: send logs as messages into a [Slack](https://slack.com/intl/en-nl/) channel.
* slug: `slack`
* build tag: `h_clean,h_slack`
* core: [slackrus](https://github.com/johntdyer/slackrus)
* fields:
  * `hook_url` (string)
  * `icon_url` (string)
  * `channel` (string)
  * `icon_emoji` (string)
  * `username` (string, `"logit"`)

## Syslog

* desc: send logs using [syslog](https://en.wikipedia.org/wiki/Syslog) protocol (for example, into [rsyslog](https://www.rsyslog.com/)).
* slug: `syslog`
* build tag: `h_clean,h_syslog`
* core: [logrus/syslog](https://github.com/sirupsen/logrus/tree/master/hooks/syslog)
* fields:
  * `network` (string, `"udp"`)
  * `address` (string, `"localhost:514"`)
  * `tag` (string, `"logit"`)
  * `priority` (string, `"info"`)
