package format

import (
	"github.com/Danile71/joy4/av/avutil"
	"github.com/Danile71/joy4/format/aac"
	"github.com/Danile71/joy4/format/flv"
	"github.com/Danile71/joy4/format/mp4"
	"github.com/Danile71/joy4/format/rtmp"
	"github.com/Danile71/joy4/format/rtsp"
	"github.com/Danile71/joy4/format/ts"
	"joy4/format/mjpeg"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
	avutil.DefaultHandlers.Add(mjpeg.Handler)
}
