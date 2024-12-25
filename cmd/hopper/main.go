// Copyright 2024 The Contributors to go-eth-app
// This file is part of the go-eth-app project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific
// language governing permissions and limitations under the License.
//
// For a list of contributors, refer to the CONTRIBUTORS file or the
// repository's commit history.

package main

import (
	"log/slog"
	"os"
	"path"

	"github.com/paulwizviz/go-eth-app/internal/hoppercli"
	"github.com/spf13/viper"
)

const (
	DefaultConfigPathName = ".hopper"
	SQLiteDBFileName      = "sqlite.db"
)

func initApp() error {
	configDir, err := hoppercli.CreateConfigDir(DefaultConfigPathName)
	if err != nil {
		return err
	}
	viper.Set(SQLiteDBFileName, path.Join(configDir, SQLiteDBFileName))
	return nil
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	hoppercli.Execute()
}
