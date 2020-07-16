RELEASE=${RELEASE:-$2}
PREVIOUS_RELEASE=${PREVIOUS_RELEASE:-$1}

## Ensure Correct Usage
if [[ -z "${PREVIOUS_RELEASE}" || -z "${RELEASE}" ]]; then
  echo Usage:
  echo ./scripts/release-notes.sh v0.4.0 v0.5.0
  echo or
  echo PREVIOUS_RELEASE=v0.4.0
  echo RELEASE=v0.5.0
  echo ./scripts/release-notes.sh
  exit 1
fi

## validate git tags
for tag in $RELEASE $PREVIOUS_RELEASE; do
  OK=$(git tag -l ${tag} | wc -l)
  if [[ "$OK" == "0" ]]; then
    echo ${tag} is not a valid release version
    exit 1
  fi
done

## Generate CHANGELOG from git log
CHANGELOG=$(git log --no-merges --pretty=format:'- %s %H (%aN)' ${PREVIOUS_RELEASE}..${RELEASE})
if [[ ! $? -eq 0 ]]; then
  echo "Error creating changelog"
  echo "try running \`git log --no-merges --pretty=format:'- %s %H (%aN)' ${PREVIOUS_RELEASE}..${RELEASE}\`"
  exit 1
fi


## Print release notes to stdout
cat <<EOF
## SMI Spec Version Compatibility
- make notes about smi spec version support here

## Changelog
${CHANGELOG}
EOF

