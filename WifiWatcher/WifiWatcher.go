// WifiWatcher.go
// Keeps you connected to wifi even if you have a solid ethernet connection
// Author: Ron Egli - ron.egli@tvrms.com
// Version 0.2

package main

import (
    "fmt"
    "os"
    "log"
    "os/exec"
    "strings"
    "time"
    "github.com/vaughan0/go-ini"
)

var config = "wifiwatcher.ini"
var SSID = ""
var SSID_RAW string
var PROFILE = ""
var INTERFACE = ""
var ok bool


type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func readIni(filename string) {
	// Load the ini file
	file, err := ini.LoadFile(filename)

	// Double check to make sure the filename could be located
	if err != nil {
		// This appears to be caught by the ini plugin.. but we'll keep this here just in case
	    log.Fatal(err)
	    fmt.Print("Unable to find the file: " + filename)
	    os.Exit(1)
	}

	// Read items into memory
	SSID, ok = file.Get("Connection", "SSID")
	PROFILE, ok = file.Get("Connection", "PROFILE")
	INTERFACE, ok = file.Get("Connection", "INTERFACE")

	SSID_RAW = SSID
	SSID = "SSID=" + SSID
	INTERFACE = "INTERFACE=" + INTERFACE
	PROFILE = "NAME=" + PROFILE

	_ = ok // Avoid unused errors with this variable

}

func main(){
	checkIfAdmin()
	readIni(config)
	var wifiStatus = getStatus()

	if wifiStatus {
		fmt.Println("Wifi Status: Connected")
	}else{
		fmt.Println("Wifi Status: Disconnected")

		var inRange = checkArea()

		if inRange {
			fmt.Println("Wifi Status: In Range....")
			fmt.Println("Wifi Status: Connecting....")
			connect()
		}else{
			fmt.Println("Wifi Status: Expected SSID Not In Range: " + SSID_RAW)
		}

		//connect()
		//checkArea()
		time.Sleep(3 * time.Second)
		main()
	}

	time.Sleep(60 * time.Second)
	main()
}

func getStatus() bool{
	// NETSH WLAN SHOW INTERFACE
	out, err := exec.Command("NETSH", "WLAN", "SHOW", "INTERFACE").Output()
    if err != nil {
        log.Fatal(err)
    } 

	if strings.Contains(string(out), "disconnected") {
    	//fmt.Println("Wifi Status: Disconnected")
    	return false
    }else {
    	//fmt.Println("Wifi Status: Connected")
    	return true
    }
}

func connect(){
	//command = "NETSH WLAN CONNECT ssid=" + TVRMS + "interface=" + "W" + name=TVRMS"
	fmt.Println("NETSH", "WLAN", "CONNECT", SSID, INTERFACE, PROFILE)
	out, err := exec.Command("NETSH", "WLAN", "CONNECT", SSID, INTERFACE, PROFILE).Output()
    if err != nil {
        log.Fatal(err)
    } 
    fmt.Println(out)
}

func checkArea() bool{
	out, err := exec.Command("NETSH", "WLAN", "SHOW", "NETWORKS").Output()
	if err != nil {
        log.Fatal(err)
    } 
    if strings.Contains(string(out), SSID_RAW) {
    	return true
    }else{
    	return false
    }
    
}

func checkIfAdmin(){
	Block{
		Try: func() {
			fo, err := os.Create("c:\\test.txt")
			if err != nil {
				Throw("Needs to be run with Administrative Priviledges")
			}
			defer fo.Close()
		},
		Catch: func(e Exception) {
			fmt.Printf("%v\n", e)
			os.Exit(1)
		},
		Finally: func() {
		},
	}.Do()
}

func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}
