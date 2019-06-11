package cmd

import (
	"os"

	"github.com/sepuka/gocademy/resources"
	"github.com/sepuka/gocademy/tree"
	"github.com/spf13/cobra"
)

var vocabularyCmd = &cobra.Command{
	Use:   "freq_vocabulary file.path",
	Short: "Build vocabulary",
	Run: func(cmd *cobra.Command, args []string) {
		dict := tree.BuildFreqVocabulary(resources.DemoText)
		tree.Print(dict, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(vocabularyCmd)
}

