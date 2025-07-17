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

import "flag"

var (
	logLevelFlag    = flag.String("log-level", "info", "Log level (\"debug\", \"info\", \"warn\", \"error\")")
	cpuProfileFlag  = flag.String("cpu-profile", "", "Write CPU profile to file")
	memProfileFlag  = flag.String("mem-profile", "", "Write memory profile to file")
	httpProfileFlag = flag.Bool("http-profile", false, "Serve profiling data via HTTP server on port :6060 (see https://pkg.go.dev/net/http/pprof)")
)
