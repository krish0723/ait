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
	return { AIT_BIN: "", GIT_BIN: "", PROJECT_PATH: "", PROJECT_ROOT: "" };
}

function loadSettings() {
	try {
		const raw = fs.readFileSync(SETTINGS_PATH, "utf8");
		const parsed = JSON.parse(raw);
		const proj =
			typeof parsed.PROJECT_PATH === "string"
				? parsed.PROJECT_PATH
				: typeof parsed.PROJECT_ROOT === "string"
					? parsed.PROJECT_ROOT
					: "";
		return {
			AIT_BIN: typeof parsed.AIT_BIN === "string" ? parsed.AIT_BIN : "",
			GIT_BIN: typeof parsed.GIT_BIN === "string" ? parsed.GIT_BIN : "",
			PROJECT_PATH: proj,
			PROJECT_ROOT: proj,
		};
	} catch {
		return defaultSettings();
	}
}

function saveSettings(settings) {
	fs.mkdirSync(path.dirname(SETTINGS_PATH), { recursive: true });
	const proj = settings.PROJECT_PATH || settings.PROJECT_ROOT || "";
	fs.writeFileSync(
		SETTINGS_PATH,
		JSON.stringify(
			{
				AIT_BIN: settings.AIT_BIN || "",
				GIT_BIN: settings.GIT_BIN || "",
				PROJECT_PATH: proj,
				PROJECT_ROOT: proj,
			},
			null,
			2,
		),
		"utf8",
	);
}

/** Git `git -C` target and `ait --path` root: manual path or process.cwd(). */
function resolveProjectRoot(settings) {
	const manual = String(
		settings.PROJECT_PATH || settings.PROJECT_ROOT || "",
	).trim();
	if (manual) {
		return path.isAbsolute(manual) ? manual : path.resolve(manual);
	}
	return process.cwd();
}

function emitPaths() {
	const s = loadSettings();
	const raw = s.PROJECT_PATH || s.PROJECT_ROOT || "";
	maxAPI.outlet(0, "ait_path", s.AIT_BIN);
	maxAPI.outlet(0, "git_path", s.GIT_BIN);
	maxAPI.outlet(0, "project_root", raw);
	maxAPI.outlet(0, "status", "settings_loaded");
}

maxAPI.addHandler("load_settings", () => {
	emitPaths();
});

maxAPI.addHandler("save_settings", (...args) => {
	const ait = String(args[0] ?? "").trim();
	const git = String(args[1] ?? "").trim();
	const proj = String(args[2] ?? "").trim();
	saveSettings({
		AIT_BIN: ait,
		GIT_BIN: git,
		PROJECT_PATH: proj,
		PROJECT_ROOT: proj,
	});
	maxAPI.outlet(0, "status", "settings_saved");
	maxAPI.outlet(0, "toast", "Settings saved.");
	emitPaths();
});

function truncate(s, max) {
	if (s.length <= max) {
		return s;
	}
	return s.slice(0, max) + "\n…(truncated)";
}

function asBool(v) {
	if (v === 0 || v === "0" || v === false || v === "false") {
		return false;
	}
	return Boolean(v);
}

function requireAitBin() {
	const s = loadSettings();
	const bin = (s.AIT_BIN || "").trim();
	if (!bin) {
		maxAPI.outlet(0, "exit", -1);
		maxAPI.outlet(
			0,
			"preview",
			"AIT_BIN is empty. Set an absolute path to the ait binary, Save, then retry.",
		);
		maxAPI.outlet(0, "toast", "AIT_BIN missing");
		maxAPI.outlet(0, "status", "ait_fail");
		return null;
	}
	return bin;
}

function emitRunOutcome(label, code, previewText, statusTag) {
	const exitVal = code == null ? -1 : code;
	maxAPI.outlet(0, "exit", exitVal);
	maxAPI.outlet(0, "preview", truncate(previewText, 8000));
	const ok = exitVal === 0;
	maxAPI.outlet(
		0,
		"toast",
		ok ? `${label}: ok (exit 0)` : `${label}: failed (exit ${exitVal})`,
	);
	maxAPI.outlet(0, "status", ok ? statusTag.ok : statusTag.fail);
}

