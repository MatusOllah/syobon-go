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

package game

import (
	"fmt"
	"log/slog"

	"github.com/MatusOllah/syobon-go/assets"
	"github.com/MatusOllah/syobon-go/internal/config"
	"github.com/MatusOllah/syobon-go/internal/controls"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	input "github.com/quasilyte/ebitengine-input"
)

const (
	Width  = 480
	Height = 420
)

type Game struct {
	grap  [][8]*ebiten.Image
	mgrap []*ebiten.Image
	otom  []*audio.Player
	oto   []*audio.Player

	anx  []int
	any_ []int
	ne   []int
	nf   []int

	fonts []*text.GoTextFace

	audioCtx    *audio.Context
	cfg         *config.Config
	inputSystem input.System
	input       *input.Handler
}

func New() (*Game, error) {
	g := &Game{
		grap:  make([][8]*ebiten.Image, 161),
		mgrap: make([]*ebiten.Image, 51),
		otom:  make([]*audio.Player, 6),
		oto:   make([]*audio.Player, 19),
		anx:   make([]int, 160),
		any_:  make([]int, 160),
		ne:    make([]int, 40),
		nf:    make([]int, 40),
		fonts: make([]*text.GoTextFace, 64),
	}

	// Config
	var err error
	g.cfg, err = config.Parse()
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	// Input
	g.inputSystem.Init(input.SystemConfig{DevicesEnabled: input.AnyDevice})
	keymap, err := controls.LoadKeymapFromConfig(*g.cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to load keymap: %w", err)
	}
	g.input = g.inputSystem.NewHandler(0, keymap)

	// Audio
	g.audioCtx = audio.NewContext(g.cfg.Audio.SampleRate)

	// Resources
	if err := g.loadg(); err != nil {
		return nil, err
	}

	if err := g.setFontSize(16); err != nil {
		return nil, fmt.Errorf("unable to load font: %w", err)
	}

	return g, nil
}

func (g *Game) setFontSize(size int) error {
	f, err := assets.FS.Open("res/sazanami-gothic.ttf")
	if err != nil {
		return err
	}
	defer f.Close()

	src, err := text.NewGoTextFaceSource(f)
	if err != nil {
		return err
	}

	g.fonts[size] = &text.GoTextFace{
		Source: src,
		Size:   float64(size),
	}

	return nil
}

func (g *Game) InitEbiten() {
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Syobon Action (しょぼんのアクション)")
}

func (g *Game) Start() error {
	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	g.inputSystem.Update()

	if g.input.ActionIsJustPressed(controls.ActionFullscreen) {
		g.cfg.Graphics.Fullscreen = !g.cfg.Graphics.Fullscreen
		ebiten.SetFullscreen(g.cfg.Graphics.Fullscreen)
		slog.Info("toggled fullscreen mode", "enabled", g.cfg.Graphics.Fullscreen)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}
