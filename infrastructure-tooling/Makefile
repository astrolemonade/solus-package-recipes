SHELL = /bin/bash
TOPLVL = "../."

MEMORY ?= 4096M
DISK_SIZE ?= 10G
ISOFILE ?= Solus-Daily.iso

build:
	sudo python $(TOPLVL)/infrastructure-tooling/image/solus-image-creator.py

boot:
	qemu-system-x86_64 -m ${MEMORY} -cdrom ${ISOFILE} -smp cpus=4 -device intel-hda -device hda-duplex

boot-usb:
	qemu-system-x86_64 -enable-kvm -m ${MEMORY} -cpu host -smp cpus=4 -drive if=none,id=usbstick,file=${ISOFILE},readonly=on,format=raw -usb -device usb-ehci,id=ehci -device usb-tablet,bus=usb-bus.0 -device usb-storage,bus=ehci.0,drive=usbstick -device intel-hda -device hda-duplex

boot-efi:
	qemu-system-x86_64 -enable-kvm -m ${MEMORY} -cdrom ${ISOFILE} -cpu host -smp cpus=4 -bios "/usr/share/edk2-ovmf/x64/OVMF.fd" -device intel-hda -device hda-duplex

boot-disk:
	if [[ ! -e ./disk.img ]]; then \
		qemu-img create -f raw disk.img ${DISK_SIZE}; \
	fi; \
	qemu-system-x86_64 -drive file=disk.img,if=virtio,aio=threads,format=raw -enable-kvm -m ${MEMORY} -cpu host -smp cpus=4 -cdrom ${ISOFILE} -device intel-hda -device hda-duplex

boot-efi-disk:
	if [[ ! -e ./disk.img ]]; then \
		qemu-img create -f raw disk.img ${DISK_SIZE}; \
	fi; \
	qemu-system-x86_64 -drive file=disk.img,if=virtio,aio=threads,format=raw -enable-kvm -m ${MEMORY} -cpu host -smp cpus=4 -bios "/usr/share/edk2-ovmf/x64/OVMF.fd" -cdrom ${ISOFILE} -device intel-hda -device hda-duplex

pull:
	git pull --rebase

help:
	@echo "boot      - Boot the ISO directly"
	@echo "boot-disk - Boot with a working disk for installation"
	@echo "build     - Build the current ISO"
	@echo "pull      - Pull/rebase latest changes"


.PHONY:
	build
