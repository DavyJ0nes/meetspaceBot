```
                        __                             ____        __ 
   ____ ___  ___  ___  / /__________  ____ _________  / __ )____  / /_
  / __ `__ \/ _ \/ _ \/ __/ ___/ __ \/ __ `/ ___/ _ \/ __  / __ \/ __/
 / / / / / /  __/  __/ /_(__  ) /_/ / /_/ / /__/  __/ /_/ / /_/ / /_  
/_/ /_/ /_/\___/\___/\__/____/ .___/\__,_/\___/\___/_____/\____/\__/  
                            /_/                                       
```

# Overview
This is a simple bot that links [Hipchat](https://www.hipchat.com/) with  [Meetspace](http://www.meetspaceapp.com) Video chat

The deployment artifact is a GO binary that is run from a docker container. To help keep the container size to a minimum I'm using a static binary by [disabling cgo](https://golang.org/cmd/cgo/) and rebuilding all dependencies as well with cgo disabled. I would have used "FROM scratch" but needed environment variables. [More Infomation](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)

If you want to run it outside of Docker, then within [Releases](./releases) there is the latest stable go binary that is compiled to run on Linux.

# Usage

## Run in Docker
You can run in Docker with:
```
docker run -p 80:8081 -d --name msb-1 -e HIPCHAT_API_TOKEN="" -e MEETSPACE_API_TOKEN="" davyj0nes/meetspacebot"
```
There is also a helper script for running on remote machine. For testing I am just using a basic AWS EC2 instance with Docker installed.

## How to deploy with script
1. Declare Environment variables or change in deploy script
2. Run `./deploy deploy`

## Set up in Hipchat
Please follow this guide: https://blog.hipchat.com/2015/02/11/build-your-own-integration-with-hipchat/
- For the Slash command, it needs to be set as "/meetspace"
- For the url, it needs to follow this schema: "<host>/api/v0/hipchat"

## Run Demo Meetspace API Server
To help with development while the Meetspace API is being added, there is a simple static version of the return data. To run if you can use the following:
```
sudo docker run --restart=unless-stopped -p 8080:8080 -d --name demo-api davyj0nes/meetspacedemoapi
```

# Use of Environment Variables
The Bot requires the following environment variables to be set:
- `HIPCHAT_API_TOKEN` - This is the Hipchat API token. You can generate one from [here](https://www.hipchat.com/account/api) 
- `MEETSPACE_API_TOKEN` - This is the API token for meetspace.
- `MEETSPACE_API_HOST` - host of the Meetspace API. When running Demo API, this will need to be changed.
- `MEETSPACEBOT_TEST` - Set to "true" if you want to use local Demo API Container

# Roadmap
This package was made for personal use but would like to add the following in the future. 
Any contributions are welcome - Just open a PR.
- Slack Support
- Deeper integration with HipChat

# License
This package is distributed under the BSD-style license found in the [LICENSE](./LICENSE) file.
