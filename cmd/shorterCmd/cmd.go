package shorterCmd

import (
	"app/pkg/config"
	"app/pkg/shorter"
	"fmt"
	"log"
	"net/url"

	"github.com/gomodule/redigo/redis"
	"github.com/sclevine/agouti"
	"github.com/spf13/cobra"
)

var myDB *shorter.Redis
var app *config.AppConfig

func NewConfig(a *config.AppConfig) {
	app = a
}

var LoadDataFromRedisCmd = &cobra.Command{
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

var PrintAllKeysFromRedisCmd = &cobra.Command{
	Use:   "keys",
	Short: "Printing all kays form db(redis)",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		db := shorter.NewRedis(myDB)

		_, err := db.PrintAll()
		if err != nil {
			log.Fatalln(err)
		}

	},
}

var EncodeCmd = &cobra.Command{
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

var AddEncodedUrlToDBCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Addeing encoded url to db",
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

var MarecheckScriptCmd = &cobra.Command{
	Use:   "marecheck",
	Short: "Using Adam's scrip to checking pages",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		// Start a new WebDriver session
		driver := agouti.ChromeDriver()

		// Start the driver
		if err := driver.Start(); err != nil {
			fmt.Println("Failed to start WebDriver:", err)
			return
		}
		defer driver.Stop()

		// Open a new page using the WebDriver
		page, err := driver.NewPage()
		if err != nil {
			fmt.Println("Failed to open page:", err)
			return
		}

		// Navigate to a URL
		if err := page.Navigate("https://www.example.com"); err != nil {
			fmt.Println("Failed to navigate:", err)
			return
		}

		// Perform actions on the page
		// For example, filling out a form and submitting it
		inputField := page.Find("#inputFieldID")
		if err := inputField.Fill("Hello, World!"); err != nil {
			fmt.Println("Failed to fill input field:", err)
			return
		}

		submitButton := page.Find("#submitButtonID")
		if err := submitButton.Submit(); err != nil {
			fmt.Println("Failed to submit form:", err)
			return
		}

		// Extract and print page content
		content, err := page.HTML()
		if err != nil {
			fmt.Println("Failed to get page content:", err)
			return
		}
		fmt.Println("Page content:", content)
	},
}

func init() {
	rootCmd.AddCommand(EncodeCmd, PrintAllKeysFromRedisCmd, AddEncodedUrlToDBCmd, LoadDataFromRedisCmd, MarecheckScriptCmd)
}
