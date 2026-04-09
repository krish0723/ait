{
  "patcher": {
    "fileversion": 1,
    "appversion": {
      "major": 8,
      "minor": 5,
      "revision": 0,
      "architecture": "x64"
    },
    "classnamespace": "box",
    "rect": [
      100.0,
      100.0,
      780.0,
      720.0
    ],
    "bglocked": 0,
    "openinpresentation": 1,
    "default_fontsize": 12.0,
    "default_fontface": 0,
    "default_fontname": "Arial",
    "gridsize": [
      15.0,
      15.0
    ],
    "boxes": [
      {
        "box": {
          "id": "obj-loadbang",
          "maxclass": "loadbang",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            30.0,
            30.0,
            40.0,
            22.0
          ],
          "presentation": 0
        }
      },
      {
        "box": {
          "id": "obj-msg-load",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            30.0,
            70.0,
            95.0,
            22.0
          ],
          "presentation": 0,
          "text": "load_settings"
        }
      },
      {
        "box": {
          "id": "obj-node",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 2,
          "outlettype": [
            "",
            ""
          ],
          "patching_rect": [
            30.0,
            120.0,
            320.0,
            22.0
          ],
          "presentation": 0,
          "text": "node.script ait-control.node.js @autostart 1"
        }
      },
      {
        "box": {
          "id": "obj-route",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 12,
          "outlettype": [
            "",
            "",
            "",
            "",
            "",
            "",
            "",
            "",
            "",
            "",
            "",
            ""
          ],
          "patching_rect": [
            30.0,
            180.0,
            560.0,
            22.0
          ],
          "presentation": 0,
          "text": "route exit preview ait_path git_path project_root status toast git_effective_root git_branch git_status git_clear_branches git_append_branch"
        }
      },
      {
        "box": {
          "id": "obj-num-exit",
          "maxclass": "number",
          "numinlets": 1,
          "numoutlets": 2,
          "outlettype": [
            "",
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            30.0,
            230.0,
            50.0,
            22.0
          ],
          "presentation": 1,
          "presentation_rect": [
            320.0,
            114.0,
            50.0,
            22.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-msg-preview",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            100.0,
            230.0,
            450.0,
            22.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            468.0,
            400.0,
            100.0
          ],
          "text": "(stdout/stderr preview)"
        }
      },
      {
        "box": {
          "id": "obj-pre-ait",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            280.0,
            220.0,
            65.0,
            22.0
          ],
          "presentation": 0,
          "text": "prepend set"
        }
      },
      {
        "box": {
          "id": "obj-pre-git",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            280.0,
            250.0,
            65.0,
            22.0
          ],
          "presentation": 0,
          "text": "prepend set"
        }
      },
      {
        "box": {
          "id": "obj-msg-ait",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            400.0,
            220.0,
            300.0,
            22.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            28.0,
            354.0,
            22.0
          ],
          "text": "/usr/local/bin/ait"
        }
      },
      {
        "box": {
          "id": "obj-msg-git",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            400.0,
            250.0,
            300.0,
            22.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            56.0,
            354.0,
            22.0
          ],
          "text": "/usr/bin/git"
        }
      },
      {
        "box": {
          "id": "obj-b-save",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            400.0,
            300.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            114.0,
            24.0,
            24.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-t-save",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 3,
          "outlettype": [
            "bang",
            "bang",
            "bang"
          ],
          "patching_rect": [
            400.0,
            340.0,
            42.0,
            22.0
          ],
          "presentation": 0,
          "text": "t b b b"
        }
      },
      {
        "box": {
          "id": "obj-join",
          "maxclass": "newobj",
          "numinlets": 3,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            400.0,
            390.0,
            95.0,
            22.0
          ],
          "presentation": 0,
          "text": "join 3"
        }
      },
      {
        "box": {
          "id": "obj-pre-save",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            400.0,
            430.0,
            125.0,
            22.0
          ],
          "presentation": 0,
          "text": "prepend save_settings"
        }
      },
      {
        "box": {
          "id": "obj-pre-toast",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            200.0,
            210.0,
            75.0,
            22.0
          ],
          "presentation": 0,
          "text": "prepend set"
        }
      },
      {
        "box": {
          "id": "obj-msg-toast",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            200.0,
            240.0,
            480.0,
            22.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            276.0,
            400.0,
            22.0
          ],
          "text": "\u2014"
        }
      },
      {
        "box": {
          "id": "obj-lbl-toast",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            200.0,
            218.0,
            140.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            258.0,
            200.0,
            20.0
          ],
          "text": "Status / toast"
        }
      },
      {
        "box": {
          "id": "obj-lbl-cmd",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            520.0,
            480.0,
            160.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            302.0,
            220.0,
            20.0
          ],
          "text": "ait commands"
        }
      },
      {
        "box": {
          "id": "obj-b-version",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            520.0,
            510.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            324.0,
            22.0,
            22.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-msg-cmd-version",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            560.0,
            510.0,
            55.0,
            22.0
          ],
          "presentation": 0,
          "text": "version"
        }
      },
      {
        "box": {
          "id": "obj-b-version-json",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            520.0,
            540.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            44.0,
            324.0,
            22.0,
            22.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-msg-cmd-version-json",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            560.0,
            540.0,
            95.0,
            22.0
          ],
          "presentation": 0,
          "text": "version_json"
        }
      },
      {
        "box": {
          "id": "obj-b-doctor",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            520.0,
            570.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            72.0,
            324.0,
            22.0,
            22.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-msg-cmd-doctor",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            560.0,
            570.0,
            50.0,
            22.0
          ],
          "presentation": 0,
          "text": "doctor"
        }
      },
      {
        "box": {
          "id": "obj-b-doctor-json",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            520.0,
            600.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            100.0,
            324.0,
            22.0,
            22.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-msg-cmd-doctor-json",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            560.0,
            600.0,
            85.0,
            22.0
          ],
          "presentation": 0,
          "text": "doctor_json"
        }
      },
      {
        "box": {
          "id": "obj-msg-cmd-init",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            520.0,
            640.0,
            320.0,
            22.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            352.0,
            400.0,
            22.0
          ],
          "text": "init_run ableton samples-ignored 0 0 0"
        }
      },
      {
        "box": {
          "id": "obj-b-init",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            480.0,
            640.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            378.0,
            22.0,
            22.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-lbl-init-hint",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            520.0,
            670.0,
            420.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            44.0,
            378.0,
            380.0,
            20.0
          ],
          "text": "init_run <daw> <preset> <dry 0|1> <force 0|1> <json 0|1>"
        }
      },
      {
        "box": {
          "id": "obj-b-hooks-in",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            520.0,
            700.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            404.0,
            22.0,
            22.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-msg-hooks-in",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            560.0,
            700.0,
            95.0,
            22.0
          ],
          "presentation": 0,
          "text": "hooks_install 0"
        }
      },
      {
        "box": {
          "id": "obj-b-hooks-in-json",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            520.0,
            730.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            44.0,
            404.0,
            22.0,
            22.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-msg-hooks-in-json",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            560.0,
            730.0,
            95.0,
            22.0
          ],
          "presentation": 0,
          "text": "hooks_install 1"
        }
      },
      {
        "box": {
          "id": "obj-b-hooks-out",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            520.0,
            760.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            72.0,
            404.0,
            22.0,
            22.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-msg-hooks-out",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            560.0,
            760.0,
            105.0,
            22.0
          ],
          "presentation": 0,
          "text": "hooks_uninstall 0"
        }
      },
      {
        "box": {
          "id": "obj-b-hooks-out-json",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            520.0,
            790.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            100.0,
            404.0,
            22.0,
            22.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-msg-hooks-out-json",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            560.0,
            790.0,
            105.0,
            22.0
          ],
          "presentation": 0,
          "text": "hooks_uninstall 1"
        }
      },
      {
        "box": {
          "id": "obj-b-smoke",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            520.0,
            300.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            52.0,
            114.0,
            24.0,
            24.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-msg-smoke",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            560.0,
            300.0,
            105.0,
            22.0
          ],
          "presentation": 0,
          "text": "smoke_version"
        }
      },
      {
        "box": {
          "id": "obj-lbl-ait",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            400.0,
            200.0,
            150.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            8.0,
            220.0,
            20.0
          ],
          "text": "AIT_BIN (absolute path)"
        }
      },
      {
        "box": {
          "id": "obj-lbl-git2",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            400.0,
            260.0,
            150.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            40.0,
            220.0,
            20.0
          ],
          "text": "GIT_BIN (absolute path)"
        }
      },
      {
        "box": {
          "id": "obj-lbl-proj",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            400.0,
            270.0,
            280.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            70.0,
            360.0,
            20.0
          ],
          "text": "PROJECT_ROOT (optional; empty = process.cwd Live set folder)"
        }
      },
      {
        "box": {
          "id": "obj-msg-proj",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            400.0,
            290.0,
            300.0,
            22.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            84.0,
            354.0,
            22.0
          ],
          "text": ""
        }
      },
      {
        "box": {
          "id": "obj-pre-proj",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            280.0,
            280.0,
            65.0,
            22.0
          ],
          "presentation": 0,
          "text": "prepend set"
        }
      },
      {
        "box": {
          "id": "obj-lbl-eff",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            520.0,
            360.0,
            220.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            430.0,
            142.0,
            200.0,
            20.0
          ],
          "text": "Effective git -C root"
        }
      },
      {
        "box": {
          "id": "obj-msg-eff",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            520.0,
            390.0,
            320.0,
            22.0
          ],
          "presentation": 1,
          "presentation_rect": [
            430.0,
            160.0,
            330.0,
            22.0
          ],
          "text": "(git refresh fills this)"
        }
      },
      {
        "box": {
          "id": "obj-pre-eff",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            600.0,
            420.0,
            65.0,
            22.0
          ],
          "presentation": 0,
          "text": "prepend set"
        }
      },
      {
        "box": {
          "id": "obj-lbl-git-panel",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            30.0,
            520.0,
            200.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            132.0,
            220.0,
            20.0
          ],
          "text": "Git panel (subprocess git -C)"
        }
      },
      {
        "box": {
          "id": "obj-msg-git-branch",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            30.0,
            550.0,
            220.0,
            22.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            154.0,
            200.0,
            22.0
          ],
          "text": "(branch)"
        }
      },
      {
        "box": {
          "id": "obj-pre-git-br",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            100.0,
            575.0,
            65.0,
            22.0
          ],
          "presentation": 0,
          "text": "prepend set"
        }
      },
      {
        "box": {
          "id": "obj-msg-git-status",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            30.0,
            590.0,
            360.0,
            22.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            178.0,
            400.0,
            36.0
          ],
          "text": "(git status -sb)"
        }
      },
      {
        "box": {
          "id": "obj-pre-git-st",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            30.0,
            560.0,
            65.0,
            22.0
          ],
          "presentation": 0,
          "text": "prepend set"
        }
      },
      {
        "box": {
          "id": "obj-b-git-refresh",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            200.0,
            520.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            230.0,
            150.0,
            24.0,
            24.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-msg-git-refresh",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            240.0,
            520.0,
            80.0,
            22.0
          ],
          "presentation": 0,
          "text": "git_refresh"
        }
      },
      {
        "box": {
          "id": "obj-msg-umenu-clear",
          "maxclass": "message",
          "numinlets": 2,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            400.0,
            520.0,
            50.0,
            22.0
          ],
          "presentation": 0,
          "text": "clear"
        }
      },
      {
        "box": {
          "id": "obj-pre-git-append",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            460.0,
            550.0,
            85.0,
            22.0
          ],
          "presentation": 0,
          "text": "prepend append"
        }
      },
      {
        "box": {
          "id": "obj-umenu-br",
          "maxclass": "umenu",
          "numinlets": 1,
          "numoutlets": 3,
          "outlettype": [
            "int",
            "",
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            30.0,
            640.0,
            140.0,
            22.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            222.0,
            200.0,
            22.0
          ],
          "items": [
            "-",
            ",",
            "-"
          ]
        }
      },
      {
        "box": {
          "id": "obj-pre-git-co",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            200.0,
            680.0,
            115.0,
            22.0
          ],
          "presentation": 0,
          "text": "prepend git_checkout"
        }
      },
      {
        "box": {
          "id": "obj-b-git-co",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            200.0,
            640.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            230.0,
            220.0,
            24.0,
            24.0
          ]
        }
      },
      {
        "box": {
          "id": "obj-lbl-co",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            230.0,
            640.0,
            120.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            260.0,
            220.0,
            140.0,
            20.0
          ],
          "text": "Checkout selected"
        }
      },
      {
        "box": {
          "id": "obj-text-commit",
          "maxclass": "textedit",
          "numinlets": 1,
          "numoutlets": 4,
          "outlettype": [
            "",
            "int",
            "",
            ""
          ],
          "parameter_enable": 0,
          "patching_rect": [
            30.0,
            720.0,
            320.0,
            50.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            258.0,
            400.0,
            40.0
          ],
          "fontsize": 12.0
        }
      },
      {
        "box": {
          "id": "obj-zl-join-commit",
          "maxclass": "newobj",
          "numinlets": 2,
          "numoutlets": 2,
          "outlettype": [
            "",
            ""
          ],
          "patching_rect": [
            30.0,
            790.0,
            75.0,
            22.0
          ],
          "presentation": 0,
          "text": "zl join 512"
        }
      },
      {
        "box": {
          "id": "obj-pre-git-commit",
          "maxclass": "newobj",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            ""
          ],
          "patching_rect": [
            30.0,
            830.0,
            125.0,
            22.0
          ],
          "presentation": 0,
          "text": "prepend git_commit"
        }
      },
      {
        "box": {
          "id": "obj-b-commit",
          "maxclass": "button",
          "numinlets": 1,
          "numoutlets": 1,
          "outlettype": [
            "bang"
          ],
          "parameter_enable": 0,
          "patching_rect": [
            380.0,
            720.0,
            24.0,
            24.0
          ],
          "presentation": 1,
          "presentation_rect": [
            430.0,
            258.0,
            24.0,
            24.0
          ],
          "hint": "Stages everything in the repo (git add -A) then git commit -m. Nothing is committed if the message is empty or git errors."
        }
      },
      {
        "box": {
          "id": "obj-lbl-commit",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            30.0,
            700.0,
            280.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            16.0,
            238.0,
            420.0,
            20.0
          ],
          "text": "Commit message (uses git add -A \u2014 see tooltip on Commit button)"
        }
      },
      {
        "box": {
          "id": "obj-lbl-git-refresh",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            230.0,
            520.0,
            80.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            260.0,
            150.0,
            80.0,
            20.0
          ],
          "text": "Refresh"
        }
      },
      {
        "box": {
          "id": "obj-lbl-save",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            430.0,
            300.0,
            80.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            88.0,
            114.0,
            36.0,
            20.0
          ],
          "text": "Save"
        }
      },
      {
        "box": {
          "id": "obj-lbl-smoke",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            550.0,
            280.0,
            120.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            88.0,
            98.0,
            140.0,
            20.0
          ],
          "text": "Smoke: ait version"
        }
      },
      {
        "box": {
          "id": "obj-lbl-exit",
          "maxclass": "comment",
          "numinlets": 1,
          "numoutlets": 0,
          "patching_rect": [
            30.0,
            210.0,
            80.0,
            20.0
          ],
          "presentation": 1,
          "presentation_rect": [
            260.0,
            98.0,
            60.0,
            20.0
          ],
          "text": "exit"
        }
      }
    ],
    "lines": [
      {
        "patchline": {
          "destination": [
            "obj-msg-load",
            0
          ],
          "source": [
            "obj-loadbang",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-load",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-route",
            0
          ],
          "source": [
            "obj-node",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-num-exit",
            0
          ],
          "source": [
            "obj-route",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-preview",
            0
          ],
          "source": [
            "obj-route",
            1
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-pre-ait",
            0
          ],
          "source": [
            "obj-route",
            2
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-pre-git",
            0
          ],
          "source": [
            "obj-route",
            3
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-pre-proj",
            0
          ],
          "source": [
            "obj-route",
            4
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-pre-eff",
            0
          ],
          "source": [
            "obj-route",
            7
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-pre-git-br",
            0
          ],
          "source": [
            "obj-route",
            8
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-pre-git-st",
            0
          ],
          "source": [
            "obj-route",
            9
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-umenu-clear",
            0
          ],
          "source": [
            "obj-route",
            10
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-pre-git-append",
            0
          ],
          "source": [
            "obj-route",
            11
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-proj",
            0
          ],
          "source": [
            "obj-pre-proj",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-pre-toast",
            0
          ],
          "source": [
            "obj-route",
            6
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-toast",
            0
          ],
          "source": [
            "obj-pre-toast",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-eff",
            0
          ],
          "source": [
            "obj-pre-eff",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-git-branch",
            0
          ],
          "source": [
            "obj-pre-git-br",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-git-status",
            0
          ],
          "source": [
            "obj-pre-git-st",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-umenu-br",
            0
          ],
          "source": [
            "obj-msg-umenu-clear",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-umenu-br",
            0
          ],
          "source": [
            "obj-pre-git-append",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-ait",
            0
          ],
          "source": [
            "obj-pre-ait",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-git",
            0
          ],
          "source": [
            "obj-pre-git",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-t-save",
            0
          ],
          "source": [
            "obj-b-save",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-proj",
            0
          ],
          "source": [
            "obj-t-save",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-git",
            0
          ],
          "source": [
            "obj-t-save",
            1
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-ait",
            0
          ],
          "source": [
            "obj-t-save",
            2
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-join",
            0
          ],
          "source": [
            "obj-msg-ait",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-join",
            1
          ],
          "source": [
            "obj-msg-git",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-join",
            2
          ],
          "source": [
            "obj-msg-proj",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-git-refresh",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-git-refresh",
            0
          ],
          "source": [
            "obj-b-git-refresh",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-umenu-br",
            0
          ],
          "source": [
            "obj-b-git-co",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-pre-git-co",
            0
          ],
          "source": [
            "obj-umenu-br",
            1
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-pre-git-co",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-text-commit",
            0
          ],
          "source": [
            "obj-b-commit",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-zl-join-commit",
            0
          ],
          "source": [
            "obj-text-commit",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-pre-git-commit",
            0
          ],
          "source": [
            "obj-zl-join-commit",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-pre-git-commit",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-pre-save",
            0
          ],
          "source": [
            "obj-join",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-pre-save",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-smoke",
            0
          ],
          "source": [
            "obj-b-smoke",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-smoke",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-cmd-version",
            0
          ],
          "source": [
            "obj-b-version",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-cmd-version",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-cmd-version-json",
            0
          ],
          "source": [
            "obj-b-version-json",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-cmd-version-json",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-cmd-doctor",
            0
          ],
          "source": [
            "obj-b-doctor",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-cmd-doctor",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-cmd-doctor-json",
            0
          ],
          "source": [
            "obj-b-doctor-json",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-cmd-doctor-json",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-cmd-init",
            0
          ],
          "source": [
            "obj-b-init",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-cmd-init",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-hooks-in",
            0
          ],
          "source": [
            "obj-b-hooks-in",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-hooks-in",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-hooks-in-json",
            0
          ],
          "source": [
            "obj-b-hooks-in-json",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-hooks-in-json",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-hooks-out",
            0
          ],
          "source": [
            "obj-b-hooks-out",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-hooks-out",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-msg-hooks-out-json",
            0
          ],
          "source": [
            "obj-b-hooks-out-json",
            0
          ]
        }
      },
      {
        "patchline": {
          "destination": [
            "obj-node",
            0
          ],
          "source": [
            "obj-msg-hooks-out-json",
            0
          ]
        }
      }
    ],
    "dependency_cache": [],
    "autosave": 0
  }
}
