package wordmixer

import "fmt"

func ToHtml(puzzle *Puzzle) string {
	const page = `<html>
	<head><style>body { font-family: Courier; }</style></head>
	<body>
	<h1>Find %s</h1>
	<div id="puzzle-div">
	%s
	</div>
	<div  id="solution-div" hidden>
	Location: %s
	</div>
	</body>
	</html>
`
	html := fmt.Sprintf(page, puzzle.word.Word, BoardToString(puzzle, "<br />\n"), SolutionToString(puzzle))
	return html
}