function runAit(argv, opts, done) {
	const bin = requireAitBin();
	if (!bin) {
		return;
	}
	const cwd = opts && opts.cwd ? opts.cwd : undefined;
	const child = spawn(bin, argv, {
		shell: false,
		env: process.env,
		...(cwd ? { cwd } : {}),
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
		done(-1, "", `spawn error: ${err.message}`, stderr);
	});

	child.on("close", (code) => {
		done(code == null ? -1 : code, stdout, stderr, "");
	});
}

function formatDoctorFindings(obj) {
	const sevRank = { error: 0, warn: 1, info: 2 };
	const findings = Array.isArray(obj.findings) ? [...obj.findings] : [];
	findings.sort((a, b) => {
		const dr =
			(sevRank[a.severity] ?? 9) - (sevRank[b.severity] ?? 9);
		if (dr !== 0) {
			return dr;
		}
		const cr = String(a.code || "").localeCompare(String(b.code || ""));
		if (cr !== 0) {
			return cr;
		}
		return String(a.path || "").localeCompare(String(b.path || ""));
	});

	const lines = [];
	lines.push(
		`schema v${obj.schema_version} | ait ${obj.ait_version} | profile ${obj.profile} | preset ${obj.preset}`,
	);
	lines.push(`cwd: ${obj.cwd || "(unknown)"}`);
	lines.push("");
	if (findings.length === 0) {
		lines.push("(no findings)");
		return lines.join("\n");
	}
	for (const f of findings) {
		lines.push(`[${f.severity}] ${f.code}`);
		lines.push(`  ${f.message || ""}`);
		if (f.path) {
			lines.push(`  path: ${f.path}`);
		}
		if (f.hint) {
			lines.push(`  hint: ${f.hint}`);
		}
		if (f.doc_anchor) {
			lines.push(`  doc: ${f.doc_anchor}`);
		}
		lines.push("");
	}
	return lines.join("\n").trimEnd();
}

function runLabeledCommand(label, argv, repoRoot, transformPreview) {
	const statusTag = { ok: "ait_ok", fail: "ait_fail" };
	runAit(argv, { cwd: repoRoot }, (code, stdout, stderr, extraErr) => {
		const errPart = extraErr || stderr;
		let preview;
		if (transformPreview) {
			try {
				preview = transformPreview(stdout, stderr, code);
			} catch (e) {
				preview = `Error formatting output: ${e.message}\n\n--- stdout ---\n${stdout}\n--- stderr ---\n${stderr}`;
			}
		} else {
			preview = [
				`exit: ${code}`,
				"",
				"--- stdout ---",
				stdout.trimEnd() || "(empty)",
				"",
				"--- stderr ---",
				errPart.trimEnd() || "(empty)",
			].join("\n");
		}
		emitRunOutcome(label, code, preview, statusTag);
	});
}

maxAPI.addHandler("version", () => {
	const repoRoot = resolveProjectRoot(loadSettings());
	runLabeledCommand("version", ["version"], repoRoot, null);
});

maxAPI.addHandler("smoke_version", () => {
	const repoRoot = resolveProjectRoot(loadSettings());
	runLabeledCommand("version", ["version"], repoRoot, null);
});

maxAPI.addHandler("version_json", () => {
	const repoRoot = resolveProjectRoot(loadSettings());
	runLabeledCommand(
		"version --json",
		["version", "--json"],
		repoRoot,
		(stdout, stderr, code) => {
			const raw = stdout.trim();
			try {
				const o = JSON.parse(raw);
				return JSON.stringify(o, null, 2);
			} catch {
				return [
					`exit: ${code}`,
					"(stdout is not valid JSON)",
					"",
					raw || "(empty stdout)",
					"",
					"--- stderr ---",
					stderr.trimEnd() || "(empty)",
				].join("\n");
			}
		},
	);
});

maxAPI.addHandler("doctor", () => {
	const s = loadSettings();
	const repoRoot = resolveProjectRoot(s);
	const argv = ["doctor", "--path", repoRoot];
	runLabeledCommand("doctor", argv, repoRoot, null);
});

