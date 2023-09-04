package shorterCmd

import (


	"app/pkg/shorter"
	"log"

	"github.com/spf13/cobra"
)

var addr string
var myDB *shorter.Redis




var PrintAllKeysFromRedisCmd = &cobra.Command{
	Use: "Keys",
	Aliases: []string{"keys, all, a, k"},
	Short: "Printing all kays form db(redis)",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		db := shorter.NewRedis(myDB)

		err := db.PrintAll()
		if err != nil {
			log.Println(err)
		}

	},
}

var EncodeCmd = &cobra.Command{
	Use: "encode",
	Aliases: []string{"enc"},
	Short: "Encode a url using sha256 hash function",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		addr := args[0]
		link := shorter.NewLink(addr)
		s := link.Encode()

		log.Printf("Input URL: %s\tShort ULR: %s", addr, s)
	},
}


var AddEncodedUrlToDBCmd = &cobra.Command{
	Use: "add",
	Aliases: []string{"a"},
	Short: "Addeing encoded url to db",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		db := shorter.NewRedis(myDB)
		addr = args[0]
		link := shorter.NewLink(addr)

		db.AddShortUrlToRedis(link)

	},
}

func init(){
	EncodeCmd.Flags().StringVar(&addr,"Address","","Urt to short")
	rootCmd.AddCommand(EncodeCmd, PrintAllKeysFromRedisCmd, AddEncodedUrlToDBCmd)
}