package catalog

const ProjectHTMLTemplate = `
<!doctype html>
<html lang="en"><head>
<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<div id='content'>
<h1 id="-title-">{{.Title}}</h1>
<table>
  <tr>
    <th> Packages </th>
  </tr>
{{range $Package := .AlphabeticalRows}}<tr><td>
<a href="{{$Package.PackageName}}/{{$Package.OutputFile}}">{{$Package.PackageName}}</a></td></tr>{{end}}</p>
</table>
<h2>Integration diagram:</h2>
<p><img src="{{.RootLevelIntegrationDiagram.OutputFileName__}}" alt="alt text"></p>
<h2>Integration diagram with end point analysis:</h2>
<p><img src="{{.RootLevelIntegrationDiagramEPA.OutputFileName__}}" alt="alt text"></p>
</div>
<style type="text/css">body {
  font: 400 16px/1.5 "Helvetica Neue", Helvetica, Arial, sans-serif;
  color: #111;
  background-color: #fdfdfd;
  -webkit-text-size-adjust: 100%;
  -webkit-font-feature-settings: "kern" 1;
  -moz-font-feature-settings: "kern" 1;
  -o-font-feature-settings: "kern" 1;
  font-feature-settings: "kern" 1;
  font-kerning: normal;
  padding: 30px;
}

@media only screen and (max-width: 600px) {
  body {
    padding: 5px;
  }

  body > #content {
    padding: 0px 20px 20px 20px !important;
  }
}

body > #content {
  margin: 0px;
  max-width: 900px;
  border: 1px solid #e1e4e8;
  padding: 10px 40px;
  padding-bottom: 20px;
  border-radius: 2px;
  margin-left: auto;
  margin-right: auto;
}

hr {
  color: #bbb;
  background-color: #bbb;
  height: 1px;
  flex: 0 1 auto;
  margin: 1em 0;
  padding: 0;
  border: none;
}

/**
 * Links
 */
a {
  color: #0366d6;
  text-decoration: none; }
  a:visited {
    color: #0366d6; }
  a:hover {
    color: #0366d6;
    text-decoration: underline; }

pre {
  background-color: #f6f8fa;
  border-radius: 3px;
  font-size: 85%;
  line-height: 1.45;
  overflow: auto;
  padding: 16px;
}

/**
  * Code blocks
  */

code {
  background-color: rgba(27,31,35,.05);
  border-radius: 3px;
  font-size: 85%;
  margin: 0;
  word-wrap: break-word;
  padding: .2em .4em;
  font-family: SFMono-Regular,Consolas,Liberation Mono,Menlo,Courier,monospace;
}

pre > code {
  background-color: transparent;
  border: 0;
  display: inline;
  line-height: inherit;
  margin: 0;
  overflow: visible;
  padding: 0;
  word-wrap: normal;
  font-size: 100%;
}


/**
 * Blockquotes
 */
blockquote {
  margin-left: 30px;
  margin-top: 0px;
  margin-bottom: 16px;
  border-left-width: 3px;
  padding: 0 1em;
  color: #828282;
  border-left: 4px solid #e8e8e8;
  padding-left: 15px;
  font-size: 18px;
  letter-spacing: -1px;
  font-style: italic;
}
blockquote * {
  font-style: normal !important;
  letter-spacing: 0;
  color: #6a737d !important;
}

/**
 * Tables
 */
table {
  border-spacing: 2px;
  display: block;
  font-size: 14px;
  overflow: auto;
  width: 100%;
  margin-bottom: 16px;
  border-spacing: 0;
  border-collapse: collapse;
}

td {
  padding: 6px 13px;
  border: 1px solid #dfe2e5;
}

th {
  font-weight: 600;
  padding: 6px 13px;
  border: 1px solid #dfe2e5;
}

tr {
  background-color: #fff;
  border-top: 1px solid #c6cbd1;
}

table tr:nth-child(2n) {
  background-color: #f6f8fa;
}

/**
 * Others
 */

img {
  max-width: 100%;
}

p {
  line-height: 24px;
  font-weight: 400;
  font-size: 16px;
  color: #24292e; }

ul {
  margin-top: 0; }

li {
  color: #24292e;
  font-size: 16px;
  font-weight: 400;
  line-height: 1.5; }

li + li {
  margin-top: 0.25em; }

* {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
  color: #24292e; }

a:visited {
  color: #0366d6; }

h1, h2, h3 {
  border-bottom: 1px solid #eaecef;
  color: #111;
  /* Darker */ }</style>
</body>
</html>
`

