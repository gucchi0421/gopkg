package esbuild

import (
	"context"
	"errors"
	"log/slog"

	"github.com/evanw/esbuild/pkg/api"
)

var BuildOptions api.BuildOptions

type Esbuild struct {
	InpPath  string
	OutPath  string
	WatchDir string
}

func New(i, o, w string) *Esbuild {
	return &Esbuild{
		InpPath:  i,
		OutPath:  o,
		WatchDir: w,
	}
}

func (es *Esbuild) Init() error {
	BuildOptions = api.BuildOptions{
		EntryPoints:       []string{es.InpPath},
		Outfile:           es.OutPath,
		Bundle:            true,
		Write:             true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
	}

	result := api.Build(BuildOptions)
	if len(result.Errors) > 0 {
		err := errors.New(result.Errors[0].Text)
		slog.Error("result to error esbuild api", "error", err)
		return err
	}

	return es.BuildJS()
}

func (es *Esbuild) BuildJS() error {
	result := api.Build(BuildOptions)
	if len(result.Errors) > 0 {
		err := errors.New(result.Errors[0].Text)
		slog.Error("failed to build js", "error", err)
		return err
	}

	slog.Info("esbuild js build successfully")
	return nil
}

func (es *Esbuild) Run(ctx context.Context) error {
	if err := es.Init(); err != nil {
		slog.Error("failed to initialize esbuild", "error", err)
		return err
	}

	go func() {
		slog.Info("started esbuild watch..")
		if err := es.Watch(ctx); err != nil {
			slog.Error("failed to esbuild watch", "error", err)
		}
	}()

	return nil
}
