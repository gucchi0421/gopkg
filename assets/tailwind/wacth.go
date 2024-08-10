package tailwind

import (
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var (
	watchFile = ".templ"
)

func (tw *Tailwind) Watch(ctx context.Context) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %w", err)
	}
	defer func() {
		if err := watcher.Close(); err != nil {
			slog.Error("failed to watcher clodsed", "error", err)
			return
		}
	}()

	err = filepath.Walk(tw.WatchDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, watchFile) {
			return watcher.Add(filepath.Dir(path))
		}
		return nil
	})
	if err != nil {
		slog.Error("failed to add watch paths", "error", err)
		return err
	}

	for {
		select {
		case <-ctx.Done():
			slog.Info("stopping tailwind watcher")
			return nil
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			if event.Op&fsnotify.Write == fsnotify.Write && strings.HasSuffix(event.Name, ".templ") {
				slog.Info("detected change in css", "path", event.Name)
				if err := tw.BuildCSS(); err != nil {
					slog.Error("failed to build css", "error", err)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			slog.Error("watcher error", "error", err)
		}
	}
}
