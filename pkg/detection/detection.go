package detection

type Detector struct {
}

func (D *Detector) GetAllTld(payload string) ([]string, error) {
	// todo: implement me
	return nil, nil
}

func (D *Detector) WordlistContainsTld(tld string, wordlist map[string]bool) (bool, error) {
	// todo: implement me
	return false, nil
}

func (D *Detector) GetFirstLevelDomain(payload, tld string) ([]string, error) {
	// todo: implement me
	return nil, nil
}

func (D *Detector) GetSubdomain(payload, domain string) ([]string, error) {
	// todo: implement me
	return nil, nil
}

func (D *Detector) GetPath(payload, domain string) ([]string, error) {
	// todo: implement me
	return nil, nil
}
