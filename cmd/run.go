/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	crmCmd "crm/cmd/crm"
	"crm/config"
	"crm/internal/dao"

	metricsMlwr "github.com/geeksmy/go-libs/goa-libs/middleware/metrics"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
func runCmd() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "启动服务",
		Long:  "管理系统服务启动命令",
		RunE: func(cmd *cobra.Command, args []string) error {
			var metrics *metricsMlwr.Prometheus
			if config.C.Metrics.Enabled {
				metrics = metricsMlwr.NewPrometheus("starter", nil)
				go metrics.Start(config.C.Metrics.Addr)
			}

			// pprof 这里启动失败会直接 panic
			if config.C.Pprof.Enabled {
				go crmCmd.RunDebugPprofServer(config.C.Pprof.Addr)
			}

			crmCmd.RunServer(config.C, metrics)

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			dao.InitDB(config.C)
			return nil
		},
	}

	return runCmd
}
