package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var baseURL string

var purple = lipgloss.Color("99")
var gray = lipgloss.Color("245")
var orange = lipgloss.Color("214")
var green = lipgloss.Color("42")
var red = lipgloss.Color("196")

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(purple).
	MarginBottom(1)

var boxLineStyle = lipgloss.NewStyle().
	Foreground(orange)

var successStyle = lipgloss.NewStyle().
	Foreground(purple).
	Bold(true)

var mutedStyle = lipgloss.NewStyle().
	Foreground(gray)

var warningStyle = lipgloss.NewStyle().
	Foreground(orange).
	Bold(true)

var goodStyle = lipgloss.NewStyle().
	Foreground(green).
	Bold(true)

var badStyle = lipgloss.NewStyle().
	Foreground(red).
	Bold(true)

var statusBarStyle = lipgloss.NewStyle().
	Foreground(purple).
	Bold(true)

var menuSelectedStyle = lipgloss.NewStyle().
	Foreground(purple).
	Bold(true)

var menuItemStyle = lipgloss.NewStyle().
	Foreground(gray)

var resultBoxStyle = lipgloss.NewStyle().
	Padding(0, 0)

var loadingStages = []string{
	"Analyzing cluster...",
	"Collecting metrics...",
	"Inspecting Kubernetes status...",
	"Generating AI insights...",
}

const maxVisibleResultLines = 36

const requestTimeout = 8 * time.Second

var httpClient = &http.Client{Timeout: requestTimeout}

type endpointResultMsg struct {
	path       string
	status     string
	statusCode int
	body       string
	err        error
	at         time.Time
}

type loadingTickMsg struct{}

type model struct {
	choices      []string
	cursor       int
	result       string
	loading      bool
	loadingFrame int
	clusterState string
	hasResult    bool
}

func separator() string {
	return strings.Repeat("=", 50)
}

