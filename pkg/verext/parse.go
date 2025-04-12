package verext

func Parse(filepath string) (Changelog, error) {

	fileBytes, err := getFile(filepath)

	if err != nil {
		return Changelog{}, err
	}

	return extract(&fileBytes)

}
