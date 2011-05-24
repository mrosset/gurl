gurlpkg:
	make -C pkg	

gurlcmd:
	make -C cmd

install: clean
	make -C pkg install
	make -C cmd install

clean:
	make -C pkg clean
	make -C cmd clean

all: gurlpkg gurlcmd

