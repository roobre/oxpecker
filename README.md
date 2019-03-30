# Oxpecker

Oxpecker is a tool you can run to mirror posted tweets to mastodon-compliant servers.

It currently supports reading from multiple twitter users (also through multiple apps) and posting to multiple masotodon servers.

## Usage

`./oxpecker <path/to/config/file.toml>`

## Roadmap

### Near future

* [x] N:M Twitter to Mastodon echoing
* [ ] Posting also media from Twitter to Mastodon
* [ ] More robust behavior
  * [ ] Retry any of the accounts on failure
* [ ] Allow interactive login instead of relying on config-defined tokens

### Future work

* [ ] Keep track of threads and try to replicate them on the other side
* [ ] N:M Mastodon to Twitter
* [ ] Media from Mastodon to Twiiter 

