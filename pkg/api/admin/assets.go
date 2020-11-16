package admin

import (
	"cloudinary-labs/cloudinary-go/pkg/api"
	"context"
	"time"
)

const (
	Assets        api.EndPoint = "resources"
	DerivedAssets api.EndPoint = "derived_resources"
	Tags          api.EndPoint = "tags"
	Context       api.EndPoint = "context"
	Moderations   api.EndPoint = "moderations"
	Restore       api.EndPoint = "restore"
)

func (a *Api) AssetTypes(ctx context.Context) (*AssetTypesResult, error) {
	res := &AssetTypesResult{}
	_, err := a.get(ctx, Assets, nil, res)

	return res, err
}

type AssetTypesResult struct {
	AssetTypes []string      `json:"resource_types"`
	Error      api.ErrorResp `json:"error,omitempty"`
}

type AssetsParams struct {
	AssetType   api.AssetType `json:"-"`
	Prefix      string        `json:"prefix,omitempty"`
	StartAt     *time.Time    `json:"start_at,omitempty"`
	NextCursor  string        `json:"next_cursor,omitempty"`
	MaxResults  int           `json:"max_results,omitempty"`
	Tags        bool          `json:"tags,omitempty"`
	Context     bool          `json:"context,omitempty"`
	Moderations bool          `json:"moderations,omitempty"`
	Direction   string        `json:"direction,omitempty"`
}

func (a *Api) Assets(ctx context.Context, params AssetsParams) (*AssetsResult, error) {
	res := &AssetsResult{}
	_, err := a.get(ctx, api.BuildPath(Assets, params.AssetType.ToString()), params, res)

	return res, err
}

type ListAssetResult struct {
	AssetID     string       `json:"asset_id"`
	PublicID    string       `json:"public_id"`
	Format      string       `json:"format"`
	Version     int          `json:"version"`
	AssetType   string       `json:"resource_type"`
	Type        string       `json:"type"`
	CreatedAt   time.Time    `json:"created_at"`
	Bytes       int          `json:"bytes"`
	Width       int          `json:"width"`
	Height      int          `json:"height"`
	Backup      bool         `json:"backup"`
	AccessMode  string       `json:"access_mode"`
	URL         string       `json:"url"`
	SecureURL   string       `json:"secure_url"`
	Tags        []string     `json:"tags,omitempty"`
	Context     api.Context  `json:"context,omitempty"`
	Metadata    api.Metadata `json:"metadata,omitempty"`
	Placeholder bool         `json:"placeholder,omitempty"`
	Error       string       `json:"error,omitempty"`
}

type AssetsResult struct {
	Assets     []ListAssetResult `json:"resources"`
	NextCursor string            `json:"next_cursor"`
	Error      api.ErrorResp     `json:"error,omitempty"`
}

type AssetsByTagParams struct {
	AssetType   api.AssetType `json:"-"`
	Tag         string        `json:"-"`
	NextCursor  string        `json:"next_cursor,omitempty"`
	MaxResults  int           `json:"max_results,omitempty"`
	Tags        bool          `json:"tags,omitempty"`
	Context     bool          `json:"context,omitempty"`
	Moderations bool          `json:"moderations,omitempty"`
	Direction   string        `json:"direction,omitempty"`
}

func (a *Api) AssetsByTag(ctx context.Context, params AssetsByTagParams) (*AssetsResult, error) {
	res := &AssetsResult{}
	_, err := a.get(ctx, api.BuildPath(Assets, params.AssetType.ToString(), Tags, params.Tag), params, res)

	return res, err
}

type AssetsByContextParams struct {
	AssetType   api.AssetType `json:"-"`
	Key         string        `json:"key"`
	Value       string        `json:"value,omitempty"`
	NextCursor  string        `json:"next_cursor,omitempty"`
	MaxResults  int           `json:"max_results,omitempty"`
	Tags        bool          `json:"tags,omitempty"`
	Context     bool          `json:"context,omitempty"`
	Moderations bool          `json:"moderations,omitempty"`
	Direction   string        `json:"direction,omitempty"`
}

func (a *Api) AssetsByContext(ctx context.Context, params AssetsByContextParams) (*AssetsResult, error) {
	res := &AssetsResult{}
	_, err := a.get(ctx, api.BuildPath(Assets, params.AssetType.ToString(), Context), params, res)

	return res, err
}

type AssetsByModerationParams struct {
	AssetType   api.AssetType `json:"-"`
	Kind        string        `json:"-"`
	Status      string        `json:"-"`
	NextCursor  string        `json:"next_cursor,omitempty"`
	MaxResults  int           `json:"max_results,omitempty"`
	Tags        bool          `json:"tags,omitempty"`
	Context     bool          `json:"context,omitempty"`
	Moderations bool          `json:"moderations,omitempty"`
	Direction   string        `json:"direction,omitempty"`
}

func (a *Api) AssetsByModeration(ctx context.Context, params AssetsByModerationParams) (*AssetsResult, error) {
	res := &AssetsResult{}
	_, err := a.get(ctx, api.BuildPath(Assets, params.AssetType.ToString(), Moderations, params.Kind, params.Status), params, res)

	return res, err
}

type AssetsByIDsParams struct {
	AssetType    api.AssetType    `json:"-"`
	DeliveryType api.DeliveryType `json:"-"`
	PublicIDs    api.CldApiArray  `json:"public_ids"`
	Tags         bool             `json:"tags,omitempty"`
	Context      bool             `json:"context,omitempty"`
	Moderations  bool             `json:"moderations,omitempty"`
}

