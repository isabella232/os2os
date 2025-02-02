// Prototype (☺) 2017 Manuel Ramirez Lopez <ramz@zhaw.ch>
// Copyright (©) 2017 Zürcher Hochschule für Angewandte Wissenschaften
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
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
)


// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Deploy a project in your new cluster",
	Long: `Deploy a project in your new cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		up(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func up(cmd *cobra.Command, args []string) {

	getAllValue()
	PathTemplate += "/" + ProjectTo
	loginCluster(ClusterTo, UsernameTo, PasswordTo)
	changeProject(ProjectTo)
	ObjectsOc = getTypeObjects(ObjectsOc)


	for _, typeObject := range ObjectsOc {
		fmt.Println(typeObject)
		fullPath := PathTemplate + "/" + typeObject
		create(fullPath)
	}
}

func create(path string){
	CmdCreate := exec.Command("oc", "create", "-f",  path)
	CmdCreateOut, err := CmdCreate.Output()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error creating " + path)
	}
	//checkErrorMessage(err, "Error running create with path " + path)
	fmt.Println(string(CmdCreateOut))
}