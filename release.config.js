const isWin = process.platform === "win32";
const isCI = process.env.CI == "true";
const isGHOk = !!(process.env.GH_TOKEN || process.env.GITHUB_TOKEN);
const changelogName = `docs/CHANGELOG.md`;

const ref = process.env.GITHUB_REF;
const branch = ref ? ref.split("/").pop() : "_none_";

console.log("OS:");
console.log("Branch name:", branch);
console.log("Is CI?:", isCI);
console.log("Is GH config ok?:", isGHOk);

if (!isGHOk) {
  console.warn("GITHUB_TOKEN is not set, GH release won't be created");
}

const plugs = [
  [
    "@semantic-release/commit-analyzer",
    {
      preset: "angular",
    },
  ],
  "@semantic-release/release-notes-generator",
  branch === "main"
    ? [
        "@semantic-release/changelog",
        {
          changelogFile: `${changelogName}`,
        },
      ]
    : undefined,
  [
    "@semantic-release/exec",
    {
      prepareCmd: isWin
        ? "task set-version VER=${nextRelease.version}; task build-all"
        : "task set-version VER=${nextRelease.version} && task build-all",
    },
  ],
  [
    "@semantic-release/git",
    {
      assets: ["docs"],
      message:
        "chore(release): ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}",
    },
  ],
  isGHOk
    ? [
        "@semantic-release/github",
        {
          assets: "_build/dist/**",
        },
      ]
    : undefined,
];

const cfg = {
  branches: [
    "+([0-9])?(.{+([0-9]),x}).x",
    "main",
    "next",
    {
      name: "beta",
      prerelease: "beta",
    },
    {
      name: "alpha",
      prerelease: "alpha",
    },
  ],
  tagFormat: "${version}",
  plugins: plugs.filter((x) => x !== null && x !== undefined),
};

module.exports = cfg;
