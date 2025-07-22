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

package config

import _ "embed"

//go:embed default_config.toml
var defaultConfig []byte

type Config struct {
	Controls Controls `toml:"controls"`
	Audio    Audio    `toml:"audio"`
	Graphics Graphics `toml:"graphics"`
}

type Controls struct {
	Left          []string `toml:"left"`
	Right         []string `toml:"right"`
	Jump          []string `toml:"jump"`
	DoubleSpeed   []string `toml:"double_speed"`
	SelfDestruct  []string `toml:"self_destruct"`
	Exit          []string `toml:"exit"`
	ReturnToTitle []string `toml:"return_to_title"`
	Fullscreen    []string `toml:"fullscreen"`
}

type Audio struct {
	Mute         bool    `toml:"mute"`
	MasterVolume float64 `toml:"master_volume"`
	SampleRate   int     `toml:"sample_rate"`
}

type Graphics struct {
	Fullscreen       bool `toml:"fullscreen"`
	VSync            bool `toml:"vsync"`
	EnableFPSCounter bool `toml:"enable_fps_counter"`
}
