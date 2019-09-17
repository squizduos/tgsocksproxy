# Lightweight SOCKS5 proxy for Telegram

Really simple and fast SOCKS5 proxy server for Telegram, with automatic restricting access to other web.

Inspired by [ex3ndr](https://github.com/ex3ndr/telegram-proxy) project. 

## Install

Best variant for proxy launch is Docker. 
```
docker pull squizduos/tgsocksproxy:latest
```

Binary packages for Windows, Mac OS X and Linux are available at [Releases](/releases) page.


## Launch

```
docker run -it -p 1080:1080 squizduos/tgsocksproxy
```

By default, bot binds to localhost with port `1080` without user authorization. 

You can test it at local machine, by adding next proxy: [Click here](https://t.me/socks?server=localhost&port=1080)

For remote server testing, we can use any of domains, that redirects to `localhost`, and set authorization credentials:

```
docker run -it -p 1080:1080 \
    -e SOCKS_HOST bot.localtest.me \
    -e SOCKS_USER test \
    -e SOCKS_PASSWORD=testtest \ 
    squizduos/tgsocksproxy
```

Link on your proxy will be printed at log:

```
2019/09/17 19:48:37 [Config] Debug: true
2019/09/17 19:48:37 [Config] Auth: true
2019/09/17 19:48:37 [Config] Restict to white list: true
2019/09/17 19:48:37 [Proxy] Listening at bot.localtest.me:1080
2019/09/17 19:48:37 [Proxy] Use this URL to connect: socks5://test:testtest@bot.localtest.me:1080
```

## Configuration

Bot can be configured using environment variables:

| OPTION           	| Description                                  	| Default     	|
|------------------	|----------------------------------------------	|-------------	|
| `SOCKS_DEBUG`    	| Enables debug mode                           	| `false`     	|
| `SOCKS_HOST`     	| IP or hostname                               	| `127.0.0.1` 	|
| `SOCKS_PORT`     	| Proxy port                                   	| `1080`      	|
| `SOCKS_USER`     	| Default username                             	| ``          	|
| `SOCKS_PASSWORD` 	| Default password                             	| ``          	|
| `SOCKS_RESTRICT`  | Enables restricting access to another sites. 	| `true`      	|

## Restrictions

Proxy uses pre-defined Telegram IP-s and domains, that listed in `rules.json`.

Format of file is described below:

```json
{
    "adresses": [
        "8.8.8.8",
        "4.4.4.4"
    ],
    "networks": [
        "8.8.8.0/24",
        "4.4.0.0/16"
    ],
    "domains": [
        "example.org",
        "google.com"
    ]
}
```

You can disable proxy restrictions with `SOCKS_RESTRICT=false`. Note, that SOCKS5 proxies is not secure.

If you want to change restrictions list, you can build custom Docker image:

```
FROM squizduos/tgsocksproxy:latest

WORKDIR /app
COPY your-rules.json /app/rules.json
```


## Contact

You can [email](mailto:squizduos@gmail.com) or send message to [Telegram](https://t.me/squizduos).

## Licensing

Public Domain