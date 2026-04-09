"use strict";

const maxAPI = require("max-api");
const { spawn } = require("child_process");
const fs = require("fs");
const path = require("path");

const SETTINGS_PATH = path.join(
	process.env.HOME || "",
	"Library/Application Support/ait/m4l-ait-control.json",
);

function defaultSettings() {
	return { AIT_BIN: "", GIT_BIN: "" };
}

function loadSettings() {
	try {
		const raw = fs.readFileSync(SETTINGS_PATH, "utf8");
		const parsed = JSON.parse(raw);
		return {
			AIT_BIN: typeof parsed.AIT_BIN === "string" ? parsed.AIT_BIN : "",
			GIT_BIN: typeof parsed.GIT_BIN === "string" ? parsed.GIT_BIN : "",
		};
	} catch {
		return defaultSettings();
	}
}

function saveSettings(settings) {
	fs.mkdirSync(path.dirname(SETTINGS_PATH), { recursive: true });
	fs.writeFileSync(
		SETTINGS_PATH,
		JSON.stringify(
			{
				AIT_BIN: settings.AIT_BIN || "",
				GIT_BIN: settings.GIT_BIN || "",
			},
			null,
			2,
		),
		"utf8",
	);
}

function emitPaths() {
	const s = loadSettings();
	maxAPI.outlet(0, "ait_path", s.AIT_BIN);
	maxAPI.outlet(0, "git_path", s.GIT_BIN);
	maxAPI.outlet(0, "status", "settings_loaded");
}

maxAPI.addHandler("load_settings", () => {
	emitPaths();
});

maxAPI.addHandler("save_settings", (...args) => {
	const ait = String(args[0] ?? "").trim();
	const git = String(args[1] ?? "").trim();
	saveSettings({ AIT_BIN: ait, GIT_BIN: git });
	maxAPI.outlet(0, "status", "settings_saved");
	emitPaths();
});

function truncate(s, max) {
	if (s.length <= max) {
		return s;
	}
	return s.slice(0, max) + "\n…(truncated)";
}

maxAPI.addHandler("smoke_version", () => {
	const s = loadSettings();
	const bin = (s.AIT_BIN || "").trim();
	if (!bin) {
		maxAPI.outlet(0, "exit", -1);
		maxAPI.outlet(
			0,
			"preview",
			"AIT_BIN is empty. Set an absolute path to the ait binary, Save, then retry.",
		);
		maxAPI.outlet(0, "status", "smoke_failed");
		return;
	}

	const child = spawn(bin, ["version"], {
		shell: false,
		env: process.env,
	});

	let stdout = "";
	let stderr = "";

	child.stdout.on("data", (chunk) => {
		stdout += chunk.toString();
	});
	child.stderr.on("data", (chunk) => {
		stderr += chunk.toString();
	});

	child.on("error", (err) => {
		maxAPI.outlet(0, "exit", -1);
		maxAPI.outlet(
			0,
			"preview",
			truncate(`spawn error: ${err.message}\n${stderr}`, 8000),
		);
		maxAPI.outlet(0, "status", "smoke_failed");
	});

	child.on("close", (code) => {
		const out = [
			`exit: ${code}`,
			"",
			"--- stdout ---",
			stdout.trimEnd() || "(empty)",
			"",
			"--- stderr ---",
			stderr.trimEnd() || "(empty)",
		].join("\n");
		maxAPI.outlet(0, "exit", code == null ? -1 : code);
		maxAPI.outlet(0, "preview", truncate(out, 8000));
		maxAPI.outlet(0, "status", "smoke_done");
	});
});
