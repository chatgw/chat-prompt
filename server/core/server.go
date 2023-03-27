package core

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"

	"server/global"
	"server/initialize"
)

func RunServer() {
	if global.CHPT_CONFIG.System.UseMultipoint {
		initialize.Redis()
	}

	addr := fmt.Sprintf("%s:%d", global.CHPT_CONFIG.System.Host, global.CHPT_CONFIG.System.Port)
	router := initialize.Routers()
	srv := http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.CHPT_LOG.Error("listen", zap.Error(err))
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.CHPT_LOG.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.CHPT_LOG.Error("Server Shutdown", zap.Error(err))
	}
	global.CHPT_LOG.Info("Server exiting")
}
