set -e

if [[ -z $GOOS ]] ; then GOOS=$(go env GOOS) ; fi
if [[ -z $GOARCH ]] ; then GOARCH=$(go env GOARCH) ; fi

OSARCH=$GOOS
OSARCH+=_$GOARCH
LIBC=libc_$OSARCH.go

MUSLARCH=$GOARCH
case $GOARCH in
	"386")
		MUSLARCH=i386 ;;
	"amd64")
		MUSLARCH=x86_64 ;;
esac

echo target arch: $MUSLARCH, file: $LIBC
git checkout ../$LIBC
# eg. $ CRTBOOTSTRAP='-tags crt.bootstrap' ./mk.sh
go install $CRTBOOTSTRAP -v modernc.org/ccgo/v2/ccgo
rm -f log-ccgo
make distclean
make clean
./configure CC=ccgo CFLAGS='-D__typeof=typeof --ccgo-define-values' \
	--target=$MUSLARCH --disable-shared |& tee log-configure
make AR=ar RANLIB=ranlib |& tee log-make
mv -v obj/include/bits/*.h arch/$MUSLARCH/bits/
ccgo -ffreestanding -D_XOPEN_SOURCE=700 -I./arch/$MUSLARCH -I./arch/generic \
       	-Iobj/src/internal -I./src/internal -Iobj/include -I./include \
	-D__typeof=typeof --ccgo-import os,runtime/debug,sync/atomic \
	--ccgo-pkg-name crt -o ../$LIBC ccgo.c lib/libc.a |& tee -a log-ccgo
go install -v modernc.org/ccgo/v2/ccgo
rm -f *.o
date
