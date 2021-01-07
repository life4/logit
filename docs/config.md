# Configuration basics

## Levels

Supported levels:

1. `trace`: finer-grained informational events than the `debug`.
1. `debug`: fine-grained informational events that are most useful to debug an application.
1. `info`: informational messages that highlight the progress of the application at coarse-grained level.
1. `warning`: potentially harmful situations.
1. `error`: error events that might still allow the application to continue running.
1. `fatal`: very severe error events that will presumably lead the application to abort.
1. `panic`: very severe error events that will definitely lead the application to abort.

Levels are case-insensitive. The log entry level should be specified in the entry itself. Otherwise, level from `level` config section will be used:

```toml
[levels]

# `default` level is used when the message is not a JSON but plain text.
default = "info"

# `error` level is used when the message is supposedly JSON (starts with `{`) but cannot be parsed (invalid JSON, missed message, wrong time format, unknown level, and so on).
error = "error"
```

## Fields

Logit expects 3 fields to be represented in every log record:

* Message. Default: `msg`. Type: string. Required.
* Level. Default: `level`. Type: string. Required.
* Time when the event happened. Default: `time`. Type: string. Optional, the current time will be used if the field is missed.

Names for fields can be configured:

```toml
[fields]
message = "msg"
level   = "level"
time    = "time"
```

## Handlers

* You can specify as many handlers as you want, and every handler will be called for every log entry.
* Every `[[handler]]` section must have `format` option which specifies the handler type.
* Every format has it's own options which you can find in [docs/handlers.toml](./docs/handlers.toml). Some options are common and secribed below.
* Options `level_from` and `level_to` allow to filter out by level records that the handler will process. Example:

    ```toml
    # `text` handler will handle `trace`, `debug`, and `info`
    [[handler]]
    format = "text"
    level_to = "info"

    # `json` handler will handle `warning`, `error`, `fatal`, and `panic`
    [[handler]]
    format = "json"
    level_from = "warning"
    ```

* Option `workers = N` makes logit to run concurrently run N workers for the handler. It is useful for some network-based handlers. Some handlers handle concurrency on their side without providing an option to turn it off.

See [handlers.md](./handlers.md) for handler-specific configuration options.
