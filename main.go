// @title Shopee API Mock Server
// @version 1.0
// @description A mock server for Shopee API endpoints
// @host localhost:3001
// @BasePath /
package main

import (
	"log"
	"strings"

	_ "shopee-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type QueryItem struct {
	OrderSN string `json:"order_sn" example:"220314U0G6UNMN"`
}

type OrderDetailQueryItem struct {
	OrderSNList []string `json:"order_sn_list" example:"220314U0G6UNMN,220315ABC123XYZ"`
}

type GetBuyerInvoiceInfoRequest struct {
	Queries     []QueryItem `json:"queries"`
	PartnerID   int64       `json:"partner_id" query:"partner_id" example:"123456"`
	ShopID      int64       `json:"shop_id" query:"shop_id" example:"789012"`
	Timestamp   int64       `json:"timestamp" query:"timestamp" example:"1640995200"`
	AccessToken string      `json:"access_token" query:"access_token" example:"your_access_token"`
	Sign        string      `json:"sign" query:"sign" example:"ABCD1234567890EFGH"`
}

type GetOrderDetailRequest struct {
	OrderSNList                 string `json:"order_sn_list" query:"order_sn_list" example:"201214JAJXU6G7,201214JASXYXY6"`
	PartnerID                   int64  `json:"partner_id" query:"partner_id" example:"123456"`
	ShopID                      int64  `json:"shop_id" query:"shop_id" example:"789012"`
	Timestamp                   int64  `json:"timestamp" query:"timestamp" example:"1640995200"`
	AccessToken                 string `json:"access_token" query:"access_token" example:"your_access_token"`
	Sign                        string `json:"sign" query:"sign" example:"ABCD1234567890EFGH"`
	RequestOrderStatusPending   bool   `json:"request_order_status_pending" query:"request_order_status_pending" example:"true"`
	ResponseOptionalFields      string `json:"response_optional_fields" query:"response_optional_fields" example:"total_amount"`
}

type GetItemBaseInfoRequest struct {
	ItemIDList          string `json:"item_id_list" query:"item_id_list" example:"[34001,34002]"`
	NeedTaxInfo         bool   `json:"need_tax_info" query:"need_tax_info" example:"true"`
	NeedComplaintPolicy bool   `json:"need_complaint_policy" query:"need_complaint_policy" example:"true"`
	PartnerID           int64  `json:"partner_id" query:"partner_id" example:"123456"`
	ShopID              int64  `json:"shop_id" query:"shop_id" example:"789012"`
	Timestamp           int64  `json:"timestamp" query:"timestamp" example:"1640995200"`
	AccessToken         string `json:"access_token" query:"access_token" example:"your_access_token"`
	Sign                string `json:"sign" query:"sign" example:"ABCD1234567890EFGH"`
}

type AddressBreakdown struct {
	Region          string `json:"region" example:"Poland"`
	State           string `json:"state" example:""`
	City            string `json:"city" example:"Warszawa"`
	District        string `json:"district" example:""`
	Town            string `json:"town" example:"Warszawa"`
	Postcode        string `json:"postcode" example:"51120"`
	DetailedAddress string `json:"detailed_address" example:"Ordona 7B Warszawa"`
	AdditionalInfo  string `json:"additional_info" example:""`
	FullAddress     string `json:"full_address" example:"Ordona 7B Warszawa, Warszawa, 51120"`
}

type InvoiceDetail struct {
	Name             string           `json:"name" example:"Kevin Yap"`
	Email            string           `json:"email" example:"testing.just@op.pl"`
	Address          string           `json:"address" example:"Ordona 7B Warszawa, Warszawa, 51120"`
	PhoneNumber      string           `json:"phone_number" example:"0886761062"`
	TaxID            string           `json:"tax_id" example:"0745561010054"`
	AddressBreakdown AddressBreakdown `json:"address_breakdown"`
}

type InvoiceInfo struct {
	OrderSN       string        `json:"order_sn" example:"2209160VNPKXF7"`
	InvoiceType   string        `json:"invoice_type" example:"personal"`
	InvoiceDetail InvoiceDetail `json:"invoice_detail"`
	IsRequested   bool          `json:"is_requested" example:"false"`
	Error         string        `json:"error" example:""`
}

