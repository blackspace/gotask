package main

import (
	"io"
	"net/http"
	"log"
	"github.com/blackspace/gotasks/task"
)


var task_pool *task.TaskPool

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	t:=task.TaskFromString("HelloWorld")
	task_pool.AddTask(t)

	r:=t.ReceiveResult()

	io.WriteString(w, r.(string))
}

func init() {
	task_pool=task.NewTaskPool()

	go func(){
		for {
			t:=task_pool.GetTask()


			if t!=nil {
				log.Println(t)
				r:=t.Exec()
				t.SendResult(r)
			}


		}
	}()
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}