package types

// Mapping defines a mapping from a short link to a long link.
type Mapping struct {
	ShortLink string `json:"shortLink" dynamodbav:"shortLink"`
	LongLink  string `json:"longLink" dynamodbav:"longLink"`
}
