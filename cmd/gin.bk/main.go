package main

import (
	"flag"
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/codegangsta/envy/lib"
	"github.com/golang/glog"
	"github.com/zykzhang/site-demo/cmd/gin/lib"

	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var (
	startTime  = time.Now()
	logger     = log.New(os.Stdout, "[gin] ", 0)
	buildError error
)

func main() {
	app := cli.NewApp()
	app.Name = "gin"
	app.Usage = "A live reload utility for Go web applications."
	app.Action = MainAction
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port,p",
			Value: 3000,
			Usage: "port for the proxy server",
		},
		cli.IntFlag{
			Name:  "appPort,a",
			Value: 3001,
			Usage: "port for the Go web server",
		},
		cli.StringFlag{
			Name:  "bin,b",
			Value: "gin-bin",
			Usage: "name of generated binary file",
		},
		cli.StringFlag{
			Name:  "path,t",
			Value: ".",
			Usage: "Path to watch files from",
		},
		cli.StringFlag{
			Name:  "watch",
			Value: "./",
			Usage: "paths to watch",
		},
		cli.BoolFlag{
			Name:  "immediate,i",
			Usage: "run the server immediately after it's built",
		},
		cli.BoolFlag{
			Name:  "godep,g",
			Usage: "use godep when building",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "run",
			ShortName: "r",
			Usage:     "Run the gin proxy in the current working directory",
			Action:    MainAction,
		},
		{
			Name:      "env",
			ShortName: "e",
			Usage:     "Display environment variables set by the .env file",
			Action:    EnvAction,
		},
	}

	app.Run(os.Args)
}

func MainAction(c *cli.Context) {
	flag.Set("logtostderr", "true")
	port := c.GlobalInt("port")
	appPort := strconv.Itoa(c.GlobalInt("appPort"))
	immediate := c.GlobalBool("immediate")
	immediate = true

	// Bootstrap the environment
	envy.Bootstrap()

	// Set the PORT env
	os.Setenv("PORT", appPort)

	var err error
	// wd, err := os.Getwd()
	// if err != nil {
	// 	logger.Fatal(err)
	// }

	builder := gin.NewBuilder(c.GlobalString("path"), c.GlobalString("bin"), c.GlobalBool("godep"))
	runner := gin.NewRunner(filepath.Join(c.GlobalString("path"), builder.Binary()), c.Args()...)
	runner.SetWriter(os.Stdout)
	proxy := gin.NewProxy(builder, runner)

	config := &gin.Config{
		Port:    port,
		ProxyTo: "http://localhost:" + appPort,
	}

	err = proxy.Run(config)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("listening on port %d\n", port)
	shutdown(runner)

	// build right now
	build(builder, runner, logger, immediate)

	// scan for changes
	for {
		scanChanges(c.GlobalString("watch"), func() {
			err := runner.Kill()
			if err != nil {
				glog.Infof("Kill failed: %v", err)
			}
			build(builder, runner, logger, immediate)
		})
		time.Sleep(5000 * time.Millisecond)
	}
}

func EnvAction(c *cli.Context) {
	// Bootstrap the environment
	env, err := envy.Bootstrap()
	if err != nil {
		logger.Fatalln(err)
	}

	for k, v := range env {
		fmt.Printf("%s: %s\n", k, v)
	}

}

func build(builder gin.Builder, runner gin.Runner, logger *log.Logger, immediate bool) {
	err := builder.Build()
	if err != nil {
		buildError = err
		logger.Println("ERROR! Build failed.")
		fmt.Println(builder.Errors())
	} else {
		// print success only if there were errors before
		if buildError != nil {
			logger.Println("Build Successful")
		}
		buildError = nil
		if immediate {
			_, err := runner.Run()
			if err != nil {
				glog.Infof("Run failed: %v", err)
			}
		}
	}

	time.Sleep(100 * time.Millisecond)
}

func scanChanges(watchPath string, cb func()) {
	var fileToMonitor []string
	err := filepath.Walk(watchPath, func(path string, info os.FileInfo, err error) error {
		if path == ".git" {
			return filepath.SkipDir
		}

		basePath := filepath.Base(path)

		// ignore hidden files
		if basePath[0] == '.' {
			return nil
		}

		if basePath == "Godeps" || basePath == "node_modules" {
			return filepath.SkipDir
		}
		ext := filepath.Ext(path)
		if ext != ".amber" && ext != ".go" {
			return nil
		}
		fileToMonitor = append(fileToMonitor, path)
		return nil
	})
	if err != nil {
		glog.Errorf("Error walking dir: %s", watchPath)
		return
	}

	if len(fileToMonitor) == 0 {
		glog.Errorf("No file to monitor!")
		return
	}
	glog.Infof("Monitoring %d files.", len(fileToMonitor))
	for {
		var maxModTime time.Time
		for _, path := range fileToMonitor {
			stat, err := os.Stat(path)
			if err != nil {
				glog.Warningf("Failed to stat: %s. Assumed changed.", path)
				cb()
				startTime = time.Now()
				return
			}

			t := stat.ModTime()
			if t.After(maxModTime) && t.After(startTime) {
				glog.Infof("Found %s changed", path)
				maxModTime = t
			}
		}
		if maxModTime.After(startTime) {
			cb()
			startTime = maxModTime
			return
		}
		time.Sleep(2000 * time.Millisecond)
	}
}

func shutdown(runner gin.Runner) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-c
		log.Println("Got signal: ", s)
		err := runner.Kill()
		if err != nil {
			log.Print("Error killing: ", err)
		}
		os.Exit(1)
	}()
}