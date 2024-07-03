package echo

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"

	"github.com/templwind/templwind/tools/twctl/internal/types"
	"github.com/templwind/templwind/tools/twctl/internal/util"
	"github.com/templwind/templwind/tools/twctl/pkg/site/spec"

	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
)

//go:embed templates/controller.tpl
var controllerTemplate string

//go:embed templates/controller.templ.tpl
var controllerTemplTemplate string

//go:embed templates/props.tpl
var propsTemplate string

func genController(dir, rootPkg string, cfg *config.Config, site *spec.SiteSpec) error {
	for _, server := range site.Servers {
		for _, service := range server.Services {
			for _, handler := range service.Handlers {
				err := genControllerByHandler(dir, rootPkg, cfg, server, handler)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func addMissingMethods(methods []MethodConfig, dir, subDir, fileName string) error {
	// Read the file and look for all the methods and compare with the defined methods
	filePath := path.Join(dir, subDir, fileName)
	fbytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read file failed: %w", err)
	}

	fileContent := string(fbytes)
	var newMethods []string

	for _, method := range methods {
		if !strings.Contains(fileContent, method.Call) {

			// Add the method definition to the newMethods slice
			newMethods = append(newMethods, generateMethodDefinition(method))
		}
	}

	// If there are new methods to add, append them to the file
	if len(newMethods) > 0 {
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return fmt.Errorf("open file for writing failed: %w", err)
		}
		defer f.Close()

		for _, newMethod := range newMethods {
			if _, err := f.WriteString(newMethod); err != nil {
				return fmt.Errorf("write to file failed: %w", err)
			}
		}
	}

	return nil
}

// This is the function to generate the method definition based on your template
func generateMethodDefinition(method MethodConfig) string {
	tmpl := `{{if .HasDoc}}{{.Doc}}{{end}}
func (l *{{.ControllerType}}) {{.Call}}({{.Request}}) {{.ResponseType}} {
	// todo: add your logic here and delete this line

	{{.ReturnString}}
}
`
	t, err := template.New("method").Parse(tmpl)
	if err != nil {
		panic(fmt.Sprintf("parsing template failed: %v", err))
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, method)
	if err != nil {
		panic(fmt.Sprintf("executing template failed: %v", err))
	}

	return buf.String()
}

func genControllerByHandler(dir, rootPkg string, cfg *config.Config, server spec.Server, handler spec.Handler) error {

	controllerLayout := server.GetAnnotation("template")

	subDir := getControllerFolderPath(server, handler)
	filename := path.Join(dir, subDir, strings.ToLower(handler.Name)+".go")
	// fmt.Println("filename::", filename)

	fileExists := false
	// check if the file exists
	if pathx.FileExists(filename) {
		fileExists = true
	}

	requiresTempl := false
	hasSocket := false

	methods := []MethodConfig{}
	for _, method := range handler.Methods {

		if !method.IsSocket {
			requiresTempl = true
		} else {
			hasSocket = true
		}

		// skip this method if it is static
		if method.IsStatic {
			continue
		}

		if method.Page != nil {
			if key, ok := method.Page.Annotation.Properties["template"]; ok {
				if layoutName, ok := key.(string); ok {
					controllerLayout = layoutName
				}
			}
		}

		var responseString string
		var returnString string
		var requestString string
		var controllerName string
		var hasResp bool
		var hasReq bool
		var requestType string
		var handlerName string
		var controllerType string
		var call string

		controllerType = strings.Title(getControllerName(handler))

		if method.IsSocket && method.SocketNode != nil {
			for _, topic := range method.SocketNode.Topics {
				call = util.ToPascal(topic.Topic)

				requestString = ""
				responseString = ""
				returnString = ""
				hasReq = false

				if topic.InitiatedByClient {
					resp := util.TopicResponseGoTypeName(topic, types.TypesPacket)
					responseString = "(resp " + resp + ", err error)"
					returnString = "return"

					if topic.RequestType != nil && len(topic.RequestType.GetName()) > 0 {
						hasReq = true
						requestString = "req " + util.TopicRequestGoTypeName(topic, types.TypesPacket)
					}
				}

				methods = append(methods, MethodConfig{
					HandlerName:    handlerName,
					RequestType:    requestType,
					ResponseType:   responseString,
					Request:        requestString,
					ReturnString:   returnString,
					ResponseString: responseString,
					HasResp:        hasResp,
					HasReq:         hasReq,
					HasDoc:         method.Doc != nil,
					HasPage:        method.Page != nil,
					Doc:            "",
					ControllerName: controllerName,
					ControllerType: controllerType,
					Call:           call,
					IsSocket:       method.IsSocket,
					Topic: Topic{
						InitiatedByServer: !topic.InitiatedByClient,
						InitiatedByClient: topic.InitiatedByClient,
						Const:             "Topic" + util.ToPascal(topic.Topic),
						ResponseType:      strings.ReplaceAll(util.TopicResponseGoTypeName(topic, types.TypesPacket), "*", "&"),
					},
				})
			}
		} else {
			if method.ResponseType != nil && len(method.ResponseType.GetName()) > 0 {
				resp := util.ResponseGoTypeName(method, types.TypesPacket)
				responseString = "(resp " + resp + ", err error)"
				returnString = "return"
			} else {
				responseString = "(templ.Component, error)"
				returnString = fmt.Sprintf(`return New(
				WithConfig(l.svcCtx.Config),
				WithRequest(c.Request()),
				WithTitle("%s"),
			), nil`, util.ToTitle(handler.Name))
			}

			if method.RequestType != nil && len(method.RequestType.GetName()) > 0 {
				requestString = "req " + util.RequestGoTypeName(method, types.TypesPacket)
			}

			hasResp = method.ResponseType != nil && len(method.ResponseType.GetName()) > 0
			hasReq := method.RequestType != nil && len(method.RequestType.GetName()) > 0

			requestType = ""
			if hasReq {
				requestType = util.ToTitle(method.RequestType.GetName())
			}

			handlerName = util.ToTitle(getHandlerName(handler, &method))

			requestStringParts := []string{
				requestString,
				"c echo.Context",
			}
			requestString = func(parts []string) string {
				rParts := make([]string, 0)
				for _, part := range parts {
					if len(part) == 0 {
						continue
					}
					rParts = append(rParts, strings.TrimSpace(part))
				}
				return strings.Join(rParts, ", ")
			}(requestStringParts)

			controllerName = strings.ToLower(util.ToCamel(handler.Name))
			call = strings.Title(strings.TrimSuffix(handlerName, "Handler"))

			// fmt.Println("handlerName:", handlerName)
			methods = append(methods, MethodConfig{
				HandlerName:    handlerName,
				RequestType:    requestType,
				ResponseType:   responseString,
				Request:        requestString,
				ReturnString:   returnString,
				ResponseString: responseString,
				HasResp:        hasResp,
				HasReq:         hasReq,
				HasDoc:         method.Doc != nil,
				HasPage:        method.Page != nil,
				Doc:            "",
				ControllerName: controllerName,
				ControllerType: controllerType,
				Call:           call,
				IsSocket:       method.IsSocket,
			})
		}
	}

	if fileExists {
		return addMissingMethods(methods,
			dir,
			subDir,
			strings.ToLower(handler.Name)+".go")
	}

	if requiresTempl {
		templImports := genTemplImports(rootPkg, strings.ToLower(util.ToCamel(controllerLayout+"Layout")))

		// fmt.Println("templImports", templImports)
		// templ file first
		if err := genFile(fileGenConfig{
			dir:             dir,
			subdir:          subDir,
			filename:        strings.ToLower(handler.Name) + ".templ",
			templateName:    "controllerTemplTemplate",
			category:        category,
			templateFile:    controllerTemplTemplateFile,
			builtinTemplate: controllerTemplTemplate,
			data: map[string]any{
				"pkgName":      subDir[strings.LastIndex(subDir, "/")+1:],
				"templImports": templImports,
				"templName":    util.ToCamel(handler.Name + "View"),
				"pageTitle":    util.ToTitle(handler.Name),
				// "controllerLayout": strings.ToLower(util.ToCamel(controllerLayout + "Layout")),
			},
		}); err != nil {
			return err
		}

		propsImports := genPropsImports(rootPkg)

		if err := genFile(fileGenConfig{
			dir:             dir,
			subdir:          subDir,
			filename:        "props.go",
			templateName:    "controllerPropsTemplate",
			category:        category,
			templateFile:    propsTemplateFile,
			builtinTemplate: propsTemplate,
			data: map[string]any{
				"pkgName":   subDir[strings.LastIndex(subDir, "/")+1:],
				"Imports":   propsImports,
				"templName": util.ToCamel(handler.Name + "View"),
			},
		}); err != nil {
			return err
		}
	}

	imports := genControllerImports(handler, rootPkg)
	controllerType := strings.Title(getControllerName(handler))

	// sort.Slice(methods, func(i, j int) bool {
	// 	return methods[i].Call < methods[j].Call
	// })

	err := genFile(fileGenConfig{
		dir:             dir,
		subdir:          subDir,
		filename:        strings.ToLower(handler.Name) + ".go",
		templateName:    "controllerTemplate",
		category:        category,
		templateFile:    controllerTemplateFile,
		builtinTemplate: controllerTemplate,
		data: map[string]any{
			"PkgName":        subDir[strings.LastIndex(subDir, "/")+1:],
			"Imports":        imports,
			"ControllerType": controllerType,
			"Methods":        methods,
			"HasSocket":      hasSocket,
		},
	})

	// os.Exit(0)
	return err
}

func getControllerFolderPath(server spec.Server, handler spec.Handler) string {
	folder := server.GetAnnotation(types.GroupProperty)
	if len(folder) == 0 || folder == "/" {
		return types.ControllerDir
	}
	folder = strings.TrimPrefix(folder, "/")
	folder = strings.TrimSuffix(folder, "/")

	return path.Join(types.ControllerDir, folder, strings.ToLower(handler.Name))
}

func genTemplImports(parentPkg, fileName string) string {
	imports := []string{
		// fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.LayoutsDir, fileName)),
	}
	return strings.Join(imports, "\n\t")
}

func genPropsImports(parentPkg string) string {
	imports := []string{
		fmt.Sprintf("\"%s\"\n", "net/http"),
		fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.ConfigDir)),
		fmt.Sprintf("\"%s\"\n", "github.com/a-h/templ"),
		fmt.Sprintf("\"%s\"", "github.com/templwind/templwind"),
	}
	return strings.Join(imports, "\n\t")
}

