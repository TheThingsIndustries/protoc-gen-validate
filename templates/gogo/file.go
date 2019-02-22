package gogo

const fileTpl = `package {{ pkg . }}

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"

	{{ range $path, $pkg := enumPackages (externalEnums .) }}
		{{ $pkg }} "{{ $path }}"
	{{ end }}
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}

	{{ range (externalEnums .) }}
		_ = {{ pkg . }}.{{ name . }}(0)
	{{ end }}
)

{{ range .AllMessages }}
	{{ template "msg" . }}
{{ end }}
`
