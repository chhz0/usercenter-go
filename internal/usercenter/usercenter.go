package usercenter

import (
	"context"
	"fmt"
	"log"

	"github.com/chhz0/usercenter-go/internal/pkg/conf"
	"github.com/chhz0/usercenter-go/internal/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type UserCenter struct {
	cmd *cobra.Command
}

func NewUserCenter() *UserCenter {
	cobra.OnInitialize(initLog)

	cfg := newConf()
	cmd := &cobra.Command{
		Use:   "usercenter",
		Short: "this is a simple usercenter application.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cfg.Load()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run(cfg)
		},
		SilenceErrors:              true,
		SilenceUsage:               true,
		SuggestionsMinimumDistance: 2,
	}

	bindFlags(cmd, cfg)

	return &UserCenter{
		cmd: cmd,
	}
}

func (u *UserCenter) Run(ctx context.Context) error {
	return u.cmd.ExecuteContext(ctx)
}

func run(cfg *Config) {
	// validation
	// config
	fmt.Printf("%+v\n", cfg.HTTP)
	// new server
	if err := cfg.Server(); err != nil {
		panic(err)
	}
	// run server
}

func bindFlags(cmd *cobra.Command, cfg *Config) {
	pfs := cmd.PersistentFlags()
	lfs := cmd.LocalFlags()

	conf.BindConfigFlag(pfs, cfg.V)

	cfg.BindFlags(lfs)
	cfg.HTTP.BindFlags(lfs)
}

func initLog() {
	zlog, err := logger.NewZapDevelopment()
	if err != nil {
		log.Fatalf("init log error: %v", err)
	}
	zap.ReplaceGlobals(zlog)
}
