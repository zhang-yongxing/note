package study_note

type AddNoteForm struct {
	NoteName string `json:"note_name" binding:"required"`
	NoteDes string `json:"note_des"`
	NoteContent string `json:"note_content" binding:"required"`
}

func (nf AddNoteForm) AddNoteFormVerify() map[string]string{
	err := make(map[string]string)
	status := false
	NoteNameLen := len(nf.NoteName)
	NoteDesLen := len(nf.NoteDes)
	if NoteNameLen > 50{
		err["note_name"] = "Length greater than 50"
	}
	if  NoteDesLen > 50{
		err["note_des"] = "Length greater than 50"
		status = true
	}
	if status == true{
		return err
	}
	return nil
}

type AlterNoteForm struct {
	NoteName string `json:"note_name" binding:"required"`
	NoteDes string `json:"note_des"`
	NoteContent string `json:"note_content" binding:"required"`
}

func (nf AlterNoteForm) AlterNoteFormVerify() map[string]string{
	err := make(map[string]string)
	status := false
	NoteNameLen := len(nf.NoteName)
	NoteDesLen := len(nf.NoteDes)
	if NoteNameLen > 50{
		err["note_name"] = "Length greater than 50"
	}
	if  NoteDesLen > 50{
		err["note_des"] = "Length greater than 50"
		status = true
	}
	if status == true{
		return err
	}
	return nil
}
