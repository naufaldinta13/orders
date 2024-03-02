#!/usr/bin/env bash

set -e

library=$1
lastVersion=''
lastTag=''
libraryPath=''

gitBranch="$(git symbolic-ref --short -q HEAD 2>/dev/null)" ||
gitBranch="(unnamed branch)"     # detached HEAD
gitBranch=${gitBranch##refs/heads/}

gitHash="$(git rev-parse --short HEAD 2>/dev/null)"

if [[ "$library" && "$library" != "." ]]; then
    libraryPath="$library/"
fi

function getLastVersion() {
    lastTag=`git describe --match "${libraryPath}v[0-9].[0-9].[0-9]" --tags 2> /dev/null` || true

    if [[ -z "$lastTag" ]]; then
        lastTag="${libraryPath}v0.0.0"
    fi

    lastTag=${lastTag/"$libraryPath"/""}
    lastTag=${lastTag:1}

    parts=(${lastTag//./ })
    partsPatch=(${parts[2]//-/ })

    if [[ ${partsPatch[0]} = "0"  ]]; then
        patch=${partsPatch[0]} | bc
    else
        patch=${partsPatch[0]}
    fi

    major=${parts[0]}
    minor=${parts[1]}

    lastVersion=$libraryPath"v$major.$minor.$patch"
}

## ini untuk checkout -b dari branch develop
branch=(${gitBranch//-/ })

if [[ "$branch" = "develop" ]]; then
    major=0
    minor=0
    patch="2-develop-$gitHash"
elif [[ "$branch" = "master" ]]; then
    getLastVersion

    case "$2" in
        'current')
            printf "%s\n" "$lastVersion"
            exit 0;;
        'patch')
            patch=$((patch + 1))
            ;;
        'minor')
            minor=$((minor + 1))
            patch=0
            ;;
        'major')
            major=$((major + 1))
            minor=0
            patch=0
            ;;
        *)
            printf "%s\n" "$lastVersion"
            exit 0
    esac
else
    printf "%s\n" "You are not in the branch that allowed to release!"
    exit 0
fi

printf "%s%s.%s.%s\n" "$libraryPath" "v$major" "$minor" "$patch"