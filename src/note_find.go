package note

func FindNote(note string) (int, string) {
	paths := ListNotePaths()
	notes := ListNoteNames()
	noteNoExt := note + ".md"
	for i, n := range notes {
		if noteNoExt == n {
			return i, paths[i]
		}
	}
	return -1, ""
}
