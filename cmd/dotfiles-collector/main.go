package main

func main() {
	app := &App{NAME: "dotfiles-collector", BufferSize: 4096}
	app.setupDirectories()
	app.setupDB()
	app.setupCollectPaths()
	app.setupIgnorePaths()
	app.copyFiles()
}
