package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/humbl3LilaS/pokedexcli/pkg/callback"
	"github.com/humbl3LilaS/pokedexcli/pkg/util"
)

func StartRepl(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")
		scanner.Scan()
		input := scanner.Text();

		clean := util.CleanInput(input)

		if len(clean) == 0 {
			continue
		}

		inputCmd := clean[0]
		availCmd := callback.GetCommands()

		cmd, ok := availCmd[inputCmd]

		if !ok {
			fmt.Println("Invalid Command. Please enter `help` to see all available command")
			continue
		}

		cmd.CallBack();

	}
}


