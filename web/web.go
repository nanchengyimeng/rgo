package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Web struct {
	engine *gin.Engine
	config Config
}

func NewWeb(c Config) *Web {
	if c.Mode != "pro" {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Web{
		engine: gin.New(),
		config: c,
	}
}

// getLogInfo
// @Description: 格式化日志内容
// @receiver w
//func (w *Web) getLogInfo(c *gin.Context) []string {
//	//替换响应writer为自定义结构体
//	blw := &bodyLogWriter{
//		c.Writer,
//		bytes.NewBufferString(""),
//	}
//	c.Writer = blw
//
//	// 记录请求时间
//	start := time.Now()
//
//	// 请求信息获取
//	reqBody, _ := c.GetRawData()
//	if len(reqBody) > 0 {
//		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
//	}
//
//	// 执行请求处理程序和其他中间件函数
//	c.Next()
//
//	_ := blw.body.String()
//	// 记录回包内容和处理时间
//	end := time.Now()
//	_ = end.Sub(start)
//
//	return
//}

func (w *Web) SetLoggerToFile(c *gin.Context) {

	return
}

// Router
// @Description: 路由注册
// @receiver w
// @param handler
func (w *Web) Router(handler func(engine *gin.Engine)) {
	handler(w.engine)
}

func (w *Web) Uniqid() {
	w.engine.Use(func(ctx *gin.Context) {
		if _, ok := ctx.Get("uniqId"); !ok {
			id := xid.New()
			ctx.Set("uniqId", id.String())
		}
	})
}

// Start
// @Description: 服务启动
// @receiver w
func (w *Web) Start() {
	w.engine.Use(gin.Recovery())

	w.Uniqid()

	if len(w.config.Port) == 0 {
		panic("lost port of web in config")
	}
	w.engine.Run(":" + w.config.Port)
}
