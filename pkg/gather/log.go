// SPDX-FileCopyrightText: The kubectl-gather authors
// SPDX-License-Identifier: Apache-2.0

package gather

import (
	"fmt"
	"io"
	"log"
	stdlog "log"
	"os"
)

func NewLogger(name string, opts *Options) *log.Logger {
	if opts.Verbose {
		prefix := fmt.Sprintf("%s/%s: ", opts.Context, name)
		return stdlog.New(os.Stderr, prefix, log.LstdFlags|log.Lmicroseconds)
	}
	return stdlog.New(io.Discard, "", 0)
}
