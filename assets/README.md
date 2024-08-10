```
var (
	cssInp   = filepath.Join("web", "view", "assets", "styles", "app.css")
	cssOut   = filepath.Join("web", "public", "styles", "app.css")
	cssWatch = filepath.Join("web", "view")
	jsInp    = filepath.Join("web", "view", "assets", "scripts", "app.js")
	jsOut    = filepath.Join("web", "public", "scripts", "app.js")
	jsWatch  = filepath.Join("web", "view", "assets", "scripts")
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sig := <-app.WaitForShutdownSignal()
		slog.Info("catch signal to canceling", "signal", sig)
		cancel()
	}()

	if app.IsDevelopment() {
		// development server only tailwind watcher run
		tw := tailwind.New(cssInp, cssOut, cssWatch)
		if err := tw.Run(ctx); err != nil {
			slog.Error("failed to tailwind run", "error", err)
			cancel()
			os.Exit(1)
		}
		// development server only esbuild watcher run
		es := esbuild.New(jsInp, jsOut, jsWatch)
		if err := es.Run(ctx); err != nil {
			slog.Error("failed to esbuild run", "error", err)
			cancel()
			os.Exit(1)
		}
	}
}
```