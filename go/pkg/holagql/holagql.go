// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package holagql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// GetCollectionActivityCollection includes the requested fields of the GraphQL type Collection.
type GetCollectionActivityCollection struct {
	Activities []GetCollectionActivityCollectionActivitiesNftActivity `json:"activities"`
}

// GetActivities returns GetCollectionActivityCollection.Activities, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollection) GetActivities() []GetCollectionActivityCollectionActivitiesNftActivity {
	return v.Activities
}

// GetCollectionActivityCollectionActivitiesNftActivity includes the requested fields of the GraphQL type NftActivity.
type GetCollectionActivityCollectionActivitiesNftActivity struct {
	ActivityType              string                                                              `json:"activityType"`
	Price                     string                                                              `json:"price"`
	Metadata                  string                                                              `json:"metadata"`
	CreatedAt                 string                                                              `json:"createdAt"`
	MarketplaceProgramAddress string                                                              `json:"marketplaceProgramAddress"`
	Nft                       GetCollectionActivityCollectionActivitiesNftActivityNft             `json:"nft"`
	Wallets                   []GetCollectionActivityCollectionActivitiesNftActivityWalletsWallet `json:"wallets"`
}

// GetActivityType returns GetCollectionActivityCollectionActivitiesNftActivity.ActivityType, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollectionActivitiesNftActivity) GetActivityType() string {
	return v.ActivityType
}

// GetPrice returns GetCollectionActivityCollectionActivitiesNftActivity.Price, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollectionActivitiesNftActivity) GetPrice() string { return v.Price }

// GetMetadata returns GetCollectionActivityCollectionActivitiesNftActivity.Metadata, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollectionActivitiesNftActivity) GetMetadata() string {
	return v.Metadata
}

// GetCreatedAt returns GetCollectionActivityCollectionActivitiesNftActivity.CreatedAt, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollectionActivitiesNftActivity) GetCreatedAt() string {
	return v.CreatedAt
}

// GetMarketplaceProgramAddress returns GetCollectionActivityCollectionActivitiesNftActivity.MarketplaceProgramAddress, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollectionActivitiesNftActivity) GetMarketplaceProgramAddress() string {
	return v.MarketplaceProgramAddress
}

// GetNft returns GetCollectionActivityCollectionActivitiesNftActivity.Nft, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollectionActivitiesNftActivity) GetNft() GetCollectionActivityCollectionActivitiesNftActivityNft {
	return v.Nft
}

// GetWallets returns GetCollectionActivityCollectionActivitiesNftActivity.Wallets, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollectionActivitiesNftActivity) GetWallets() []GetCollectionActivityCollectionActivitiesNftActivityWalletsWallet {
	return v.Wallets
}

// GetCollectionActivityCollectionActivitiesNftActivityNft includes the requested fields of the GraphQL type Nft.
type GetCollectionActivityCollectionActivitiesNftActivityNft struct {
	Name        string `json:"name"`
	MintAddress string `json:"mintAddress"`
	Image       string `json:"image"`
}

// GetName returns GetCollectionActivityCollectionActivitiesNftActivityNft.Name, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollectionActivitiesNftActivityNft) GetName() string { return v.Name }

// GetMintAddress returns GetCollectionActivityCollectionActivitiesNftActivityNft.MintAddress, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollectionActivitiesNftActivityNft) GetMintAddress() string {
	return v.MintAddress
}

// GetImage returns GetCollectionActivityCollectionActivitiesNftActivityNft.Image, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollectionActivitiesNftActivityNft) GetImage() string { return v.Image }

// GetCollectionActivityCollectionActivitiesNftActivityWalletsWallet includes the requested fields of the GraphQL type Wallet.
type GetCollectionActivityCollectionActivitiesNftActivityWalletsWallet struct {
	Address string `json:"address"`
}

