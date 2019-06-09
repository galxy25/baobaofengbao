#! /bin/bash
# use like
# ./transform.sh *.jpeg Episode3.mp3 output.mp4 1
set -x
cd input
# TODO: Make sure the user passed in the right (4) number of arguments
# XXX: $#  Expands to the number of positional parameters in decimal.
# slideshow images glob pattern
IMAGEGLOB=$1
# audio source filename
AUDIOFILE=$2
# output video file name
OUTPUTFILE=$3
# how many seconds to show each picture.
# Default = 1
# https://en.wikibooks.org/wiki/FFMPEG_An_Intermediate_Guide/image_sequence#Framerate
SECONDSPERIMAGE=$4
# https://stackoverflow.com/questions/46328198/ffmpeg-image-music-video
# https://trac.ffmpeg.org/wiki/Slideshow
ffmpeg -loop 1 -framerate 1/${SECONDSPERIMAGE:-1} -pattern_type glob -i "$IMAGEGLOB" -i "$AUDIOFILE" -c:v libx264 -preset veryfast -tune stillimage -c:a copy -shortest "$OUTPUTFILE"
# https://alpine-dash-hls.gq
docker run -v $PWD:/video majamee/alpine-dash-hls $OUTPUTFILE
cd ../
