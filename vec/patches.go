//go:build linux || darwin || freebsd || openbsd || windows

package vec

import (
	sqlite3 "github.com/next-bin/go-sqlite/lib"
	"github.com/next-bin/go-sqlite/libc"
)

func init() {
	tls := libc.NewTLS()
	defer tls.Close()
	sqlite3.Xsqlite3_auto_extension(tls, __ccgo_fp(Xsqlite3_vec_init))
}
