package context

type ArticlePost struct {
	UniqueTitle  bool
	UniqueSource bool
}

func NewArticlePost(t, s bool) *ArticlePost {
	return &ArticlePost{UniqueTitle: t, UniqueSource: s}
}
