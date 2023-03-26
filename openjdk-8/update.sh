#!/bin/bash
VERSION="$1"
REVISION="ga"
BASEURI="http://hg.openjdk.java.net/jdk8u/jdk8u"
TARGETDIR="jdk8u${VERSION}-${REVISION}"
set -e

if [[ -z $VERSION ]]; then
    echo "Usage: $0 <version number>"
    exit 1
fi

# Main tarball
wget -nv -nc "${BASEURI}/archive/jdk8u${VERSION}-${REVISION}.tar.bz2"
sha256sum jdk8u${VERSION}-${REVISION}.tar.bz2
mkdir ${TARGETDIR}
tar -xf jdk8u${VERSION}-${REVISION}.tar.bz2 --strip-components=1 -C ${TARGETDIR}
rm -f jdk8u${VERSION}-${REVISION}.tar.bz2
echo

function fetch_and_emit()
{
    wget -nv -nc ${BASEURI}/$1/archive/jdk8u${VERSION}-${REVISION}.tar.bz2 -O $1-u${VERSION}-${REVISION}.tar.bz2
    sha256sum $1-u${VERSION}-${REVISION}.tar.bz2
    mkdir -p ${TARGETDIR}/$1
    tar -xf $1-u${VERSION}-${REVISION}.tar.bz2 --strip-components=1 -C ${TARGETDIR}/$1
    rm -f $1-u${VERSION}-${REVISION}.tar.bz2
    echo
}

# subproject tarballs
for i in corba hotspot jaxp jaxws jdk langtools nashorn; do
    fetch_and_emit $i
done

tar -cf jdk8u${VERSION}-${REVISION}-all.tar.gz ${TARGETDIR}
rm -rf ${TARGETDIR}
echo
sha256sum jdk8u${VERSION}-${REVISION}-all.tar.gz
