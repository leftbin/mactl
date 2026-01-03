# Automated Release Pipeline with GoReleaser and Homebrew Casks

**Date**: January 3, 2026

## Summary

Implemented a fully automated release pipeline for mactl using GoReleaser v2 and GitHub Actions. Running `make release` now auto-calculates the next semantic version, creates a git tag, and triggers a workflow that builds darwin binaries and auto-updates the Homebrew Cask. Since mactl is a macOS-only tool, builds are limited to darwin (amd64 + arm64).

## Problem Statement

The release process for mactl required significant manual intervention and was error-prone.

### Pain Points

- **Manual version specification**: Every release required editing the Makefile with a hardcoded `version?=v0.0.17` value
- **Local tooling dependency**: Required `gh` CLI and authentication on developer machines
- **Multi-step manual process**: Build binaries, create checksums, create GitHub release, attach assets
- **Manual Homebrew updates**: Formula had to be manually edited after each release with new version and checksums
- **macOS Gatekeeper warnings**: Users saw security warnings when running unsigned binaries

## Solution

A three-part release automation system following the patterns established with gitr, karayaml, and project-planton:

### 1. GoReleaser Configuration

Created `.goreleaser.yaml` with GoReleaser v2 features:
- darwin-only builds (amd64 + arm64) - no linux/windows since mactl is macOS-specific
- tar.gz archives with SHA256 checksums
- Homebrew Cask integration with auto-update
- macOS Gatekeeper fix via post-install hook that removes quarantine attribute

### 2. Auto-Version Bumping

Updated `Makefile` with semver logic:
- Automatically calculates next version from latest git tag
- Supports `patch`, `minor`, `major` bump types (default: patch)
- Preview capability with `make next-version`

### 3. GitHub Actions Workflow

Created `.github/workflows/release.yml`:
- Triggers on `v*` tag pushes
- Uses GoReleaser v2 action
- Creates GitHub releases with auto-generated notes
- Pushes Cask updates to homebrew-tap repository

## Implementation Details

### GoReleaser Configuration

```yaml
version: 2
project_name: mactl

builds:
  - id: mactl
    binary: mactl
    ldflags:
      - -s -w -X github.com/leftbin/mactl/internal/version.Version={{.Version}}
    goos:
      - darwin  # macOS only
    goarch:
      - amd64
      - arm64

homebrew_casks:
  - name: mactl
    repository:
      owner: leftbin
      name: homebrew-tap
      token: "{{ .Env.HOMEBREW_TAP_GITHUB_TOKEN }}"
    hooks:
      post:
        install: |
          if OS.mac?
            system_command "/usr/bin/xattr", args: ["-dr", "com.apple.quarantine", "#{staged_path}/mactl"]
          end
```

### Makefile Release Targets

| Command | Description |
|---------|-------------|
| `make release` | Auto-bump patch version, tag & push |
| `make release bump=minor` | Bump minor version |
| `make release bump=major` | Bump major version |
| `make next-version` | Preview what the next version would be |
| `make snapshot` | Local GoReleaser test build |

### Version Bump Behavior

| Command | Current | Result |
|---------|---------|--------|
| `make release` | v0.0.21 | v0.0.22 |
| `make release bump=minor` | v0.0.21 | v0.1.0 |
| `make release bump=major` | v0.0.21 | v1.0.0 |

### Homebrew Tap Migration

Created `tap_migrations.json` in homebrew-tap for seamless migration from Formula to Cask:

```json
{
  "mactl": "leftbin/tap/mactl"
}
```

This ensures existing users running `brew upgrade` automatically migrate to the new Cask.

## Benefits

### For Maintainers

- **Zero-friction releases**: `make release` with no arguments needed
- **No local tooling**: No `gh` CLI or authentication required
- **Version preview**: `make next-version` shows what would be released
- **Consistent versioning**: No typos or version confusion
- **Reproducible builds**: GoReleaser ensures consistent artifacts

### For Users

- **No security warnings**: macOS Gatekeeper quarantine auto-removed
- **Easy installation**: `brew install --cask leftbin/tap/mactl`
- **GitHub Releases**: Direct downloads without extra authentication
- **Automatic updates**: Homebrew Cask updated with each release

### For CI/CD

- **Parallel builds**: darwin amd64 + arm64 via GoReleaser
- **Automatic checksums**: SHA256 checksums in release artifacts
- **Idempotent**: Homebrew update skips if no changes needed

## Impact

### Release Process Changes

| Before | After |
|--------|-------|
| Hardcoded version in Makefile | Auto-calculated from git tags |
| Manual `gh release create` | Automatic via GitHub Actions |
| Manual formula update | Automatic Cask push |
| Gatekeeper warnings | No warnings (quarantine removed) |
| Multi-step process | Single `make release` command |

### Files Changed

| File | Action | Description |
|------|--------|-------------|
| `.goreleaser.yaml` | Created | GoReleaser v2 configuration (darwin only) |
| `.github/workflows/release.yml` | Created | GitHub Actions workflow |
| `Makefile` | Updated | Auto-versioning with semver bump |
| `homebrew-tap/tap_migrations.json` | Created | Formula → Cask migration |

## Migration Steps

Before the first automated release:

1. **Add GitHub Secret**: Add `HOMEBREW_TAP_GITHUB_TOKEN` to the mactl repository secrets (Settings → Secrets → Actions). This needs to be a Personal Access Token with `repo` scope.

2. **Push tap_migrations.json**: Commit and push to leftbin/homebrew-tap

3. **Run first release**:
   ```bash
   make release  # Creates v0.0.22
   ```

4. **Delete old Formula**: After successful release, remove `Formula/mactl.rb` from homebrew-tap

## Related Work

- Follows patterns from gitr, karayaml, and project-planton release automation
- Part of standardizing release processes across Leftbin tools
- Uses same GoReleaser v2 and Homebrew Cask approach

---

**Status**: ✅ Production Ready (pending `HOMEBREW_TAP_GITHUB_TOKEN` secret setup)
**Timeline**: Single session implementation