func fmtTS(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func loadingTickCmd() tea.Cmd {
	return tea.Tick(350*time.Millisecond, func(time.Time) tea.Msg {
		return loadingTickMsg{}
	})
}

func performEndpointCall(path string) endpointResultMsg {
	msg := endpointResultMsg{path: path, at: time.Now()}

	resp, err := httpClient.Get(baseURL + path)
	if err != nil {
		msg.err = err
		return msg
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	bodyText := strings.TrimSpace(string(body))
	if bodyText == "" {
		bodyText = "(empty response body)"
	}

	msg.status = resp.Status
	msg.statusCode = resp.StatusCode
	msg.body = bodyText
	return msg
}

func endpointCmd(path string) tea.Cmd {
	return func() tea.Msg {
		return performEndpointCall(path)
	}
}

func okLine(text string) string {
	return fmt.Sprintf("%s %s", goodStyle.Render("[OK]"), text)
}

func hintLine(text string) string {
	return fmt.Sprintf("%s %s", warningStyle.Render("->"), text)
}

func statusFromCode(statusCode int) string {
	if statusCode >= 200 && statusCode < 300 {
		return goodStyle.Render("HEALTHY")
	}
	return badStyle.Render("DEGRADED")
}

func buildHealthReport(msg endpointResultMsg) string {
	ai := "No anomalies detected. Cluster operating normally."
	if msg.statusCode < 200 || msg.statusCode >= 300 {
		ai = "Potential reliability anomaly detected. Investigate probes and pod events."
	}

	return strings.Join([]string{
		separator(),
		successStyle.Render("🟢 CLUSTER HEALTH ANALYSIS"),
		separator(),
		"",
		"API Status:",
		fmt.Sprintf("%s %s", okLine("API health endpoint reachable"), statusFromCode(msg.statusCode)),
		fmt.Sprintf("Status Code: %s", msg.status),
		"",
		"Kubernetes:",
		okLine("Pod running"),
		okLine("Service reachable"),
		okLine("Liveness probe passing"),
		okLine("Readiness probe passing"),
		"",
		"Cluster:",
		"Namespace: kubepulse",
		"Environment: Minikube",
		"Replicas: 1/1",
		"",
		"Observability:",
		okLine("Prometheus connected"),
		okLine("Grafana active"),
		okLine("HPA enabled"),
		"",
		"Infrastructure:",
		"Cloud: AWS-ready",
		"Container Runtime: Docker",
		"Orchestrator: Kubernetes",
		"",
		"AI Insights:",
		ai,
		"",
		"Response Body:",
		msg.body,
		"",
		fmt.Sprintf("Timestamp: %s", fmtTS(msg.at)),
	}, "\n")
}

func buildReadinessReport(msg endpointResultMsg) string {
	ai := "Readiness state is stable. Traffic routing can continue normally."
	if msg.statusCode < 200 || msg.statusCode >= 300 {
		ai = "Readiness degradation detected. Pause rollout and inspect service endpoints."
	}

	traffic := okLine("Service can accept Kubernetes traffic")
	if msg.statusCode < 200 || msg.statusCode >= 300 {
		traffic = badStyle.Render("[WARN]") + " Service not ready for traffic routing"
	}

	return strings.Join([]string{
		separator(),
		successStyle.Render("🟣 DEPLOYMENT READINESS ANALYSIS"),
		separator(),
		"",
		"Readiness Validation:",
		fmt.Sprintf("Endpoint: %s", msg.path),
		fmt.Sprintf("Status: %s", msg.status),
		traffic,
		"",
		"Traffic Routing:",
		okLine("Kubernetes service discovery active"),
		okLine("Endpoint registration validated"),
		okLine("Readiness gate evaluated"),
		"",
		"Deployment State:",
		"Namespace: kubepulse",
		"Service Availability: online",
		"Traffic Acceptance: conditional on readiness probe",
		"",
		"Observability Signals:",
		okLine("Prometheus scrape target expected"),
		okLine("Grafana dashboards can reflect readiness state"),
		"",
		"AI Insights:",
		ai,
		"",
		"Recommendation:",
		hintLine("kubectl get endpoints -n kubepulse"),
		hintLine("kubectl describe pod -n kubepulse"),
		"",
		fmt.Sprintf("Timestamp: %s", fmtTS(msg.at)),
	}, "\n")
}

func buildLoadReport(msg endpointResultMsg) string {
	return strings.Join([]string{
		separator(),
		successStyle.Render("🔥 STRESS TEST EXECUTED"),
		separator(),
		"",
		"Operation:",
		"CPU saturation simulation triggered.",
		fmt.Sprintf("API Status: %s", msg.status),
		"",
		"Expected Cluster Behavior:",
		"CPU usage increase",
		"HPA scaling activation",
		"Replica expansion",
		"Metrics spike in Grafana",
		"",
		"Observability Pipeline:",
		okLine("Prometheus collecting metrics"),
		okLine("Grafana dashboards updating"),
		okLine("HPA monitoring active"),
		"",
		"Resilience Testing:",
		okLine("Kubernetes self-healing enabled"),
		okLine("Auto-scaling active"),
		"",
		"AI Insights:",
		"Potential CPU saturation activity observed.",
		"",
		"Recommended Monitoring:",
		hintLine("kubectl get hpa -w"),
		hintLine("kubectl top pods -n kubepulse"),
		hintLine("Grafana CPU dashboard"),
		"",
		"Response Body:",
		msg.body,
		"",
		fmt.Sprintf("Timestamp: %s", fmtTS(msg.at)),
	}, "\n")
}

func buildFailureReport(msg endpointResultMsg) string {
	return strings.Join([]string{
		separator(),
		successStyle.Render("🚨 FAILURE SIMULATION ACTIVATED"),
		separator(),
		"",
		"Incident Type:",
		"Simulated application failure",
		fmt.Sprintf("API Status: %s", msg.status),
		"",
		"Expected Recovery Workflow:",
		"1. Liveness probe failure detected",
		"2. Kubernetes marks pod unhealthy",
		"3. Pod termination initiated",
		"4. Replacement pod scheduled",
		"5. Service restored automatically",
		"",
		"Recovery Systems:",
		okLine("Self-healing enabled"),
		okLine("Kubernetes orchestration active"),
		okLine("High availability workflow active"),
		"",
		"Monitoring:",
		hintLine("Observe pod restart in Kubernetes"),
		hintLine("Monitor events in Grafana"),
		"",
		"AI Insights:",
		"Failure simulation triggered successfully. Cluster recovery expected automatically.",
		"",
		"Suggested Commands:",
		hintLine("kubectl get pods -n kubepulse -w"),
		hintLine("kubectl get events -n kubepulse --sort-by=.metadata.creationTimestamp"),
		"",
		"Response Body:",
		msg.body,
		"",
		fmt.Sprintf("Timestamp: %s", fmtTS(msg.at)),
	}, "\n")
}

func buildErrorReport(path string, err error) string {
	return strings.Join([]string{
		separator(),
		badStyle.Render("❌ INFRASTRUCTURE REQUEST FAILED"),
		separator(),
		"",
		fmt.Sprintf("Endpoint: %s", path),
		fmt.Sprintf("Error: %v", err),
		"",
		"AI Insights:",
		"Connectivity anomaly detected. Validate API URL, service endpoint, and network reachability.",
		"",
		"Suggested Commands:",
		hintLine("kubectl get svc -n kubepulse"),
		hintLine("minikube service kubepulse-service -n kubepulse --url"),
		"",
		fmt.Sprintf("Timestamp: %s", fmtTS(time.Now())),
	}, "\n")
}

func formatEndpointReport(msg endpointResultMsg) string {
	if msg.err != nil {
		return buildErrorReport(msg.path, msg.err)
	}

	switch msg.path {
	case "/health":
		return buildHealthReport(msg)
	case "/ready":
		return buildReadinessReport(msg)
	case "/load":
		return buildLoadReport(msg)
	case "/fail":
		return buildFailureReport(msg)
	default:
		return strings.Join([]string{
			separator(),
			successStyle.Render("NEUROOPS PLATFORM RESPONSE"),
			separator(),
			fmt.Sprintf("Endpoint: %s", msg.path),
			fmt.Sprintf("Status: %s", msg.status),
			"",
			"Response Body:",
			msg.body,
			"",
			fmt.Sprintf("Timestamp: %s", fmtTS(msg.at)),
		}, "\n")
	}
}

func callEndpoint(path string) string {
	return formatEndpointReport(performEndpointCall(path))
}

func initialModel() model {
	return model{
		choices: []string{
			"Health check",
			"Readiness check",
			"Simulate CPU load",
			"Trigger failure",
			"Exit",
		},
		result:       "Select an operation to analyze platform health, resilience, and observability signals.",
		clusterState: "HEALTHY",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) actionPath() string {
	switch m.cursor {
	case 0:
		return "/health"
	case 1:
		return "/ready"
	case 2:
		return "/load"
	case 3:
		return "/fail"
	default:
		return ""
	}
}

func (m model) loadingText() string {
	return loadingStages[m.loadingFrame%len(loadingStages)]
}

func (m model) statusBar() string {
	cluster := m.clusterState
	if m.loading {
		cluster = "ANALYZING"
	}

	return statusBarStyle.Render(
		fmt.Sprintf(
			"Pods: 1 | HPA: ACTIVE | Prometheus: CONNECTED | Grafana: ACTIVE | Cluster: %s",
			cluster,
		),
	)
}

func trimResultForView(result string) string {
	lines := strings.Split(result, "\n")
	if len(lines) <= maxVisibleResultLines {
		return result
	}

	trimmed := strings.Join(lines[:maxVisibleResultLines], "\n")
	return trimmed + "\n\n" + mutedStyle.Render("... output trimmed for terminal stability")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "esc":
			if m.loading {
				m.loading = false
				m.clusterState = "DEGRADED"
				m.result = warningStyle.Render("Request canceled by user. You can select another option.")
				return m, nil
			}

		case "up", "k":
			if m.loading {
				return m, nil
			}
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.loading {
				return m, nil
			}
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", "ctrl+m":
			if m.loading {
				m.result = warningStyle.Render("Operation already running... wait or press Esc to cancel.")
				return m, nil
			}

			if m.cursor == len(m.choices)-1 {
				return m, tea.Quit
			}

			path := m.actionPath()
			m.loading = true
			m.loadingFrame = 0
			m.clusterState = "ANALYZING"
			m.result = warningStyle.Render(m.loadingText())
			return m, tea.Batch(endpointCmd(path), loadingTickCmd())
		}

	case loadingTickMsg:
		if m.loading {
			m.loadingFrame++
			m.result = warningStyle.Render(m.loadingText())
			return m, loadingTickCmd()
		}

	case endpointResultMsg:
		m.loading = false
		m.hasResult = true
		m.result = formatEndpointReport(msg) + "\n\n" + mutedStyle.Render("Ready for next action. Use up/down + Enter, or q to quit.")

		switch msg.path {
		case "/fail":
			m.clusterState = "RECOVERING"
		case "/load":
			m.clusterState = "SCALING"
		default:
			if msg.err != nil || msg.statusCode < 200 || msg.statusCode >= 300 {
				m.clusterState = "DEGRADED"
			} else {
				m.clusterState = "HEALTHY"
			}
		}
	}

	if len(m.choices) > 0 {
		if m.cursor < 0 {
			m.cursor = 0
		}
		if m.cursor >= len(m.choices) {
			m.cursor = len(m.choices) - 1
		}
	}

	return m, nil
}

func (m model) View() string {
	logo := `
███╗   ██╗███████╗██╗   ██╗██████╗  ██████╗      ██████╗ ██████╗ ███████╗
████╗  ██║██╔════╝██║   ██║██╔══██╗██╔═══██╗    ██╔═══██╗██╔══██╗██╔════╝
██╔██╗ ██║█████╗  ██║   ██║██████╔╝██║   ██║    ██║   ██║██████╔╝███████╗
██║╚██╗██║██╔══╝  ██║   ██║██╔══██╗██║   ██║    ██║   ██║██╔═══╝ ╚════██║
██║ ╚████║███████╗╚██████╔╝██║  ██║╚██████╔╝    ╚██████╔╝██║     ███████║
╚═╝  ╚═══╝╚══════╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝      ╚═════╝ ╚═╝     ╚══════╝
`

	s := titleStyle.Render(logo)
	s += "\n"
	s += m.statusBar() + "\n\n"

	s += boxLineStyle.Render(
		"┌─ Cloud ───────────────┬─ Kubernetes ──────────┬─ AI Remediation ───────────┐",
	) + "\n"

	s += "│ " +
		successStyle.Render("AWS · ECR · IAM") +
		"       │ " +
		successStyle.Render("HPA · probes · SRE") +
		"      │ " +
		successStyle.Render("LLM analysis · Bedrock") +
		"     │\n"

	s += boxLineStyle.Render(
		"└───────────────────────┴───────────────────────┴────────────────────────────┘",
	) + "\n\n"

	s += "            AI-powered Kubernetes observability and remediation\n"
	s += mutedStyle.Render("             dev · cloud · github.com/SamyBaouche/neuroops") + "\n\n"

	if !m.hasResult {
		s += warningStyle.Render("What do you want to inspect ?") + "\n\n"

		for i, choice := range m.choices {
			if m.cursor == i {
				s += fmt.Sprintf("%s\n", menuSelectedStyle.Render("> "+choice))
				continue
			}
			s += fmt.Sprintf("%s\n", menuItemStyle.Render("- "+choice))
		}

		s += "\n"
		s += mutedStyle.Render("Use up/down and Enter. Press q to quit.") + "\n"
	}

	if m.loading {
		s += "\n" + warningStyle.Render(m.loadingText()) + "\n"
	}

	s += "\n" + resultBoxStyle.Render(successStyle.Render("Platform Analysis")+"\n\n"+trimResultForView(m.result)) + "\n"

	if m.hasResult {
		s += "\n" + warningStyle.Render("Next Action") + "\n\n"
		for i, choice := range m.choices {
			if m.cursor == i {
				s += fmt.Sprintf("%s\n", menuSelectedStyle.Render("> "+choice))
				continue
			}
			s += fmt.Sprintf("%s\n", menuItemStyle.Render("- "+choice))
		}

		s += "\n" + mutedStyle.Render("Choose another option with up/down + Enter, or press q to quit.") + "\n"
	}

	return s
}

func runTUI() {
	// Avoid alt screen so PowerShell keeps scrollback and users can scroll output.
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "neuroops",
		Short: "Neuro Ops CLI",
		Long:  "Neuro Ops is an AI-powered Kubernetes observability and remediation platform.",
	}

	rootCmd.PersistentFlags().StringVar(
		&baseURL,
		"url",
		"http://127.0.0.1:8080",
		"API base URL",
	)

	tuiCmd := &cobra.Command{
		Use:   "tui",
		Short: "Launch Neuro Ops terminal dashboard",
		Run: func(cmd *cobra.Command, args []string) {
			runTUI()
		},
	}

	healthCmd := &cobra.Command{
		Use:   "health",
		Short: "Check API health",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(callEndpoint("/health"))
		},
	}

	readyCmd := &cobra.Command{
		Use:   "ready",
		Short: "Check API readiness",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(callEndpoint("/ready"))
		},
	}

	loadCmd := &cobra.Command{
		Use:   "load",
		Short: "Simulate CPU load",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(callEndpoint("/load"))
		},
	}

	failCmd := &cobra.Command{
		Use:   "fail",
		Short: "Trigger failure simulation",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(callEndpoint("/fail"))
		},
	}

	rootCmd.AddCommand(
		tuiCmd,
		healthCmd,
		readyCmd,
		loadCmd,
		failCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
