{
	"patcher" : 	{
		"fileversion" : 1,
		"appversion" : 		{
			"major" : 8,
			"minor" : 5,
			"revision" : 0,
			"architecture" : "x64"
		},
		"classnamespace" : "box",
		"rect" : [ 100.0, 100.0, 720.0, 420.0 ],
		"bglocked" : 0,
		"openinpresentation" : 1,
		"default_fontsize" : 12.0,
		"default_fontface" : 0,
		"default_fontname" : "Arial",
		"gridsize" : [ 15.0, 15.0 ],
		"boxes" : [ 			{
				"box" : 				{
					"id" : "obj-loadbang",
					"maxclass" : "loadbang",
					"numinlets" : 1,
					"numoutlets" : 1,
					"outlettype" : [ "bang" ],
					"parameter_enable" : 0,
					"patching_rect" : [ 30.0, 30.0, 40.0, 22.0 ],
					"presentation" : 0
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-msg-load",
					"maxclass" : "message",
					"numinlets" : 2,
					"numoutlets" : 1,
					"outlettype" : [ "" ],
					"parameter_enable" : 0,
					"patching_rect" : [ 30.0, 70.0, 95.0, 22.0 ],
					"presentation" : 0,
					"text" : "load_settings"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-node",
					"maxclass" : "newobj",
					"numinlets" : 1,
					"numoutlets" : 2,
					"outlettype" : [ "", "" ],
					"patching_rect" : [ 30.0, 120.0, 320.0, 22.0 ],
					"presentation" : 0,
					"text" : "node.script ait-control.node.js @autostart 1"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-route",
					"maxclass" : "newobj",
					"numinlets" : 1,
					"numoutlets" : 5,
					"outlettype" : [ "", "", "", "", "" ],
					"patching_rect" : [ 30.0, 180.0, 220.0, 22.0 ],
					"presentation" : 0,
					"text" : "route exit preview ait_path git_path status"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-num-exit",
					"maxclass" : "number",
					"numinlets" : 1,
					"numoutlets" : 2,
					"outlettype" : [ "", "bang" ],
					"parameter_enable" : 0,
					"patching_rect" : [ 30.0, 230.0, 50.0, 22.0 ],
					"presentation" : 1,
					"presentation_rect" : [ 320.0, 86.0, 50.0, 22.0 ]
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-msg-preview",
					"maxclass" : "message",
					"numinlets" : 2,
					"numoutlets" : 1,
					"outlettype" : [ "" ],
					"parameter_enable" : 0,
					"patching_rect" : [ 100.0, 230.0, 450.0, 22.0 ],
					"presentation" : 1,
					"presentation_rect" : [ 16.0, 118.0, 400.0, 80.0 ],
					"text" : "(stdout/stderr preview)"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-pre-ait",
					"maxclass" : "newobj",
					"numinlets" : 1,
					"numoutlets" : 1,
					"outlettype" : [ "" ],
					"patching_rect" : [ 280.0, 220.0, 65.0, 22.0 ],
					"presentation" : 0,
					"text" : "prepend set"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-pre-git",
					"maxclass" : "newobj",
					"numinlets" : 1,
					"numoutlets" : 1,
					"outlettype" : [ "" ],
					"patching_rect" : [ 280.0, 250.0, 65.0, 22.0 ],
					"presentation" : 0,
					"text" : "prepend set"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-msg-ait",
					"maxclass" : "message",
					"numinlets" : 2,
					"numoutlets" : 1,
					"outlettype" : [ "" ],
					"parameter_enable" : 0,
					"patching_rect" : [ 400.0, 220.0, 300.0, 22.0 ],
					"presentation" : 1,
					"presentation_rect" : [ 16.0, 28.0, 354.0, 22.0 ],
					"text" : "/usr/local/bin/ait"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-msg-git",
					"maxclass" : "message",
					"numinlets" : 2,
					"numoutlets" : 1,
					"outlettype" : [ "" ],
					"parameter_enable" : 0,
					"patching_rect" : [ 400.0, 250.0, 300.0, 22.0 ],
					"presentation" : 1,
					"presentation_rect" : [ 16.0, 56.0, 354.0, 22.0 ],
					"text" : "/usr/bin/git"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-b-save",
					"maxclass" : "button",
					"numinlets" : 1,
					"numoutlets" : 1,
					"outlettype" : [ "bang" ],
					"parameter_enable" : 0,
					"patching_rect" : [ 400.0, 300.0, 24.0, 24.0 ],
					"presentation" : 1,
					"presentation_rect" : [ 16.0, 86.0, 24.0, 24.0 ]
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-t-save",
					"maxclass" : "newobj",
					"numinlets" : 1,
					"numoutlets" : 2,
					"outlettype" : [ "bang", "bang" ],
					"patching_rect" : [ 400.0, 340.0, 42.0, 22.0 ],
					"presentation" : 0,
					"text" : "t b b"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-join",
					"maxclass" : "newobj",
					"numinlets" : 2,
					"numoutlets" : 1,
					"outlettype" : [ "" ],
					"patching_rect" : [ 400.0, 390.0, 95.0, 22.0 ],
					"presentation" : 0,
					"text" : "join 2"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-pre-save",
					"maxclass" : "newobj",
					"numinlets" : 1,
					"numoutlets" : 1,
					"outlettype" : [ "" ],
					"patching_rect" : [ 400.0, 430.0, 125.0, 22.0 ],
					"presentation" : 0,
					"text" : "prepend save_settings"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-b-smoke",
					"maxclass" : "button",
					"numinlets" : 1,
					"numoutlets" : 1,
					"outlettype" : [ "bang" ],
					"parameter_enable" : 0,
					"patching_rect" : [ 520.0, 300.0, 24.0, 24.0 ],
					"presentation" : 1,
					"presentation_rect" : [ 52.0, 86.0, 24.0, 24.0 ]
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-msg-smoke",
					"maxclass" : "message",
					"numinlets" : 2,
					"numoutlets" : 1,
					"outlettype" : [ "" ],
					"parameter_enable" : 0,
					"patching_rect" : [ 560.0, 300.0, 105.0, 22.0 ],
					"presentation" : 0,
					"text" : "smoke_version"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-lbl-ait",
					"maxclass" : "comment",
					"numinlets" : 1,
					"numoutlets" : 0,
					"patching_rect" : [ 400.0, 200.0, 150.0, 20.0 ],
					"presentation" : 1,
					"presentation_rect" : [ 16.0, 8.0, 220.0, 20.0 ],
					"text" : "AIT_BIN (absolute path)"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-lbl-git2",
					"maxclass" : "comment",
					"numinlets" : 1,
					"numoutlets" : 0,
					"patching_rect" : [ 400.0, 260.0, 150.0, 20.0 ],
					"presentation" : 1,
					"presentation_rect" : [ 16.0, 40.0, 220.0, 20.0 ],
					"text" : "GIT_BIN (absolute path)"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-lbl-save",
					"maxclass" : "comment",
					"numinlets" : 1,
					"numoutlets" : 0,
					"patching_rect" : [ 430.0, 300.0, 80.0, 20.0 ],
					"presentation" : 1,
					"presentation_rect" : [ 88.0, 86.0, 36.0, 20.0 ],
					"text" : "Save"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-lbl-smoke",
					"maxclass" : "comment",
					"numinlets" : 1,
					"numoutlets" : 0,
					"patching_rect" : [ 550.0, 280.0, 120.0, 20.0 ],
					"presentation" : 1,
					"presentation_rect" : [ 88.0, 70.0, 140.0, 20.0 ],
					"text" : "Smoke: ait version"
				}

			}
, 			{
				"box" : 				{
					"id" : "obj-lbl-exit",
					"maxclass" : "comment",
					"numinlets" : 1,
					"numoutlets" : 0,
					"patching_rect" : [ 30.0, 210.0, 80.0, 20.0 ],
					"presentation" : 1,
					"presentation_rect" : [ 260.0, 70.0, 60.0, 20.0 ],
					"text" : "exit"
				}

			}
 ],
		"lines" : [ 			{
				"patchline" : 				{
					"destination" : [ "obj-msg-load", 0 ],
					"source" : [ "obj-loadbang", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-node", 0 ],
					"source" : [ "obj-msg-load", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-route", 0 ],
					"source" : [ "obj-node", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-num-exit", 0 ],
					"source" : [ "obj-route", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-msg-preview", 0 ],
					"source" : [ "obj-route", 1 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-pre-ait", 0 ],
					"source" : [ "obj-route", 2 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-pre-git", 0 ],
					"source" : [ "obj-route", 3 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-msg-ait", 0 ],
					"source" : [ "obj-pre-ait", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-msg-git", 0 ],
					"source" : [ "obj-pre-git", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-t-save", 0 ],
					"source" : [ "obj-b-save", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-msg-git", 0 ],
					"source" : [ "obj-t-save", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-msg-ait", 0 ],
					"source" : [ "obj-t-save", 1 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-join", 0 ],
					"source" : [ "obj-msg-ait", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-join", 1 ],
					"source" : [ "obj-msg-git", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-pre-save", 0 ],
					"source" : [ "obj-join", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-node", 0 ],
					"source" : [ "obj-pre-save", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-msg-smoke", 0 ],
					"source" : [ "obj-b-smoke", 0 ]
				}

			}
, 			{
				"patchline" : 				{
					"destination" : [ "obj-node", 0 ],
					"source" : [ "obj-msg-smoke", 0 ]
				}

			}
 ],
		"dependency_cache" : [ 		],
		"autosave" : 0
	}

}
