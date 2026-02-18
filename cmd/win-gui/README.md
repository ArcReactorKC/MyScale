# win-gui

`win-gui` is a minimal Windows GUI wrapper that hosts the existing systray UI.

## Build

From this repository root:

```bash
go build tailscale.com/cmd/win-gui
```

## Run

```powershell
.\win-gui.exe -app-name "MyScale"
```

It connects to a running `tailscaled` instance and lets you manage connection state,
profiles, and exit nodes from the Windows notification area.
