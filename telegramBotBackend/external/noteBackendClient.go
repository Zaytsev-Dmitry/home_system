package external

import (
	"fmt"
	noteSpec "github.com/Zaytsev-Dmitry/home_system_open_api/noteServerBackend"
	"strconv"
	"telegramCLient/util"
)

type NoteBackendClient struct {
	noteBackendClientUrk string
}

func NewNoteBackendClient(noteBackendClientUrk string) *NoteBackendClient {
	return &NoteBackendClient{noteBackendClientUrk}
}

func (c *NoteBackendClient) GetAllNotesByAccount(accId int64) noteSpec.NoteResponseList {
	var result noteSpec.NoteResponseList
	resp, err := client.Get(c.noteBackendClientUrk + "/note/" + strconv.Itoa(int(accId)))
	if err != nil {
		fmt.Println(err)
	}
	util.ParseResponseToStruct(resp, &result)
	return result
}

func (c *NoteBackendClient) Save(source noteSpec.CreateNoteRequest) noteSpec.NoteResponse {
	var result noteSpec.NoteResponse
	resp, err := post(source, c.noteBackendClientUrk+"/note")
	if err != nil {
		fmt.Println(err)
	}
	util.ParseResponseToStruct(resp, &result)
	return result
}
