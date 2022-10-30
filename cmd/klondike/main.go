package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nerg4l/gambol/internal/klondiketui"
	"os"
)

func main() {
	p := tea.NewProgram(klondiketui.New(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Klondike experienced an error: %v", err)
		os.Exit(1)
	}
}
