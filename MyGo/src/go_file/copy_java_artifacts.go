package main
import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)
func main() {
	for {
		arg := os.Args
		totalArgs := len(arg)
		fmt.Println(arg)
		fmt.Println(totalArgs)
		filestoBeProcessed := ""
		if arg == nil {
			fmt.Println("arg is nill so processing jar and war files")
			filestoBeProcessed = "JAVA"
			fmt.Println("filestoBeProcessed :" + filestoBeProcessed)
		}

		cli, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}

		containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}
	datas := make(map[string]types.Container)
	for _, container := range containers {
		fmt.Println("Container Id : " + container.ID)
		fmt.Println("==========================================")
		fmt.Println("Container CMD : " + container.Command)
		fmt.Println("==========================================")
		fmt.Println("Container Image : "+ container.Image)
		fmt.Println("========Container network settings========")
		fmt.Println(container.NetworkSettings.Networks)
		fmt.Println(container.NetworkSettings.Networks)
		s := container.Command
		fmt.Println("---:Info regarding contains:---");
		//if (strings.Contains(container.Command, "JAVA")) || (strings.Contains(container.Command, "java")) {
			if strings.Contains(s, "java") || strings.Contains(s, "JAVA") {
				split := strings.Split(s, " ")
				for _, str := range split {
					if strings.HasSuffix(str, ".jar") || strings.HasSuffix(str, ".war") {
						datas[str] = container
						copyFilefromContainerToHost(str, container)
					}
				}
			}else {
				fmt.Println("filestoBeProcessed :" + filestoBeProcessed)
				fmt.Println("filestoBeProcessed  not yet implemented :" + filestoBeProcessed)
			}
		}
		res2B, _ := json.Marshal(datas)
		fmt.Println(string(res2B))
		f, err := os.Create("f:/Docker_Applications/information.txt")
		_, err = f.Write(res2B)
		check(err)
		time.Sleep(5 * time.Minute)
	}
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func copyFilefromContainerToHost(filename string, container types.Container) {
	fmt.Println(filename)
	destDir := "f:/Docker_Applications/"
	fmt.Printf("A:"+container.ID+":"+filename+"\n")
	fmt.Printf("B:" + destDir+filename+"\n")
	cmd := exec.Command("docker", "cp", container.ID+":"+filename, destDir+filename)
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	fmt.Println(err);
	log.Printf("Command finished with error: %v", err)
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
