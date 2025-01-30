package loader

import _ "embed"

var (
	//go:embed files/tutorial_start_text.txt
	StartText string

	//go:embed files/note_command_desc_text.txt
	NoteCommandDescText string

	//go:embed files/profile_command_desc_text.txt
	ProfileCommandDescText string

	//go:embed files/menu_command_desc_text.txt
	MenuCommandDescText string

	//go:embed files/expense_accounting_command_desc_text.txt
	ExpenseAccCommandDescText string
)
