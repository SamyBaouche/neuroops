{{- define "neuroops.name" -}}
{{- printf "%s-api" .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "neuroops.selectorLabels" -}}
app: {{ include "neuroops.name" . }}
{{- end -}}

{{- define "neuroops.labels" -}}
{{ include "neuroops.selectorLabels" . }}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Values.image.tag | quote }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ printf "%s-%s" .Chart.Name .Chart.Version }}
{{- end -}}
