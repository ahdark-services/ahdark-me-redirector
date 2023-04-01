package wpgo

type Context string

const (
	ContextView  Context = "view"
	ContextEmbed Context = "embed"
	ContextEdit  Context = "edit"
)

type GetPostParams struct {
	Context  Context `url:"context,omitempty"`
	Password string  `url:"password,omitempty"`
}

type ListPostParams struct {
	// A comma-separated list of post IDs to retrieve.
	// Default value is empty.
	Context           Context `url:"context,omitempty"`
	Page              int     `url:"page,omitempty"`
	PerPage           int     `url:"per_page,omitempty"`
	Search            string  `url:"search,omitempty"`
	After             string  `url:"after,omitempty"`
	Author            []int   `url:"author,omitempty"`
	AuthorExclude     []int   `url:"author_exclude,omitempty"`
	Before            string  `url:"before,omitempty"`
	Exclude           []int   `url:"exclude,omitempty"`
	Include           []int   `url:"include,omitempty"`
	Offset            int     `url:"offset,omitempty"`
	Order             Order   `url:"order,omitempty"`
	OrderBy           string  `url:"orderby,omitempty"`
	Slug              string  `url:"slug,omitempty"`
	Status            string  `url:"status,omitempty"`
	TaxRelation       string  `url:"tax_relation,omitempty"`
	Categories        []int   `url:"categories,omitempty"`
	CategoriesExclude []int   `url:"categories_exclude,omitempty"`
	Tags              []int   `url:"tags,omitempty"`
	TagsExclude       []int   `url:"tags_exclude,omitempty"`
	Sticky            bool    `url:"sticky,omitempty"`
}

type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)
