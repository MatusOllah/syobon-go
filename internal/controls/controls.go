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

package controls

import (
	"log/slog"

	"github.com/MatusOllah/syobon-go/internal/config"
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionLeft input.Action = iota
	ActionRight
	ActionJump
	ActionDoubleSpeed
	ActionSelfDestruct
	ActionExit
	ActionReturnToTitle
	ActionFullscreen
)

func LoadKeymapFromConfig(cfg config.Config) (input.Keymap, error) {
	slog.Info("loading keymap")

	keymap := input.Keymap{}

	getKeys := func(s []string, action input.Action) error {
		var keys []input.Key
		for _, ss := range s {
			key, err := input.ParseKey(ss)
			if err != nil {
				return err
			}
			keys = append(keys, key)
		}
		keymap[action] = keys
		return nil
	}

	if err := getKeys(cfg.Controls.Left, ActionLeft); err != nil {
		return nil, err
	}
	if err := getKeys(cfg.Controls.Right, ActionRight); err != nil {
		return nil, err
	}
	if err := getKeys(cfg.Controls.Jump, ActionJump); err != nil {
		return nil, err
	}
	if err := getKeys(cfg.Controls.DoubleSpeed, ActionDoubleSpeed); err != nil {
		return nil, err
	}
	if err := getKeys(cfg.Controls.DoubleSpeed, ActionSelfDestruct); err != nil {
		return nil, err
	}
	if err := getKeys(cfg.Controls.SelfDestruct, ActionSelfDestruct); err != nil {
		return nil, err
	}
	if err := getKeys(cfg.Controls.Exit, ActionExit); err != nil {
		return nil, err
	}
	if err := getKeys(cfg.Controls.ReturnToTitle, ActionReturnToTitle); err != nil {
		return nil, err
	}
	if err := getKeys(cfg.Controls.Fullscreen, ActionFullscreen); err != nil {
		return nil, err
	}

	slog.Debug("loading keymap OK", "keymap", keymap)

	return keymap, nil
}
