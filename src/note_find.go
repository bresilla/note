package note

func FindNote(note string) (int, string) {
	notes_dir := ListNotes_Dir()
	notes_name := ListNotes_Name()
	for i, n := range notes_name {
		if note == n {
			//	fmt.Println(i, n)
			//	fmt.Println(notes_dir[i])
			return i, notes_dir[i]
		}
	}
	return -1, ""
}
