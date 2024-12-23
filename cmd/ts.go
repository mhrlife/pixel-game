package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"nevissGo/app/serializer"
)

var tsCmd = &cobra.Command{
	Use:   "ts",
	Short: "Convert serializers to type script",
	Run: func(cmd *cobra.Command, args []string) {
		converter := typescriptify.New().
			Add(serializer.UserWithToken{}).
			Add(serializer.User{}).
			Add(serializer.PixelSerializer{}).
			Add(serializer.PixelWithUserSerializer{}).
			Add(serializer.BoardSerializer{}).
			Add(serializer.HypeSerializer{}).
			Add(serializer.UpdatedBoardSerializer{}).
			WithInterface(true)
		err := converter.ConvertToFile("./ui/src/types/serializer.ts")
		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(tsCmd)
}
