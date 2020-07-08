// markdown-template.go: the markdown template used to template the sysl module
package catalog

const ProjectTemplate = `
{{/* Automatically generated by https://github.com/anz-bank/sysl-catalog it is strongly recommended not to edit this file */}}
{{range $name, $link := .Links}} [{{$name}}]({{$link}}) | {{end}} [Chat with us]({{.ChatLink}}) | [New bug or feature request]({{.FeedbackLink}})
# {{Base .Title}}

| Package |
----|{{range $val := Packages .Module}}
[{{$val}}]({{$val}}/README.md)|{{end}}

## Integration Diagram
<img src="{{CreateIntegrationDiagram .Module .Title false}}">

## End Point Analysis Integration Diagram
<img src="{{CreateIntegrationDiagram .Module .Title true}}">

`

const MacroPackageProject = `
{{/* Automatically generated by https://github.com/anz-bank/sysl-catalog it is strongly recommended not to edit this file */}}
[Chat with us]({{.ChatLink}}) | [New bug or feature request]({{.FeedbackLink}})
# {{Base .Title}}

| Package |
----|{{range $val := MacroPackages .Module}}
[{{$val}}]({{$val}}/README.md)|{{end}}

## Integration Diagram
<img src="{{CreateIntegrationDiagram .Module .Title false}}">

## End Point Analysis Integration Diagram
<img src="{{CreateIntegrationDiagram .Module .Title true}}">

`