type GetBuyerInvoiceInfoResponse struct {
	RequestID       string        `json:"request_id" example:"a2c45ca2683caf1651ecab5a4d5942ce"`
	Error           string        `json:"error" example:""`
	Message         string        `json:"message" example:""`
	InvoiceInfoList []InvoiceInfo `json:"invoice_info_list"`
}

type ImageInfo struct {
	ImageURL string `json:"image_url" example:"https://cf.shopee.vn/file/vn-11134207-7qukw-lf6guphtf6oad3_tn"`
}

type OrderItem struct {
	AddOnDeal                bool        `json:"add_on_deal" example:"false"`
	AddOnDealID              int64       `json:"add_on_deal_id" example:"0"`
	ImageInfo                ImageInfo   `json:"image_info"`
	IsB2COwnedItem           bool        `json:"is_b2c_owned_item" example:"false"`
	IsPrescriptionItem       bool        `json:"is_prescription_item" example:"false"`
	ItemID                   int64       `json:"item_id" example:"23620853561"`
	ItemName                 string      `json:"item_name" example:"Minecraft NFA"`
	ItemSKU                  string      `json:"item_sku" example:""`
	MainItem                 bool        `json:"main_item" example:"false"`
	ModelDiscountedPrice     int64       `json:"model_discounted_price" example:"48000"`
	ModelID                  int64       `json:"model_id" example:"221404189791"`
	ModelName                string      `json:"model_name" example:"60g（Papaya）"`
	ModelOriginalPrice       int64       `json:"model_original_price" example:"300000"`
	ModelQuantityPurchased   int         `json:"model_quantity_purchased" example:"1"`
	ModelSKU                 string      `json:"model_sku" example:"QAZ-SADOER-05"`
	OrderItemID              int64       `json:"order_item_id" example:"23620853561"`
	ProductLocationID        []string    `json:"product_location_id" example:"VN10XX2UZ"`
	PromotionGroupID         int64       `json:"promotion_group_id" example:"0"`
	PromotionID              int64       `json:"promotion_id" example:"779222207758537"`
	PromotionType            string      `json:"promotion_type" example:"flash_sale"`
	Weight                   float64     `json:"weight" example:"0.01"`
	Wholesale                bool        `json:"wholesale" example:"false"`
}

type PackageItemDetail struct {
	ItemID              int64  `json:"item_id" example:"23620853561"`
	ModelID             int64  `json:"model_id" example:"221404189791"`
	ModelQuantity       int    `json:"model_quantity" example:"1"`
	OrderItemID         int64  `json:"order_item_id" example:"23620853561"`
	ProductLocationID   string `json:"product_location_id" example:"VN10XX2UZ"`
	PromotionGroupID    int64  `json:"promotion_group_id" example:"0"`
}

type PackageDetail struct {
	GroupShipmentID              *string              `json:"group_shipment_id" example:"null"`
	ItemList                     []PackageItemDetail  `json:"item_list"`
	LogisticsStatus              string               `json:"logistics_status" example:"LOGISTICS_DELIVERY_DONE"`
	PackageNumber                string               `json:"package_number" example:"OFG166300791210964"`
	ParcelChargeableWeightGram   int                  `json:"parcel_chargeable_weight_gram" example:"10"`
	ShippingCarrier              string               `json:"shipping_carrier" example:"5-Day Delivery (SPX)"`
	LogisticsChannelID           int64                `json:"logistics_channel_id" example:"18080"`
	AllowSelfDesignAWB           bool                 `json:"allow_self_design_awb" example:"true"`
	SortingGroup                 string               `json:"sorting_group" example:"North"`
}

type RecipientAddress struct {
	City        string `json:"city" example:"มะกอก"`
	District    string `json:"district" example:"แม่แรง"`
	FullAddress string `json:"full_address" example:"บ้าน******"`
	Name        string `json:"name" example:"P******n"`
	Phone       string `json:"phone" example:"******64"`
	Region      string `json:"region" example:"VN"`
	State       string `json:"state" example:"Bạc Liêu"`
	Town        string `json:"town" example:""`
	Zipcode     string `json:"zipcode" example:""`
}

