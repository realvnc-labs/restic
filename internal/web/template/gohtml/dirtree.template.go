// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots
package gohtml

var dirtree = `{{ if .Node.IsDir }}
    <li>
        <label for="{{ .Node.Name }}"><a href="{{ .NodesContext.Url }}">{{ .Node.Name }}</a></label>
        {{ if .Node.IsExpanded }}
        <input type="checkbox" checked="checked" id="{{ .Node.Name }}">
        {{ else }}
            <a href="{{ .NodesContext.Url }}" class="dir-tree-expanded">&nbsp;</a>
        {{end}}
        <ol>
            {{ .Node.Nodes.Render .NodesContext }}
        </ol>
    </li>
{{ else }}
    <li class="file">
        <a href="/dump?id={{.NodesContext.Params.id}}&path={{ .Node.Path }}&size={{ .Node.Size }}">{{ .Node.Name }}</a>
    </li>
{{ end }}
`