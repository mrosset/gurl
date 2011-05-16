gurlcmd:
	make -C cmd

gurlpkg:
	make -C pkg	

install: clean
	make -C pkg install
	make -C cmd install

clean:
	make -C pkg clean
	make -C cmd clean