func genControllerImports(handler spec.Handler, parentPkg string) string {
	var imports []string

	requireTemplwind := false
	requireEcho := false
	hasType := false
	hasSocket := false
	hasEvents := false
	for _, method := range handler.Methods {
		// show when the response type is empty
		if (method.ResponseType == nil || method.ReturnsPartial) && !method.IsSocket {
			requireTemplwind = true
			requireEcho = true
		}

		if method.ResponseType != nil || method.RequestType != nil {
			hasType = true
			requireEcho = true
		}

		if method.IsSocket {
			hasSocket = true
			for _, topic := range method.SocketNode.Topics {
				if topic.ResponseType != nil || topic.RequestType != nil {
					hasType = true
				}
				if !topic.InitiatedByClient {
					hasEvents = true
				}
			}
		}
	}

	imports = append(imports, fmt.Sprintf("\"%s\"", "context"))
	if hasSocket {
		imports = append(imports, fmt.Sprintf("\"%s\"", "net"))
	}
	imports = append(imports, "\n\n")

	if hasEvents {
		imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.EventsDir)))
	}

	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.ContextDir)))

	if hasType {
		imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.TypesDir)))
	}

	imports = append(imports, "\n\n")

	if requireTemplwind {
		imports = append(imports, fmt.Sprintf("\n\n\"%s\"", "github.com/a-h/templ"))
	}
	if requireEcho {
		imports = append(imports, fmt.Sprintf("\"%s\"", "github.com/labstack/echo/v4"))
	}
	// TODO: method fix

	// if requireTemplwind {
	// 	imports = append(imports, "\"github.com/templwind/templwind\"")
	// }
	imports = append(imports, fmt.Sprintf("\"%s/core/logx\"", vars.ProjectOpenSourceURL))
	return strings.Join(imports, "\n\t")
}

func onlyPrimitiveTypes(val string) bool {
	fields := strings.FieldsFunc(val, func(r rune) bool {
		return r == '[' || r == ']' || r == ' '
	})

	for _, field := range fields {
		if field == "map" {
			continue
		}
		// ignore array dimension number, like [5]int
		if _, err := strconv.Atoi(field); err == nil {
			continue
		}
		if !api.IsBasicType(field) {
			return false
		}
	}

	return true
}

func shallImportTypesPackage(method spec.Method) bool {

	if method.RequestType != nil && len(method.RequestType.GetName()) > 0 {
		return true
	}

	// fmt.Println("method.RequestType.GetName()", method.RequestType.GetName())

	respTypeName := method.ResponseType
	if method.ResponseType == nil || len(respTypeName.GetName()) == 0 {
		return false
	}

	if onlyPrimitiveTypes(respTypeName.GetName()) {
		return false
	}

	return true
}
