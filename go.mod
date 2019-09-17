module github.com/moethu/webg3n

require (
	github.com/disintegration/imaging v1.6.1
	github.com/g3n/engine v0.1.0
	github.com/gin-gonic/gin v1.4.0
	github.com/gorilla/websocket v1.4.1
	github.com/prometheus/common v0.7.0
	github.com/satori/go.uuid v1.2.0
)

replace github.com/moethu/webg3n/renderer => ./renderer

replace github.com/g3n/engine => github.com/moethu/engine v0.0.0-20190917085758-471685a1006f