type OrderDetail struct {
	ActualShippingFeeConfirmed bool              `json:"actual_shipping_fee_confirmed" example:"true"`
	BuyerCancelReason          string            `json:"buyer_cancel_reason" example:""`
	BuyerCpfID                 *string           `json:"buyer_cpf_id" example:"null"`
	BuyerUserID                int64             `json:"buyer_user_id" example:"1170319091"`
	BuyerUsername              string            `json:"buyer_username" example:"xt4fdsf96j"`
	CancelBy                   string            `json:"cancel_by" example:""`
	CancelReason               string            `json:"cancel_reason" example:""`
	COD                        bool              `json:"cod" example:"true"`
	CreateTime                 int64             `json:"create_time" example:"1712601591"`
	Currency                   string            `json:"currency" example:"VND"`
	DaysToShip                 int               `json:"days_to_ship" example:"2"`
	Dropshipper                *string           `json:"dropshipper" example:"null"`
	DropshipperPhone           *string           `json:"dropshipper_phone" example:"null"`
	EstimatedShippingFee       int64             `json:"estimated_shipping_fee" example:"5000"`
	FulfillmentFlag            string            `json:"fulfillment_flag" example:"fulfilled_by_local_seller"`
	GoodsToDeclare             bool              `json:"goods_to_declare" example:"false"`
	InvoiceData                *string           `json:"invoice_data" example:"null"`
	ItemList                   []OrderItem       `json:"item_list"`
	MessageToSeller            string            `json:"message_to_seller" example:""`
	Note                       string            `json:"note" example:""`
	NoteUpdateTime             int64             `json:"note_update_time" example:"0"`
	OrderSN                    string            `json:"order_sn" example:"2404098R48U37H"`
	OrderStatus                string            `json:"order_status" example:"COMPLETED"`
	PackageList                []PackageDetail   `json:"package_list"`
	PayTime                    int64             `json:"pay_time" example:"1712817766"`
	PaymentMethod              string            `json:"payment_method" example:"Cash on Delivery"`
	PickupDoneTime             int64             `json:"pickup_done_time" example:"1712726577"`
	RecipientAddress           RecipientAddress  `json:"recipient_address"`
	Region                     string            `json:"region" example:"VN"`
	ReverseShippingFee         int64             `json:"reverse_shipping_fee" example:"0"`
	ShipByDate                 int64             `json:"ship_by_date" example:"1712671200"`
	ShippingCarrier            string            `json:"shipping_carrier" example:"Thunder Express"`
	SplitUp                    bool              `json:"split_up" example:"false"`
	TotalAmount                int64             `json:"total_amount" example:"32119"`
	UpdateTime                 int64             `json:"update_time" example:"1713139948"`
}

type OrderListResponse struct {
	OrderList []OrderDetail `json:"order_list"`
}

type GetOrderDetailResponse struct {
	Error     string            `json:"error" example:""`
	Message   string            `json:"message" example:""`
	RequestID string            `json:"request_id" example:"023c50ace933ba38473a5fb2a7dc8821"`
	Response  OrderListResponse `json:"response"`
}

type AttributeValue struct {
	ValueID           int64  `json:"value_id" example:"0"`
	OriginalValueName string `json:"original_value_name" example:"Default"`
	ValueUnit         string `json:"value_unit" example:"g"`
}

type Attribute struct {
	AttributeID           int64            `json:"attribute_id" example:"4811"`
	OriginalAttributeName string           `json:"original_attribute_name" example:"Brand: L2 Default [14644]"`
	IsMandatory           bool             `json:"is_mandatory" example:"true"`
	AttributeValueList    []AttributeValue `json:"attribute_value_list"`
}

type PriceInfo struct {
	Currency                        string  `json:"currency" example:"SGD"`
	OriginalPrice                   float64 `json:"original_price" example:"122.02"`
	CurrentPrice                    float64 `json:"current_price" example:"122.02"`
	InflatedPriceOfOriginalPrice    float64 `json:"inflated_price_of_original_price" example:"222.02"`
	InflatedPriceOfCurrentPrice     float64 `json:"inflated_price_of_current_price" example:"111.02"`
	SipItemPrice                    float64 `json:"sip_item_price" example:"100.02"`
	SipItemPriceSource              string  `json:"sip_item_price_source" example:"auto"`
}

