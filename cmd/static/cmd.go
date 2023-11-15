/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package static

import (
	"github.com/baoyxing/ncgo/cmd/dynamic"
	"github.com/baoyxing/ncgo/config"
	"github.com/baoyxing/ncgo/meta"
	"github.com/baoyxing/ncgo/pkg/client"
	"github.com/baoyxing/ncgo/pkg/consts"
	"github.com/baoyxing/ncgo/pkg/fallback"
	"github.com/baoyxing/ncgo/pkg/model"
	"github.com/baoyxing/ncgo/pkg/server"
	"github.com/urfave/cli/v2"
)

func Init() *cli.App {
	globalArgs := config.GetGlobalArgs()
	verboseFlag := cli.BoolFlag{Name: "verbose,vv", Usage: "turn on verbose mode"}

	app := cli.NewApp()
	app.Name = meta.Name
	app.Usage = AppUsage
	app.Version = meta.Version
	// The default separator for multiple parameters is modified to ";"
	app.SliceFlagSeparator = consts.Comma

	// global flags
	app.Flags = []cli.Flag{
		&verboseFlag,
	}

	// Commands
	app.Commands = []*cli.Command{
		{
			Name:   InitName,
			Usage:  InitUsage,
			Action: dynamic.Terminal,
		},
		{
			Name:  ServerName,
			Usage: ServerUsage,
			Flags: serverFlags(),
			Action: func(c *cli.Context) error {
				err := globalArgs.ServerArgument.ParseCli(c)
				if err != nil {
					return err
				}

				return server.Server(globalArgs.ServerArgument)
			},
		},
		{
			Name:  ClientName,
			Usage: ClientUsage,
			Flags: clientFlags(),
			Action: func(c *cli.Context) error {
				err := globalArgs.ClientArgument.ParseCli(c)
				if err != nil {
					return err
				}
				return client.Client(globalArgs.ClientArgument)
			},
		},
		{
			Name:  ModelName,
			Usage: ModelUsage,
			Flags: modelFlags(),
			Action: func(c *cli.Context) error {
				if err := globalArgs.ModelArgument.ParseCli(c); err != nil {
					return err
				}
				return model.Model(globalArgs.ModelArgument)
			},
		},
		{
			Name:  FallbackName,
			Usage: FallbackUsage,
			Action: func(c *cli.Context) error {
				if err := globalArgs.FallbackArgument.ParseCli(c); err != nil {
					return err
				}
				return fallback.Fallback(globalArgs.FallbackArgument)
			},
		},
	}
	return app
}

const (
	AppUsage = "All in one tools for CloudWeGo"

	ServerName  = "server"
	ServerUsage = `generate RPC or HTTP server

Examples:
  # Generate RPC server code 
  ncgo server --type RPC --idl  {{path/to/IDL_file.thrift}} --service {{svc_name}}
  
  # Generate HTTP server code 
  ncgo server --type HTTP --idl  {{path/to/IDL_file.thrift}} --service {{svc_name}}
`

	ClientName  = "client"
	ClientUsage = `generate RPC or HTTP client

Examples:
  # Generate RPC client code 
  ncgo client --type RPC --idl  {{path/to/IDL_file.thrift}} --service {{svc_name}}
  
  # Generate HTTP client code 
  ncgo client --type HTTP --idl  {{path/to/IDL_file.thrift}} --service {{svc_name}}
`

	ModelName  = "model"
	ModelUsage = `generate DB model

Examples:
  # Generate DB model code
  ncgo  model --db_type mysql --dsn "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
`
	InitName  = "init"
	InitUsage = `interactive terminals provide a more user-friendly experience for generating code`

	FallbackName  = "fallback"
	FallbackUsage = "fallback to hz or kitex"
)
