package token

type Type int

const (
	Headline Type = iota
	BulletList
	NumberList
	Text
)
