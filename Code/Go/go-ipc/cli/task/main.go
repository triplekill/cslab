package main

import (
	"io"
	"log"
	"net/rpc"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/exfly/cslab/Code/Go/go-ipc/api"
	"github.com/exfly/cslab/Code/Go/go-ipc/conf"
)

func main() {
	// start rpc server
	var rpcSvrCmd, _, _, _, err = startRPC()
	if err != nil {
		log.Fatalf("start rpc server err %v", err)
	}

	interactWithRPC()

	// shutdown rpc server and exit
	err = stopProcess(rpcSvrCmd)
	if err != nil {
		log.Fatalf("stop child process failed, error:%s", err)
	}
}

func startRPC() (*exec.Cmd, io.WriteCloser, io.ReadCloser, io.ReadCloser, error) {
	rpcSvrCmd := exec.Command(conf.RPCSvrBinPath)
	rpcSvrStdinPipe, err := rpcSvrCmd.StdinPipe()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	rpcSvrStdoutPipe, err := rpcSvrCmd.StdoutPipe() // get stdout pipe
	if err != nil {
		return nil, nil, nil, nil, err
	}
	rpcSvrStderrPipe, err := rpcSvrCmd.StderrPipe() // get stderr pipe, log write here
	if err != nil {
		return nil, nil, nil, nil, err
	}
	// star the rpc server
	if err = rpcSvrCmd.Start(); err != nil {
		return nil, nil, nil, nil, err
	}
	log.Print("sleep 1s to wait rpc server startup")
	time.Sleep(1 * time.Second)

	go ReadFromPipe(rpcSvrStdoutPipe, "rpc_srv_stdout")
	go ReadFromPipe(rpcSvrStderrPipe, "rpc_srv_stderr")

	return rpcSvrCmd, rpcSvrStdinPipe, rpcSvrStdoutPipe, rpcSvrStderrPipe, err
}

func interactWithRPC() {
	// interact with rpc server
	var reply api.ToDo

	addr := conf.RPCSvrListener
	client, err := rpc.DialHTTP(addr.NetWork, addr.Address)
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	finishApp := api.ToDo{"Finish App", "Started"}
	makeDinner := api.ToDo{"Make Dinner", "Not Started"}
	walkDog := api.ToDo{"Walk the dog", "Not Started"}

	client.Call("Task.MakeToDo", finishApp, &reply)
	client.Call("Task.MakeToDo", makeDinner, &reply)
	client.Call("Task.MakeToDo", walkDog, &reply)

	client.Call("Task.DeleteToDo", makeDinner, &reply)

	client.Call("Task.MakeToDo", makeDinner, &reply)

	client.Call("Task.GetToDo", "Finish App", &reply)
	log.Println("Finish App: ", reply)

	if err = client.Call("Task.EditToDo", api.EditToDo{"Finish App", "Finish App", "Completed"}, &reply); err != nil {
		log.Fatal("Problem editing ToDo: ", err)
	}
}

func stopProcess(cmd *exec.Cmd) error {
	pro, err := os.FindProcess(cmd.Process.Pid)
	if err != nil {
		return err
	}
	err = pro.Signal(syscall.SIGINT)
	if err != nil {
		return err
	}
	cmd.Wait()
	log.Printf("stop subproc %s success\n", cmd.Path)
	return nil
}

// ReadFromPipe read from pipe
func ReadFromPipe(pipe io.ReadCloser, name string) {
	// read the rpc server stdout
	var buffer = make([]byte, 4096)
	for {
		n, err := pipe.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Printf("pipe %v has Closed\n", name)
				break
			} else {
				log.Println("Read content failed")
			}
		}
		log.Print(string(buffer[:n]))
	}
}
