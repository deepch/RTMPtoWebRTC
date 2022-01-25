# RTMPtoWebRTC

RTMP Stream to WebBrowser over WebRTC based on Pion (full native! not using ffmpeg or gstreamer).

**Note:** [RTSPtoWeb](https://github.com/deepch/RTSPtoWeb) is an improved service that provides the same functionality, an improved API, and supports even more protocols. *RTSPtoWeb is recommended over using this service.*


if you need RTSPtoWSMP4f use https://github.com/deepch/RTSPtoWSMP4f


![RTMPtoWebRTC image](doc/demo4.png)

### Download Source

1. Download source
   ```bash 
   $ git clone https://github.com/deepch/RTMPtoWebRTC  
   ```
3. CD to Directory
   ```bash
    $ cd RTMPtoWebRTC/
   ```
4. Test Run
   ```bash
    $ GO111MODULE=on go run *.go
   ```
5. Open Browser
    ```bash
    open web browser http://127.0.0.1:8083 work chrome, safari, firefox
    ```

## Configuration

### Edit file config.json

format:

```bash
{
  "server": {
    "http_port": ":8083"
  },
  "streams": {
    "demo1": {
      "on_demand" : false
      "url": "rtmp://170.93.143.139/rtplive/470011e600ef003a004ee33696235daa"
    },
    "demo2": {
      "on_demand" : true
      "url": "rtmp://admin:admin123@10.128.18.224/mpeg4"
    },
    "demo3": {
      "on_demand" : false
      "url": "rtmp://170.93.143.139/rtplive/470011e600ef003a004ee33696235daa"
    }
  }
}
```

## Livestreams

Use option ``` "on_demand": false ``` otherwise you will get choppy jerky streams and performance issues when multiple clients connect. 

## Limitations

Video Codecs Supported: H264

Audio Codecs Supported: pcm alaw and pcm mulaw 

## Team

Deepch - https://github.com/deepch streaming developer

Dmitry - https://github.com/vdalex25 web developer

Now test work on (chrome, safari, firefox) no MAC OS

## Other Example

Examples of working with video on golang

- [RTSPtoWeb](https://github.com/deepch/RTSPtoWeb)
- [RTMPtoWebRTC](https://github.com/deepch/RTMPtoWebRTC)
- [RTSPtoWSMP4f](https://github.com/deepch/RTSPtoWSMP4f)
- [RTSPtoImage](https://github.com/deepch/RTSPtoImage)
- [RTSPtoHLS](https://github.com/deepch/RTSPtoHLS)
- [RTSPtoHLSLL](https://github.com/deepch/RTSPtoHLSLL)

[![paypal.me/AndreySemochkin](https://ionicabizau.github.io/badges/paypal.svg)](https://www.paypal.me/AndreySemochkin) - You can make one-time donations via PayPal. I'll probably buy a ~~coffee~~ tea. :tea:
