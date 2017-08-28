package ScrapeLinks

func Format(c chan RefsList, chanLength int) map[string][]Ref {
	result := make(map[string][]Ref)

	for ref := range(c) {
		result[ref.Url] = ref.Links
		if (chanLength <= 1) {
			close(c)
		}
		chanLength = chanLength - 1
	}
	return result;
 }
