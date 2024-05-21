# Useful Stuff

This file contains information, links, tools/commands that can be useful for the video team.

## Testing

### Test Presentation

http://tochev.net/presentation-test/

### Sample RTMP Stream

```sh
ffmpeg -re -f lavfi -i testsrc=size=1920x1080:rate=30 -f lavfi -i "sine=frequency=220" -c:v libx264 -tune zerolatency -af "volume=1.0" -c:a aac -f flv rtmp://localhost/test.stream
```

The above command will send a 1080p test pattern with sine audio.

## Tools

### Play Video with Audio Bar

```sh
mpv --mute --demuxer-lavf-analyzeduration=30 --lavfi-complex='[aid1] asplit [t1] [ao] ; [t1] showvolume=w=1000:h=100 [t2] ; [vid1]  [t2]  overlay  [vo]' SOME_VIDEO_SOURCE
```

Remove `--mute` if you want to start the video with an audible audio.

### ffprobe

```sh
ffprobe SOME_VIDEO_SOURCE
```

Test the resolution, length, codec, etc. of a video source.

## Resources

### Royalty-Free Music

- https://pixabay.com/music/ (do check the license of the file)
