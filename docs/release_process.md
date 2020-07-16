# Release Process

This document describes the release process for this project. Releases are published via GitHub releases.

All releases will have a [semantic version](https://semver.org/) associated with them.
This project will be versioned at 1.0.0 once the
[Service Mesh Interface(SMI) specification](https://github.com/servicemeshinterface/smi-spec)
is at 1.0.0 and all APIs defined in this project are at `v1`.
All releases until then will be marked `v0.x.0` where `x` is incremented with each release.

To perform a release of the smi-sdk-go project:

1. Create and push a git tag for release version
Start by creating and pushing a git tag in the form: `v0.x.0`.
```console
$ git tag -a v0.x.0 -m "version 0.x.0"
$ git push origin v0.x.0
```

2. Create a GitHub release
Visit the [releases page](https://github.com/servicemeshinterface/smi-sdk-go/releases)
to `Draft a new release` using the tag you just created and pushed.

3. Add release notes
Run script `scripts/release-notes.sh` to generate release notes. Copy and paste output
to release description section. Fill in smi spec version compatability information.

4. Click the "Publish release" button
