package format

import (
	"github.com/kittoa/vdk/av/avutil"
	"github.com/kittoa/vdk/format/aac"
	"github.com/kittoa/vdk/format/flv"
	"github.com/kittoa/vdk/format/mp4"
	"github.com/kittoa/vdk/format/rtmp"
	"github.com/kittoa/vdk/format/rtsp"
	"github.com/kittoa/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
