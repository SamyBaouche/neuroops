package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var baseURL string

var purple = lipgloss.Color("99")
var gray = lipgloss.Color("245")
var orange = lipgloss.Color("214")

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

var resultBoxStyle = lipgloss.NewStyle().
	Padding(0, 0)

type model struct {
	choices []string
	cursor  int
	result  string
	loading bool
}

type endpointResultMsg struct {
	content string
}

func endpointCmd(path string) tea.Cmd {
	return func() tea.Msg {
		return endpointResultMsg{content: callEndpoint(path)}
	}
}

func callEndpoint(path string) string {
	resp, err := http.Get(baseURL + path)
	if err != nil {
		return warningStyle.Render("Error: " + err.Error())
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	bodyText := strings.TrimSpace(string(body))
	if bodyText == "" {
		bodyText = "(empty response body)"
	}

	return fmt.Sprintf(
		"%s\nendpoint: %s\nstatus: %s\n\nresponse:\n%s",
		successStyle.Render("Neuro Ops response"),
		path,
		resp.Status,
		bodyText,
	)
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
		result: "waiting for action...",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", "ctrl+m":
			if m.loading {
				return m, nil
			}

			switch m.cursor {

			case 0:
				m.loading = true
				m.result = mutedStyle.Render("Calling /health ...")
				return m, endpointCmd("/health")

			case 1:
				m.loading = true
				m.result = mutedStyle.Render("Calling /ready ...")
				return m, endpointCmd("/ready")

			case 2:
				m.loading = true
				m.result = mutedStyle.Render("Calling /load ...")
				return m, endpointCmd("/load")

			case 3:
				m.loading = true
				m.result = mutedStyle.Render("Calling /fail ...")
				return m, endpointCmd("/fail")

			case 4:
				return m, tea.Quit
			}
		}

	case endpointResultMsg:
		m.loading = false
		m.result = msg.content
	}

	// Keep the cursor index in bounds even if the terminal sends unexpected key sequences.
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

	s += mutedStyle.Render(
		"             dev · cloud · github.com/SamyBaouche/neuroops",
	) + "\n\n"

	s += warningStyle.Render(
		"What do you want to inspect ?",
	) + "\n\n"

	for i, choice := range m.choices {
		if m.cursor == i {
			s += fmt.Sprintf("%s\n", successStyle.Render("> "+choice))
			continue
		}

		s += fmt.Sprintf("%s\n", mutedStyle.Render("- "+choice))
	}

	s += "\n"

	s += mutedStyle.Render(
		"Use up/down and Enter. Press q to quit.",
	) + "\n"

	if m.loading {
		s += "\n" + warningStyle.Render("Loading...") + "\n"
	}

	s += "\n" + resultBoxStyle.Render(successStyle.Render("Result")+"\n\n"+m.result) + "\n"

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
