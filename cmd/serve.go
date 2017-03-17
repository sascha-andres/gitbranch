// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"net/http"

	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sascha-andres/gitbranch/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start rest api",
	Long:  `Fire up a webserver and answer requests`,
	Run: func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()
		if "" == viper.GetString("serve.secret") {
			r.HandleFunc("/api/branches", app.BranchHandler).
				Methods("POST")
		} else {
			r.HandleFunc("/api/branches", app.BranchHandler).
				Methods("POST").Headers("X-BranchSecret", viper.GetString("serve.secret"))
		}
		r.Headers("Content-Type", "application/json")
		if err := http.ListenAndServe(viper.GetString("serve.listen"), handlers.LoggingHandler(os.Stdout, r)); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringP("listen", "l", ":8080", "Provide binding definition")
	serveCmd.Flags().StringP("secret", "s", "", "Provide a secret that has to be sent with X-BranchSecret")
	viper.BindPFlag("serve.listen", serveCmd.Flags().Lookup("listen"))
	viper.BindPFlag("serve.secret", serveCmd.Flags().Lookup("secret"))
}
