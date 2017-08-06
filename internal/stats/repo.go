package stats

type Repo struct {
	apiKey string
	listId string
}

func NewRepo(apiKey, listId string) *Repo {
	return &Repo{apiKey: apiKey, listId: listId}
}

func (r *Repo) MemberCount() (int, error) {
	// Use go's http implementations & mailchimp's rest apis to get the count for you list in this method.
	return 0, nil
}