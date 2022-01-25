package main

import (
	"errors"
	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/format/rtmp"
	"log"
	"time"
)

var (
	ErrorStreamExitNoVideoOnStream = errors.New("Stream Exit No Video On Stream")
	ErrorStreamExitRtspDisconnect  = errors.New("Stream Exit Rtsp Disconnect")
	ErrorStreamExitNoViewer        = errors.New("Stream Exit On Demand No Viewer")
)

func serveStreams() {
	for k, v := range Config.Streams {
		if !v.OnDemand {
			go RTSPWorkerLoop(k, v.URL, v.OnDemand, v.DisableAudio, v.Debug)
		}
	}
}
func RTSPWorkerLoop(name, url string, OnDemand, DisableAudio, Debug bool) {
	defer Config.RunUnlock(name)
	for {
		log.Println("Stream Try Connect", name)
		err := RTSPWorker(name, url, OnDemand, DisableAudio, Debug)
		if err != nil {
			log.Println(err)
			Config.LastError = err
		}
		if OnDemand && !Config.HasViewer(name) {
			log.Println(ErrorStreamExitNoViewer)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
func RTSPWorker(name, url string, OnDemand, DisableAudio, Debug bool) error {

	RTMPClient, err := rtmp.Dial(url)
	if err != nil {
		return err
	}

	defer RTMPClient.Close()

	codec, err := RTMPClient.Streams()

	if err != nil {
		return err
	}
	var videoIDX int

	if DisableAudio {

		for i, data := range codec {
			if data.Type() != av.H264 {
				codec = remove(codec, i)
			} else {
				videoIDX = i
			}

		}

	}

	if len(codec) > 0 {

		Config.coAd(name, codec)

	}

	timeLine := make([]time.Duration, len(codec))

	for {

		pkt, err := RTMPClient.ReadPacket()

		if err != nil {
			return err
		}

		//video and disabled audio
		if DisableAudio && pkt.Idx == int8(videoIDX) {

			pkt.Duration = pkt.Time - timeLine[pkt.Idx]
			pkt.Idx = 0
			Config.cast(name, pkt)

		} else if !DisableAudio {

			pkt.Duration = pkt.Time - timeLine[pkt.Idx]
			Config.cast(name, pkt)

		}
		if pkt.Idx <= int8(len(codec))-1 {

			timeLine[pkt.Idx] = pkt.Time

		}

	}

}
func remove(slice []av.CodecData, s int) []av.CodecData {
	return append(slice[:s], slice[s+1:]...)
}
