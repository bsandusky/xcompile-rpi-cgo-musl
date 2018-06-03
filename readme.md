#Instructions to cross-compile from macOS High Sierra to Raspberry Pi 3 Model B armv7l against musl libc

1. Run `brew install FiloSottile/musl-cross/musl-cross --without-x86_64 --with-arm-hf`

This will take some time (my installation took ~45 minutes), but afterwards, a number of compilers for musl will now be in /usr/local/bin. Specifically, the one that we're concerned with is `arm-linux-musleabihf-gcc` which will replace the regular `gcc` when we call `go build` or `vgo build` to generate a binary.

2. Create go code inclusive of cgo (i.e. import "C")

3. To compile a dynamic binary against musl libc for Rpi use the following build command:

`env CC=arm-linux-musleabihf-gcc GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 vgo build -o some_name /some/path/to/main.go`

4. To compile a static binary against musl libc for Rpi use the following build command:

`env CC=arm-linux-musleabihf-gcc GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 vgo build -o some_name --ldflags '-linkmode external -extldflags "-static"' /some/path/to/main.go`

(Yes, `vgo build` works, too!)

The difference in size between these two will be minimal, unlike when using glibc where there is a greater delta between dynamic binaries linking to glibc directly and static.

Also, if you fail to include `CGO_ENABLED=1` in the build command, compilation will fail. The compiler will ignore any files containing the `import "C"` directive and an error on compilation will come up stating that go symbols are undefined. This bit was confusing and threw me off. I couldn't figure out why the compiler was complaining a go func was undefined. I ran the build using the `-x` switch and realized that the particular file importing C was ignored. To ensure that these go files with cgo in them are included, you must explicitly tell the build command to enable cgo.

5. Copy or move binary output to Rpi via `scp` or samba. 

6. Execute binary on Rpi. Ta da! It works! And your cgo code will call C using musl libc.
