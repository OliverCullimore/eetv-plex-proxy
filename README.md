# EETV Plex Proxy

An application to emulate the HDHomeRun API which allows Plex Media Server's [DVR feature](https://www.plex.tv/features/live-tv-dvr/) to connect to EETV Boxes.

## Usage

This application is distributed as a docker image, please ensure you have docker [set-up and configured](https://www.digitalocean.com/community/tutorial_collections/how-to-install-and-use-docker) before continuing.

### Quick start

Run the following command replacing `YOUREETVBOXIP` with your EETV box IP address:

`docker run -d -p 5004:5004 --name eetv-plex-proxy --env PROXY_HOST=localhost --env PROXY_PORT=5004 --env EETV_IP=YOUREETVBOXIP --restart unless-stopped -v eetv_plex_proxy:/config olivercullimore/eetv-plex-proxy`

## EETV Proxy Configuration

Set the parameters and environment variables below to configure the proxy:

| Parameter | Function |
| :----: | --- |
| `-p 5004:5004` | Web API |
| `-e PROXY_HOST=localhost` | Specify the host domain/IP to use e.g. 192.168.1.50 |
| `-e PROXY_PORT=5004` | Specify the port to use e.g. 5004 |
| `-e EETV_IP=192.168.1.52` | Specify the IP of the EETV Box e.g. 192.168.1.50 |
| `-e EETV_APP_KEY=` | Specify the AppKey for the EETV Box. Leave blank to use default AppKey |

## Plex configuration
Enter the IP address or hostname of the host running eetv-plex-proxy including port 5004. E.g. ```192.168.1.50:5004```

## For development / running standalone

Build and run the application by running `make run`

Or build the application by running `make build` and then run the binary produced

## Credits
https://github.com/jkaberg/tvhProxy

https://github.com/TheJF/antennas

https://github.com/kevjs1982/python-eetv

https://github.com/kevjs1982/eetv-webui