// GetAddress returns GetCollectionActivityCollectionActivitiesNftActivityWalletsWallet.Address, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityCollectionActivitiesNftActivityWalletsWallet) GetAddress() string {
	return v.Address
}

// GetCollectionActivityResponse is returned by GetCollectionActivity on success.
type GetCollectionActivityResponse struct {
	// Returns collection data along with collection activities
	Collection GetCollectionActivityCollection `json:"collection"`
}

// GetCollection returns GetCollectionActivityResponse.Collection, and is useful for accessing the field via an interface.
func (v *GetCollectionActivityResponse) GetCollection() GetCollectionActivityCollection {
	return v.Collection
}

// GetCollectionNFTsNftsNft includes the requested fields of the GraphQL type Nft.
type GetCollectionNFTsNftsNft struct {
	Name        string                                      `json:"name"`
	MintAddress string                                      `json:"mintAddress"`
	Image       string                                      `json:"image"`
	Listings    []GetCollectionNFTsNftsNftListingsAhListing `json:"listings"`
}

// GetName returns GetCollectionNFTsNftsNft.Name, and is useful for accessing the field via an interface.
func (v *GetCollectionNFTsNftsNft) GetName() string { return v.Name }

// GetMintAddress returns GetCollectionNFTsNftsNft.MintAddress, and is useful for accessing the field via an interface.
func (v *GetCollectionNFTsNftsNft) GetMintAddress() string { return v.MintAddress }

// GetImage returns GetCollectionNFTsNftsNft.Image, and is useful for accessing the field via an interface.
func (v *GetCollectionNFTsNftsNft) GetImage() string { return v.Image }

// GetListings returns GetCollectionNFTsNftsNft.Listings, and is useful for accessing the field via an interface.
func (v *GetCollectionNFTsNftsNft) GetListings() []GetCollectionNFTsNftsNftListingsAhListing {
	return v.Listings
}

// GetCollectionNFTsNftsNftListingsAhListing includes the requested fields of the GraphQL type AhListing.
type GetCollectionNFTsNftsNftListingsAhListing struct {
	Price string `json:"price"`
}

// GetPrice returns GetCollectionNFTsNftsNftListingsAhListing.Price, and is useful for accessing the field via an interface.
func (v *GetCollectionNFTsNftsNftListingsAhListing) GetPrice() string { return v.Price }

// GetCollectionNFTsResponse is returned by GetCollectionNFTs on success.
type GetCollectionNFTsResponse struct {
	Nfts []GetCollectionNFTsNftsNft `json:"nfts"`
}

// GetNfts returns GetCollectionNFTsResponse.Nfts, and is useful for accessing the field via an interface.
func (v *GetCollectionNFTsResponse) GetNfts() []GetCollectionNFTsNftsNft { return v.Nfts }

// GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollection includes the requested fields of the GraphQL type Collection.
type GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollection struct {
	Nft GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft `json:"nft"`
}

// GetNft returns GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollection.Nft, and is useful for accessing the field via an interface.
func (v *GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollection) GetNft() GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft {
	return v.Nft
}

// GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft includes the requested fields of the GraphQL type Nft.
type GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft struct {
	Name        string                                                                                   `json:"name"`
	Address     string                                                                                   `json:"address"`
	MintAddress string                                                                                   `json:"mintAddress"`
	Image       string                                                                                   `json:"image"`
	Creators    []GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNftCreatorsNftCreator `json:"creators"`
}

// GetName returns GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft.Name, and is useful for accessing the field via an interface.
func (v *GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft) GetName() string {
	return v.Name
}

// GetAddress returns GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft.Address, and is useful for accessing the field via an interface.
func (v *GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft) GetAddress() string {
	return v.Address
}

// GetMintAddress returns GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft.MintAddress, and is useful for accessing the field via an interface.
func (v *GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft) GetMintAddress() string {
	return v.MintAddress
}

