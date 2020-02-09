#!/usr/bin/env bash

set -e

PROJECT="sswarm"
SCRIPTPATH="$(dirname "$(readlink -f "${0}")")"
DOCKERFILE="./ci/build.Dockerfile"
DOCKERTAG="${PROJECT}-build:tmp"
BINARY="./${PROJECT,,}"
SILENT="${1}" # ./build -s

spinner()
{
    # http://fitnr.com/showing-a-bash-spinner.html
    local pid=$1
    local delay=0.15
    local spinstr='|/-\'
    while [ "$(ps a | awk '{print $1}' | grep $pid)" ]; do
        local temp=${spinstr#?}
        printf " [%c]  " "$spinstr"
        local spinstr=$temp${spinstr%"$temp"}
        sleep $delay
        printf "\b\b\b\b\b\b"
    done
    printf "    \b\b\b\b"
}

build()
{
    pushd "${SCRIPTPATH}/.." > /dev/null 2>&1
    # rm -f "${BINARY}"
    echo
    echo "Using temporary docker build environment..."
    echo
    docker build -t "${DOCKERTAG}" -f "${DOCKERFILE}" .
    docker run --rm "${DOCKERTAG}" > "${BINARY}"
    docker rmi "${DOCKERTAG}"
    chmod a+rx "${BINARY}"
    echo
    echo "Removed temporary build environment"
    echo
    popd > /dev/null 2>&1
    echo -n "Build time:"
}

echo "Building..."
if [ "${SILENT}" != "-s" ] ; then
    (
        time build
        echo
        echo "${PROJECT^} project executable: ${BINARY}"
        echo
    ) &
else
    (build) > /dev/null 2>&1 &
fi
spinner $!
echo "Done!"