type ItemImage struct {
	ImageURLList []string `json:"image_url_list" example:"-"`
	ImageIDList  []string `json:"image_id_list" example:"-"`
}

type Dimension struct {
	PackageLength int `json:"package_length" example:"11"`
	PackageWidth  int `json:"package_width" example:"12"`
	PackageHeight int `json:"package_height" example:"13"`
}

type LogisticInfo struct {
	LogisticID            int64   `json:"logistic_id" example:"80012"`
	LogisticName          string  `json:"logistic_name" example:"-"`
	Enabled               bool    `json:"enabled" example:"true"`
	ShippingFee           float64 `json:"shipping_fee" example:"5.02"`
	SizeID                int64   `json:"size_id" example:"0"`
	IsFree                bool    `json:"is_free" example:"true"`
	EstimatedShippingFee  float64 `json:"estimated_shipping_fee" example:"4.02"`
}

type PreOrder struct {
	IsPreOrder  bool `json:"is_pre_order" example:"true"`
	DaysToShip  int  `json:"days_to_ship" example:"3"`
}

type Wholesale struct {
	MinCount                    int     `json:"min_count" example:"1"`
	MaxCount                    int     `json:"max_count" example:"2"`
	UnitPrice                   float64 `json:"unit_price" example:"4.02"`
	InflatedPriceOfUnitPrice    float64 `json:"inflated_price_of_unit_price" example:"5.02"`
}

type VideoInfo struct {
	VideoURL     string `json:"video_url" example:"-"`
	ThumbnailURL string `json:"thumbnail_url" example:"-"`
	Duration     int    `json:"duration" example:"0"`
}

type Brand struct {
	BrandID            int64  `json:"brand_id" example:"123"`
	OriginalBrandName  string `json:"original_brand_name" example:"nike"`
}

type ComplaintPolicy struct {
	WarrantyTime                 string `json:"warranty_time" example:"ONE_YEAR"`
	ExcludeEntrepreneurWarranty  bool   `json:"exclude_entrepreneur_warranty" example:"true"`
	ComplaintAddressID           int64  `json:"complaint_address_id" example:"0"`
	AdditionalInformation        string `json:"additional_information" example:"-"`
}

type TaxInfo struct {
	NCM            string `json:"ncm" example:"-"`
	DiffStateCFOP  string `json:"diff_state_cfop" example:"-"`
	CSOSN          string `json:"csosn" example:"-"`
	Origin         string `json:"origin" example:"-"`
	CEST           string `json:"cest" example:"-"`
	MeasureUnit    string `json:"measure_unit" example:"-"`
	InvoiceOption  string `json:"invoice_option" example:"-"`
	VATRate        string `json:"vat_rate" example:"-"`
	HSCode         string `json:"hs_code" example:"-"`
	TaxCode        string `json:"tax_code" example:"-"`
}

type DescriptionFieldImage struct {
	ImageID  string `json:"image_id" example:"-"`
	ImageURL string `json:"image_url" example:"-"`
}

type DescriptionField struct {
	FieldType string                `json:"field_type" example:"-"`
	Text      string                `json:"text" example:"-"`
	ImageInfo DescriptionFieldImage `json:"image_info"`
}

type ExtendedDescription struct {
	FieldList []DescriptionField `json:"field_list"`
}

type DescriptionInfo struct {
	ExtendedDescription ExtendedDescription `json:"extended_description"`
}

type SummaryInfo struct {
	TotalReservedStock  int `json:"total_reserved_stock" example:"100"`
	TotalAvailableStock int `json:"total_available_stock" example:"100"`
}

type StockLocation struct {
	LocationID string `json:"location_id" example:"-"`
	Stock      int    `json:"stock" example:"10"`
}

type StockInfoV2 struct {
	SummaryInfo  SummaryInfo     `json:"summary_info"`
	SellerStock  []StockLocation `json:"seller_stock"`
	ShopeeStock  []StockLocation `json:"shopee_stock"`
}

