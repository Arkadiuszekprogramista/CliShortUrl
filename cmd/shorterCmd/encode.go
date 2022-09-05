package shorterCmd

import (
	"app/pkg/shorter"
	"fmt"

	"github.com/spf13/cobra"
)

var addr string

var EncodeCmd = &cobra.Command{
	Use: "encode",
	Aliases: []string{"enc"},
	Short: "Encode a url using sha256 hash function",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		link := shorter.NewLink(addr)
		link.Encode()
		fmt.Println(link.Encode())

	},
}

func init(){
	EncodeCmd.Flags().StringVar(&addr,"Address","","Urt to short")
	rootCmd.AddCommand(EncodeCmd)
}