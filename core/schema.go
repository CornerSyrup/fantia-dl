package core

import "time"

type User struct {
	ID                     int    `json:"id"`
	ToranoanaIdentifyToken string `json:"toranoana_identify_token"`
	Name                   string `json:"name"`
	Image                  struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"image"`
	ProfileText string `json:"profile_text"`
	HasFanclub  bool   `json:"has_fanclub"`
}

type Fanclub struct {
	ID                         int             `json:"id"`
	User                       User            `json:"user"`
	Category                   FanclubCategory `json:"category"`
	Name                       string          `json:"name"`
	CreatorName                string          `json:"creator_name"`
	FanclubName                string          `json:"fanclub_name"`
	FanclubNameWithCreatorName string          `json:"fanclub_name_with_creator_name"`
	FanclubNameOrCreatorName   string          `json:"fanclub_name_or_creator_name"`
	Title                      string          `json:"title"`
	Cover                      struct {
		Thumb    string `json:"thumb"`
		Medium   string `json:"medium"`
		Main     string `json:"main"`
		Ogp      string `json:"ogp"`
		Original string `json:"original"`
	} `json:"cover"`
	Icon struct {
		Thumb    string `json:"thumb"`
		Main     string `json:"main"`
		Original string `json:"original"`
	} `json:"icon"`
	IsJoin        bool `json:"is_join"`
	FanCount      int  `json:"fan_count"`
	PostsCount    int  `json:"posts_count"`
	ProductsCount int  `json:"products_count"`
	URI           struct {
		Show     string `json:"show"`
		Posts    string `json:"posts"`
		Plans    string `json:"plans"`
		Products string `json:"products"`
	} `json:"uri"`
	IsBlocked      bool          `json:"is_blocked"`
	RecentPosts    []interface{} `json:"recent_posts"`
	RecentProducts []interface{} `json:"recent_products"`
	Plans          []FanclubPlan `json:"plans"`
}

type FanclubCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	URI  struct {
		Fanclub  string `json:"fanclub"`
		Products string `json:"products"`
		Posts    string `json:"posts"`
	} `json:"uri"`
}

type FanclubPlan struct {
	ID          int    `json:"id"`
	Price       int    `json:"price"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Limit       int    `json:"limit"`
	Thumb       string `json:"thumb"`
}

type PostContentPhoto struct {
	ID  int `json:"id"`
	URL struct {
		Thumb    string `json:"thumb"`
		Medium   string `json:"medium"`
		Large    string `json:"large"`
		Main     string `json:"main"`
		Micro    string `json:"micro"`
		Original string `json:"original"`
	} `json:"url"`
	Comment         interface{} `json:"comment"`
	ShowOriginalURI string      `json:"show_original_uri"`
	IsConverted     bool        `json:"is_converted"`
}

type BacknumberContent struct {
	ID               int         `json:"id"`
	Title            string      `json:"title"`
	VisibleStatus    string      `json:"visible_status"`
	PublishedState   string      `json:"published_state"`
	Category         string      `json:"category"`
	Comment          string      `json:"comment"`
	EmbedURL         interface{} `json:"embed_url"`
	ContentType      string      `json:"content_type"`
	CommentEndpoints struct {
		PostURI   string `json:"post_uri"`
		DeleteURI string `json:"delete_uri"`
		GetURL    string `json:"get_url"`
	} `json:"comment_endpoints"`
	CommentsReactions struct {
		PostURI   string `json:"post_uri"`
		DeleteURI string `json:"delete_uri"`
		GetURL    string `json:"get_url"`
	} `json:"comments_reactions"`
	EmbedAPIURL string `json:"embed_api_url"`
	Reactions   struct {
		GetURL    string `json:"get_url"`
		PostURI   string `json:"post_uri"`
		DeleteURI string `json:"delete_uri"`
	} `json:"reactions"`
	ReactionTypesURL       string      `json:"reaction_types_url"`
	PostContentPhotosMicro []string    `json:"post_content_photos_micro"`
	Plan                   FanclubPlan `json:"plan"`
	Product                interface{} `json:"product"`
	OnsaleBacknumber       interface{} `json:"onsale_backnumber"`
	BacknumberLink         string      `json:"backnumber_link"`
	JoinStatus             interface{} `json:"join_status"`
	ParentPost             struct {
		Title    string    `json:"title"`
		URL      string    `json:"url"`
		Date     time.Time `json:"date"`
		Deadline time.Time `json:"deadline"`
	} `json:"parent_post"`
	PostContentCommentData struct {
		Comments [][]struct {
			ID              int         `json:"id"`
			Text            string      `json:"text"`
			ParentCommentID interface{} `json:"parent_comment_id"`
			ByOwner         bool        `json:"by_owner"`
			PostedAt        time.Time   `json:"posted_at"`
			Contributor     struct {
				Name          string      `json:"name"`
				IdentifyToken string      `json:"identify_token"`
				Icon          interface{} `json:"icon"`
			} `json:"contributor"`
			Replies []interface{} `json:"replies"`
		} `json:"comments"`
		IsFinish bool `json:"is_finish"`
	} `json:"post_content_comment_data"`
	CommentCount int `json:"comment_count"`

	PostContentPhotos []PostContentPhoto `json:"post_content_photos"`

	IsConverted bool   `json:"is_converted"`
	Filename    string `json:"filename"`
	DownloadURI string `json:"download_uri"`
	HlsURI      string `json:"hls_uri"`
}

