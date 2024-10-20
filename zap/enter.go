package zaplog

import (
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//设置日志编译器，什么类型的日志
func getEncoder() zapcore.Encoder {
	//encoder配置
	encoderConfig := zap.NewProductionEncoderConfig()
	//设置时间格式为2024-9-1-12.32
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//json格式
	// JsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	//终端形式
	ConsoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	return ConsoleEncoder
}

//设置输出位置
func getwriteSyncer(logfilename string) zapcore.WriteSyncer {
	//日志文件
	// logfile, _ :=os.OpenFile(logfilename,os.O_APPEND | os.O_CREATE|os.O_RDWR,0666)

	// 分割日志文件
	l, _ :=rotatelogs.New(logfilename+".%Y%m%d%H%M.log", 
		rotatelogs.WithMaxAge(30*24*time.Hour), // 最长保存时间 30天
		rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
	)
	//只输出到日志文件
	// return zapcore.AddSync(logfile)

	//也输出到终端
	wc := io.MultiWriter(l,os.Stdout)
	return zapcore.AddSync(wc)
}


func InitLogger() *zap.Logger {
	//编码器
	// encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	encoder := getEncoder()

	//输出位置
	// writeSyncer := getwriteSyncer("zap_log.log")
	// core := zapcore.NewCore(encoder,writeSyncer,zap.DebugLevel)
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l > zapcore.WarnLevel
	})
	// 记录不同级别的日志
	core1 := zapcore.NewCore(
		encoder,
		getwriteSyncer("log.all.log"),
		levelEnabler,//全记录
	)
		//错误日志
		core2 :=  zapcore.NewCore(
			encoder,
			getwriteSyncer("log.err.log"),
			zapcore.ErrorLevel,
		)
		
	// logger,_ := zap.NewProduction()
	// return logger
	//创建单个logger
	// logger:= zap.New(core1,zap.AddCaller(), zap.AddCallerSkip(1)) //AddCaller详细记录调用的代码行，AddCallerSkip(1)调用链很多时直接跳过
	// return logger.Sugar()
	//创建双日志，全日志和错误日志

	core :=zapcore.NewTee(core1, core2)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

// gin默认路由使用了自带的两个中间件
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
// func Default(opts ...OptionFunc) *Engine {
// 	debugPrintWARNINGDefault()
// 	engine := New()
// 	engine.Use(Logger(), Recovery())
// 	return engine.With(opts...)
// }

// 在gin框架中使用zap日志记录器
// 要想将zap集成到gin中，就需要重写logger和recovery两个中间件。
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool

				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)

				if brokenPipe {
					logger.Error(c.Request.URL.Path,
					 zap.Any("error", err),
					 zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
					 zap.Any("error", err),
					 zap.String("request", string(httpRequest)),
					 zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
					 zap.Any("error", err),
					 zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}

var Logger *zap.Logger
func MockBuniness() {
	//初始化zap日志记录器
	Logger = InitLogger()
	zap.ReplaceGlobals(Logger)


	defer Logger.Sync()
	// Simplefunc("https://www.baidu.com")

	// Simplefunc("https://sexcn.me")
	r := gin.New()
	zap.L().Info("info:非标准插件，请按照文档自动迁移使用")
	zap.L().Error("非标准插件，请按照文档自动迁移使用")
	//将zap自定义的logger嵌入到gin中。
	r.Use(GinLogger(Logger), GinRecovery(Logger, true))
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(ctx.Writer.Status(),gin.H{"ok": "11"})
	})
	r.Run(":8088")

}

func Simplefunc(url string) {
	res,err:=http.Get(url)
	if err != nil {
		//记录错误日志
		Logger.Error(
			"http get failed..",
			zap.String("url:",url),
			zap.Error(err),
		)
	} else {
		//使用info记录成功日志。
		Logger.Info(
			"get success",
			zap.String("status:",res.Status),
			zap.String("url:",url),
		)
	}
	res.Body.Close()
}