const PackageHTMLTemplate = `
<!doctype html>
<html lang="en"><head>
<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<div id='content'>
<p><a href="../index.html">Back</a></p>
<h1 id="package-packagename-">Package {{.PackageName}}</h1>
<h2 id="index">Index</h2>

<table>
  <tr>
    <th>Service Name</th>
	<th>Method</th>
  </tr>
 {{range $appName, $Diagrams := .SequenceDiagrams}}{{range $Diagram := $Diagrams}}
<tr><td>{{$Diagram.AppName}} </td> <td><a href="#{{$Diagram.AppName}}-{{$Diagram.EndpointName}}">{{$Diagram.EndpointName}}</a></td></tr>{{end}}{{end}}
{{range $appName, $Diagrams := .DatabaseModel}}
<tr><td>{{$appName}} </td> <td><a href="#Database-{{$appName}}">Database</a></td></tr>{{end}}
</table>
<hr>

<h2>Integration diagram:</h2>
<p><img src="{{.Integration.OutputFileName__}}" alt="alt text"></p>

{{range $appName, $Diagrams := .SequenceDiagrams}}
{{$first := true}}
{{range $Diagram := $Diagrams}}
{{if $first}}
<h2> {{$Diagram.AppName}} </h2>
 {{$Diagram.AppComment}}
{{end}}
{{$first = false}}
<h2 id="{{$Diagram.AppName}}-{{$Diagram.EndpointName}}">{{$Diagram.AppName}} {{$Diagram.EndpointName}}</h2>
<h3 id="sequence-diagram">Sequence Diagram</h3>
<p><img src="{{.OutputFileName__}}.svg" alt="alt text"></p>
<h3 id="request-types">Request types</h3>
<p>{{range $DataModelDiagram := $Diagram.InputDataModel}}
<img src="{{$DataModelDiagram.OutputFileName__}}.svg" alt="alt text">
{{end}}</p>
<h3 id="response-types">Response types</h3>
<p>{{range $DataModelDiagram := $Diagram.OutputDataModel}}
<img src="{{$DataModelDiagram.OutputFileName__}}.svg" alt="alt text"></p>
<h2 id="-end-">{{end}}</h2>
{{end}}{{end}}


<p>{{range $appName, $Diagrams := .DatabaseModel}}
<h3 id="Database-{{$appName}}">Database {{$appName}}</h3>
<img src="{{$Diagrams.OutputFileName__}}.svg" alt="alt text"></p>
<h2 id="-end-">{{end}}</h2>
</div>
<style type="text/css">body {
  font: 400 16px/1.5 "Helvetica Neue", Helvetica, Arial, sans-serif;
  color: #111;
  background-color: #fdfdfd;
  -webkit-text-size-adjust: 100%;
  -webkit-font-feature-settings: "kern" 1;
  -moz-font-feature-settings: "kern" 1;
  -o-font-feature-settings: "kern" 1;
  font-feature-settings: "kern" 1;
  font-kerning: normal;
  padding: 30px;
}

@media only screen and (max-width: 600px) {
  body {
    padding: 5px;
  }

  body > #content {
    padding: 0px 20px 20px 20px !important;
  }
}

body > #content {
  margin: 0px;
  max-width: 900px;
  border: 1px solid #e1e4e8;
  padding: 10px 40px;
  padding-bottom: 20px;
  border-radius: 2px;
  margin-left: auto;
  margin-right: auto;
}

hr {
  color: #bbb;
  background-color: #bbb;
  height: 1px;
  flex: 0 1 auto;
  margin: 1em 0;
  padding: 0;
  border: none;
}

/**
 * Links
 */
a {
  color: #0366d6;
  text-decoration: none; }
  a:visited {
    color: #0366d6; }
  a:hover {
    color: #0366d6;
    text-decoration: underline; }

pre {
  background-color: #f6f8fa;
  border-radius: 3px;
  font-size: 85%;
  line-height: 1.45;
  overflow: auto;
  padding: 16px;
}

/**
  * Code blocks
  */

code {
  background-color: rgba(27,31,35,.05);
  border-radius: 3px;
  font-size: 85%;
  margin: 0;
  word-wrap: break-word;
  padding: .2em .4em;
  font-family: SFMono-Regular,Consolas,Liberation Mono,Menlo,Courier,monospace;
}

pre > code {
  background-color: transparent;
  border: 0;
  display: inline;
  line-height: inherit;
  margin: 0;
  overflow: visible;
  padding: 0;
  word-wrap: normal;
  font-size: 100%;
}


/**
 * Blockquotes
 */
blockquote {
  margin-left: 30px;
  margin-top: 0px;
  margin-bottom: 16px;
  border-left-width: 3px;
  padding: 0 1em;
  color: #828282;
  border-left: 4px solid #e8e8e8;
  padding-left: 15px;
  font-size: 18px;
  letter-spacing: -1px;
  font-style: italic;
}
blockquote * {
  font-style: normal !important;
  letter-spacing: 0;
  color: #6a737d !important;
}

/**
 * Tables
 */
table {
  border-spacing: 2px;
  display: block;
  font-size: 14px;
  overflow: auto;
  width: 100%;
  margin-bottom: 16px;
  border-spacing: 0;
  border-collapse: collapse;
}

td {
  padding: 6px 13px;
  border: 1px solid #dfe2e5;
}

th {
  font-weight: 600;
  padding: 6px 13px;
  border: 1px solid #dfe2e5;
}

tr {
  background-color: #fff;
  border-top: 1px solid #c6cbd1;
}

table tr:nth-child(2n) {
  background-color: #f6f8fa;
}

/**
 * Others
 */

img {
  max-width: 100%;
}

p {
  line-height: 24px;
  font-weight: 400;
  font-size: 16px;
  color: #24292e; }

ul {
  margin-top: 0; }

li {
  color: #24292e;
  font-size: 16px;
  font-weight: 400;
  line-height: 1.5; }

li + li {
  margin-top: 0.25em; }

* {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
  color: #24292e; }

a:visited {
  color: #0366d6; }

h1, h2, h3 {
  border-bottom: 1px solid #eaecef;
  color: #111;
  /* Darker */ }</style>
</body>
</html>
`