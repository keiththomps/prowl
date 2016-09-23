// Copyright © 2016 Keith Thompson <me@keiththomps.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api [API_KEY]",
	Short: "Configure your prowlapp.com API key.",
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			fmt.Println("API key:", viper.Get("apiKey"))
		case 1:
			viper.Set("apiKey", args[0])
			settings, _ := json.Marshal(viper.AllSettings())

			if _, err := os.Stat(defaultCfgFile); err != nil {
				f, err := os.Create(defaultCfgFile)

				if err != nil {
					fmt.Println("Error: couldn't create config file", err)
					os.Exit(1)
				}

				defer f.Close()
				f.Write(settings)
			} else {
				err = ioutil.WriteFile(defaultCfgFile, settings, 0644)

				if err != nil {
					fmt.Println("Error: unable to write config file")
					os.Exit(1)
				}
			}
		default:
			fmt.Println("Error: api only takes 1 parameter")
			os.Exit(1)
		}
	},
}

func init() {
	configCmd.AddCommand(apiCmd)
}
