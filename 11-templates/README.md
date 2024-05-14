1.2 - present {{ . }} that prints the value, explain {{- and -}}

1.3 - present {{range $index, $value := . }} ... {{ end }} to iterate over a slice and extract index + value

2 - introduce {{range . }} ... {{ end }} and accessors with {{ .Name }}

3 - present {{ if func }} {{ else }} {{ end }} 

4 - {{ if . }} is executed if len(x) > 0 (not zero value, not empty, not nil) 

5 - {{ len . }} to return the length of a slice + using two sources for data

6 - funcmap  + calling methods. And some CSS.
