# bearish_conscience

## Build
Run in project root directory
```
docker build -t bearish_conscience_web_root .
docker run -p1323:1323  -v /var/run/docker.sock:/var/run/docker.sock -it bearish_conscience_main:latest
```