maxAPI.addHandler("doctor_json", () => {
	const s = loadSettings();
	const repoRoot = resolveProjectRoot(s);
	const argv = ["doctor", "--path", repoRoot, "--json"];
	runLabeledCommand("doctor --json", argv, repoRoot, (stdout, stderr, code) => {
		const raw = stdout.trim();
		try {
			const o = JSON.parse(raw);
			const body = formatDoctorFindings(o);
			const tail =
				stderr.trimEnd() ?
					`\n\n--- stderr ---\n${stderr.trimEnd()}`
				:	"";
			return `exit: ${code}\n\n${body}${tail}`;
		} catch (e) {
			return [
				`exit: ${code}`,
				`JSON parse error: ${e.message}`,
				"",
				"--- stdout ---",
				raw || "(empty)",
				"",
				"--- stderr ---",
				stderr.trimEnd() || "(empty)",
			].join("\n");
		}
	});
});

maxAPI.addHandler(
	"init_run",
	(daw, preset, dryRun, force, jsonOut) => {
		const s = loadSettings();
		const repoRoot = resolveProjectRoot(s);
		const argv = [
			"init",
			"--path",
			repoRoot,
			"--daw",
			String(daw ?? "ableton").trim() || "ableton",
			"--preset",
			String(preset ?? "samples-ignored").trim() || "samples-ignored",
		];
		if (asBool(dryRun)) {
			argv.push("--dry-run");
		}
		if (asBool(force)) {
			argv.push("--force");
		}
		if (asBool(jsonOut)) {
			argv.push("--json");
		}
		runLabeledCommand(
			"init",
			argv,
			repoRoot,
			asBool(jsonOut) ?
				(stdout, stderr, code) => {
					const raw = stdout.trim();
					try {
						const o = JSON.parse(raw);
						return JSON.stringify(o, null, 2);
					} catch {
						return [
							`exit: ${code}`,
							"(stdout is not valid JSON)",
							"",
							raw || "(empty)",
							"",
							"--- stderr ---",
							stderr.trimEnd() || "(empty)",
						].join("\n");
					}
				}
			:	null,
		);
	},
);

maxAPI.addHandler("hooks_install", (jsonOut) => {
	const s = loadSettings();
	const repoRoot = resolveProjectRoot(s);
	const argv = ["hooks", "install", "--path", repoRoot];
	if (asBool(jsonOut)) {
		argv.push("--json");
	}
	runLabeledCommand(
		"hooks install",
		argv,
		repoRoot,
		asBool(jsonOut) ?
			(stdout, stderr, code) => {
				const raw = stdout.trim();
				try {
					const o = JSON.parse(raw);
					return JSON.stringify(o, null, 2);
				} catch {
					return [
						`exit: ${code}`,
						"(stdout is not valid JSON)",
						"",
						raw || "(empty)",
						"",
						"--- stderr ---",
						stderr.trimEnd() || "(empty)",
					].join("\n");
				}
			}
		:	null,
	);
});

maxAPI.addHandler("hooks_uninstall", (jsonOut) => {
	const s = loadSettings();
	const repoRoot = resolveProjectRoot(s);
	const argv = ["hooks", "uninstall", "--path", repoRoot];
	if (asBool(jsonOut)) {
		argv.push("--json");
	}
	runLabeledCommand(
		"hooks uninstall",
		argv,
		repoRoot,
		asBool(jsonOut) ?
			(stdout, stderr, code) => {
				const raw = stdout.trim();
				try {
					const o = JSON.parse(raw);
					return JSON.stringify(o, null, 2);
				} catch {
					return [
						`exit: ${code}`,
						"(stdout is not valid JSON)",
						"",
						raw || "(empty)",
						"",
						"--- stderr ---",
						stderr.trimEnd() || "(empty)",
					].join("\n");
				}
			}
		:	null,
	);
});

function resolveGitBin(settings) {
	const g = (settings.GIT_BIN || "").trim();
	return g || "git";
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
