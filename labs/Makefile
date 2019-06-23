obj-m += helloworld.o

KERNEL_DIR ?= /home/elias/embedded-raspian/linux-modules/lib/modules/4.9.80-v7+/build

all:
	make -C $(KERNEL_DIR) M=$(shell pwd) \
		ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- \
		modules
		
clean:
	make -C $(KERNEL_DIR) \
		ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- \
		SUBDIRS=$(PWD) clean