package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

const (
	bottlerocketUserData = `
[settings.kubernetes]
api-server = "{{.Cluster.Endpoint}}"
{{if .Cluster.CABundle}}{{if len .Cluster.CABundle}}cluster-certificate = "{{.Cluster.CABundle}}"{{end}}{{end}}
cluster-name = "{{if .Cluster.Name}}{{.Cluster.Name}}{{end}}"
{{if .Constraints.Labels }}[settings.kubernetes.node-labels]{{ end }}
{{ range $Key, $Value := .Constraints.Labels }}"{{ $Key }}" = "{{ $Value }}"
{{ end }}
{{if .Constraints.Taints }}[settings.kubernetes.node-taints]{{ end }}
{{ range $Taint := .Constraints.Taints }}"{{ $Taint.Key }}" = "{{ $Taint.Value}}:{{ $Taint.Effect }}"
{{ end }}
{{.Cluster.Mergatroid}}
`
)

type cluster struct {
	Endpoint string
	CABundle *string
	Name     string
}

func (c cluster) Mergatroid() string {
	return "MergatroidWasHere"
}

type constraints struct {
	Labels []string
	Taints map[string]string
}

type Provisioner struct {
	Cluster     cluster
	Constraints constraints
}

var empty string

func main() {
	fmt.Printf("-> 'template' running\n")
	t := template.Must(template.New("").Parse(bottlerocketUserData))
	if err := t.Execute(os.Stdout, Provisioner{
		Cluster: cluster{
			Endpoint: "",
			CABundle: &empty,
			Name:     "",
		},
		Constraints: constraints{
			Labels: []string{},
			Taints: map[string]string{},
		},
	}); err != nil {
		log.Fatalf("execute failed, %v", err)
	}
}
