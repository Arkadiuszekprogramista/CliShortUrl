package shorterCmd

import (
	"app/pkg/config"
	"app/pkg/models"
	"app/pkg/shorter"
	"encoding/json"
	"fmt"
	"io"

	"log"
	"net/http"
	"net/url"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

var myDB *shorter.Redis
var app *config.AppConfig

func NewConfig(a *config.AppConfig) {
	app = a
}

var LoadDataFromRedisCMD = &cobra.Command{
	Use:     "get",
	Aliases: []string{"get, load, g"},
	Short:   "Searching for rcord of given key",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		input := args[0]

		db := shorter.NewRedis(myDB)

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

var PrintAllKeysFromRedisCMD = &cobra.Command{
	Use:   "keys",
	Short: "Printing all kays form db(redis)",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		db := shorter.NewRedis(myDB)

		_, err := db.PrintAllKeys()
		if err != nil {
			log.Fatalln(err)
		}

	},
}

var EncodeCMD = &cobra.Command{
	Use:     "encode",
	Aliases: []string{"enc"},
	Short:   "Encode a url using sha256 hash function",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		input := args[0]

		url, err := url.ParseRequestURI(input)

		if err != nil {
			log.Fatalf("%s is not a URL", input)

		} else {

			if app.UseMySite != true {
				s, err := shorter.Encode(url)
				if err != nil {
					log.Printf("Encoded: %s\t%s", input, s)
				}

			} else {

				if app.MySite != args[0] {

					log.Printf("address you trying to encode is not a %v page", app.MySite)
				}
			}
		}
	},
}

var AddEncodedUrlToDBCMD = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Adding encoded url to db",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		input := args[0]

		db := shorter.NewRedis(myDB)

		url, err := url.ParseRequestURI(input)
		if err != nil {
			log.Fatalf("%s is not a URL", input)
		} else {

			if app.UseMySite != true {

				db.AddShortUrlToRedis(url)

			} else {

				if app.MySite != input {

					log.Printf("address you trying to encode is not a %v page", app.MySite)
				}
			}
		}
	},
}

var PopulateDbByUsersCMD = &cobra.Command{
	Use: "populate",
	Short: "Adding random fake users data to redis",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var Response models.Response

		db := shorter.NewRedis(myDB)

		numberOfRecords := args[0]
		url := fmt.Sprint("https://randomuser.me/api/?results=" + string(numberOfRecords))

		c := &http.Client{}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := c.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		b, err := io.ReadAll(resp.Body)


		err = json.Unmarshal(b, &Response)
		if err != nil {
			log.Fatalln(err)
		}
		
		log.Println(len(Response.Results))
		for i := range Response.Results {
			err = db.AddUserToRedis(&Response.Results[i])
			if err != nil {
				log.Fatalln(err)
			}
		}

	},
}

var RunCMD = &cobra.Command{
	Use: "run",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

	},

}

func init() {
	rootCmd.AddCommand(EncodeCMD, PrintAllKeysFromRedisCMD, AddEncodedUrlToDBCMD, LoadDataFromRedisCMD,PopulateDbByUsersCMD)
}
