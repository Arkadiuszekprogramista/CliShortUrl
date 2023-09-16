package shorterCmd

import (
	"app/pkg/shorter"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

var myDB *shorter.Redis

var LoadDataFromRedisCmd = &cobra.Command{
	Use: "get",
	Aliases: []string{"get, load, g"},
	Short: "Searching for rcord of given key",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		db := shorter.NewRedis(myDB)

		input := args[0]
		
		get, err := db.LoadDataFromRedis(input)

		if err == redis.ErrNil {
			log.Printf("no key: %s", input)
			return
		}

		if err != nil {
			log.Fatalln(err)
			return
		}
		log.Printf("Key: %s\tValue: %s", input, get)
		return
	},
}


var PrintAllKeysFromRedisCmd = &cobra.Command{
	Use: "keys",
	Aliases: []string{"keys, all, a, k"},
	Short: "Printing all kays form db(redis)",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		db := shorter.NewRedis(myDB)

		_, err := db.PrintAll()
		if err != nil {
			return
		}
	},
}

var EncodeCmd = &cobra.Command{
	Use: "encode",
	Aliases: []string{"enc"},
	Short: "Encode a url using sha256 hash function",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		input := args[0]

		s, err := shorter.Encode(args[0])
		if err != nil {
			return
		}

		log.Printf("Input URL: %s\tShort ULR: %s", input, s)
	},
}


var AddEncodedUrlToDBCmd = &cobra.Command{
	Use: "add",
	Aliases: []string{"a"},
	Short: "Addeing encoded url to db",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		db := shorter.NewRedis(myDB)

		db.AddShortUrlToRedis(args[0])

	},
}

func init(){
	rootCmd.AddCommand(EncodeCmd, PrintAllKeysFromRedisCmd, AddEncodedUrlToDBCmd, LoadDataFromRedisCmd)
}