type BackNumberApi struct {
	Backnumber struct {
		ID             int         `json:"id"`
		Title          string      `json:"title"`
		Comment        interface{} `json:"comment"`
		Rating         string      `json:"rating"`
		Thumb          interface{} `json:"thumb"`
		ThumbMicro     string      `json:"thumb_micro"`
		ShowAdultThumb bool        `json:"show_adult_thumb"`
		PostedAt       string      `json:"posted_at"`
		LikesCount     int         `json:"likes_count"`
		Liked          bool        `json:"liked"`
		IsContributor  bool        `json:"is_contributor"`
		URI            struct {
			Show string      `json:"show"`
			Edit interface{} `json:"edit"`
		} `json:"uri"`
		IsPulishOpen        bool                `json:"is_pulish_open"`
		IsBlog              bool                `json:"is_blog"`
		ConvertedAt         time.Time           `json:"converted_at"`
		FanclubBrand        int                 `json:"fanclub_brand"`
		SpecialReaction     interface{}         `json:"special_reaction"`
		RedirectURLFromSave string              `json:"redirect_url_from_save"`
		Fanclub             Fanclub             `json:"fanclub"`
		BacknumberContents  []BacknumberContent `json:"backnumber_contents"`
	} `json:"backnumber"`
}

type PostApi struct {
	Post struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Comment string `json:"comment"`
		Rating  string `json:"rating"`
		Thumb   struct {
			Thumb    string `json:"thumb"`
			Medium   string `json:"medium"`
			Large    string `json:"large"`
			Main     string `json:"main"`
			Ogp      string `json:"ogp"`
			Micro    string `json:"micro"`
			Original string `json:"original"`
		} `json:"thumb"`
		ThumbMicro     string `json:"thumb_micro"`
		ShowAdultThumb bool   `json:"show_adult_thumb"`
		PostedAt       string `json:"posted_at"`
		LikesCount     int    `json:"likes_count"`
		Liked          bool   `json:"liked"`
		IsContributor  bool   `json:"is_contributor"`
		URI            struct {
			Show string      `json:"show"`
			Edit interface{} `json:"edit"`
		} `json:"uri"`
		IsPulishOpen        bool        `json:"is_pulish_open"`
		IsBlog              bool        `json:"is_blog"`
		ConvertedAt         time.Time   `json:"converted_at"`
		FanclubBrand        int         `json:"fanclub_brand"`
		SpecialReaction     interface{} `json:"special_reaction"`
		RedirectURLFromSave string      `json:"redirect_url_from_save"`
		Fanclub             struct {
			ID   int `json:"id"`
			User struct {
				ID                     int    `json:"id"`
				ToranoanaIdentifyToken string `json:"toranoana_identify_token"`
				Name                   string `json:"name"`
				Image                  struct {
					Small  string `json:"small"`
					Medium string `json:"medium"`
					Large  string `json:"large"`
				} `json:"image"`
				ProfileText string `json:"profile_text"`
				HasFanclub  bool   `json:"has_fanclub"`
			} `json:"user"`
			Category struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Slug string `json:"slug"`
				URI  struct {
					Fanclub  string `json:"fanclub"`
					Products string `json:"products"`
					Posts    string `json:"posts"`
				} `json:"uri"`
			} `json:"category"`
			Name                       string `json:"name"`
			CreatorName                string `json:"creator_name"`
			FanclubName                string `json:"fanclub_name"`
			FanclubNameWithCreatorName string `json:"fanclub_name_with_creator_name"`
			FanclubNameOrCreatorName   string `json:"fanclub_name_or_creator_name"`
			Title                      string `json:"title"`
			Cover                      struct {
				Thumb    string `json:"thumb"`
				Medium   string `json:"medium"`
				Main     string `json:"main"`
				Ogp      string `json:"ogp"`
				Original string `json:"original"`
			} `json:"cover"`
			Icon struct {
				Thumb    string `json:"thumb"`
				Main     string `json:"main"`
				Original string `json:"original"`
			} `json:"icon"`
			IsJoin        bool `json:"is_join"`
			FanCount      int  `json:"fan_count"`
			PostsCount    int  `json:"posts_count"`
			ProductsCount int  `json:"products_count"`
			URI           struct {
				Show     string `json:"show"`
				Posts    string `json:"posts"`
				Plans    string `json:"plans"`
				Products string `json:"products"`
			} `json:"uri"`
			IsBlocked   bool `json:"is_blocked"`
			RecentPosts []struct {
				ID      int    `json:"id"`
				Title   string `json:"title"`
				Comment string `json:"comment"`
				Rating  string `json:"rating"`
				Thumb   struct {
					Thumb    string `json:"thumb"`
					Medium   string `json:"medium"`
					Large    string `json:"large"`
					Main     string `json:"main"`
					Ogp      string `json:"ogp"`
					Micro    string `json:"micro"`
					Original string `json:"original"`
				} `json:"thumb"`
				ThumbMicro     string `json:"thumb_micro"`
				ShowAdultThumb bool   `json:"show_adult_thumb"`
				PostedAt       string `json:"posted_at"`
				LikesCount     int    `json:"likes_count"`
				Liked          bool   `json:"liked"`
				IsContributor  bool   `json:"is_contributor"`
				URI            struct {
					Show string      `json:"show"`
					Edit interface{} `json:"edit"`
				} `json:"uri"`
				IsPulishOpen    bool      `json:"is_pulish_open"`
				IsBlog          bool      `json:"is_blog"`
				ConvertedAt     time.Time `json:"converted_at"`
				FanclubBrand    int       `json:"fanclub_brand"`
				SpecialReaction struct {
					Reaction    string `json:"reaction"`
					Kind        string `json:"kind"`
					DisplayType string `json:"display_type"`
				} `json:"special_reaction"`
				RedirectURLFromSave string `json:"redirect_url_from_save"`
			} `json:"recent_posts"`
			RecentProducts []interface{} `json:"recent_products"`
			Plans          []struct {
				ID          int         `json:"id"`
				Price       int         `json:"price"`
				Name        string      `json:"name"`
				Description string      `json:"description"`
				Limit       int         `json:"limit"`
				Thumb       string      `json:"thumb"`
				VacantSeat  interface{} `json:"vacant_seat"`
				Order       struct {
					Status     string `json:"status"`
					IsOneclick bool   `json:"is_oneclick"`
					URI        string `json:"uri"`
				} `json:"order"`
			} `json:"plans"`
		} `json:"fanclub"`
		Tags              []interface{}        `json:"tags"`
		Status            string               `json:"status"`
		PostContents      []PostApiPostContent `json:"post_contents"`
		Deadline          string               `json:"deadline"`
		PublishReservedAt string               `json:"publish_reserved_at"`
		CommentEndpoints  struct {
			PostURI   string `json:"post_uri"`
			DeleteURI string `json:"delete_uri"`
			GetURL    string `json:"get_url"`
		} `json:"comment_endpoints"`
		BlogComment       string `json:"blog_comment"`
		CommentsReactions struct {
			PostURI   string `json:"post_uri"`
			DeleteURI string `json:"delete_uri"`
			GetURL    string `json:"get_url"`
		} `json:"comments_reactions"`
		Reactions struct {
			PostURI   string `json:"post_uri"`
			DeleteURI string `json:"delete_uri"`
			GetURL    string `json:"get_url"`
		} `json:"reactions"`
		ReactionTypesURL string `json:"reaction_types_url"`
		OgpAPIURL        string `json:"ogp_api_url"`
		Links            struct {
			Previous struct {
				ID      int    `json:"id"`
				Title   string `json:"title"`
				Comment string `json:"comment"`
				Rating  string `json:"rating"`
				Thumb   struct {
					Thumb    string `json:"thumb"`
					Medium   string `json:"medium"`
					Large    string `json:"large"`
					Main     string `json:"main"`
					Ogp      string `json:"ogp"`
					Micro    string `json:"micro"`
					Original string `json:"original"`
				} `json:"thumb"`
				ThumbMicro     string `json:"thumb_micro"`
				ShowAdultThumb bool   `json:"show_adult_thumb"`
				PostedAt       string `json:"posted_at"`
				LikesCount     int    `json:"likes_count"`
				Liked          bool   `json:"liked"`
				IsContributor  bool   `json:"is_contributor"`
				URI            struct {
					Show string      `json:"show"`
					Edit interface{} `json:"edit"`
				} `json:"uri"`
				IsPulishOpen    bool      `json:"is_pulish_open"`
				IsBlog          bool      `json:"is_blog"`
				ConvertedAt     time.Time `json:"converted_at"`
				FanclubBrand    int       `json:"fanclub_brand"`
				SpecialReaction struct {
					Reaction    string `json:"reaction"`
					Kind        string `json:"kind"`
					DisplayType string `json:"display_type"`
				} `json:"special_reaction"`
				RedirectURLFromSave string `json:"redirect_url_from_save"`
			} `json:"previous"`
			Next struct {
				ID      int         `json:"id"`
				Title   string      `json:"title"`
				Comment interface{} `json:"comment"`
				Rating  string      `json:"rating"`
				Thumb   struct {
					Thumb    string `json:"thumb"`
					Medium   string `json:"medium"`
					Large    string `json:"large"`
					Main     string `json:"main"`
					Ogp      string `json:"ogp"`
					Micro    string `json:"micro"`
					Original string `json:"original"`
				} `json:"thumb"`
				ThumbMicro     string `json:"thumb_micro"`
				ShowAdultThumb bool   `json:"show_adult_thumb"`
				PostedAt       string `json:"posted_at"`
				LikesCount     int    `json:"likes_count"`
				Liked          bool   `json:"liked"`
				IsContributor  bool   `json:"is_contributor"`
				URI            struct {
					Show string      `json:"show"`
					Edit interface{} `json:"edit"`
				} `json:"uri"`
				IsPulishOpen        bool        `json:"is_pulish_open"`
				IsBlog              bool        `json:"is_blog"`
				ConvertedAt         time.Time   `json:"converted_at"`
				FanclubBrand        int         `json:"fanclub_brand"`
				SpecialReaction     interface{} `json:"special_reaction"`
				RedirectURLFromSave string      `json:"redirect_url_from_save"`
			} `json:"next"`
		} `json:"links"`
		IsFanclubTipAccept  bool            `json:"is_fanclub_tip_accept"`
		IsFanclubJoined     bool            `json:"is_fanclub_joined"`
		PostComments        [][]interface{} `json:"post_comments"`
		PostContentComments []struct {
			PostContentID int             `json:"post_content_id"`
			Comments      [][]interface{} `json:"comments"`
			IsFinish      bool            `json:"is_finish"`
		} `json:"post_content_comments"`
		CommentCount       int  `json:"comment_count"`
		CommentBlockStatus bool `json:"comment_block_status"`
	} `json:"post"`
}

