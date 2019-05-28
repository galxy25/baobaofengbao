# baobaofengbao

This repo exists for the purpose of taking a collection of images, an audio file, and transforming them into a single video file that is then segmented for VOD streaming using HLS or MPEG-DASH for delivery.

This repo also contains a (Golang) HTTP server implementation for listing and serving access to these mixes, using the [fluid video player](https://www.fluidplayer.com) for playing of the files in a mobile or desktop browser.

## Pre-requisites

* Docker
* Go 1.11+
* ffmpeg
* bash

# Media Pipeline

## Inputs

## Running

```
./transform.sh *.jpeg Episode2.m4a output.mp4
```

Setup directory for serving the new mix

```
mkdir ./520/mixes/episode2
```

Move all the `ts, m3u8, mp4, mpd` files in the output (or whatever the filename minus the extension was for the third argument passed to the previous `transform.sh` invocation) directory of the inputs folder along with the thumbnails directory to the directory you created in the previous step.

If there is an Apple Music playlist for the mix, create a file with the embed link to the playlist (Itunes > Playlist > Share > Embed Code)

```
echo "https://embed.music.apple.com/us/playlist/episode-2/pl.u-PDb4YVpTZm3x19" > ./520/mixes/episode2/playlist.txt
```

Run the server

```
go build main.go
./main

# OR for local development

go run main.go
```

Enjoy

```
http://localhost:520
```
