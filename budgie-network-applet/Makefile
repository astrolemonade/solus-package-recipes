all: build

build: BUILDDIR
	ninja -C BUILDDIR

BUILDDIR:
	meson --buildtype debug BUILDDIR

reconfigure:
	meson --reconfigure --buildtype debug BUILDDIR

clean:
	rm -rf BUILDDIR || exit 0

test:
	./BUILDDIR/budgie-network-applet