// GetImage returns GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft.Image, and is useful for accessing the field via an interface.
func (v *GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft) GetImage() string {
	return v.Image
}

// GetCreators returns GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft.Creators, and is useful for accessing the field via an interface.
func (v *GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNft) GetCreators() []GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNftCreatorsNftCreator {
	return v.Creators
}

// GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNftCreatorsNftCreator includes the requested fields of the GraphQL type NftCreator.
type GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNftCreatorsNftCreator struct {
	Address         string `json:"address"`
	MetadataAddress string `json:"metadataAddress"`
}

// GetAddress returns GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNftCreatorsNftCreator.Address, and is useful for accessing the field via an interface.
func (v *GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNftCreatorsNftCreator) GetAddress() string {
	return v.Address
}

// GetMetadataAddress returns GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNftCreatorsNftCreator.MetadataAddress, and is useful for accessing the field via an interface.
func (v *GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollectionNftCreatorsNftCreator) GetMetadataAddress() string {
	return v.MetadataAddress
}

// GetCollectionsByMarketCapResponse is returned by GetCollectionsByMarketCap on success.
type GetCollectionsByMarketCapResponse struct {
	// Returns featured collection NFTs ordered by market cap (floor price * number of NFTs in collection)
	CollectionsFeaturedByMarketCap []GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollection `json:"collectionsFeaturedByMarketCap"`
}

// GetCollectionsFeaturedByMarketCap returns GetCollectionsByMarketCapResponse.CollectionsFeaturedByMarketCap, and is useful for accessing the field via an interface.
func (v *GetCollectionsByMarketCapResponse) GetCollectionsFeaturedByMarketCap() []GetCollectionsByMarketCapCollectionsFeaturedByMarketCapCollection {
	return v.CollectionsFeaturedByMarketCap
}

// GetCollectionsByVolumeCollectionsFeaturedByVolumeCollection includes the requested fields of the GraphQL type Collection.
type GetCollectionsByVolumeCollectionsFeaturedByVolumeCollection struct {
	Nft GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft `json:"nft"`
}

// GetNft returns GetCollectionsByVolumeCollectionsFeaturedByVolumeCollection.Nft, and is useful for accessing the field via an interface.
func (v *GetCollectionsByVolumeCollectionsFeaturedByVolumeCollection) GetNft() GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft {
	return v.Nft
}

// GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft includes the requested fields of the GraphQL type Nft.
type GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft struct {
	Name        string                                                                             `json:"name"`
	Address     string                                                                             `json:"address"`
	MintAddress string                                                                             `json:"mintAddress"`
	Image       string                                                                             `json:"image"`
	Creators    []GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNftCreatorsNftCreator `json:"creators"`
}

// GetName returns GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft.Name, and is useful for accessing the field via an interface.
func (v *GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft) GetName() string {
	return v.Name
}

// GetAddress returns GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft.Address, and is useful for accessing the field via an interface.
func (v *GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft) GetAddress() string {
	return v.Address
}

// GetMintAddress returns GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft.MintAddress, and is useful for accessing the field via an interface.
func (v *GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft) GetMintAddress() string {
	return v.MintAddress
}

// GetImage returns GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft.Image, and is useful for accessing the field via an interface.
func (v *GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft) GetImage() string {
	return v.Image
}

// GetCreators returns GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft.Creators, and is useful for accessing the field via an interface.
func (v *GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNft) GetCreators() []GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNftCreatorsNftCreator {
	return v.Creators
}

// GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNftCreatorsNftCreator includes the requested fields of the GraphQL type NftCreator.
type GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNftCreatorsNftCreator struct {
	Address         string `json:"address"`
	MetadataAddress string `json:"metadataAddress"`
}

// GetAddress returns GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNftCreatorsNftCreator.Address, and is useful for accessing the field via an interface.
func (v *GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNftCreatorsNftCreator) GetAddress() string {
	return v.Address
}

