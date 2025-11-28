package cmd

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/jzelinskie/cobrautil/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap"

	v1 "github.com/tupyy/assisted-migration-agent/api/v1"
	"github.com/tupyy/assisted-migration-agent/internal/config"
	"github.com/tupyy/assisted-migration-agent/internal/handlers"
	"github.com/tupyy/assisted-migration-agent/internal/server"
)

func NewRunCommand(cfg *config.Configuration) *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run agent",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
			wg := sync.WaitGroup{}
			wg.Add(1)

			// init handlers
			h := handlers.New()

			srv, err := server.NewServer(cfg, func(router *gin.RouterGroup) {
				v1.RegisterHandlers(router, h)
			})
			if err != nil {
				zap.S().Errorw("failed to create http server", "error", err)
				return err
			}

			go func() {
				defer func() {
					wg.Done()
					cancel()
				}()
				zap.S().Infof("Starting HTTP server on port %d", cfg.HTTPPort)

				if err := srv.Start(ctx); err != nil {
					if !errors.Is(err, http.ErrServerClosed) {
						zap.S().Errorw("failed to start http server", "error", err)
					}
				}
			}()

			go func() {
				<-ctx.Done()
				stopCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()
				srv.Stop(stopCtx)
			}()

			<-ctx.Done()
			wg.Wait()

			zap.S().Info("server shutdown")

			return nil
		},
	}

	registerFlags(runCmd, cfg)

	return runCmd
}

func registerFlags(cmd *cobra.Command, config *config.Configuration) {
	nfs := cobrautil.NewNamedFlagSets(cmd)

	serverFlagSet := nfs.FlagSet(color.New(color.FgCyan, color.Bold).Sprint("Server"))
	registerServerFlags(serverFlagSet, config)

	authenticationFlagSet := nfs.FlagSet(color.New(color.FgBlue, color.Bold).Sprint("Authentication"))
	registerAuthenticationFlags(authenticationFlagSet, config)

	nfs.AddFlagSets(cmd)
}

func registerServerFlags(flagSet *pflag.FlagSet, config *config.Configuration) {
	flagSet.StringVar(&config.Mode, "mode", config.Mode, "agent starting mode. Cound be connected to disconnected. Defaults to disconnected")
	flagSet.IntVar(&config.HTTPPort, "http-port", config.HTTPPort, "port on which the HTTP server is listening")
	flagSet.StringVar(&config.StaticsFolder, "statics-folder", config.StaticsFolder, "path to statics")
	flagSet.StringVar(&config.DataFolder, "data-folder", config.DataFolder, "path to the root folder container media")
	flagSet.StringVar(&config.ServerMode, "server-mode", config.ServerMode, "server mode: either prod or dev. If prod it statics folder must be set")
	flagSet.StringVar(&config.ConsoleURL, "console-url", config.ConsoleURL, "url of console.redhat.com.Defaults to localhost:7443")
}

func registerAuthenticationFlags(flagSet *pflag.FlagSet, config *config.Configuration) {
	flagSet.BoolVar(&config.Auth.Enabled, "authentication-enabled", config.Auth.Enabled, "enable authentication when connecting to console")
	flagSet.StringVar(&config.Auth.JWTFilePath, "authentication-jwt-filepath", config.Auth.JWTFilePath, "path of the jwt file")
}

func registerHandlerFn() func(router *gin.RouterGroup) {
	return func(router *gin.RouterGroup) {
	}
}
