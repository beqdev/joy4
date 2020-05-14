package format

import (
	"github.com/beqdev/joy4/av/avutil"
	"github.com/beqdev/joy4/format/aac"
	"github.com/beqdev/joy4/format/flv"
	"github.com/beqdev/joy4/format/mp4"
	"github.com/beqdev/joy4/format/rtmp"
	"github.com/beqdev/joy4/format/rtsp"
	"github.com/beqdev/joy4/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