// GetMetadataAddress returns GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNftCreatorsNftCreator.MetadataAddress, and is useful for accessing the field via an interface.
func (v *GetCollectionsByVolumeCollectionsFeaturedByVolumeCollectionNftCreatorsNftCreator) GetMetadataAddress() string {
	return v.MetadataAddress
}

// GetCollectionsByVolumeResponse is returned by GetCollectionsByVolume on success.
type GetCollectionsByVolumeResponse struct {
	// Returns featured collection NFTs ordered by volume (sum of purchase prices)
	CollectionsFeaturedByVolume []GetCollectionsByVolumeCollectionsFeaturedByVolumeCollection `json:"collectionsFeaturedByVolume"`
}

// GetCollectionsFeaturedByVolume returns GetCollectionsByVolumeResponse.CollectionsFeaturedByVolume, and is useful for accessing the field via an interface.
func (v *GetCollectionsByVolumeResponse) GetCollectionsFeaturedByVolume() []GetCollectionsByVolumeCollectionsFeaturedByVolumeCollection {
	return v.CollectionsFeaturedByVolume
}

// GetCreatorNameMetadataJsonsMetadataJson includes the requested fields of the GraphQL type MetadataJson.
type GetCreatorNameMetadataJsonsMetadataJson struct {
	Name string `json:"name"`
}

// GetName returns GetCreatorNameMetadataJsonsMetadataJson.Name, and is useful for accessing the field via an interface.
func (v *GetCreatorNameMetadataJsonsMetadataJson) GetName() string { return v.Name }

// GetCreatorNameResponse is returned by GetCreatorName on success.
type GetCreatorNameResponse struct {
	// returns metadata_jsons matching the term
	MetadataJsons []GetCreatorNameMetadataJsonsMetadataJson `json:"metadataJsons"`
}

// GetMetadataJsons returns GetCreatorNameResponse.MetadataJsons, and is useful for accessing the field via an interface.
func (v *GetCreatorNameResponse) GetMetadataJsons() []GetCreatorNameMetadataJsonsMetadataJson {
	return v.MetadataJsons
}

// Sorts results ascending or descending
type OrderDirection string

const (
	OrderDirectionDesc OrderDirection = "DESC"
	OrderDirectionAsc  OrderDirection = "ASC"
)

// __GetCollectionActivityInput is used internally by genqlient
type __GetCollectionActivityInput struct {
	MintAddress string `json:"MintAddress"`
	Limit       int    `json:"Limit"`
	Offset      int    `json:"Offset"`
}

// GetMintAddress returns __GetCollectionActivityInput.MintAddress, and is useful for accessing the field via an interface.
func (v *__GetCollectionActivityInput) GetMintAddress() string { return v.MintAddress }

// GetLimit returns __GetCollectionActivityInput.Limit, and is useful for accessing the field via an interface.
func (v *__GetCollectionActivityInput) GetLimit() int { return v.Limit }

// GetOffset returns __GetCollectionActivityInput.Offset, and is useful for accessing the field via an interface.
func (v *__GetCollectionActivityInput) GetOffset() int { return v.Offset }

// __GetCollectionNFTsInput is used internally by genqlient
type __GetCollectionNFTsInput struct {
	MintAddress string `json:"MintAddress"`
	Limit       int    `json:"Limit"`
	Offset      int    `json:"Offset"`
}

// GetMintAddress returns __GetCollectionNFTsInput.MintAddress, and is useful for accessing the field via an interface.
func (v *__GetCollectionNFTsInput) GetMintAddress() string { return v.MintAddress }

// GetLimit returns __GetCollectionNFTsInput.Limit, and is useful for accessing the field via an interface.
func (v *__GetCollectionNFTsInput) GetLimit() int { return v.Limit }

// GetOffset returns __GetCollectionNFTsInput.Offset, and is useful for accessing the field via an interface.
func (v *__GetCollectionNFTsInput) GetOffset() int { return v.Offset }

