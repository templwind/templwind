package echo

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/templwind/templwind/tools/twctl/internal/util"
	"github.com/templwind/templwind/tools/twctl/pkg/site/spec"
)

//go:embed templates/readme.tpl
var readmeTemplate string

//go:embed templates/package.tpl
var packageTemplate string

//go:embed templates/postcss.tpl
var postcssConfigTemplate string

//go:embed templates/tailwind.tpl
var tailwindConfigTemplate string

//go:embed templates/tsconfig.tpl
var tsconfigTemplate string

//go:embed templates/vite.tpl
var viteConfigTemplate string

//go:embed templates/vite-env.tpl
var viteEnvTemplate string

//go:embed templates/styles.tpl
var stylesTemplate string

//go:embed templates/main.ts.tpl
var mainTsTemplate string

//go:embed templates/error.ts.tpl
var errorTsTemplate string

//go:embed templates/gitignore.tpl
var gitignoreTemplate string

//go:embed templates/schema.sql
var schemaTemplate string

//go:embed templates/local.env.tpl
var localEnvTemplate string

//go:embed templates/makefile.tpl
var makefileTemplate string

var (
	srcDir       = "src"
	migrationDir = "db/migrations"
	componentDir = filepath.Join(srcDir, "components")
)

func genNpmFiles(dir string, site *spec.SiteSpec) error {
	filename := util.ToCamel(site.Name)
	// fmt.Println("filename:", filename)

	files := []struct {
		filename        string
		subdir          string
		template        string
		templateName    string
		templateFile    string
		builtinTemplate string
		data            map[string]string
	}{
		{
			filename:        "README.md",
			templateName:    "readmeTemplate",
			templateFile:    readmeTemplateFile,
			builtinTemplate: readmeTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		},
		{
			filename:        "package.json",
			templateName:    "packageTemplate",
			templateFile:    packageTemplateFile,
			builtinTemplate: packageTemplate,
			data: map[string]string{
				"serviceName": strings.ToLower(filename),
			},
		},
		{
			filename:        "postcss.config.js",
			templateName:    "postcssConfigTemplate",
			templateFile:    postcssConfigTemplateFile,
			builtinTemplate: postcssConfigTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		},
		{
			filename:        "tailwind.config.js",
			templateName:    "tailwindConfigTemplate",
			templateFile:    tailwindConfigTemplateFile,
			builtinTemplate: tailwindConfigTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		},
		{
			filename:        "tsconfig.json",
			templateName:    "tsconfigTemplate",
			templateFile:    tsconfigTemplateFile,
			builtinTemplate: tsconfigTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		},
		{
			filename:        "vite.config.js",
			templateName:    "viteConfigTemplate",
			templateFile:    viteConfigTemplateFile,
			builtinTemplate: viteConfigTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		},
		{
			filename:        "vite-env.d.ts",
			subdir:          srcDir,
			templateName:    "viteEnvTemplate",
			templateFile:    viteEnvTemplateFile,
			builtinTemplate: viteEnvTemplate,
		},
		{
			filename:        "styles.scss",
			subdir:          srcDir,
			templateName:    "stylesTemplate",
			templateFile:    stylesTemplateFile,
			builtinTemplate: stylesTemplate,
		},
		{
			filename:        "main.ts",
			subdir:          srcDir,
			templateName:    "mainTsTemplate",
			templateFile:    mainTsTemplateFile,
			builtinTemplate: mainTsTemplate,
		},
		{
			filename:        "error.ts",
			subdir:          componentDir,
			templateName:    "errorTsTemplate",
			templateFile:    errorTsTemplateFile,
			builtinTemplate: errorTsTemplate,
		},
		{
			filename:        ".gitignore",
			templateName:    "gitignoreTemplate",
			templateFile:    gitignoreTemplateFile,
			builtinTemplate: gitignoreTemplate,
		},
		{
			filename:        "1_schema.sql",
			subdir:          migrationDir,
			templateName:    "schemaTemplate",
			templateFile:    schemaTemplateFile,
			builtinTemplate: schemaTemplate,
		},
		{
			filename:        ".env",
			templateName:    "localEnvTemplate",
			templateFile:    localEnvTemplateFile,
			builtinTemplate: localEnvTemplate,
			data: map[string]string{
				"dsnName": strings.ToLower(filename),
			},
		},
		{
			filename:        "Makefile",
			templateName:    "makefileTemplate",
			templateFile:    makefileTemplateFile,
			builtinTemplate: makefileTemplate,
			data: map[string]string{
				"serviceName": strings.ToLower(filename),
			},
		},
	}

	for _, file := range files {
		if err := genFile(fileGenConfig{
			dir:             dir,
			subdir:          file.subdir,
			filename:        file.filename,
			templateName:    file.templateName,
			category:        category,
			templateFile:    file.templateFile,
			builtinTemplate: file.builtinTemplate,
			data:            file.data,
		}); err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}
