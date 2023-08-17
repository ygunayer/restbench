default: all

.DEFAULT:
	cd build && $(MAKE) $@

generate::
	conan install . --output-folder=build --build=missing
	cmake -S . -B build
.PHONY: generate
