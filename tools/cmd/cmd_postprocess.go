package cmd

import (
	"github.com/spf13/cobra"
)

const usePost = "post"
const shortPost = "Post processes whatever files need to be post processed"
const longPost = `
This subcommand post processes the other commands.
`

var postCmd = &cobra.Command{
	Use:   usePost,
	Short: shortPost,
	Long:  colorHelp(longPost),
	RunE:  runPost,
}

func init() {
	rootCmd.AddCommand(postCmd)
}
