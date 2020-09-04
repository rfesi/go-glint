package dynamiccli

// Components are the individual items that are rendered within a document.
type Component interface {
	// Body
	Body() Component

	// Render is called to render this element.
	//
	// The rows/cols given are advisory. If the cols are ignored, the return
	// value may be wrapped or truncated (depending on layout settings). This
	// behavior may be undesirable and so it is recommended you remain within
	// the advisory amounts. If the rows are ignored, the output will be
	// truncated.
}

// ComponentTerminalSizer can be implemented to receive the terminal size.
// See the function docs for more information.
type ComponentTerminalSizer interface {
	Component

	// SetTerminalSize is called with the full terminal size. This may
	// exceed the size given by Render in certain cases. This will be called
	// before Render and Layout.
	SetTerminalSize(rows, cols uint)
}

type ComponentLayout interface {
	Component

	Layout() *Layout
}

type terminalComponent struct{}

func (terminalComponent) Body() Component { return nil }
