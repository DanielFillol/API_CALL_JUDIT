package models

import "time"

type ResponseToCSV struct {
	Document  string
	TotalTime time.Duration
	R         ResponseJudit
}

type FirstResponse struct {
	RequestId string       `json:"request_id"`
	Search    SearchParams `json:"search"`
	Origin    string       `json:"origin"`
	OriginId  string       `json:"origin_id"`
	Status    string       `json:"status"`
	Tags      Tag
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SearchParams struct {
	SearchType   string   `json:"search_type"`
	SearchKey    string   `json:"search_key"`
	SearchParams struct{} `json:"search_params"`
}

type ResponseJudit struct {
	Page          string        `json:"page"`
	PageData      []PageDataStr `json:"page_data"`
	PageCount     int           `json:"page_count"`
	AllCount      int           `json:"all_count"`
	AllPagesCount int           `json:"all_pages_count"`
}

type PageDataStr struct {
	RequestId    string    `json:"request_id"`
	ResponseId   string    `json:"response_id"`
	Origin       string    `json:"origin"`
	OriginId     string    `json:"origin_id"`
	ResponseType string    `json:"response_type"`
	ResponseData Response  `json:"response_data"`
	UserId       string    `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	Tags         Tag       `json:"tags"`
}

type Response struct {
	Code             string           `json:"code"`
	Justice          string           `json:"justice"`
	Tribunal         string           `json:"tribunal"`
	Instance         int              `json:"instance"`
	DistributionDate time.Time        `json:"distribution_date"`
	TribunalAcronym  string           `json:"tribunal_acronym"`
	SecrecyLevel     int              `json:"secrecy_level"`
	Labels           Label            `json:"tags"`
	Subjects         []Subject        `json:"subjects"`
	Classifications  []Classification `json:"classifications"`
	Courts           []Court          `json:"courts"`
	Parties          []P              `json:"parties"`
	Steps            []Step           `json:"steps"`
	Attachments      []interface{}    `json:"attachments"`      //ignored field
	RelatedLawsuits  []interface{}    `json:"related_lawsuits"` //ignored field
	Crawler          Crawler          `json:"crawler"`
	Amount           float64          `json:"amount,omitempty"`
	LastStep         LastStep         `json:"last_step"`
	Phase            string           `json:"phase"`
	Status           string           `json:"status"`
	Name             string           `json:"name"`
	Judge            interface{}      `json:"judge"` //ignored field
	FreeJustice      bool             `json:"free_justice,omitempty"`
}

type Label struct {
	IsFallbackSource    bool      `json:"is_fallback_source"`
	CrawlId             string    `json:"crawl_id"`
	DictionaryUpdatedAt time.Time `json:"dictionary_updated_at"`
	Criminal            bool      `json:"criminal,omitempty"`
}

type Subject struct {
	Code string    `json:"code"`
	Name string    `json:"name,omitempty"`
	Date time.Time `json:"date,omitempty"`
}

type Classification struct {
	Code string    `json:"code"`
	Name string    `json:"name"`
	Date time.Time `json:"date,omitempty"`
}

type Court struct {
	Code string    `json:"code"`
	Name string    `json:"name"`
	Date time.Time `json:"date,omitempty"`
}

type P struct {
	Name         string     `json:"name"`
	Side         string     `json:"side"`
	PersonType   string     `json:"person_type"`
	MainDocument string     `json:"main_document,omitempty"`
	Documents    []Document `json:"documents,omitempty"`
	Lawyers      []Lawyer   `json:"lawyers,omitempty"`
	Tags         Tag2       `json:"tags"`
	EntityType   string     `json:"entity_type,omitempty"`
}

type Document struct {
	Document     string `json:"document"`
	DocumentType string `json:"document_type"`
}

type Lawyer struct {
	Name         string     `json:"name"`
	Documents    []Document `json:"documents"`
	MainDocument string     `json:"main_document,omitempty"`
}

type Tag2 struct {
	CrawlId string `json:"crawl_id"`
}

type Step struct {
	LawsuitCnj      string    `json:"lawsuit_cnj"`
	LawsuitInstance int       `json:"lawsuit_instance"`
	StepId          string    `json:"step_id"`
	StepDate        time.Time `json:"step_date"`
	Content         string    `json:"content"`
	Private         bool      `json:"private"`
	Tags            Tag2      `json:"tags"`
	StepsCount      int       `json:"steps_count,omitempty"`
}

type Crawler struct {
	SourceName string    `json:"source_name"`
	CrawlId    string    `json:"crawl_id"`
	Weight     int       `json:"weight"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type LastStep struct {
	LawsuitCnj      string    `json:"lawsuit_cnj"`
	LawsuitInstance int       `json:"lawsuit_instance"`
	StepId          string    `json:"step_id"`
	StepDate        time.Time `json:"step_date"`
	Private         bool      `json:"private"`
	StepsCount      int       `json:"steps_count"`
	Content         string    `json:"content,omitempty"`
	Tags            Tag2      `json:"tags,omitempty"`
}

type Tag struct {
	DashboardId      *string `json:"dashboard_id"`
	Name             string  `json:"name,omitempty"`
	HistoricId       string  `json:"historic_id,omitempty"`
	DocumentRelated  string  `json:"document_related,omitempty"`
	ResponseOriginId string  `json:"response_origin_id,omitempty"`
}

//type T struct {
//	Page     int `json:"page"`
//	PageData []struct {
//		RequestId    string `json:"request_id"`
//		ResponseId   string `json:"response_id"`
//		Origin       string `json:"origin"`
//		OriginId     string `json:"origin_id"`
//		ResponseType string `json:"response_type"`
//		ResponseData struct {
//			Code             string    `json:"code"`
//			Justice          string    `json:"justice"`
//			Tribunal         string    `json:"tribunal"`
//			Instance         int       `json:"instance"`
//			DistributionDate time.Time `json:"distribution_date"`
//			TribunalAcronym  string    `json:"tribunal_acronym"`
//			SecrecyLevel     int       `json:"secrecy_level"`
//			Tags             struct {
//				IsFallbackSource    bool      `json:"is_fallback_source"`
//				CrawlId             string    `json:"crawl_id"`
//				DictionaryUpdatedAt time.Time `json:"dictionary_updated_at"`
//				Criminal            bool      `json:"criminal,omitempty"`
//			} `json:"tags"`
//			Subjects []struct {
//				Code string    `json:"code"`
//				Name string    `json:"name,omitempty"`
//				Date time.Time `json:"date,omitempty"`
//			} `json:"subjects"`
//			Classifications []struct {
//				Code string    `json:"code"`
//				Name string    `json:"name"`
//				Date time.Time `json:"date,omitempty"`
//			} `json:"classifications"`
//			Courts []struct {
//				Code string    `json:"code"`
//				Name string    `json:"name"`
//				Date time.Time `json:"date,omitempty"`
//			} `json:"courts"`
//			Parties []struct {
//				Name         string `json:"name"`
//				Side         string `json:"side"`
//				PersonType   string `json:"person_type"`
//				MainDocument string `json:"main_document,omitempty"`
//				Documents    []struct {
//					Document     string `json:"document"`
//					DocumentType string `json:"document_type"`
//				} `json:"documents,omitempty"`
//				Lawyers []struct {
//					Name      string `json:"name"`
//					Documents []struct {
//						Document     string `json:"document"`
//						DocumentType string `json:"document_type"`
//					} `json:"documents"`
//					MainDocument string `json:"main_document,omitempty"`
//				} `json:"lawyers,omitempty"`
//				Tags struct {
//					CrawlId string `json:"crawl_id"`
//				} `json:"tags"`
//				EntityType string `json:"entity_type,omitempty"`
//			} `json:"parties"`
//			Steps []struct {
//				LawsuitCnj      string    `json:"lawsuit_cnj"`
//				LawsuitInstance int       `json:"lawsuit_instance"`
//				StepId          string    `json:"step_id"`
//				StepDate        time.Time `json:"step_date"`
//				Content         string    `json:"content"`
//				Private         bool      `json:"private"`
//				Tags            struct {
//					CrawlId string `json:"crawl_id"`
//				} `json:"tags"`
//				StepsCount int `json:"steps_count,omitempty"`
//			} `json:"steps"`
//			Attachments     []interface{} `json:"attachments"`
//			RelatedLawsuits []interface{} `json:"related_lawsuits"`
//			Crawler         struct {
//				SourceName string    `json:"source_name"`
//				CrawlId    string    `json:"crawl_id"`
//				Weight     int       `json:"weight"`
//				UpdatedAt  time.Time `json:"updated_at"`
//			} `json:"crawler"`
//			Amount   *float64 `json:"amount,omitempty"`
//			LastStep struct {
//				LawsuitCnj      string    `json:"lawsuit_cnj"`
//				LawsuitInstance int       `json:"lawsuit_instance"`
//				StepId          string    `json:"step_id"`
//				StepDate        time.Time `json:"step_date"`
//				Private         bool      `json:"private"`
//				StepsCount      int       `json:"steps_count"`
//				Content         string    `json:"content,omitempty"`
//				Tags            struct {
//					CrawlId string `json:"crawl_id"`
//				} `json:"tags,omitempty"`
//			} `json:"last_step"`
//			Phase       string      `json:"phase"`
//			Status      string      `json:"status"`
//			Name        string      `json:"name"`
//			Judge       interface{} `json:"judge"`
//			FreeJustice bool        `json:"free_justice,omitempty"`
//		} `json:"response_data"`
//		UserId    string    `json:"user_id"`
//		CreatedAt time.Time `json:"created_at"`
//		Tags      struct {
//			DashboardId      *string `json:"dashboard_id"`
//			Name             string  `json:"name,omitempty"`
//			HistoricId       string  `json:"historic_id,omitempty"`
//			DocumentRelated  string  `json:"document_related,omitempty"`
//			ResponseOriginId string  `json:"response_origin_id,omitempty"`
//		} `json:"tags"`
//	} `json:"page_data"`
//	PageCount     s   `json:"page_count"`
//	AllCount      int `json:"all_count"`
//	AllPagesCount int `json:"all_pages_count"`
//}
