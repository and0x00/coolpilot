package core

import (
	"context"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

type Runner struct {
	options *Options
}

// NewRunner init runner
func NewRunner(opts *Options) (*Runner, error) {
	runner := &Runner{options: opts}
	return runner, nil
}

func (r *Runner) Run(ctx context.Context) error {

	u := launcher.New().
		Headless(false). // disable headless mode
		MustLaunch()
	browser := rod.New().ControlURL(u).Timeout(time.Minute)
	// browser := rod.New().Timeout(time.Minute)
	err := browser.Connect()
	if err != nil {
		fmt.Printf("error connecting to the browser: %v\n", err)
		return nil
	}
	defer browser.Close()

	page, err := stealth.Page(browser)
	if err != nil {
		fmt.Printf("error creating stealth page: %v\n", err)
		return nil
	}
	defer page.Close()

	err = page.Navigate("http://copilot.microsoft.com/")
	if err != nil {
		fmt.Printf("Erro ao navegar: %v\n", err)
		return nil
	}

	select {}
}
