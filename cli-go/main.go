package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var baseURL string

var blue = lipgloss.Color("39")
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
	Border(lipgloss.RoundedBorder()).
	Padding(1, 2).
	BorderForeground(blue)

type model struct {
	choices []string
	cursor  int
	result  string
}

func callEndpoint(path string) string {
	resp, err := http.Get(baseURL + path)
	if err != nil {
		return warningStyle.Render("Error: " + err.Error())
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return fmt.Sprintf(
		"%s\n\nendpoint: %s\nstatus: %s\nresponse: %s",
		successStyle.Render("Neuro Ops response"),
		path,
		resp.Status,
		string(body),
	)
}

func initialModel() model {
	return model{
		choices: []string{
			"health check",
			"readiness check",
			"simulate cpu load",
			"trigger failure",
			"exit",
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

		case "enter":

			switch m.cursor {

			case 0:
				m.result = callEndpoint("/health")

			case 1:
				m.result = callEndpoint("/ready")

			case 2:
				m.result = callEndpoint("/load")

			case 3:
				m.result = callEndpoint("/fail")

			case 4:
				return m, tea.Quit
			}
		}
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
		"┌─ cloud ───────────────┬─ kubernetes ──────────┬─ ai remediation ───────────┐",
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

	s += "            ai-powered kubernetes observability and remediation\n"

	s += mutedStyle.Render(
		"             dev · cloud · github.com/SamyBaouche/neuroops",
	) + "\n\n"

	s += warningStyle.Render(
		"What do you want to inspect ?",
	) + "\n\n"

	for i, choice := range m.choices {

		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n"

	s += mutedStyle.Render(
		"Use ↑/↓ and Enter. Press q to quit.",
	) + "\n"

	return s
}

func runTUI() {

	// Alt screen prevents redraw noise/flooding in terminals like PowerShell.
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

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
