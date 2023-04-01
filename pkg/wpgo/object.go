package wpgo

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status int `json:"status"`
	} `json:"data"`
}

type PostList []Post

type Post struct {
	ID            int64         `json:"id"`
	Date          string        `json:"date"`
	DateGmt       string        `json:"date_gmt"`
	GUID          GUID          `json:"guid"`
	Modified      string        `json:"modified"`
	ModifiedGmt   string        `json:"modified_gmt"`
	Slug          string        `json:"slug"`
	Status        PostStatus    `json:"status"`
	Type          PostType      `json:"type"`
	Link          string        `json:"link"`
	Title         GUID          `json:"title"`
	Content       Content       `json:"content"`
	Excerpt       Content       `json:"excerpt"`
	Author        int64         `json:"author"`
	FeaturedMedia int64         `json:"featured_media"`
	CommentStatus CommentStatus `json:"comment_status"`
	PingStatus    PingStatus    `json:"ping_status"`
	Sticky        bool          `json:"sticky"`
	Template      string        `json:"template"`
	Format        PostFormat    `json:"format"`
	Meta          []interface{} `json:"meta"`
	Categories    []int64       `json:"categories"`
	Tags          []int64       `json:"tags"`
	Links         Links         `json:"_links"`
}

type PostStatus string

const (
	PostStatusPublish PostStatus = "publish"
	PostStatusFuture  PostStatus = "future"
	PostStatusDraft   PostStatus = "draft"
	PostStatusPending PostStatus = "pending"
	PostStatusPrivate PostStatus = "private"
)

type PingStatus string

const (
	PingStatusOpen  PingStatus = "open"
	PingStatusClose PingStatus = "closed"
)

type CommentStatus string

const (
	CommentStatusOpen  CommentStatus = "open"
	CommentStatusClose CommentStatus = "closed"
)

type PostType string

const (
	PostTypePost PostType = "post"
)

type Content struct {
	Rendered  string `json:"rendered"`
	Protected bool   `json:"protected"`
}

type GUID struct {
	Rendered string `json:"rendered"`
}

type Links struct {
	Self            []About          `json:"self"`
	Collection      []About          `json:"collection"`
	About           []About          `json:"about"`
	Author          []AuthorElement  `json:"author"`
	Replies         []AuthorElement  `json:"replies"`
	VersionHistory  []VersionHistory `json:"version-history"`
	WpAttachment    []About          `json:"wp:attachment"`
	WpTerm          []Term           `json:"wp:term"`
	Curies          []Cury           `json:"curies"`
	WpFeaturedmedia []AuthorElement  `json:"wp:featuredmedia,omitempty"`
}

type About struct {
	Href string `json:"href"`
}

type AuthorElement struct {
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}

type Cury struct {
	Name      string `json:"name"`
	Href      string `json:"href"`
	Templated bool   `json:"templated"`
}

type VersionHistory struct {
	Count int64  `json:"count"`
	Href  string `json:"href"`
}

type Term struct {
	Taxonomy   Taxonomy `json:"taxonomy"`
	Embeddable bool     `json:"embeddable"`
	Href       string   `json:"href"`
}

type Taxonomy string

const (
	Category Taxonomy = "category"
	PostTag  Taxonomy = "post_tag"
)

type PostFormat string

const (
	PostFormatStandard PostFormat = "standard"
	PostFormatAside    PostFormat = "aside"
	PostFormatChat     PostFormat = "chat"
	PostFormatGallery  PostFormat = "gallery"
	PostFormatLink     PostFormat = "link"
	PostFormatImage    PostFormat = "image"
	PostFormatQuote    PostFormat = "quote"
	PostFormatStatus   PostFormat = "status"
	PostFormatVideo    PostFormat = "video"
	PostFormatAudio    PostFormat = "audio"
)
