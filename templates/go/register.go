package golang

import (
	"text/template"

	"github.com/envoyproxy/protoc-gen-validate/templates/goshared"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	goshared.Register(tpl, params, pgsgo.InitContext)
	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("required").Parse(requiredTpl))
	template.Must(tpl.New("timestamp").Parse(timestampTpl))
	template.Must(tpl.New("duration").Parse(durationTpl))
	template.Must(tpl.New("message").Parse(messageTpl))
}
