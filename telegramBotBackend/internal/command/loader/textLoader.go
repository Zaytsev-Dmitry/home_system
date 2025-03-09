package loader

import (
	"embed"
	_ "embed"
)

var (
	//go:embed files/tutorial_start_text.txt
	TutorialStartText string

	//go:embed files/note_command_text.txt
	NoteCommandText string

	//go:embed files/profile_command_text.txt
	ProfileCommandText string

	//go:embed files/menu_command_text.txt
	MenuCommandText string

	//go:embed files/expense_accounting_command_text.txt
	ExpenseAccountingCommandText string

	//go:embed files/profile_command_content_text.txt
	ProfileCommandContentText string

	//go:embed files/start_command_text.txt
	StartCommandText string

	//go:embed files/unnecessary_action_text.txt
	UnnecessaryActionText string

	//go:embed files/register_confirm_text.txt
	RegisterConfirmText string

	//go:embed files/register_complete_text.txt
	RegisterCompleteText string

	//go:embed files/note_command_content_text.txt
	NoteCommandContentText string

	//go:embed files/add_note_start_command_text.txt
	AddNoteStartCommandText string

	//go:embed files/add_note_confirm_command_text.txt
	AddNoteConfirmCommandText string

	//go:embed files/add_note_complete_command_text.txt
	AddNoteCompleteCommandText string

	//go:embed files/ebat_ty_loh.jpg
	EnterEmailMistakeMem embed.FS

	//go:embed files/ricardo_milos.gif
	RicardoMilasMemGif embed.FS
)
