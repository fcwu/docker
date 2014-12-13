Debug
==================

Build and run

```bash
make binary shell
/go/src/github.com/docker/docker/bundles/1.4.0-dev/binary/docker -d -d -H tcp://0.0.0.0:4243 -H unix:///var/run/docker.sock -D
```

test

```bash
docker exec -it $(docker ps -l -q) curl http://127.0.0.1:4243/remote/images/dorowu/ubuntu-desktop-lxde-vnc/json
```