const NewPackageTemplate = `
{{/* Automatically generated by https://github.com/anz-bank/sysl-catalog it is strongly recommended not to edit this file */}}
[Back](../README.md) | [Chat with us]({{ChatLink}}) | [New bug or feature request]({{FeedbackLink}})
{{$packageName := ModulePackageName .}}

# {{$packageName}}

## Integration Diagram
![]({{CreateIntegrationDiagram . $packageName false}})
{{$Apps := .Apps}}

{{$databases := false}}
{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}{{if and (eq (hasPattern $app.Attrs "ignore") false) (eq (hasPattern $app.Attrs "db") true)}}
{{$databases = true}}
{{end}}{{end}}

{{if $databases}}
## Database Index
| Database Application Name  | Source Location |
----|----{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}{{if and (eq (hasPattern $app.Attrs "ignore") false) (eq (hasPattern $app.Attrs "db") true)}}
[{{$appName}}](#Database-{{$appName}}) | [{{SourcePath $app}}]({{SourcePath $app}})|  {{end}}{{end}}
{{end}}

## Application Index
{{$anyApps := false}}

{{$Apps := .Apps}}{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}{{if eq (hasPattern $app.Attrs "ignore") false}}{{$Endpoints := $app.Endpoints}}{{range $endpointName := SortedKeys $Endpoints}}{{$endpoint := index $Endpoints $endpointName}}{{if eq (hasPattern $endpoint.Attrs "ignore") false}}{{if not $anyApps}}| Application Name | Method | Source Location |
|----|----|----|{{$anyApps = true}}{{end}}
| {{$appName}} | [{{$endpoint.Name}}](#{{$appName}}-{{SanitiseOutputName $endpoint.Name}}) | [{{SourcePath $app}}]({{SourcePath $app}})|  {{end}}{{end}}{{end}}{{end}}

{{if not $anyApps}}
<span style="color:grey">No Applications Defined</span>
{{end}}


## Type Index
{{$anyTypes := false}}

{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}{{$types := $app.Types}}{{if ne (hasPattern $app.Attrs "db") true}}{{range $typeName := SortedKeys $types}}{{$type := index $types $typeName}}{{if not $anyTypes}}| Application Name | Type Name | Source Location |
|----|----|----|{{$anyTypes = true}}{{end}}
| {{$appName}} | [{{$typeName}}](#{{$appName}}.{{$typeName}}) | [{{SourcePath $type}}]({{SourcePath $type}})|{{end}}{{end}}{{end}}

{{if not $anyTypes}}
<span style="color:grey">No Types Defined</span>
{{end}}


{{if $databases}}
# Databases
{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}
{{if hasPattern $app.GetAttrs "db"}}

<details>
<summary>Database {{$appName}}</summary>

{{Attribute $app "description"}}
![]({{GenerateDataModel $app}})
</details>
{{end}}{{end}}
{{end}}


{{if $anyApps}}
# Applications
{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}
{{if eq (hasPattern $app.Attrs "ignore") false}}
{{if eq (hasPattern $app.Attrs "db") false}}
{{if ne (len $app.Endpoints) 0}}

## Application {{$appName}}

{{$desc := Attribute $app "description"}}
{{if $desc}}
- {{$desc}}
{{end}}

{{ServiceMetadata $app}}

{{with CreateRedoc $app.SourceContext $appName}}
[View OpenAPI Specs in Redoc]({{CreateRedoc $app.SourceContext $appName}})
{{end}}

{{range $e := $app.Endpoints}}
{{if eq (hasPattern $e.Attrs "ignore") false}}


### <a name={{$appName}}-{{SanitiseOutputName $e.Name}}></a>{{$appName}} {{$e.Name}}
{{Attribute $e "description"}}

<details>
<summary>Sequence Diagram</summary>

![]({{CreateSequenceDiagram $appName $e}})
</details>

<details>
<summary>Request types</summary>

{{if and (not $e.Param) (not $e.RestParams) }}
<span style="color:grey">No Request types</span>
{{end}}
{{if not $e.Param}}{{if $e.RestParams }}{{if not $e.RestParams.UrlParam}}
<span style="color:grey">No Request types</span>
{{end}}{{end}}{{end}}

{{range $param := $e.Param}}
{{Attribute $param.Type "description"}}

![]({{CreateParamDataModel $app $param}})
{{end}}

{{if $e.RestParams}}{{if $e.RestParams.UrlParam}}
{{range $param := $e.RestParams.UrlParam}}
{{$pathDataModel := (CreateParamDataModel $app $param)}}
{{if ne $pathDataModel ""}}
#### Path Parameter

![]({{$pathDataModel}})
{{end}}{{end}}{{end}}

{{if $e.RestParams.QueryParam}}
{{range $param := $e.RestParams.QueryParam}}
{{$queryDataModel := (CreateParamDataModel $app $param)}}
{{if ne $queryDataModel ""}}
#### Query Parameter

![]({{$queryDataModel}})
{{end}}{{end}}{{end}}{{end}}
</details>

<details>
<summary>Response types</summary>

{{$responses := false}}
{{range $s := $e.Stmt}}{{$diagram := CreateReturnDataModel  $appName $s $e}}{{if ne $diagram ""}}
{{$responses = true}}
{{$ret := (GetReturnType $e $s)}}{{if $ret }}
{{Attribute $ret "description"}}{{end}}

![]({{$diagram}})

{{end}}{{end}}

{{if not $responses}}
<span style="color:grey">No Response Types</span>
{{end}}
</details>
{{end}}

---

{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}


{{if $anyTypes}}
# Types


{{range $appName := SortedKeys .Apps}}{{$app := index $Apps $appName}}{{$types := $app.Types}}
{{if ne (hasPattern $app.Attrs "db") true}}


{{range $typeName := SortedKeys $types}}{{$type := index $types $typeName}}
<a name={{$appName}}.{{$typeName}}></a><details>
<summary>{{$appName}}.{{$typeName}}</summary>

### {{$appName}}.{{$typeName}}
{{$typedesc := (Attribute $type "description")}}
{{if ne $typedesc ""}}- {{$typedesc}}{{end}}

![]({{CreateTypeDiagram $appName $typeName $type false}})

[Full Diagram]({{CreateTypeDiagram $appName $typeName $type true}})

{{if Fields $type}}
#### Fields
{{$fieldHeader := false}}
{{$fieldMap := Fields $type}}{{range $fieldName := SortedKeys $fieldMap}}{{$field := index $fieldMap $fieldName}}{{if not $fieldHeader}}| Field name | Type | Description |
|----|----|----|{{$fieldHeader = true}}{{end}}
| {{$fieldName}} | {{FieldType $field}} | {{$desc := Attribute $field "description"}}{{if ne $desc $typedesc}}{{$desc}}{{end}}|{{end}}
{{end}}

</details>{{end}}{{end}}{{end}}
{{end}}

<div class="footer">

`
