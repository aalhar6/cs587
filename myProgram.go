package main
import (
"ethos/syscall"
"ethos/ethos"
ethosLog "ethos/log"
"ethos/efmt"
"fmt" 
)
func main () {
me := syscall.GetUser()
path := "/user/" + me + "/myDir/"
status := ethosLog.RedirectToLog("myProgram")
if status != syscall.StatusOk {
efmt.Fprintf(syscall.Stderr, "Error opening %v: %v\n", path, 
status)
syscall.Exit(syscall.StatusOk)
}

fmt.Println("===============================")
data := Box { "foo", "bar" }
data.Field1="Ahmed---------------"
data.Field2="Alharhi-------------"
fmt.Println(data.Field1)
fmt.Println(data.Field2)
fmt.Println("===============================")

fd, status := ethos.OpenDirectoryPath(path)
data.Write(fd)
data.WriteVar(path +"foobar")
efmt.Fprint(syscall.Stderr, "/n *****this will print in the log***** /n")
}
