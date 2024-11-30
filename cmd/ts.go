/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"nevissGo/app/serializer"
)

// tsCmd represents the ts command
var tsCmd = &cobra.Command{
	Use:   "ts",
	Short: "Convert serializers to type script",
	Run: func(cmd *cobra.Command, args []string) {
		converter := typescriptify.New().
			Add(serializer.UserWithToken{}).
			Add(serializer.User{}).WithInterface(true)
		err := converter.ConvertToFile("./ui/src/types/serializer.ts")
		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(tsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
