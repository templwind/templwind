package {{.PkgName}}

import (
	{{.Imports}}
)

type {{.ControllerType}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	{{if .HasSocket}}conn   net.Conn{{end}}
}

func New{{.ControllerType}}(ctx context.Context, svcCtx *svc.ServiceContext{{if .HasSocket}}, conn net.Conn{{end}}) *{{.ControllerType}} {
	return &{{.ControllerType}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		{{if .HasSocket}}conn:   conn,{{end}}
		svcCtx: svcCtx,
	}
}

{{- range .Methods}}
{{if .HasDoc}}{{.Doc}}{{end}}
func (l *{{.ControllerType}}) {{.Call}}({{.Request}}) {{.ResponseType}} {
	// todo: add your logic here and delete this line
	{{- if .IsSocket}}
	{{- if .Topic.InitiatedByServer}}
	resp := {{.Topic.ResponseType}}{}

	// send the response to the client via the events engine
	events.Next(types.{{.Topic.Const}}, resp)
	{{else}}
	return
	{{end -}}
	{{else}}
	{{.ReturnString}}
	{{end -}}
}
{{end}}