type ItemDetail struct {
	ItemID              int64             `json:"item_id" example:"34002"`
	CategoryID          int64             `json:"category_id" example:"14646"`
	ItemName            string            `json:"item_name" example:"seller discount"`
	Description         string            `json:"description" example:"first product 001first product"`
	ItemSKU             string            `json:"item_sku" example:"-"`
	CreateTime          int64             `json:"create_time" example:"1600572637"`
	UpdateTime          int64             `json:"update_time" example:"1600572640"`
	AttributeList       []Attribute       `json:"attribute_list"`
	PriceInfo           []PriceInfo       `json:"price_info"`
	Image               ItemImage         `json:"image"`
	Weight              string            `json:"weight" example:"10.02"`
	Dimension           Dimension         `json:"dimension"`
	LogisticInfo        []LogisticInfo    `json:"logistic_info"`
	PreOrder            PreOrder          `json:"pre_order"`
	Wholesales          []Wholesale       `json:"wholesales"`
	Condition           string            `json:"condition" example:"NEW/USED"`
	SizeChart           string            `json:"size_chart" example:"-"`
	ItemStatus          string            `json:"item_status" example:"NORMAL"`
	Deboost             string            `json:"deboost" example:"false"`
	HasModel            bool              `json:"has_model" example:"true"`
	PromotionID         int64             `json:"promotion_id" example:"13123"`
	VideoInfo           []VideoInfo       `json:"video_info"`
	Brand               Brand             `json:"brand"`
	ItemDangerous       int               `json:"item_dangerous" example:"0"`
	ComplaintPolicy     ComplaintPolicy   `json:"complaint_policy"`
	TaxInfo             TaxInfo           `json:"tax_info"`
	DescriptionInfo     DescriptionInfo   `json:"description_info"`
	DescriptionType     string            `json:"description_type" example:"-"`
	StockInfoV2         StockInfoV2       `json:"stock_info_v2"`
}

type ItemListResponse struct {
	ItemList []ItemDetail `json:"item_list"`
}

type GetItemBaseInfoResponse struct {
	Error     string           `json:"error" example:"-"`
	Message   string           `json:"message" example:"-"`
	Warning   string           `json:"warning" example:"-"`
	RequestID string           `json:"request_id" example:"7b9da0c6926642199c33ee9dd3a266f5"`
	Response  ItemListResponse `json:"response"`
}

// getBuyerInvoiceInfo retrieves buyer invoice information
// @Summary Get buyer invoice information
// @Description Retrieves buyer invoice information for a specific order
// @Tags Order
// @Accept json
// @Produce json
// @Param partner_id query int64 true "Partner ID" example(123456)
// @Param shop_id query int64 false "Shop ID" example(789012)
// @Param timestamp query int64 true "Request timestamp" example(1640995200)
// @Param access_token query string true "Access token" example("your_access_token")
// @Param sign query string true "Signature" example("ABCD1234567890EFGH")
// @Param request body GetBuyerInvoiceInfoRequest true "Request body with queries array"
// @Success 200 {object} GetBuyerInvoiceInfoResponse "Success response"
// @Failure 400 {object} GetBuyerInvoiceInfoResponse "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /api/v2/order/get_buyer_invoice_info [post]
func getBuyerInvoiceInfo(c *fiber.Ctx) error {
	var req GetBuyerInvoiceInfoRequest

	// Parse query parameters for auth
	if err := c.QueryParser(&req); err != nil {
		return c.Status(400).JSON(GetBuyerInvoiceInfoResponse{
			RequestID: "",
			Error:     "invalid_request",
			Message:   "Invalid query parameters",
		})
	}

	// Parse body for queries
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(GetBuyerInvoiceInfoResponse{
			RequestID: "",
			Error:     "invalid_request",
			Message:   "Invalid request body",
		})
	}

	if len(req.Queries) == 0 || req.Queries[0].OrderSN == "" {
		return c.Status(400).JSON(GetBuyerInvoiceInfoResponse{
			RequestID: "",
			Error:     "missing_order_sn",
			Message:   "Order SN is required in queries array",
		})
	}

	response := GetBuyerInvoiceInfoResponse{
		RequestID: "a2c45ca2683caf1651ecab5a4d5942ce",
		Error:     "",
		Message:   "",
		InvoiceInfoList: []InvoiceInfo{
			{
				OrderSN:     req.Queries[0].OrderSN,
				InvoiceType: "personal",
				InvoiceDetail: InvoiceDetail{
					Name:        "Kevin Yap",
					Email:       "testing.just@op.pl",
					Address:     "Ordona 7B Warszawa, Warszawa, 51120",
					PhoneNumber: "",
					TaxID:       "0745561010054",
					AddressBreakdown: AddressBreakdown{
						Region:          "Poland",
						State:           "",
						City:            "Warszawa",
						District:        "",
						Town:            "Warszawa",
						Postcode:        "51120",
						DetailedAddress: "Ordona 7B Warszawa",
						AdditionalInfo:  "",
						FullAddress:     "Ordona 7B Warszawa, Warszawa, 51120",
					},
				},
				IsRequested: false,
				Error:       "",
			},
		},
	}

	return c.JSON(response)
}

