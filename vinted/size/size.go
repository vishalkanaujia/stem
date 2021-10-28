package size

type Size string

const (
	Small   Size = "S"
	Medium  Size = "M"
	Large   Size = "L"
	Unknown Size = "Unknown"
)

func (s Size) Size(inSize string) Size {
	switch inSize {
	case "S":
		return Small
	case "M":
		return Medium
	case "L":
		return Large
	default:
		return Unknown
	}
}
