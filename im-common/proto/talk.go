package proto

type TalkMessageInput struct {
	From    string
	To      string
	Content string
}

type TalkMessageSync struct {
	From    string
	Content string
}