// getOrderDetail retrieves order details
// @Summary Get order details
// @Description Retrieves detailed information about orders
// @Tags Order
// @Accept json
// @Produce json
// @Param order_sn_list query string true "Comma-separated list of order serial numbers" example("201214JAJXU6G7,201214JASXYXY6")
// @Param partner_id query int64 true "Partner ID" example(123456)
// @Param shop_id query int64 false "Shop ID" example(789012)
// @Param timestamp query int64 true "Request timestamp" example(1640995200)
// @Param access_token query string true "Access token" example("your_access_token")
// @Param sign query string true "Signature" example("ABCD1234567890EFGH")
// @Param request_order_status_pending query bool false "Request pending order status" example(true)
// @Param response_optional_fields query string false "Optional response fields" example("total_amount")
// @Success 200 {object} GetOrderDetailResponse "Success response"
// @Failure 400 {object} GetOrderDetailResponse "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /api/v2/order/get_order_detail [get]
func getOrderDetail(c *fiber.Ctx) error {
	var req GetOrderDetailRequest

	// Parse query parameters
	if err := c.QueryParser(&req); err != nil {
		return c.Status(400).JSON(GetOrderDetailResponse{
			RequestID: "",
			Error:     "invalid_request",
			Message:   "Invalid query parameters",
		})
	}

	if req.OrderSNList == "" {
		return c.Status(400).JSON(GetOrderDetailResponse{
			RequestID: "",
			Error:     "missing_order_sn_list",
			Message:   "Order SN list is required",
		})
	}

	// Parse comma-separated order SNs
	orderSNs := strings.Split(req.OrderSNList, ",")
	
	// Create mock orders for each requested order SN
	var orders []OrderDetail
	for _, orderSN := range orderSNs {
		orderSN = strings.TrimSpace(orderSN)
		order := OrderDetail{
			ActualShippingFeeConfirmed: true,
			BuyerCancelReason:          "",
			BuyerCpfID:                 nil,
			BuyerUserID:                1170319091,
			BuyerUsername:              "xt4fdsf96j",
			CancelBy:                   "",
			CancelReason:               "",
			COD:                        true,
			CreateTime:                 1712601591,
			Currency:                   "VND",
			DaysToShip:                 2,
			Dropshipper:                nil,
			DropshipperPhone:           nil,
			EstimatedShippingFee:       5000,
			FulfillmentFlag:            "fulfilled_by_local_seller",
			GoodsToDeclare:             false,
			InvoiceData:                nil,
			ItemList: []OrderItem{
				{
					AddOnDeal:          false,
					AddOnDealID:        0,
					ImageInfo:          ImageInfo{ImageURL: "https://cf.shopee.vn/file/vn-11134207-7qukw-lf6guphtf6oad3_tn"},
					IsB2COwnedItem:     false,
					IsPrescriptionItem: false,
					ItemID:             23620853561,
					ItemName:           "Minecraft NFA",
					ItemSKU:            "",
					MainItem:           false,
					ModelDiscountedPrice: 48000,
					ModelID:            221404189791,
					ModelName:          " Non-refundable",
					ModelOriginalPrice: 300000,
					ModelQuantityPurchased: 1,
					ModelSKU:           "QAZ-SADOER-05",
					OrderItemID:        23620853561,
					ProductLocationID:  []string{"VN10XX2UZ"},
					PromotionGroupID:   0,
					PromotionID:        779222207758537,
					PromotionType:      "flash_sale",
					Weight:             0.01,
					Wholesale:          false,
				},
			},
			MessageToSeller: "",
			Note:            "",
			NoteUpdateTime:  0,
			OrderSN:         orderSN,
			OrderStatus:     "COMPLETED",
			PackageList: []PackageDetail{
				{
					GroupShipmentID: nil,
					ItemList: []PackageItemDetail{
						{
							ItemID:            23620853561,
							ModelID:           221404189791,
							ModelQuantity:     1,
							OrderItemID:       23620853561,
							ProductLocationID: "VN10XX2UZ",
							PromotionGroupID:  0,
						},
					},
					LogisticsStatus:            "LOGISTICS_DELIVERY_DONE",
					PackageNumber:              "OFG166300791210964",
					ParcelChargeableWeightGram: 10,
					ShippingCarrier:            "5-Day Delivery (SPX)",
					LogisticsChannelID:         18080,
					AllowSelfDesignAWB:         true,
					SortingGroup:               "North",
				},
			},
			PayTime:        1712817766,
			PaymentMethod:  "Cash on Delivery",
			PickupDoneTime: 1712726577,
			RecipientAddress: RecipientAddress{
				City:        "มะกอก",
				District:    "แม่แรง",
				FullAddress: "บ้าน******",
				Name:        "P******n",
				Phone:       "******64",
				Region:      "VN",
				State:       "เมือง",
				Town:        "เทส",
				Zipcode:     "51120",
			},
			Region:             "VN",
			ReverseShippingFee: 0,
			ShipByDate:         1712671200,
			ShippingCarrier:    "Thunder Express",
			SplitUp:            false,
			TotalAmount:        32119,
			UpdateTime:         1713139948,
		}
		orders = append(orders, order)
	}

	response := GetOrderDetailResponse{
		Error:     "",
		Message:   "",
		RequestID: "023c50ace933ba38473a5fb2a7dc8821",
		Response: OrderListResponse{
			OrderList: orders,
		},
	}

	return c.JSON(response)
}

