package esbuild

import (
	"context"
	"log/slog"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func (es *Esbuild) Watch(ctx context.Context) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer func() {
		if err := watcher.Close(); err != nil {
			slog.Error("failed to watcher clodsed", "error", err)
			return
		}
	}()

	err = watcher.Add(es.WatchDir)
	if err != nil {
		slog.Error("failed to add watch dir", "error", err)
		return err
	}

	for {
		select {
		case <-ctx.Done():
			slog.Info("stopping esbuild watcher")
			return nil
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				if filepath.Ext(event.Name) == ".js" {
					slog.Info("detected change in js", "path", event.Name)
					if err := es.BuildJS(); err != nil {
						slog.Error("failed to build js", "error", err)
					}
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
