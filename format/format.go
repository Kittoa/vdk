package format

import (
	"github.com/Kittoa/vdk/av/avutil"
	"github.com/Kittoa/vdk/format/aac"
	"github.com/Kittoa/vdk/format/flv"
	"github.com/Kittoa/vdk/format/mp4"
	"github.com/Kittoa/vdk/format/rtmp"
	"github.com/Kittoa/vdk/format/rtsp"
	"github.com/Kittoa/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
