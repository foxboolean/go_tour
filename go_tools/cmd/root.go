package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

// Execute 暴露出 cmd 执行方法
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
}
