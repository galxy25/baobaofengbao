#! /bin/bash
set -x
cd input
# TODO: Make sure the user passed in the right number of arguments
# XXX: $#  Expands to the number of positional parameters in decimal.
# slideshow images glob pattern
IMAGEGLOB=$1
# audio source filename
AUDIOFILE=$2
# output video file name
OUTPUTFILE=$3
# https://stackoverflow.com/questions/46328198/ffmpeg-image-music-video
# https://trac.ffmpeg.org/wiki/Slideshow
# https://en.wikibooks.org/wiki/FFMPEG_An_Intermediate_Guide/image_sequence#Framerate
# TODO: allow user to specify how many seconds each image should be shown for
ffmpeg -loop 1 -framerate 1/180 -pattern_type glob -i "$IMAGEGLOB" -i "$AUDIOFILE" -c:v libx264 -crf 0 -preset veryfast -tune stillimage -c:a copy -shortest "$OUTPUTFILE"
# https://alpine-dash-hls.gq
docker run -v $PWD:/video majamee/alpine-dash-hls $OUTPUTFILE
cd ..
