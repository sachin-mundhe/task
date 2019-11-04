package cmd

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"strings"
// 	"testing"

// 	"gopkg.in/mgo.v2/bson"

// 	db "../db"
// )

// var file *os.File
// var terminalSTDOUT interface{}
// var err error
// var id1 string
// var id2 string
// var id3 string

// func init() {
// 	db.Connect()
// 	db.TaskCollection.RemoveAll(bson.M{})
// 	file, err = os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
// 	if err != nil {
// 		log.Fatalln("Error occured in init function. Error:", err)
// 	}
// 	terminalSTDOUT = os.Stdout
// 	os.Stdout = file
// 	file.Truncate(0)

// }

// func readOutput() string {
// 	file.Seek(0, 0)
// 	var content
// 	content, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		log.Fatalln("Error occured while reading file. Error Description:", err)
// 	}

// }

// func TestCommands(t *testing.T) {

// 	//Test case to list empty list of tasks
// 	t.Run("List - Empty DB", func(tc *testing.T) {

// 		listCmd.Run(listCmd, nil)

// 		output := string(content)
// 		isContain := strings.Contains(output, "No tasks found")
// 		if !isContain {
// 			tc.Error("Error while while reading empty DB\nExpected output: No tasks found\nGot:\n", output)
// 		}
// 		os.Stdout = oldStdout
// 		file.Close()

// 	})

// 	t.Run("Add Command", func(tc *testing.T) {

// 		file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
// 		oldStdout := os.Stdout
// 		os.Stdout = file
// 		file.Truncate(0)

// 		addCmd.Run(addCmd, []string{"Watch Gophercises"})
// 		addCmd.Run(addCmd, []string{"Complete training"})
// 		addCmd.Run(addCmd, []string{"Start working"})

// 		file.Seek(0, 0)
// 		content, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			tc.Error("error occured while test case : ", err)
// 		}

// 		output := string(content)
// 		count := strings.Count(output, "successfully")
// 		if count != 3 {
// 			tc.Error("Error while adding tasks.. Expected tasks: 3, Got:", count)
// 		}

// 		temp := strings.Split(strings.ReplaceAll(output, "\n", " "), " ")
// 		id1 = temp[5]
// 		id2 = temp[11]
// 		id3 = temp[17]
// 		os.Stdout = oldStdout
// 		file.Close()

// 		// for i, v := range temp {
// 		// 	fmt.Println(i, v)

// 		// }

// 		// fmt.Println("ID1", id1)
// 		// fmt.Println("ID2", id2)
// 		// fmt.Println("ID3", id3)

// 	})

// 	t.Run("List - Command", func(tc *testing.T) {

// 		file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
// 		oldStdout := os.Stdout
// 		os.Stdout = file
// 		file.Truncate(0)

// 		listCmd.Run(listCmd, nil)

// 		file.Seek(0, 0)
// 		content, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			tc.Error("error occured while test case : ", err)
// 		}

// 		output := string(content)

// 		if isContain := strings.Contains(output, id1); !isContain {
// 			tc.Error("Error while reading empty ID: ", id1)
// 		}

// 		if isContain := strings.Contains(output, id2); !isContain {
// 			tc.Error("Error while reading empty ID: ", id2)
// 		}

// 		if isContain := strings.Contains(output, id3); !isContain {
// 			tc.Error("Error while reading empty ID: ", id3)
// 		}
// 		os.Stdout = oldStdout
// 		file.Close()

// 	})

// 	t.Run("Do Command - invalid ID", func(tc *testing.T) {

// 		file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
// 		oldStdout := os.Stdout
// 		os.Stdout = file
// 		file.Truncate(0)

// 		doCmd.Run(doCmd, []string{"a2331d325kl76"})

// 		file.Seek(0, 0)
// 		content, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			tc.Error("error occured while test case : ", err)
// 		}

// 		output := string(content)

// 		if isContain := strings.Contains(output, "Invalid ID: a2331d325kl76"); !isContain {
// 			tc.Error("Error while Do command - Invalid ID: a2331d325kl76")
// 		}
// 		os.Stdout = oldStdout
// 		file.Close()
// 	})

// 	t.Run("Do Command - ID not found", func(tc *testing.T) {

// 		file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
// 		oldStdout := os.Stdout
// 		os.Stdout = file
// 		file.Truncate(0)

// 		doCmd.Run(doCmd, []string{"123456789012345678901234"})

// 		file.Seek(0, 0)
// 		content, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			tc.Error("error occured while test case : ", err)
// 		}

// 		output := string(content)

// 		if isContain := strings.Contains(output, "Task not found"); !isContain {
// 			tc.Error("Error while Do command - Task not found")
// 		}
// 		os.Stdout = oldStdout
// 		file.Close()
// 	})

// 	t.Run("Do Command - Valid ID", func(tc *testing.T) {

// 		file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
// 		oldStdout := os.Stdout
// 		os.Stdout = file
// 		file.Truncate(0)

// 		arg := []string{id1}

// 		doCmd.Run(doCmd, arg)

// 		file.Seek(0, 0)
// 		content, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			tc.Error("error occured while test case : ", err)
// 		}

// 		output := string(content)

// 		if isContain := strings.Contains(output, "successfully"); !isContain {
// 			tc.Error("Error while Do command - Valid ID")
// 		}
// 		os.Stdout = oldStdout
// 		file.Close()
// 		fmt.Println("Print from Do Command - Valid ID")
// 	})

// }

// func TestMain(m *testing.M) {
// 	fmt.Println("Before subtest runs")
// 	m.Run()
// 	fmt.Println("After subtest runs")

// }
