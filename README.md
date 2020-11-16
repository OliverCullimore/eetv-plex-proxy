# EETV Plex Proxy

An app to emulate the HDHomeRun API which allows Plex Media Server's [DVR feature](https://www.plex.tv/features/live-tv-dvr/) to connect to EETV Boxes.

## Usage

Here are some example snippets to help you get started creating a container.

### docker-compose ([recommended](https://docs.linuxserver.io/general/docker-compose))

Compatible with docker-compose v2 schemas.

```yaml
---
version: "4.1"
services:
  tvheadend:
    image: olivercullimore/eetv-plex-proxy
    container_name: eetv-plex-proxy
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Europe/London
    volumes:
      - <path to data>:/config
    ports:
      - 5004:5004
    restart: unless-stopped
```

### docker cli

```
docker run -d \
  --name=tvheadend \
  -e PUID=1000 \
  -e PGID=1000 \
  -e TZ=Europe/London \
  -e RUN_OPTS=<run options here> `#optional` \
  -p 9981:9981 \
  -p 9982:9982 \
  -v <path to data>:/config \
  -v <path to recordings>:/recordings \
  --device /dev/dri:/dev/dri `#optional` \
  --device /dev/dvb:/dev/dvb `#optional` \
  --restart unless-stopped \
  ghcr.io/linuxserver/tvheadend
```

## Parameters and enviroment variables

Container images are configured using parameters passed at runtime (such as those above). These parameters are separated by a colon and indicate `<external>:<internal>` respectively. For example, `-p 8080:80` would expose port `80` from inside the container to be accessible from the host's IP on port `8080` outside the container.

| Parameter | Function |
| :----: | --- |
| `-p 5004` | Web API |
| `-e PROXY_HOST=localhost` | Specify the host domain/IP to use e.g. 192.168.1.50 |
| `-e PROXY_PORT=5004` | Specify the port to use e.g. 5004 |
| `-e EETV_IP=192.168.1.52` | Specify the IP of the EETV Box e.g. 192.168.1.50 |
| `-e EETV_APP_KEY=` | Specify the AppKey for the EETV Box. Leave blank to use default AppKey |
| `-e PUID=1000` | for UserID - see below for explanation |
| `-e PGID=1000` | for GroupID - see below for explanation |
| `-e TZ=Europe/London` | Specify a timezone to use e.g. Europe/London |

## Configuration
Set the environment variables below to configure the proxy:

`PROXY_BASE_URL`

`EETV_IP`

`EETV_APP_KEY`

## Plex configuration
Enter the IP of the host running eetv-plex-proxy including port 5004. E.g. ```192.168.1.50:5004```

## Credits
https://github.com/jkaberg/tvhProxy

https://github.com/TheJF/antennas

https://github.com/kevjs1982/python-eetv

https://github.com/kevjs1982/eetv-webui

## User / Group Identifiers

When using volumes (`-v` flags) permissions issues can arise between the host OS and the container, we avoid this issue by allowing you to specify the user `PUID` and group `PGID`.

Ensure any volume directories on the host are owned by the same user you specify and any permissions issues will vanish like magic.

In this instance `PUID=1000` and `PGID=1000`, to find yours use `id user` as below:

```
  $ id username
    uid=1000(dockeruser) gid=1000(dockergroup) groups=1000(dockergroup)
```


&nbsp;

## Support Info

* Shell access whilst the container is running: `docker exec -it tvheadend /bin/bash`
* To monitor the logs of the container in realtime: `docker logs -f tvheadend`
* container version number
  * `docker inspect -f '{{ index .Config.Labels "build_version" }}' tvheadend`
* image version number
  * `docker inspect -f '{{ index .Config.Labels "build_version" }}' ghcr.io/linuxserver/tvheadend`

## Versions

* **16.11.20:** - Initial Release.