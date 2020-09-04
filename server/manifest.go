// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "com.mattermost.chr-reminder-plugin",
  "name": "Concept Help Request Reminder Plugin",
  "description": "This Plugin can be used to remind users about the Concept Help Request feature in Scaler from Mattermost. And provide a direct link to raise a CHR",
  "version": "0.1.0",
  "min_server_version": "5.12.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "settings_schema": {
    "header": "The settings page of CHR creation plugin",
    "footer": "Made with \u003c3 by Scaler",
    "settings": [
      {
        "key": "ChrTriggerWords",
        "display_name": "CHR Trigger words",
        "type": "longtext",
        "help_text": "The words to be considered as Triggers for the CH",
        "placeholder": "",
        "default": "Doubt ? what when how would"
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
