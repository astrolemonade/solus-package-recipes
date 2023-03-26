#!/bin/bash

GTDIR="./BUILDROOT"

if [[ -z "${IMAGE_TYPE}" ]]; then
    IMAGE_TYPE="unstable"
fi

if [[ "${IMAGE_TYPE}" == "unstable" ]]; then
    REPO="https://mirrors.rit.edu/solus/packages/unstable/eopkg-index.xml.xz"
    IMAGE_NAME="unstable-x86_64.img"
elif [[ "${IMAGE_TYPE}" == "stable" ]]; then
    REPO="https://mirrors.rit.edu/solus/packages/shannon/eopkg-index.xml.xz"
    IMAGE_NAME="main-x86_64.img"
else
    echo "Unknown IMAGE_TYPE: ${IMAGE_TYPE} (need stable or unstable)"
    exit 1
fi

# We cache packages and archives
PKGDIR="/var/lib/solbuild/packages"
ARCHIVEDIR="/var/lib/solbuild/sources"

if [[ $UID -ne 0 ]]; then
    echo "You must be root to run solbuild_create"
    exit 1
fi

# Typical die function
function do_die() {
    echo "ERROR: $1"
    do_clean
    exit 1
}

# Can't be too careful
function do_clean()
{
    if [[ `grep $PKGDIR /etc/mtab` ]]; then
        umount $GTDIR/$PKGDIR 2>/dev/null
    fi
    if [[ `grep $ARCHIVEDIR /etc/mtab` ]]; then
        umount $GTDIR/$ARCHIVEDIR 2>/dev/null
    fi
    if [[ `grep BUILDROOT/dev /etc/mtab` ]]; then
        umount $GTDIR/dev
    fi
    if [[ `grep BUILDROOT/proc /etc/mtab` ]]; then
        umount $GTDIR/proc 2>/dev/null
    fi
    PID=`cat $GTDIR/run/dbus/pid 2>/dev/null`
    if [[ $? -ne 0 ]]; then
        exit 0
    fi
    kill $PID
    sleep 1
    kill -9 $PID 2>/dev/null
    rm $GTDIR/run/dbus/pid
    if [[ `grep BUILDROOT /etc/mtab` ]]; then
        umount BUILDROOT
    fi
}

function do_img_init()
{
     fallocate -l 10GB ${IMAGE_NAME} || do_die "Unable to construct sparse image"
     mkfs.ext4 ${IMAGE_NAME} -F || do_die "Unable to format sparse image"
     tune2fs -c0 -i0 ${IMAGE_NAME}
}

function do_mount()
{
    TDIR="./BUILDROOT"
    if [[ ! -d "$TDIR" ]]; then
        mkdir -p "$TDIR";
    fi
    mount -o loop -t ext4 ${IMAGE_NAME} BUILDROOT || do_die "Unable to mount sparse image"
}

function do_unmount()
{
    sync
    umount BUILDROOT
    e2fsck -y ${IMAGE_NAME}
    e2fsck -f ${IMAGE_NAME}
}

# trap keyboard interrupt (control-c)
trap do_clean SIGINT

# Initialize a Solus chroot
function do_chroot_init() {
    if [[ ! -d "/var/evobuild" ]]; then
        mkdir /var/evobuild
    fi

    TDIR="./BUILDROOT"
    # Set up normal links before we go ahead and throw eopkg in
    mkdir -p $TDIR/run/lock
    mkdir -p $TDIR/var
    ln -s ../run/lock $TDIR/var/lock
    ln -s ../run $TDIR/var/run
    GTDIR="$TDIR"
    EIT="eopkg it --ignore-comar -D $TDIR/ -y"

    if [[ ! -d "$TDIR/var/cache/eopkg/packages" ]]; then
        mkdir -p "$TDIR/var/cache/eopkg/packages"
    fi
    if [[ ! -d "$TDIR/var/cache/eopkg/archives" ]]; then
        mkdir -p "$TDIR/var/cache/eopkg/archives"
    fi

    # Mount host package directory to save on bandwidth if it exists
    if [[ ! -d "$PKGDIR" ]]; then
        mkdir "$PKGDIR"
    fi
    mount --bind "$PKGDIR" $TDIR/var/cache/eopkg/packages || do_die "Unable to bind mount $PKGDIR"

    # Add repo first
    eopkg add-repo "Solus" "$REPO" -D $TDIR
    # Install eopkg stuffs
    $EIT --ignore-safety baselayout
    $EIT -c system.base

    # Shove in base and devel (safety catches)
    $EIT -c system.devel

    # Needed for solbuild
    $EIT abi-wizard ccache iproute2 sccache python3

    umount $TDIR/var/cache/eopkg/packages

    # Copy base layout files in for accounts
    cp -v $TDIR/usr/share/baselayout/* $TDIR/etc/.

    # Now configure it..
    chroot $TDIR systemd-sysusers

    mount --bind /proc $TDIR/proc
    # Saves us from ugly ass bind-mounts
    chroot $TDIR mknod -m 444 /dev/random c 1 8
    chroot $TDIR mknod -m 444 /dev/urandom c 1 9
    chroot $TDIR mknod -m 644 /dev/null c 1 3

    # Ensure we can let dbus start...
    mkdir $TDIR/run/dbus -p

    chroot $TDIR dbus-uuidgen --ensure
    chroot $TDIR dbus-daemon --system
    chroot $TDIR eopkg configure-pending
    umount $TDIR/proc
    # Cleanup cache
    chroot $TDIR eopkg delete-cache

    # Kill dbus
    PID=`cat $TDIR/run/dbus/pid`
    kill $PID
    sleep 1
    kill -9 $PID 2>/dev/null
    rm $TDIR/run/dbus/pid
    # Clean run out again
    rm -rf $TDIR/run/*
}

do_img_init
do_mount
do_chroot_init
do_unmount

# Got this far, XZ it up
xz -9 ${IMAGE_NAME}