func (a *Api) AssetsByIDs(ctx context.Context, params AssetsByIDsParams) (*AssetsResult, error) {
	res := &AssetsResult{}
	_, err := a.get(ctx, api.BuildPath(Assets, params.AssetType.ToString(), params.DeliveryType.ToString()), params, res)

	return res, err
}

type RestoreAssetsParams struct {
	AssetType    api.AssetType    `json:"-"`
	DeliveryType api.DeliveryType `json:"-"`
	PublicIDs    api.CldApiArray  `json:"public_ids"`
	Versions     api.CldApiArray  `json:"versions"`
}

func (a *Api) RestoreAssets(ctx context.Context, params RestoreAssetsParams) (*RestoreAssetsResult, error) {
	res := &RestoreAssetsResult{}
	_, err := a.post(ctx, api.BuildPath(Assets, params.AssetType.ToString(), params.DeliveryType.ToString(), Restore), params, res)

	return res, err
}

type RestoreAssetsResult map[string]ListAssetResult

type DeleteAssetsParams struct {
	AssetType       api.AssetType    `json:"-"`
	DeliveryType    api.DeliveryType `json:"-"`
	PublicIDs       api.CldApiArray  `json:"public_ids"`
	KeepOriginal    bool             `json:"keep_original,omitempty"`
	Invalidate      bool             `json:"invalidate,omitempty"`
	Transformations string           `json:"transformations,omitempty"`
	NextCursor      string           `json:"next_cursor,omitempty"`
}

func (a *Api) DeleteAssets(ctx context.Context, params DeleteAssetsParams) (*DeleteAssetsResult, error) {
	res := &DeleteAssetsResult{}
	_, err := a.delete(ctx, api.BuildPath(Assets, params.AssetType.ToString(), params.DeliveryType.ToString()), params, res)

	return res, err
}

type DeleteAssetsResult struct {
	Deleted       map[string]string      `json:"deleted"`
	DeletedCounts map[string]interface{} `json:"deleted_counts"`
	Partial       bool                   `json:"partial"`
	Error         api.ErrorResp          `json:"error,omitempty"`
}

type DeleteAssetsByPrefixParams struct {
	AssetType       api.AssetType    `json:"-"`
	DeliveryType    api.DeliveryType `json:"-"`
	Prefix          api.CldApiArray  `json:"prefix"`
	KeepOriginal    bool             `json:"keep_original,omitempty"`
	Invalidate      bool             `json:"invalidate,omitempty"`
	Transformations string           `json:"transformations,omitempty"`
	NextCursor      string           `json:"next_cursor,omitempty"`
}

func (a *Api) DeleteAssetsByPrefix(ctx context.Context, params DeleteAssetsByPrefixParams) (*DeleteAssetsResult, error) {
	res := &DeleteAssetsResult{}
	_, err := a.delete(ctx, api.BuildPath(Assets, params.AssetType.ToString(), params.DeliveryType.ToString()), params, res)

	return res, err
}

type DeleteAssetsByTagParams struct {
	AssetType       api.AssetType `json:"-"`
	Tag             string        `json:"-"`
	KeepOriginal    bool          `json:"keep_original,omitempty"`
	Invalidate      bool          `json:"invalidate,omitempty"`
	Transformations string        `json:"transformations,omitempty"`
	NextCursor      string        `json:"next_cursor,omitempty"`
}

func (a *Api) DeleteAssetsByTag(ctx context.Context, params DeleteAssetsByTagParams) (*DeleteAssetsResult, error) {
	res := &DeleteAssetsResult{}
	_, err := a.delete(ctx, api.BuildPath(Assets, params.AssetType.ToString(), Tags, params.Tag), params, res)

	return res, err
}

type DeleteAllAssetsParams struct {
	AssetType       api.AssetType    `json:"-"`
	DeliveryType    api.DeliveryType `json:"-"`
	All             bool             `json:"all"`
	KeepOriginal    bool             `json:"keep_original,omitempty"`
	Invalidate      bool             `json:"invalidate,omitempty"`
	Transformations string           `json:"transformations,omitempty"`
	NextCursor      string           `json:"next_cursor,omitempty"`
}

func (a *Api) DeleteAllAssets(ctx context.Context, params DeleteAllAssetsParams) (*DeleteAssetsResult, error) {
	params.All = true

	res := &DeleteAssetsResult{}
	_, err := a.delete(ctx, api.BuildPath(Assets, params.AssetType.ToString(), params.DeliveryType.ToString()), params, res)

	return res, err
}

type DeleteDerivedAssetsParams struct {
	DerivedAssetIDs api.CldApiArray `json:"derived_resource_ids"`
}

func (a *Api) DeleteDerivedAssets(ctx context.Context, params DeleteDerivedAssetsParams) (*DeleteAssetsResult, error) {
	res := &DeleteAssetsResult{}
	_, err := a.delete(ctx, api.BuildPath(DerivedAssets), params, res)

	return res, err
}

type DeleteDerivedAssetsByTransformationParams struct {
	AssetType       api.AssetType    `json:"-"`
	DeliveryType    api.DeliveryType `json:"-"`
	PublicIDs       api.CldApiArray  `json:"public_ids"`
	Transformations string           `json:"transformations"`
	KeepOriginal    bool             `json:"keep_original"`
	Invalidate      bool             `json:"invalidate,omitempty"`
}

func (a *Api) DeleteDerivedAssetsByTransformation(ctx context.Context, params DeleteDerivedAssetsByTransformationParams) (*DeleteAssetsResult, error) {
	params.KeepOriginal = true

	res := &DeleteAssetsResult{}
	_, err := a.delete(ctx, api.BuildPath(Assets, params.AssetType.ToString(), params.DeliveryType.ToString()), params, res)

	return res, err
}