// __GetCollectionsByMarketCapInput is used internally by genqlient
type __GetCollectionsByMarketCapInput struct {
	OrderDirection OrderDirection `json:"OrderDirection"`
	StartDate      string         `json:"StartDate"`
	EndDate        string         `json:"EndDate"`
	Limit          int            `json:"Limit"`
	Offset         int            `json:"Offset"`
}

// GetOrderDirection returns __GetCollectionsByMarketCapInput.OrderDirection, and is useful for accessing the field via an interface.
func (v *__GetCollectionsByMarketCapInput) GetOrderDirection() OrderDirection {
	return v.OrderDirection
}

// GetStartDate returns __GetCollectionsByMarketCapInput.StartDate, and is useful for accessing the field via an interface.
func (v *__GetCollectionsByMarketCapInput) GetStartDate() string { return v.StartDate }

// GetEndDate returns __GetCollectionsByMarketCapInput.EndDate, and is useful for accessing the field via an interface.
func (v *__GetCollectionsByMarketCapInput) GetEndDate() string { return v.EndDate }

// GetLimit returns __GetCollectionsByMarketCapInput.Limit, and is useful for accessing the field via an interface.
func (v *__GetCollectionsByMarketCapInput) GetLimit() int { return v.Limit }

// GetOffset returns __GetCollectionsByMarketCapInput.Offset, and is useful for accessing the field via an interface.
func (v *__GetCollectionsByMarketCapInput) GetOffset() int { return v.Offset }

// __GetCollectionsByVolumeInput is used internally by genqlient
type __GetCollectionsByVolumeInput struct {
	OrderDirection OrderDirection `json:"OrderDirection"`
	StartDate      string         `json:"StartDate"`
	EndDate        string         `json:"EndDate"`
	Limit          int            `json:"Limit"`
	Offset         int            `json:"Offset"`
}

// GetOrderDirection returns __GetCollectionsByVolumeInput.OrderDirection, and is useful for accessing the field via an interface.
func (v *__GetCollectionsByVolumeInput) GetOrderDirection() OrderDirection { return v.OrderDirection }

// GetStartDate returns __GetCollectionsByVolumeInput.StartDate, and is useful for accessing the field via an interface.
func (v *__GetCollectionsByVolumeInput) GetStartDate() string { return v.StartDate }

// GetEndDate returns __GetCollectionsByVolumeInput.EndDate, and is useful for accessing the field via an interface.
func (v *__GetCollectionsByVolumeInput) GetEndDate() string { return v.EndDate }

// GetLimit returns __GetCollectionsByVolumeInput.Limit, and is useful for accessing the field via an interface.
func (v *__GetCollectionsByVolumeInput) GetLimit() int { return v.Limit }

// GetOffset returns __GetCollectionsByVolumeInput.Offset, and is useful for accessing the field via an interface.
func (v *__GetCollectionsByVolumeInput) GetOffset() int { return v.Offset }

// __GetCreatorNameInput is used internally by genqlient
type __GetCreatorNameInput struct {
	Address string `json:"Address"`
}

// GetAddress returns __GetCreatorNameInput.Address, and is useful for accessing the field via an interface.
func (v *__GetCreatorNameInput) GetAddress() string { return v.Address }

