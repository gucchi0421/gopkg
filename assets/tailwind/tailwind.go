package tailwind

import (
	"context"
	"log/slog"
	"os/exec"
)

type Tailwind struct {
	InpPath  string
	OutPath  string
	WatchDir string
}

func New(i, o, w string) *Tailwind {
	return &Tailwind{
		InpPath:  i,
		OutPath:  o,
		WatchDir: w,
	}
}

func (tw *Tailwind) Init() error {
	return tw.BuildCSS()
}

func (tw *Tailwind) BuildCSS() error {
	cmd := exec.Command("tailwindcss", "-i", tw.InpPath, "-o", tw.OutPath, "--minify")

	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("failed to build css", "error", err, "output", output)
		return err
	}

	slog.Info("tailwind css build successfully")
	return nil
}

func (tw *Tailwind) Run(ctx context.Context) error {
	if err := tw.Init(); err != nil {
		slog.Error("failed to initialize esbuild", "error", err)
		return err
	}

	go func() {
		slog.Info("started tailwind watch..")
		if err := tw.Watch(ctx); err != nil {
			slog.Error("failed to tailwind watch", "error", err)
		}
	}()

	return nil
}