type PostApiPostContent struct {
	ID               int         `json:"id"`
	Title            string      `json:"title"`
	VisibleStatus    string      `json:"visible_status"`
	PublishedState   string      `json:"published_state"`
	Category         string      `json:"category"`
	Comment          interface{} `json:"comment,omitempty"`
	EmbedURL         interface{} `json:"embed_url"`
	ContentType      interface{} `json:"content_type"`
	CommentEndpoints struct {
		PostURI   string `json:"post_uri"`
		DeleteURI string `json:"delete_uri"`
		GetURL    string `json:"get_url"`
	} `json:"comment_endpoints"`
	CommentsReactions struct {
		PostURI   string `json:"post_uri"`
		DeleteURI string `json:"delete_uri"`
		GetURL    string `json:"get_url"`
	} `json:"comments_reactions"`
	EmbedAPIURL string `json:"embed_api_url"`
	Reactions   struct {
		GetURL    string `json:"get_url"`
		PostURI   string `json:"post_uri"`
		DeleteURI string `json:"delete_uri"`
	} `json:"reactions"`
	ReactionTypesURL       string             `json:"reaction_types_url"`
	PostContentPhotos      []PostContentPhoto `json:"post_content_photos"`
	PostContentPhotosMicro []string           `json:"post_content_photos_micro"`
	Plan                   struct {
		ID          int    `json:"id"`
		Price       int    `json:"price"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Limit       int    `json:"limit"`
		Thumb       string `json:"thumb"`
	} `json:"plan"`
	Product          interface{} `json:"product"`
	OnsaleBacknumber interface{} `json:"onsale_backnumber"`
	BacknumberLink   interface{} `json:"backnumber_link"`
	JoinStatus       interface{} `json:"join_status"`
	ParentPost       struct {
		Title    string    `json:"title"`
		URL      string    `json:"url"`
		Date     time.Time `json:"date"`
		Deadline time.Time `json:"deadline"`
	} `json:"parent_post"`
	PostContentCommentData struct {
		Comments [][]interface{} `json:"comments"`
		IsFinish bool            `json:"is_finish"`
	} `json:"post_content_comment_data"`
	CommentCount int    `json:"comment_count"`
	IsConverted  bool   `json:"is_converted,omitempty"`
	Filename     string `json:"filename,omitempty"`
	DownloadURI  string `json:"download_uri,omitempty"`
}
