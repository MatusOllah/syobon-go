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
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"io/fs"
	"log/slog"

	"github.com/MatusOllah/syobon-go/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) loadg() (err error) {
	// I do not understand any of these Japanese comments xD
	// But this seems to load spritesheets, slice them, and then load audio files.

	slog.Info("loading game resources")

	//プレイヤー
	g.mgrap[0], _, err = ebitenutil.NewImageFromFileSystem(assets.FS, "res/player.png")
	if err != nil {
		return err
	}

	//ブロック
	g.mgrap[1], _, err = ebitenutil.NewImageFromFileSystem(assets.FS, "res/block.png") // Original filename had a typo: "brock.png"
	if err != nil {
		return err
	}

	//アイテム
	g.mgrap[2], _, err = ebitenutil.NewImageFromFileSystem(assets.FS, "res/item.png")
	if err != nil {
		return err
	}

	//敵
	g.mgrap[3], _, err = ebitenutil.NewImageFromFileSystem(assets.FS, "res/teki.png")
	if err != nil {
		return err
	}

	//背景
	g.mgrap[4], _, err = ebitenutil.NewImageFromFileSystem(assets.FS, "res/haikei.png")
	if err != nil {
		return err
	}

	//ブロック2
	g.mgrap[5], _, err = ebitenutil.NewImageFromFileSystem(assets.FS, "res/block2.png") // Original filename had a typo: "brock2.png"
	if err != nil {
		return err
	}

	//おまけ
	g.mgrap[6], _, err = ebitenutil.NewImageFromFileSystem(assets.FS, "res/omake.png")
	if err != nil {
		return err
	}

	//おまけ2
	g.mgrap[7], _, err = ebitenutil.NewImageFromFileSystem(assets.FS, "res/omake2.png")
	if err != nil {
		return err
	}

	//タイトル
	g.mgrap[8], _, err = ebitenutil.NewImageFromFileSystem(assets.FS, "res/syobon3.png")
	if err != nil {
		return err
	}

	//プレイヤー読み込み
	g.grap[40][0] = g.mgrap[0].SubImage(image.Rect(0, 0, 36, 36)).(*ebiten.Image)
	g.grap[0][0] = g.mgrap[0].SubImage(image.Rect(31*4, 0, 30, 36)).(*ebiten.Image)
	g.grap[1][0] = g.mgrap[0].SubImage(image.Rect(31, 0, 30, 36)).(*ebiten.Image)
	g.grap[2][0] = g.mgrap[0].SubImage(image.Rect(31*2, 0, 30, 36)).(*ebiten.Image)
	g.grap[3][0] = g.mgrap[0].SubImage(image.Rect(31*3, 0, 30, 36)).(*ebiten.Image)
	g.grap[41][0] = g.mgrap[0].SubImage(image.Rect(50, 0, 51, 73)).(*ebiten.Image)

	x1 := 1
	//ブロック読み込み
	for i := 0; i <= 6; i++ {
		g.grap[i][x1] = g.mgrap[x1].SubImage(image.Rect(33*i, 0, 30, 30)).(*ebiten.Image)
		g.grap[i+30][x1] = g.mgrap[x1].SubImage(image.Rect(33*i, 33, 30, 30)).(*ebiten.Image)
		g.grap[i+60][x1] = g.mgrap[x1].SubImage(image.Rect(33*i, 66, 30, 30)).(*ebiten.Image)
		g.grap[i+90][x1] = g.mgrap[x1].SubImage(image.Rect(33*i, 99, 30, 30)).(*ebiten.Image)
	}
	g.grap[8][x1] = g.mgrap[x1].SubImage(image.Rect(33*7, 0, 30, 30)).(*ebiten.Image)
	g.grap[16][x1] = g.mgrap[2].SubImage(image.Rect(33*6, 0, 24, 27)).(*ebiten.Image)
	g.grap[10][x1] = g.mgrap[x1].SubImage(image.Rect(33*9, 0, 30, 30)).(*ebiten.Image)
	g.grap[40][x1] = g.mgrap[x1].SubImage(image.Rect(33*9, 33, 30, 30)).(*ebiten.Image)
	g.grap[70][x1] = g.mgrap[x1].SubImage(image.Rect(33*9, 66, 30, 30)).(*ebiten.Image)
	g.grap[100][x1] = g.mgrap[x1].SubImage(image.Rect(33*9, 99, 30, 30)).(*ebiten.Image)
	//ブロック読み込み2
	x1 = 5
	for i := 0; i <= 6; i++ {
		g.grap[i][x1] = g.mgrap[x1].SubImage(image.Rect(33*i, 0, 30, 30)).(*ebiten.Image)
	}
	g.grap[10][5] = g.mgrap[x1].SubImage(image.Rect(33*1, 33, 30, 30)).(*ebiten.Image)
	g.grap[11][5] = g.mgrap[x1].SubImage(image.Rect(33*2, 33, 30, 30)).(*ebiten.Image)
	g.grap[12][5] = g.mgrap[x1].SubImage(image.Rect(33*0, 66, 30, 30)).(*ebiten.Image)
	g.grap[13][5] = g.mgrap[x1].SubImage(image.Rect(33*1, 66, 30, 30)).(*ebiten.Image)
	g.grap[14][5] = g.mgrap[x1].SubImage(image.Rect(33*2, 66, 30, 30)).(*ebiten.Image)

	//アイテム読み込み
	x1 = 2
	for i := 0; i <= 5; i++ {
		g.grap[i][x1] = g.mgrap[x1].SubImage(image.Rect(33*i, 0, 30, 30)).(*ebiten.Image)
	}

	//敵キャラ読み込み
	x1 = 3
	g.grap[0][x1] = g.mgrap[x1].SubImage(image.Rect(33*0, 0, 30, 30)).(*ebiten.Image)
	g.grap[1][x1] = g.mgrap[x1].SubImage(image.Rect(33*1, 0, 30, 43)).(*ebiten.Image)
	g.grap[2][x1] = g.mgrap[x1].SubImage(image.Rect(33*2, 0, 30, 30)).(*ebiten.Image)
	g.grap[3][x1] = g.mgrap[x1].SubImage(image.Rect(33*3, 0, 30, 44)).(*ebiten.Image)
	g.grap[4][x1] = g.mgrap[x1].SubImage(image.Rect(33*4, 0, 33, 35)).(*ebiten.Image)
	g.grap[5][x1] = g.mgrap[7].SubImage(image.Rect(0, 0, 37, 55)).(*ebiten.Image)
	g.grap[6][x1] = g.mgrap[7].SubImage(image.Rect(38*2, 0, 36, 50)).(*ebiten.Image)
	g.grap[150][x1] = g.mgrap[7].SubImage(image.Rect(38*2+37*2, 0, 36, 50)).(*ebiten.Image)
	g.grap[7][x1] = g.mgrap[x1].SubImage(image.Rect(33*6+1, 0, 32, 32)).(*ebiten.Image)
	g.grap[8][x1] = g.mgrap[7].SubImage(image.Rect(38*2+37*3, 0, 37, 47)).(*ebiten.Image)
	g.grap[151][x1] = g.mgrap[7].SubImage(image.Rect(38*3+37*3, 0, 37, 47)).(*ebiten.Image)
	g.grap[9][x1] = g.mgrap[x1].SubImage(image.Rect(33*7+1, 0, 26, 30)).(*ebiten.Image)
	g.grap[10][x1] = g.mgrap[6].SubImage(image.Rect(214, 0, 46, 16)).(*ebiten.Image)

	//モララー
	g.grap[30][x1] = g.mgrap[7].SubImage(image.Rect(0, 56, 30, 36)).(*ebiten.Image)
	g.grap[155][x1] = g.mgrap[7].SubImage(image.Rect(31*3, 56, 30, 36)).(*ebiten.Image)
	g.grap[31][x1] = g.mgrap[6].SubImage(image.Rect(50, 74, 49, 79)).(*ebiten.Image)

	g.grap[80][x1] = g.mgrap[4].SubImage(image.Rect(151, 31, 70, 40)).(*ebiten.Image)
	g.grap[81][x1] = g.mgrap[4].SubImage(image.Rect(151, 72, 70, 40)).(*ebiten.Image)
	g.grap[130][x1] = g.mgrap[4].SubImage(image.Rect(151+71, 72, 70, 40)).(*ebiten.Image)
	g.grap[82][x1] = g.mgrap[5].SubImage(image.Rect(33*1, 0, 30, 30)).(*ebiten.Image)
	g.grap[83][x1] = g.mgrap[6].SubImage(image.Rect(0, 0, 49, 48)).(*ebiten.Image)
	g.grap[84][x1] = g.mgrap[x1].SubImage(image.Rect(33*5+1, 0, 30, 30)).(*ebiten.Image)
	g.grap[86][x1] = g.mgrap[6].SubImage(image.Rect(102, 66, 49, 59)).(*ebiten.Image)
	g.grap[152][x1] = g.mgrap[6].SubImage(image.Rect(152, 66, 49, 59)).(*ebiten.Image)

	g.grap[90][x1] = g.mgrap[6].SubImage(image.Rect(102, 0, 64, 63)).(*ebiten.Image)

	g.grap[100][x1] = g.mgrap[2].SubImage(image.Rect(33*1, 0, 30, 30)).(*ebiten.Image)
	g.grap[101][x1] = g.mgrap[2].SubImage(image.Rect(33*7, 0, 30, 30)).(*ebiten.Image)
	g.grap[102][x1] = g.mgrap[2].SubImage(image.Rect(33*3, 0, 30, 30)).(*ebiten.Image)

	//g.grap[104][x1] = g.mgrap[5].SubImage(image.Rect(33*2, 0, 30, 30)).(*ebiten.Image)
	g.grap[105][x1] = g.mgrap[2].SubImage(image.Rect(33*5, 0, 30, 30)).(*ebiten.Image)
	g.grap[110][x1] = g.mgrap[2].SubImage(image.Rect(33*4, 0, 30, 30)).(*ebiten.Image)

	//背景読み込み
	x1 = 4
	g.grap[0][x1] = g.mgrap[x1].SubImage(image.Rect(0, 0, 150, 90)).(*ebiten.Image)
	g.grap[1][x1] = g.mgrap[x1].SubImage(image.Rect(151, 0, 65, 29)).(*ebiten.Image)
	g.grap[2][x1] = g.mgrap[x1].SubImage(image.Rect(151, 31, 70, 40)).(*ebiten.Image)
	g.grap[3][x1] = g.mgrap[x1].SubImage(image.Rect(0, 91, 100, 90)).(*ebiten.Image)
	g.grap[4][x1] = g.mgrap[x1].SubImage(image.Rect(151, 113, 51, 29)).(*ebiten.Image)
	g.grap[5][x1] = g.mgrap[x1].SubImage(image.Rect(222, 0, 28, 60)).(*ebiten.Image)
	g.grap[6][x1] = g.mgrap[x1].SubImage(image.Rect(151, 143, 90, 40)).(*ebiten.Image)
	g.grap[30][x1] = g.mgrap[x1].SubImage(image.Rect(293, 0, 149, 90)).(*ebiten.Image)
	g.grap[31][x1] = g.mgrap[x1].SubImage(image.Rect(293, 92, 64, 29)).(*ebiten.Image)

	//中間フラグ
	g.grap[20][x1] = g.mgrap[x1].SubImage(image.Rect(40, 182, 40, 60)).(*ebiten.Image)

	//グラ
	x1 = 5
	g.grap[0][x1] = g.mgrap[6].SubImage(image.Rect(167, 0, 45, 45)).(*ebiten.Image)

	//敵サイズ収得
	x1 = 3
	for i := 0; i <= 140; i++ {
		if g.grap[i][x1] != nil {
			g.anx[i] = g.grap[i][x1].Bounds().Dx()
			g.any_[i] = g.grap[i][x1].Bounds().Dy()
			g.anx[i] *= 100
			g.any_[i] *= 100
		} else {
			g.anx[i] = 0
			g.any_[i] = 0
		}
	}
	g.anx[79] = 120 * 100
	g.any_[79] = 15 * 100
	g.anx[85] = 25 * 100
	g.any_[85] = 30 * 10 * 100

	//背景サイズ収得
	x1 = 4
	for i := range 40 {
		if g.grap[i][x1] != nil {
			g.ne[i] = g.grap[i][x1].Bounds().Dx()
			g.nf[i] = g.grap[i][x1].Bounds().Dy()
		} else {
			g.ne[i] = 0
			g.nf[i] = 0
		}
	}

	//ogg読み込み
	g.otom[1], err = g.loadAudio(assets.FS, "BGM/field.ogg")
	if err != nil {
		return err
	}

	g.otom[2], err = g.loadAudio(assets.FS, "BGM/dungeon.ogg")
	if err != nil {
		return err
	}

	g.otom[3], err = g.loadAudio(assets.FS, "BGM/star4.ogg")
	if err != nil {
		return err
	}

	g.otom[4], err = g.loadAudio(assets.FS, "BGM/castle.ogg")
	if err != nil {
		return err
	}

	g.otom[5], err = g.loadAudio(assets.FS, "BGM/puyo.ogg")
	if err != nil {
		return err
	}

	g.oto[1], err = g.loadAudio(assets.FS, "SE/jump.ogg")
	if err != nil {
		return err
	}

	g.oto[3], err = g.loadAudio(assets.FS, "SE/brockbreak.ogg")
	if err != nil {
		return err
	}

	g.oto[4], err = g.loadAudio(assets.FS, "SE/coin.ogg")
	if err != nil {
		return err
	}

	g.oto[5], err = g.loadAudio(assets.FS, "SE/humi.ogg")
	if err != nil {
		return err
	}

	g.oto[6], err = g.loadAudio(assets.FS, "SE/koura.ogg")
	if err != nil {
		return err
	}

	g.oto[7], err = g.loadAudio(assets.FS, "SE/dokan.ogg")
	if err != nil {
		return err
	}

	g.oto[8], err = g.loadAudio(assets.FS, "SE/brockkinoko.ogg")
	if err != nil {
		return err
	}

	g.oto[9], err = g.loadAudio(assets.FS, "SE/powerup.ogg")
	if err != nil {
		return err
	}

	g.oto[10], err = g.loadAudio(assets.FS, "SE/kirra.ogg")
	if err != nil {
		return err
	}

	g.oto[11], err = g.loadAudio(assets.FS, "SE/goal.ogg")
	if err != nil {
		return err
	}

	g.oto[12], err = g.loadAudio(assets.FS, "SE/death.ogg")
	if err != nil {
		return err
	}

	g.oto[13], err = g.loadAudio(assets.FS, "SE/Pswitch.ogg")
	if err != nil {
		return err
	}

	g.oto[14], err = g.loadAudio(assets.FS, "SE/jumpBlock.ogg")
	if err != nil {
		return err
	}

	g.oto[15], err = g.loadAudio(assets.FS, "SE/hintBlock.ogg")
	if err != nil {
		return err
	}

	g.oto[16], err = g.loadAudio(assets.FS, "SE/4-clear.ogg")
	if err != nil {
		return err
	}

	g.oto[17], err = g.loadAudio(assets.FS, "SE/allclear.ogg")
	if err != nil {
		return err
	}

	g.oto[18], err = g.loadAudio(assets.FS, "SE/tekifire.ogg")
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) loadAudio(fsys fs.FS, name string) (*audio.Player, error) {
	b, err := fs.ReadFile(fsys, name)
	if err != nil {
		return nil, fmt.Errorf("failed to open audio file %s: %w", name, err)
	}

	stream, err := vorbis.DecodeF32(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("failed to decode audio file %s: %w", name, err)
	}

	resample := audio.ResampleF32(stream, stream.Length(), 22500, g.audioCtx.SampleRate())

	player, err := g.audioCtx.NewPlayerF32(resample)
	if err != nil {
		return nil, fmt.Errorf("failed to create audio player for %s: %w", name, err)
	}

	return player, nil
}