// getItemBaseInfo retrieves item base information
// @Summary Get item base information
// @Description Retrieves detailed information about products
// @Tags Product
// @Accept json
// @Produce json
// @Param item_id_list query string true "Array of item IDs" example("[34001,34002]")
// @Param need_tax_info query bool false "Include tax info in response" example(true)
// @Param need_complaint_policy query bool false "Include complaint policy in response" example(true)
// @Param partner_id query int64 true "Partner ID" example(123456)
// @Param shop_id query int64 false "Shop ID" example(789012)
// @Param timestamp query int64 true "Request timestamp" example(1640995200)
// @Param access_token query string true "Access token" example("your_access_token")
// @Param sign query string true "Signature" example("ABCD1234567890EFGH")
// @Success 200 {object} GetItemBaseInfoResponse "Success response"
// @Failure 400 {object} GetItemBaseInfoResponse "Bad request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /api/v2/product/get_item_base_info [get]
func getItemBaseInfo(c *fiber.Ctx) error {
	var req GetItemBaseInfoRequest

	// Parse query parameters
	if err := c.QueryParser(&req); err != nil {
		return c.Status(400).JSON(GetItemBaseInfoResponse{
			Error:     "invalid_request",
			Message:   "Invalid query parameters",
			Warning:   "",
			RequestID: "",
		})
	}

	if req.ItemIDList == "" {
		return c.Status(400).JSON(GetItemBaseInfoResponse{
			Error:     "missing_item_id_list",
			Message:   "Item ID list is required",
			Warning:   "",
			RequestID: "",
		})
	}

	// Create mock item details
	var items []ItemDetail
	
	// For demo, create items for IDs 34001 and 34002
	itemIDs := []int64{34001, 34002}
	
	for _, itemID := range itemIDs {
		item := ItemDetail{
			ItemID:      itemID,
			CategoryID:  14646,
			ItemName:    "seller discount",
			Description: "first product 001first product",
			ItemSKU:     "-",
			CreateTime:  1600572637,
			UpdateTime:  1600572640,
			AttributeList: []Attribute{
				{
					AttributeID:           4811,
					OriginalAttributeName: "Brand: L2 Default [14644]",
					IsMandatory:           true,
					AttributeValueList: []AttributeValue{
						{
							ValueID:           0,
							OriginalValueName: "Default",
							ValueUnit:         "g",
						},
					},
				},
			},
			PriceInfo: []PriceInfo{
				{
					Currency:                     "SGD",
					OriginalPrice:                122.02,
					CurrentPrice:                 122.02,
					InflatedPriceOfOriginalPrice: 222.02,
					InflatedPriceOfCurrentPrice:  111.02,
					SipItemPrice:                 100.02,
					SipItemPriceSource:           "auto",
				},
			},
			Image: ItemImage{
				ImageURLList: []string{"-"},
				ImageIDList:  []string{"-"},
			},
			Weight: "10.02",
			Dimension: Dimension{
				PackageLength: 11,
				PackageWidth:  12,
				PackageHeight: 13,
			},
			LogisticInfo: []LogisticInfo{
				{
					LogisticID:           80012,
					LogisticName:         "-",
					Enabled:              true,
					ShippingFee:          5.02,
					SizeID:               0,
					IsFree:               true,
					EstimatedShippingFee: 4.02,
				},
			},
			PreOrder: PreOrder{
				IsPreOrder: true,
				DaysToShip: 3,
			},
			Wholesales: []Wholesale{
				{
					MinCount:                 1,
					MaxCount:                 2,
					UnitPrice:                4.02,
					InflatedPriceOfUnitPrice: 5.02,
				},
			},
			Condition:     "NEW/USED",
			SizeChart:     "-",
			ItemStatus:    "NORMAL",
			Deboost:       "false",
			HasModel:      true,
			PromotionID:   13123,
			VideoInfo: []VideoInfo{
				{
					VideoURL:     "-",
					ThumbnailURL: "-",
					Duration:     0,
				},
			},
			Brand: Brand{
				BrandID:           123,
				OriginalBrandName: "nike",
			},
			ItemDangerous: 0,
			ComplaintPolicy: ComplaintPolicy{
				WarrantyTime:                "ONE_YEAR",
				ExcludeEntrepreneurWarranty: true,
				ComplaintAddressID:          0,
				AdditionalInformation:       "-",
			},
			TaxInfo: TaxInfo{
				NCM:           "-",
				DiffStateCFOP: "-",
				CSOSN:         "-",
				Origin:        "-",
				CEST:          "-",
				MeasureUnit:   "-",
				InvoiceOption: "-",
				VATRate:       "-",
				HSCode:        "-",
				TaxCode:       "-",
			},
			DescriptionInfo: DescriptionInfo{
				ExtendedDescription: ExtendedDescription{
					FieldList: []DescriptionField{
						{
							FieldType: "-",
							Text:      "-",
							ImageInfo: DescriptionFieldImage{
								ImageID:  "-",
								ImageURL: "-",
							},
						},
					},
				},
			},
			DescriptionType: "-",
			StockInfoV2: StockInfoV2{
				SummaryInfo: SummaryInfo{
					TotalReservedStock:  100,
					TotalAvailableStock: 100,
				},
				SellerStock: []StockLocation{
					{
						LocationID: "-",
						Stock:      10,
					},
				},
				ShopeeStock: []StockLocation{
					{
						LocationID: "-",
						Stock:      0,
					},
				},
			},
		}
		items = append(items, item)
	}

	response := GetItemBaseInfoResponse{
		Error:     "-",
		Message:   "-",
		Warning:   "-",
		RequestID: "7b9da0c6926642199c33ee9dd3a266f5",
		Response: ItemListResponse{
			ItemList: items,
		},
	}

	return c.JSON(response)
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:         "Shopee API Mock Server",
		ReadBufferSize:  16384,
		WriteBufferSize: 16384,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	orderAPI := app.Group("/api/v2/order")
	productAPI := app.Group("/api/v2/product")

	// api.Use(validateTimestamp)
	// api.Use(validateShopeeSignature)

	orderAPI.Post("/get_buyer_invoice_info", getBuyerInvoiceInfo)
	orderAPI.Get("/get_order_detail", getOrderDetail)
	productAPI.Get("/get_item_base_info", getItemBaseInfo)

	log.Println("Starting server on :3001")
	log.Fatal(app.Listen(":3001"))
}
