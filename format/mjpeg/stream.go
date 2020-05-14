package mjpeg

import (
	"github.com/beqdev/joy4/av"
	"github.com/beqdev/joy4/format/rtsp/sdp"
)

type Stream struct {
	av.CodecData
	Sdp    sdp.Media
	client *Client
}
