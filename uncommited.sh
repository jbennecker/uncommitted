#!/usr/bin/env sh

for d in ./*/ ; do
    (
        cd "$d"
        if [ -d .git ] && [ -n "$(git status --porcelain)" ]; then
            echo "Uncommited changes in $d"
        fi
    )
done
