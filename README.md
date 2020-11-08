# Reddit Streamer for Telegram

## Overview

`reddit-streamer` streams posts from your favorite subreddit and posts them to Telegram. The posts can be filtered by keywords and media links specified in the configuration. The streamer does not require reddit credentials to work, and instead uses the Reddit read-only API for streaming posts.

## Getting Started

To get started, copy `config.toml.example` to `config.toml`.

### Configuration

An example configuration to post news about Trump and Biden, limited to WashingtonPost or NYTimes would look like:

```
[general]
interval=60
debug=false

[telegram]
api_key="<telegram-api-key>"
channel_ids=[<channel-id>]

[filters]
subreddit="worldnews"
keywords=["trump", "biden"] # set to [] if you want to match all posts
media_whitelist=["washingtonpost", "nytimes"] # set to [] if you do not want to filter by links
```

`reddit-streamer`, as of right now, can only stream data from a single subreddit. You may or may not specify any filters on the `keywords` and `media_whitelist`. If none of the two filters are specified, all posts from the subreddit will be posted to telegram.

## Installation

### Compiling the Binary

```shell
$ git clone git@github.com:thunderbottom/reddit-streamer.git
$ cd reddit-streamer
$ make dist
$ cp config.toml.example config.toml
$ ./reddit-streamer
```

### Docker Installation

To locally build and run the docker image, make sure you have edited `config.toml` before running:

```shell
$ docker build -t reddit-streamer -f docker/Dockerfile .
$ docker run -v config.toml:/config.toml reddit-streamer
```

If you do not want to build your own docker image:

```shell
$ docker run -v $(pwd)/config.toml:/config.toml thunderbottom/reddit-streamer
```

## License

```
MIT License

Copyright (c) 2020 Chinmay Pai

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
