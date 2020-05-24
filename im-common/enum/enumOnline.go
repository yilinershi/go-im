package enum

type EnumOnline int32

const (
	OnLine  EnumOnline = 0
	OffLine EnumOnline = 1
)

func (p EnumOnline) String() string {
	switch p {
	case OnLine:
		return "OnLine"
	case OffLine:
		return "OffLine"
	default:
		return "UNKNOWN"
	}
}
