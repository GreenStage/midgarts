package character

type AttachmentType int

const (
	AttachmentShadow = AttachmentType(iota)
	AttachmentBody
	AttachmentHead

	NumAttachments
)

func (e AttachmentType) String() (att string) {
	switch e {
	case AttachmentShadow:
		att = "AttachmentShadow"
	case AttachmentBody:
		att = "AttachmentBody"
	case AttachmentHead:
		att = "AttachmentHead"
	default:
		panic("unsupported attachment type")
	}

	return att
}

func Attachments() []AttachmentType {
	return []AttachmentType{
		AttachmentShadow,
		AttachmentBody,
		AttachmentHead,
	}
}
