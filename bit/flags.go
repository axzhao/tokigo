package bit

func setFlag(flags *uint8, isSet bool, bitFlag uint8) {
	if isSet {
		*flags = *flags | bitFlag
	}
}

func isFlag(flags uint8, typeFlag uint8) bool {
	result := flags & typeFlag
	return result == typeFlag
}

func DecodeFlags(flags uint8, args ...*bool) {

	if len(args) > 8 {
		return
	}

	for i, obj := range args {
		*obj = isFlag(flags, (uint8(1) << uint8(i)))
	}
}

func EncodeFlags(flags *uint8, args ...bool) {

	if len(args) > 8 {
		return
	}

	*flags = 0
	for i, obj := range args {
		setFlag(flags, obj, (uint8(1) << uint8(i)))
	}
}
