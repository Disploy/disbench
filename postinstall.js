#!/usr/bin/env node

const fs = require("node:fs");
const path = require("node:path");
const child_process = require("node:child_process");

const ARCH_MAPPING = {
  ia32: "386",
  x64: "amd64",
  arm: "arm",
  arm64: "arm64",
};

const PLATFORM_MAPPING = {
  darwin: "darwin",
  linux: "linux",
  win32: "windows",
  freebsd: "freebsd",
};

const platform = PLATFORM_MAPPING[process.platform];
const arch = ARCH_MAPPING[process.arch];

const folder = fs
  .readdirSync("dist")
  .find((file) => file.includes(platform) && file.includes(arch));

if (!folder) {
  console.error(`No binary found for platform ${platform} and arch ${arch}`);
  process.exit(1);
}

const binP = path.join(
  "dist",
  folder,
  `disbench${platform === "windows" ? ".exe" : ""}`
);

fs.writeFileSync(
  "disbench.js",
  [
    "const { spawn } = require('node:child_process');",
    "const { join } = require('node:path');",
    `spawn(join(__dirname, ${binP}), process.argv.slice(2), { stdio: 'inherit' })`,
    ".on('exit', (code) => process.exit(code));",
  ].join("\n")
);