func GetCollectionActivity(
	ctx context.Context,
	client graphql.Client,
	MintAddress string,
	Limit int,
	Offset int,
) (*GetCollectionActivityResponse, error) {
	req := &graphql.Request{
		OpName: "GetCollectionActivity",
		Query: `
query GetCollectionActivity ($MintAddress: String!, $Limit: Int!, $Offset: Int!) {
	collection(address: $MintAddress) {
		activities(limit: $Limit, offset: $Offset) {
			activityType
			price
			metadata
			createdAt
			marketplaceProgramAddress
			nft {
				name
				mintAddress
				image(width: 400)
			}
			wallets {
				address
			}
		}
	}
}
`,
		Variables: &__GetCollectionActivityInput{
			MintAddress: MintAddress,
			Limit:       Limit,
			Offset:      Offset,
		},
	}
	var err error

	var data GetCollectionActivityResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetCollectionNFTs(
	ctx context.Context,
	client graphql.Client,
	MintAddress string,
	Limit int,
	Offset int,
) (*GetCollectionNFTsResponse, error) {
	req := &graphql.Request{
		OpName: "GetCollectionNFTs",
		Query: `
query GetCollectionNFTs ($MintAddress: PublicKey!, $Limit: Int!, $Offset: Int!) {
	nfts(collection: $MintAddress, limit: $Limit, offset: $Offset) {
		name
		mintAddress
		image(width: 400)
		listings {
			price
		}
	}
}
`,
		Variables: &__GetCollectionNFTsInput{
			MintAddress: MintAddress,
			Limit:       Limit,
			Offset:      Offset,
		},
	}
	var err error

	var data GetCollectionNFTsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetCollectionsByMarketCap(
	ctx context.Context,
	client graphql.Client,
	OrderDirection OrderDirection,
	StartDate string,
	EndDate string,
	Limit int,
	Offset int,
) (*GetCollectionsByMarketCapResponse, error) {
	req := &graphql.Request{
		OpName: "GetCollectionsByMarketCap",
		Query: `
query GetCollectionsByMarketCap ($OrderDirection: OrderDirection!, $StartDate: DateTimeUtc!, $EndDate: DateTimeUtc!, $Limit: Int!, $Offset: Int!) {
	collectionsFeaturedByMarketCap(orderDirection: $OrderDirection, startDate: $StartDate, endDate: $EndDate, limit: $Limit, offset: $Offset) {
		nft {
			name
			address
			mintAddress
			image(width: 400)
			creators {
				address
				metadataAddress
			}
		}
	}
}
`,
		Variables: &__GetCollectionsByMarketCapInput{
			OrderDirection: OrderDirection,
			StartDate:      StartDate,
			EndDate:        EndDate,
			Limit:          Limit,
			Offset:         Offset,
		},
	}
	var err error

	var data GetCollectionsByMarketCapResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetCollectionsByVolume(
	ctx context.Context,
	client graphql.Client,
	OrderDirection OrderDirection,
	StartDate string,
	EndDate string,
	Limit int,
	Offset int,
) (*GetCollectionsByVolumeResponse, error) {
	req := &graphql.Request{
		OpName: "GetCollectionsByVolume",
		Query: `
query GetCollectionsByVolume ($OrderDirection: OrderDirection!, $StartDate: DateTimeUtc!, $EndDate: DateTimeUtc!, $Limit: Int!, $Offset: Int!) {
	collectionsFeaturedByVolume(orderDirection: $OrderDirection, startDate: $StartDate, endDate: $EndDate, limit: $Limit, offset: $Offset) {
		nft {
			name
			address
			mintAddress
			image(width: 400)
			creators {
				address
				metadataAddress
			}
		}
	}
}
`,
		Variables: &__GetCollectionsByVolumeInput{
			OrderDirection: OrderDirection,
			StartDate:      StartDate,
			EndDate:        EndDate,
			Limit:          Limit,
			Offset:         Offset,
		},
	}
	var err error

	var data GetCollectionsByVolumeResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetCreatorName(
	ctx context.Context,
	client graphql.Client,
	Address string,
) (*GetCreatorNameResponse, error) {
	req := &graphql.Request{
		OpName: "GetCreatorName",
		Query: `
query GetCreatorName ($Address: String!) {
	metadataJsons(term: $Address, limit: 1, offset: 0) {
		name
	}
}
`,
		Variables: &__GetCreatorNameInput{
			Address: Address,
		},
	}
	var err error

	var data GetCreatorNameResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
