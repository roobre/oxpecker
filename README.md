# Oxpecker

Oxpecker is a tool you can run to mirror posted tweets to mastodon-compliant servers.

It currently supports reading from multiple twitter users (also through multiple apps) and posting to multiple masotodon servers.

## Usage

`./oxpecker <path/to/config/file.toml>`

## Features

* [x] N:M Twitter to Mastodon echoing
* [x] Keep track of threads and try to replicate them on the other side

### Near future

* [ ] Post also media from Twitter to Mastodon ([Blocker](https://github.com/McKael/madon/issues/6))
* [ ] More robust behavior
  * [ ] Retry any of the accounts on failure

### Far future

* [ ] Allow interactive login instead of relying on config-defined tokens
* [ ] Posts from Mastodon to Twitter
    * [ ] Media from Mastodon to Twiiter 

