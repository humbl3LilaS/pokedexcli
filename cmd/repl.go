package cmd

import (
	"bufio"
	"fmt"
	"os"

	customTypes "github.com/humbl3LilaS/pokedexcli/pkg/types"
	"github.com/humbl3LilaS/pokedexcli/pkg/util"
)



func StartRepl(conf *customTypes.AppConfig){
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
		availCmd := util.GetCommands()

		cmd, ok := availCmd[inputCmd]

		if !ok {
			fmt.Println("Invalid Command. Please enter `help` to see all available command")
			continue
		}

		if len(clean) - 1 != cmd.MaxArg {
			fmt.Printf("Invalid Number of Arguments. `%s` only accepts %d argument(s).\n", cmd.Name, cmd.MaxArg)
			continue
		}

		args := []string{}

		if cmd.MaxArg > 0 {
			args = clean[1: cmd.MaxArg + 1]
		}

		err :=	cmd.CallBack(customTypes.CmdArg{
			AppConf: conf,
			Args: args,
		});

		if err != nil {
			fmt.Println(err)
		}

	}
}


