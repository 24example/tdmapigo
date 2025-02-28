package models

//
// General

type SettingsDocumentInfo struct {
	IDs   []int   `json:"ids"`
	Size  float64 `json:"size"`
	Count int     `json:"count"`
}

type SettingsDocuments struct {
	Brochure    SettingsDocumentInfo `json:"brochure"`
	Model3D     SettingsDocumentInfo `json:"model_3d"`
	Passport    SettingsDocumentInfo `json:"passport"`
	Software    SettingsDocumentInfo `json:"software"`
	BimModel    SettingsDocumentInfo `json:"bim_model"`
	Certificate SettingsDocumentInfo `json:"certificate"`
}

type Settings struct {
	Icon      *string            `json:"icon"`
	Documents *SettingsDocuments `json:"documents"`
}

// General
//

//
// Category

type Category struct {
	ID               int       `json:"id"`
	TdmID            string    `json:"tdm_id"`
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	ParentID         *int      `json:"parent_id"`
	Description      *string   `json:"description"`
	IsActive         bool      `json:"is_active"`
	EAC              bool      `json:"eac"`
	ImageIcon        *string   `json:"image_icon"`
	ImageIconSelect  *string   `json:"image_icon_select"`
	ImageCert        *string   `json:"image_cert"`
	ImageCertPreview *string   `json:"image_cert_preview"`
	Image            *string   `json:"image"`
	Order            int       `json:"order"`
	Settings         *Settings `json:"settings"`
	ProductsCount    int       `json:"products_count"`
	Views            int       `json:"views"`
	Depth            int       `json:"depth"`
}

// Category
//

//
// CategoryTree

type CategoryTreeChild struct {
	ID                int                 `json:"id"`
	TdmID             string              `json:"tdm_id"`
	Code              string              `json:"code"`
	Name              string              `json:"name"`
	Slug              string              `json:"slug"`
	ParentID          *int                `json:"parent_id"`
	Description       *string             `json:"description"`
	IsActive          bool                `json:"is_active"`
	EAC               bool                `json:"eac"`
	ImageIcon         *string             `json:"image_icon"`
	ImageIconSelect   *string             `json:"image_icon_select"`
	ImageCert         *string             `json:"image_cert"`
	ImageCertPreview  *string             `json:"image_cert_preview"`
	Image             *string             `json:"image"`
	Order             int                 `json:"order"`
	Settings          *Settings           `json:"settings"`
	ProductsCount     int                 `json:"products_count"`
	Views             int                 `json:"views"`
	Depth             int                 `json:"depth"`
	ChildrenRecursive []CategoryTreeChild `json:"children_recursive"`
	Parent            *Category           `json:"parent"`
}

type CategoryTree struct {
	ID                int                 `json:"id"`
	TdmID             string              `json:"tdm_id"`
	Code              string              `json:"code"`
	Name              string              `json:"name"`
	Slug              string              `json:"slug"`
	ParentID          *int                `json:"parent_id"`
	Description       *string             `json:"description"`
	IsActive          bool                `json:"is_active"`
	EAC               bool                `json:"eac"`
	ImageIcon         *string             `json:"image_icon"`
	ImageIconSelect   *string             `json:"image_icon_select"`
	ImageCert         *string             `json:"image_cert"`
	ImageCertPreview  *string             `json:"image_cert_preview"`
	Image             *string             `json:"image"`
	Order             int                 `json:"order"`
	Settings          *Settings           `json:"settings"`
	ProductsCount     int                 `json:"products_count"`
	Views             int                 `json:"views"`
	Depth             int                 `json:"depth"`
	Parent            *Category           `json:"parent"`
	ChildrenRecursive []CategoryTreeChild `json:"children_recursive"`
}

// CategoryTree
//

//
// CategoryTreeID

type CategoryTreeIDChild struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	Slug              string              `json:"slug"`
	ParentID          *int                `json:"parent_id"`
	ImageIcon         string              `json:"image_icon"`
	Image             *string             `json:"image"`
	TdmID             string              `json:"tdm_id"`
	ChildrenRecursive []CategoryTreeChild `json:"children_recursive"`
	Parent            *Category           `json:"parent"`
}

type CategoryTreeID struct {
	ID                int                   `json:"id"`
	Name              string                `json:"name"`
	Slug              string                `json:"slug"`
	ParentID          *int                  `json:"parent_id"`
	ImageIcon         string                `json:"image_icon"`
	Image             *string               `json:"image"`
	TdmID             string                `json:"tdm_id"`
	ChildrenRecursive []CategoryTreeIDChild `json:"children_recursive"`
}

// CategoryTreeID
//

//
// CategoryTreeParentsID

type CategoryTreeParentsID struct {
	ID              int                    `json:"id"`
	Name            string                 `json:"name"`
	Slug            string                 `json:"slug"`
	ParentID        *int                   `json:"parent_id"`
	ImageIcon       string                 `json:"image_icon"`
	Image           *string                `json:"image"`
	TDMID           string                 `json:"tdm_id"`
	ParentRecursive *CategoryTreeParentsID `json:"parent_recursive"`
}

//CategoryTreeParentsID
//
