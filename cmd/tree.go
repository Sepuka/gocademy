package cmd

import (
	"github.com/sepuka/gocademy/tree"
	"github.com/spf13/cobra"
)

const defaultTree = "1 2 3 4 4 3 4 4 2 3 4 4 3 4 4"

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Build and print a binary tree",
	Long:  `Print a tree like this
              1
       2               2
   3       3       3       3
 4   4   4   4   4   4   4   4
5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5`,
	Run: func(cmd *cobra.Command, args []string) {
		var source = defaultTree
		if len(args) > 0 {
			source = args[0]
		}
		builder := tree.NewTreeBuilder(source)
		view := tree.NewView(builder.Build(builder.Length()))
		for _, item := range view.Items() {
			row := ""
			for _, r := range item {
				row += r.String()
			}
			println(row)
		}
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
}
