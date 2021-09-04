package sitemap

func TraverseUrl(urlStart string, maxdepth int, maxLimit int) []string {
	visited := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStart: struct{}{},
	}
	count := 0

	for i := 0; i < maxdepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 || count > maxLimit {
			break
		}
		for u, _ := range q {
			if _, ok := visited[u]; ok {
				continue
			}
			visited[u] = struct{}{}
			baselink, alllinks := GenerateBaseUrl(&u)
			sitelinks := GenerateSiteLinks(alllinks, baselink)
			//fmt.Println(baselink)
			for _, l := range sitelinks {
				//fmt.Println(l)
				nq[l] = struct{}{}
				count = count + 1
			}
		}
	}
	var result []string
	for u, _ := range visited {
		result = append(result, u)
	}
	return result
}
