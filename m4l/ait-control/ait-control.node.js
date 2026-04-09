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
	return { AIT_BIN: "", GIT_BIN: "", PROJECT_ROOT: "" };
}

function loadSettings() {
	try {
		const raw = fs.readFileSync(SETTINGS_PATH, "utf8");
		const parsed = JSON.parse(raw);
		return {
			AIT_BIN: typeof parsed.AIT_BIN === "string" ? parsed.AIT_BIN : "",
			GIT_BIN: typeof parsed.GIT_BIN === "string" ? parsed.GIT_BIN : "",
			PROJECT_ROOT:
				typeof parsed.PROJECT_ROOT === "string" ? parsed.PROJECT_ROOT : "",
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
				PROJECT_ROOT: settings.PROJECT_ROOT || "",
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
	maxAPI.outlet(0, "project_root", s.PROJECT_ROOT);
	maxAPI.outlet(0, "status", "settings_loaded");
}

maxAPI.addHandler("load_settings", () => {
	emitPaths();
});

maxAPI.addHandler("save_settings", (...args) => {
	const ait = String(args[0] ?? "").trim();
	const git = String(args[1] ?? "").trim();
	const proj = String(args[2] ?? "").trim();
	saveSettings({ AIT_BIN: ait, GIT_BIN: git, PROJECT_ROOT: proj });
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

function resolveGitBin(settings) {
	const g = (settings.GIT_BIN || "").trim();
	return g || "git";
}

function resolveProjectRoot(settings) {
	const manual = (settings.PROJECT_ROOT || "").trim();
	if (manual) {
		return manual;
	}
	return process.cwd();
}

function runGitAsync(gitBin, projectRoot, gitArgs) {
	return new Promise((resolve) => {
		const child = spawn(gitBin, ["-C", projectRoot, ...gitArgs], {
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
			resolve({
				code: -1,
				stdout,
				stderr: (stderr + String(err.message)).trimEnd(),
				spawnError: true,
			});
		});
		child.on("close", (code) => {
			resolve({
				code: code == null ? -1 : code,
				stdout,
				stderr,
				spawnError: false,
			});
		});
	});
}

function emitGitFailure(code, stderr, stdout) {
	const detail = [
		`exit: ${code}`,
		"",
		"--- stderr ---",
		(stderr || "").trimEnd() || "(empty)",
		"",
		"--- stdout ---",
		(stdout || "").trimEnd() || "(empty)",
	].join("\n");
	maxAPI.outlet(0, "exit", code);
	maxAPI.outlet(0, "preview", truncate(detail, 8000));
	maxAPI.outlet(0, "status", "git_error");
}

maxAPI.addHandler("git_refresh", async () => {
	const s = loadSettings();
	const gitBin = resolveGitBin(s);
	const root = resolveProjectRoot(s);

	maxAPI.outlet(0, "git_effective_root", root);

	const inside = await runGitAsync(gitBin, root, [
		"rev-parse",
		"--is-inside-work-tree",
	]);
	if (inside.code !== 0 || inside.stdout.trim() !== "true") {
		emitGitFailure(
			inside.code,
			inside.stderr ||
				inside.stdout ||
				"not a git repository (or git could not read the directory)",
			inside.stdout,
		);
		return;
	}

	const [headRes, stRes, brRes] = await Promise.all([
		runGitAsync(gitBin, root, ["rev-parse", "--abbrev-ref", "HEAD"]),
		runGitAsync(gitBin, root, ["status", "-sb"]),
		runGitAsync(gitBin, root, ["branch", "--list", "--format=%(refname:short)"]),
	]);

	if (headRes.code !== 0) {
		emitGitFailure(headRes.code, headRes.stderr, headRes.stdout);
		return;
	}
	if (stRes.code !== 0) {
		emitGitFailure(stRes.code, stRes.stderr, stRes.stdout);
		return;
	}
	if (brRes.code !== 0) {
		emitGitFailure(brRes.code, brRes.stderr, brRes.stdout);
		return;
	}

	const branch = headRes.stdout.trim();
	const statusText = stRes.stdout.trimEnd();
	const branches = brRes.stdout
		.split("\n")
		.map((l) => l.trim())
		.filter(Boolean);

	maxAPI.outlet(0, "git_clear_branches");
	for (const b of branches) {
		maxAPI.outlet(0, "git_append_branch", b);
	}
	maxAPI.outlet(0, "git_branch", branch);
	maxAPI.outlet(0, "git_status", truncate(statusText, 4000));
	maxAPI.outlet(0, "preview", truncate(`git -C ${root}\nOn ${branch}\n${statusText}`, 8000));
	maxAPI.outlet(0, "exit", 0);
	maxAPI.outlet(0, "status", "git_refresh_done");
});

maxAPI.addHandler("git_checkout", (...args) => {
	const branch = args.join(" ").trim();
	const run = async () => {
		const s = loadSettings();
		const gitBin = resolveGitBin(s);
		const root = resolveProjectRoot(s);
		maxAPI.outlet(0, "git_effective_root", root);
		if (!branch) {
			emitGitFailure(-1, "No branch selected.", "");
			return;
		}
		const res = await runGitAsync(gitBin, root, ["checkout", branch]);
		if (res.code !== 0) {
			emitGitFailure(res.code, res.stderr, res.stdout);
			return;
		}
		maxAPI.outlet(
			0,
			"preview",
			truncate(
				(res.stdout.trimEnd() || `Switched to branch '${branch}'.`) +
					"\n\nReopen your .als from disk after a branch switch so Live reloads file paths.",
				8000,
			),
		);
		maxAPI.outlet(0, "exit", 0);
		maxAPI.outlet(0, "status", "git_checkout_done");
	};
	void run();
});

maxAPI.addHandler("git_commit", (...args) => {
	const message = args.join(" ").trim();
	const run = async () => {
		const s = loadSettings();
		const gitBin = resolveGitBin(s);
		const root = resolveProjectRoot(s);
		maxAPI.outlet(0, "git_effective_root", root);
		if (!message) {
			emitGitFailure(-1, "Commit message is empty.", "");
			return;
		}
		const addRes = await runGitAsync(gitBin, root, ["add", "-A"]);
		if (addRes.code !== 0) {
			emitGitFailure(addRes.code, addRes.stderr, addRes.stdout);
			return;
		}
		const commitRes = await runGitAsync(gitBin, root, [
			"commit",
			"-m",
			message,
		]);
		if (commitRes.code !== 0) {
			emitGitFailure(commitRes.code, commitRes.stderr, commitRes.stdout);
			return;
		}
		maxAPI.outlet(
			0,
			"preview",
			truncate(commitRes.stdout.trimEnd() || "commit created", 8000),
		);
		maxAPI.outlet(0, "exit", 0);
		maxAPI.outlet(0, "status", "git_commit_done");
	};
	void run();
});
