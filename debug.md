Debug
==================

Build and run

```bash
make binary shell
/go/src/github.com/docker/docker/bundles/1.4.0-dev/binary/docker -d -H tcp://0.0.0.0:4243 -H unix:///var/run/docker.sock -D
```

test

```bash
docker exec -it $(docker ps -l -q) curl http://127.0.0.1:4243/remote/images/dorowu/ubuntu-desktop-lxde-vnc/json

docker exec -it $(docker ps -l -q) curl -XPOST -N "http://127.0.0.1:4243/images/create2?fromImage=busybox:latest"
docker exec -it $(docker ps -l -q) curl -XPOST -N "http://127.0.0.1:4243/images/create2?fromImage=sequenceiq/busybox"

docker exec -it $(docker ps -l -q) /go/src/github.com/docker/docker/bundles/1.4.0-dev/binary/docker pull2 redis:latest
docker exec -it $(docker ps -l -q) /go/src/github.com/docker/docker/bundles/1.4.0-dev/binary/docker pull2 dorowu/ubuntu-lxqt-vnc:latest
```


Known issues
 1. 後面的 % 升很慢，或是會在某 % 停很久
 2. 後面還有一小段時間解壓
 3. 兼顧 client command , 所以目前輸出的欄位比較怪
 4. 沒給 tag 會全抓，%數判斷會有問題，目前 SPEC UI 一定要選 tag
 5. 有很多 error handle 的部分沒有測到
