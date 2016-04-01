# breaker
Simple airbrake 2.3 to file logger

Config via env variables:

```
BREAKER_DIR - log directory
BREAKER_DEBUG - debug mode
BREAKER_HOST - binding host
BREAKER_PORT - binding port
```

And, for each application:
```
BREAKER_[APPNAME]_KEY=[APIKEY]
```
for example:
```
BREAKER_PUMA_KEY=f9e9c42ea3636796a2d35ba1f133e186
```
