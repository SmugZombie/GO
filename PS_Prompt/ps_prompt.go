// PS_Prompt
// Utilizes PowerShell for creating interactive Prompts
// Ron Egli - github.com/smugzombie

package main

import (
	"os/exec"
	"fmt"
)

func main() {
	output, err := exec.Command("PowerShell", "[System.Reflection.Assembly]::LoadWithPartialName('Microsoft.VisualBasic') | Out-Null; $computer = [Microsoft.VisualBasic.Interaction]::InputBox(\"Enter a computer name\", \"Computer\", \"$env:computername\") ; echo $computer").Output()
	if err != nil {
        fmt.Println("Whoops:", err)
    }else{
    	fmt.Println("You input:" + string(output))
    }
}
