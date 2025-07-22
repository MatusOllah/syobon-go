// Copyright (C) 2025 Matúš Ollah
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

//go:build js

package config

import "github.com/BurntSushi/toml"

func Parse() (*Config, error) {
	// In js/wasm, we can't actually dump the default config file so we're just parsing the default config directly.
	var cfg Config
	if err := toml.Unmarshal(defaultConfig, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
