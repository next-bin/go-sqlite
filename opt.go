// Package opt implements command-line flag parsing.
package opt // import "modernc.org/opt"

import (
	"fmt"
	"strings"
)

type opt struct {
	handler func(opt, arg string) error
	name    string

	arg bool // Enable argument, e.g. -I foo or -I=foo
}

type Parser struct {
	cfg map[string]*opt
	imm []*opt
}

func NewParser() *Parser { return &Parser{cfg: map[string]*opt{}} }

// Opt defines a simple option, e.g. -f.
func (p *Parser) Opt(name string, handler func(opt string) error) {
	p.cfg[name] = &opt{
		handler: func(opt, arg string) error { return handler(opt) },
	}
}

// OptArg defines a simple option with an argumen, e.g. -Ifoo, -I=foo, -I foo.
// The imm argument enables the first form above.
func (p *Parser) OptArg(name string, imm bool, handler func(opt, arg string) error) {
	switch {
	case imm:
		p.imm = append(p.imm, &opt{
			handler: handler,
			name:    name,
		})
	default:
		p.cfg[name] = &opt{
			arg:     true,
			handler: handler,
			name:    name,
		}
	}
}

func (p *Parser) Parse(opts []string, argHandler func(string) error) error {
	for len(opts) != 0 {
		opt := opts[0]
		opts = opts[1:]
		var arg string
	out:
		switch {
		case strings.HasPrefix(opt, "-"):
			name := opt[1:]
			for _, cfg := range p.imm {
				if strings.HasPrefix(name, cfg.name) {
					switch {
					case name == cfg.name:
						if len(opts) == 0 {
							return fmt.Errorf("missing argument of %s", opt)
						}

						if err := cfg.handler(opt, opts[0]); err != nil {
							return err
						}

						opts = opts[1:]
					default:
						if err := cfg.handler(opt[:len(cfg.name)+1], name[len(cfg.name):]); err != nil {
							return err
						}
					}
					break out
				}
			}

			if n := strings.IndexByte(opt, '='); n > 0 {
				arg = opt[n+1:]
				name = opt[1:n]
				opt = opt[:n]
			}
			switch cfg := p.cfg[name]; {
			case cfg == nil:
				if err := argHandler(opt); err != nil {
					return err
				}
			default:
				switch {
				case cfg.arg:
					switch {
					case arg != "":
						if err := cfg.handler(opt, arg); err != nil {
							return err
						}
					default:
						if len(opts) == 0 {
							return fmt.Errorf("missing argument of %s", opt)
						}

						if err := cfg.handler(opt, opts[0]); err != nil {
							return err
						}

						opts = opts[1:]
					}
				default:
					cfg.handler(opt, "")
				}
			}
		default:
			if opt == "" {
				break
			}

			if err := argHandler(opt); err != nil {
				return err
			}
		}
	}
	return nil
}
