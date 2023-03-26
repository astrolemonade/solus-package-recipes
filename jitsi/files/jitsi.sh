#!/bin/sh

SCDIR=/usr/share/jitsi
LIBPATH="$SCDIR/lib"
CLASSPATH="${LIBPATH}/felix.jar:${SCDIR}/sc-bundles/sc-launcher.jar:${SCDIR}/sc-bundles/util.jar:${SCDIR}/sc-bundles/dnsjava.jar:${LIBPATH}"
FELIX_CONFIG="$LIBPATH/felix.client.run.properties"
LOG_CONFIG="$LIBPATH/logging.properties"

if [ -z "$JAVA_HOME" ]; then
    export JAVA_HOME=/usr/lib64/openjdk-8
fi

COMMAND=($JAVA_HOME/bin/java -classpath "$CLASSPATH" -Djna.library.path="$LIBPATH/native" -Dfelix.config.properties="file:$FELIX_CONFIG" -Djava.util.logging.config.file="$LOG_CONFIG" net.java.sip.communicator.launcher.SIPCommunicator)

export LD_LIBRARY_PATH="$LD_LIBRARY_PATH:$LIBPATH/native"

cd "$SCDIR"
exec "${COMMAND[@]}" "$@"
