// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build windows

// win-gui is a minimal Windows GUI wrapper around the local Tailscale daemon.
package main

import (
	"flag"

	"tailscale.com/client/local"
	"tailscale.com/client/systray"
	"tailscale.com/paths"
)

var socket = flag.String("socket", paths.DefaultTailscaledSocket(), "path to tailscaled socket")
var appName = flag.String("app-name", "MyScale", "application name shown in the Windows notification area")

func main() {
	flag.Parse()
	lc := &local.Client{Socket: *socket}
	menu := &systray.Menu{AppName: *appName}
	menu.Run(lc)
}
