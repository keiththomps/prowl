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
	"fmt"
	"os"

	"github.com/rem7/goprowl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var prowl goprowl.Goprowl
var app, priority, providerKey, url string

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send [EVENT_NAME] [DESCRIPTION]",
	Short: "Send push notification via prowlapp.com API",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Error: must provide a subject and description to send in notification")
			os.Exit(-1)
		}

		err := prowl.RegisterKey(viper.GetString("apiKey"))

		if err != nil {
			fmt.Println("Error: unable to register API key")
			os.Exit(-1)
		}

		n := &goprowl.Notification{
			Event:       args[0],
			Description: args[1],
			Application: app,
			Priority:    priority,
			Providerkey: providerKey,
			Url:         url,
		}

		err = prowl.Push(n)
		if err != nil {
			fmt.Println("Error: unable to send push notification")
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(sendCmd)

	// Local flags
	sendCmd.Flags().StringVarP(&providerKey, "provider-key", "k", "", "Provider key (only used by whitelisted applications)")
	sendCmd.Flags().StringVarP(&app, "app", "a", "Prowl CLI", "Name of application sending the notification")
	sendCmd.Flags().StringVarP(&url, "url", "u", "", "URL attached to notification. Will be shown and cause redirect when launched")
	sendCmd.Flags().StringVarP(&priority, "priority", "p", "0", "Importance of message, integer from -2 to 2. 2 being emergency")
}
