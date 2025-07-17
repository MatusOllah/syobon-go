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

package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/MatusOllah/slogcolor"
	"github.com/MatusOllah/syobon-go/internal/game"
)

// getLogLevel gets the log level from command-line flags.
func getLogLevel() slog.Leveler {
	switch s := strings.ToLower(*logLevelFlag); s {
	case "":
		return slog.LevelInfo
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		panic(fmt.Sprintf("invalid log level: \"%s\"; should be one of \"debug\", \"info\", \"warn\", \"error\"", s))
	}
}

func main() {
	flag.Parse()

	// pprof HTTP server
	if *httpProfileFlag {
		go func() {
			if err := http.ListenAndServe("localhost:6060", nil); err != nil {
				panic(err)
			}
		}()
	}

	// CPU profile
	if *cpuProfileFlag != "" {
		f, err := os.Create(*cpuProfileFlag)
		if err != nil {
			panic(fmt.Errorf("could not create CPU profile: %w", err))
		}
		defer func() {
			if err := f.Close(); err != nil {
				panic(fmt.Errorf("could not close CPU profile: %w", err))
			}
		}()
		if err := pprof.StartCPUProfile(f); err != nil {
			panic(fmt.Errorf("could not start CPU profile: %w", err))
		}
		defer pprof.StopCPUProfile()
	}

	// Logger (using slogcolor: https://github.com/MatusOllah/slogcolor)
	opts := slogcolor.DefaultOptions
	opts.Level = getLogLevel()
	opts.SrcFileLength = 32
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, opts)))

	slog.Info("syobon-g version", "version", Version)
	slog.Info("Go version", "version", runtime.Version(), "os", runtime.GOOS, "arch", runtime.GOARCH)

	//TODO: window icon

	slog.Info("initializing game")
	g := &game.Game{}

	// Ebiten init
	slog.Info("initializing ebitengine")
	g.InitEbiten()

	// Start
	slog.Info("starting game")
	if err := g.Start(); err != nil {
		slog.Error("failed to run game", "err", err)
		os.Exit(1)
	}

	// Memory profile
	if *memProfileFlag != "" {
		f, err := os.Create(*memProfileFlag)
		if err != nil {
			panic(fmt.Errorf("could not create memory profile: %w", err))
		}
		defer func() {
			if err := f.Close(); err != nil {
				panic(fmt.Errorf("could not close memory profile: %w", err))
			}
		}()
		runtime.GC() // get up-to-date statistics
		// Lookup("allocs") creates a profile similar to go test -memprofile.
		// Alternatively, use Lookup("heap") for a profile
		// that has inuse_space as the default index.
		if err := pprof.Lookup("allocs").WriteTo(f, 0); err != nil {
			panic(fmt.Errorf("could not write memory profile: %w", err))
		}
	}
}
