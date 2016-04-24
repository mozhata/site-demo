package lib

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/golang/glog"
)

type Runner interface {
	Start()
	Kill()
	Error() error
}

type builderRunner struct {
	projectDir  ProjectDir
	binName     string
	mainPath    string
	args        []string
	waitResults chan waitResult

	mu            sync.Mutex
	buildErr      error
	startErr      error
	terminatedErr error
	binFile       string
	run           *exec.Cmd
	// output        *CircularBuffer
}

type waitResult struct {
	cmd *exec.Cmd
	err error
}

func NewBuilderRunner(projectDir ProjectDir, binName string, mainPath string, args []string) Runner {
	return &builderRunner{
		projectDir:  projectDir,
		binName:     binName,
		mainPath:    mainPath,
		args:        args,
		waitResults: make(chan waitResult, 10)}
}

func (r *builderRunner) Start() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.run != nil {
		return
	}

	// r.binFile, r.buildErr = BuildA2Binary(r.projectDir, r.binName, r.mainPath)
	// if r.buildErr != nil {
	// 	glog.Error(r.buildErr)
	// 	return
	// }
	//
	// cmd := exec.Command(r.binFile, r.args...)
	// r.output = NewCircularBuffer(1024 * 1024)
	// w := io.MultiWriter(os.Stdout, r.output)
	// cmd.Stderr = w
	// cmd.Stdout = w
	// r.terminatedErr = nil
	// r.startErr = cmd.Start()
	// if r.startErr != nil {
	// 	glog.Error(r.startErr)
	// 	return
	// }
	// go func() {
	// 	err := cmd.Wait()
	// 	r.waitResults <- waitResult{cmd, err}
	// }()
	// r.run = cmd

}

func (r *builderRunner) Kill() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.run == nil {
		return
	}

	alreadyStopped := r.run.ProcessState != nil
	if !alreadyStopped {
		if err := r.run.Process.Signal(os.Interrupt); err != nil {
			glog.Error(err)
		}
	}

	//Wait for process to die before we return or hard kill after 3 sec
	// timedout := time.After(3 * time.Second)
	// waitfor:
	// for {
	// 	select {
	// 	case <-timedout:
	// 		if alreadyStopped {
	// 			r.terminatedErr = fmt.Errorf("Output:\n%s\nError: Stuck in kill", r.output.Bytes())
	// 		}
	// 		if err := r.run.Process.Kill(); err != nil {
	// 			glog.Error("failed to kill: ", err)
	// 		}
	// 	case wait := <-r.waitResults:
	// 		if wait.cmd != r.run {
	// 			continue
	// 		}
	// 		if !alreadyStopped {
	// 			break waitfor
	// 		}
	// 		r.terminatedErr = fmt.Errorf("Output:\n%s\nError: %v", r.output.Bytes(), wait.err)
	// 		break waitfor
	// 	}
	// }
	r.run = nil
}

func (r *builderRunner) Error() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.run != nil && r.run.ProcessState != nil {
		// the process finished but not cleaned up yet.
		r.mu.Unlock()
		r.Kill()
		r.mu.Lock()
	}

	if r.buildErr != nil {
		return r.buildErr
	}
	if r.startErr != nil {
		return r.startErr
	}
	if r.terminatedErr != nil {
		return r.terminatedErr
	}
	if r.run == nil {
		return fmt.Errorf("restarting or not running")
	}
	return nil
}
