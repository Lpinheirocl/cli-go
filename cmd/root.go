package cmd

import "github.com/spf13/cobra"

func Execute() {
	root := &cobra.Command{
		Short: "studemt Mananeger",
		Version: "1.0.0",
	} 

	root.AddCommand(
		
	)
}