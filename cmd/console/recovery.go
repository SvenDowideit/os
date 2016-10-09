package console

import (
  "strings"
  "github.com/abiosoft/ishell"
)

func recoveryConsole(){
    // create new shell.
    // by default, new shell includes 'exit', 'help' and 'clear' commands.
    shell := ishell.New()

    // display welcome info.
    shell.Println("Sample Interactive Shell")

    // register a function for "greet" command.
    shell.Register("greet", func(args ...string) (string, error) {
        name := "Stranger"
        if len(args) > 0 {
            name = strings.Join(args, " ")
        }
        return "Hello "+name, nil
    })

    // start shell
    shell.Start()
}
